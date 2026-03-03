package robot

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"cyberstrike-ai/internal/config"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
	"go.uber.org/zap"
)

const (
	larkReconnectInitial = 5 * time.Second  // Intervalle de reconnexion initiale
	larkReconnectMax     = 60 * time.Second // Intervalle de reconnexion maximale
)

type larkTextContent struct {
	Text string `json:"text"`
}

// StartLark démarre la connexion longue Feishu (sans besoin de réseau public), appelle le handler après réception des messages et répond.
// Se reconnecte automatiquement après déconnexion (comme la mise en veille de l'ordinateur portable, interruption réseau) ; quitte lorsque ctx est annulé, pratique pour redémarrer lors de changements de configuration.
func StartLark(ctx context.Context, cfg config.RobotLarkConfig, h MessageHandler, logger *zap.Logger) {
	if !cfg.Enabled || cfg.AppID == "" || cfg.AppSecret == "" {
		return
	}
	go runLarkLoop(ctx, cfg, h, logger)
}

// runLarkLoop maintient la connexion longue Feishu en boucle : se reconnecte selon l'intervalle de backoff lorsque déconnecté et que ctx n'est pas annulé.
func runLarkLoop(ctx context.Context, cfg config.RobotLarkConfig, h MessageHandler, logger *zap.Logger) {
	backoff := larkReconnectInitial
	for {
		larkClient := lark.NewClient(cfg.AppID, cfg.AppSecret)
		eventHandler := dispatcher.NewEventDispatcher("", "").OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
			go handleLarkMessage(ctx, event, h, larkClient, logger)
			return nil
		})
		wsClient := larkws.NewClient(cfg.AppID, cfg.AppSecret,
			larkws.WithEventHandler(eventHandler),
			larkws.WithLogLevel(larkcore.LogLevelInfo),
		)
		logger.Info("Connexion longue Feishu en cours de connexion…", zap.String("app_id", cfg.AppID))
		err := wsClient.Start(ctx)
		if ctx.Err() != nil {
			logger.Info("Connexion longue Feishu fermée selon la reconfiguration")
			return
		}
		if err != nil {
			logger.Warn("Connexion longue Feishu déconnectée (veille/déconnexion réseau), reconnexion automatique", zap.Error(err), zap.Duration("retry_after", backoff))
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(backoff):
			if backoff < larkReconnectMax {
				backoff *= 2
				if backoff > larkReconnectMax {
					backoff = larkReconnectMax
				}
			}
		}
	}
}

func handleLarkMessage(ctx context.Context, event *larkim.P2MessageReceiveV1, h MessageHandler, client *lark.Client, logger *zap.Logger) {
	if event == nil || event.Event == nil || event.Event.Message == nil || event.Event.Sender == nil || event.Event.Sender.SenderId == nil {
		return
	}
	msg := event.Event.Message
	msgType := larkcore.StringValue(msg.MessageType)
	if msgType != larkim.MsgTypeText {
		logger.Debug("Feishu traite uniquement les messages texte pour le moment", zap.String("msg_type", msgType))
		return
	}
	var textBody larkTextContent
	if err := json.Unmarshal([]byte(larkcore.StringValue(msg.Content)), &textBody); err != nil {
		logger.Warn("Échec de l'analyse du contenu du message Feishu", zap.Error(err))
		return
	}
	text := strings.TrimSpace(textBody.Text)
	if text == "" {
		return
	}
	userID := ""
	if event.Event.Sender.SenderId.UserId != nil {
		userID = *event.Event.Sender.SenderId.UserId
	}
	messageID := larkcore.StringValue(msg.MessageId)
	reply := h.HandleMessage("lark", userID, text)
	contentBytes, _ := json.Marshal(larkTextContent{Text: reply})
	_, err := client.Im.Message.Reply(ctx, larkim.NewReplyMessageReqBuilder().
		MessageId(messageID).
		Body(larkim.NewReplyMessageReqBodyBuilder().
			MsgType(larkim.MsgTypeText).
			Content(string(contentBytes)).
			Build()).
		Build())
	if err != nil {
		logger.Warn("Échec de la réponse Feishu", zap.String("message_id", messageID), zap.Error(err))
		return
	}
	logger.Debug("Réponse Feishu envoyée", zap.String("message_id", messageID))
}
