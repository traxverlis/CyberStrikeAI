package robot

// MessageHandler interface de traitement des messages pour les connexions longues Feishu/DingTalk (implémenté par handler.RobotHandler)
type MessageHandler interface {
	HandleMessage(platform, userID, text string) string
}
