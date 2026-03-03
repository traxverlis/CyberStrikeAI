package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// ExternalMCPClient interface client MCP externe (implémenté par client_sdk.go basé sur le SDK officiel)
type ExternalMCPClient interface {
	Initialize(ctx context.Context) error
	ListTools(ctx context.Context) ([]Tool, error)
	CallTool(ctx context.Context, name string, args map[string]interface{}) (*ToolResult, error)
	Close() error
	IsConnected() bool
	GetStatus() string
}

// MCP Types de messages
const (
	MessageTypeRequest  = "request"
	MessageTypeResponse = "response"
	MessageTypeError    = "error"
	MessageTypeNotify   = "notify"
)

// MCP Version du protocole
const ProtocolVersion = "2024-11-05"

// MessageID représente le champ id de JSON-RPC 2.0, peut être chaîne, nombre ou null
type MessageID struct {
	value interface{}
}

// UnmarshalJSON désérialisation personnalisée, supporte chaîne, nombre et null
func (m *MessageID) UnmarshalJSON(data []byte) error {
	// Tenter d'analyser comme null
	if string(data) == "null" {
		m.value = nil
		return nil
	}

	// Tenter d'analyser comme chaîne
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		m.value = str
		return nil
	}

	// Tenter d'analyser comme nombre
	var num json.Number
	if err := json.Unmarshal(data, &num); err == nil {
		m.value = num
		return nil
	}

	return fmt.Errorf("invalid id type")
}

// MarshalJSON sérialisation personnalisée
func (m MessageID) MarshalJSON() ([]byte, error) {
	if m.value == nil {
		return []byte("null"), nil
	}
	return json.Marshal(m.value)
}

// String retourne la représentation en chaîne
func (m MessageID) String() string {
	if m.value == nil {
		return ""
	}
	return fmt.Sprintf("%v", m.value)
}

// Value retourne la valeur brute
func (m MessageID) Value() interface{} {
	return m.value
}

// Message représente un message MCP (conforme à la spécification JSON-RPC 2.0)
type Message struct {
	ID      MessageID       `json:"id,omitempty"`
	Type    string          `json:"-"` // Usage interne, non sérialisé en JSON
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *Error          `json:"error,omitempty"`
	Version string          `json:"jsonrpc,omitempty"` // Identifiant de version JSON-RPC 2.0
}

// Error représente une erreur MCP
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Tool représente la définition d'un outil MCP
type Tool struct {
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`                // Description détaillée
	ShortDescription string                 `json:"shortDescription,omitempty"` // Description courte (pour liste d'outils, réduit consommation de tokens)
	InputSchema      map[string]interface{} `json:"inputSchema"`
}

// ToolCall représente un appel d'outil
type ToolCall struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

// ToolResult représente le résultat d'exécution d'un outil
type ToolResult struct {
	Content []Content `json:"content"`
	IsError bool      `json:"isError,omitempty"`
}

// Content représente le contenu
type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// InitializeRequest requête d'initialisation
type InitializeRequest struct {
	ProtocolVersion string                 `json:"protocolVersion"`
	Capabilities    map[string]interface{} `json:"capabilities"`
	ClientInfo      ClientInfo             `json:"clientInfo"`
}

// ClientInfo informations client
type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// InitializeResponse réponse d'initialisation
type InitializeResponse struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      ServerInfo         `json:"serverInfo"`
}

// ServerCapabilities capacités du serveur
type ServerCapabilities struct {
	Tools     map[string]interface{} `json:"tools,omitempty"`
	Prompts   map[string]interface{} `json:"prompts,omitempty"`
	Resources map[string]interface{} `json:"resources,omitempty"`
	Sampling  map[string]interface{} `json:"sampling,omitempty"`
}

// ServerInfo informations du serveur
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ListToolsRequest requête de liste d'outils
type ListToolsRequest struct{}

// ListToolsResponse réponse de liste d'outils
type ListToolsResponse struct {
	Tools []Tool `json:"tools"`
}

// ListPromptsResponse réponse de liste de prompts
type ListPromptsResponse struct {
	Prompts []Prompt `json:"prompts"`
}

// ListResourcesResponse réponse de liste de ressources
type ListResourcesResponse struct {
	Resources []Resource `json:"resources"`
}

// CallToolRequest requête d'appel d'outil
type CallToolRequest struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

// CallToolResponse réponse d'appel d'outil
type CallToolResponse struct {
	Content []Content `json:"content"`
	IsError bool      `json:"isError,omitempty"`
}

// ToolExecution enregistrement d'exécution d'outil
type ToolExecution struct {
	ID        string                 `json:"id"`
	ToolName  string                 `json:"toolName"`
	Arguments map[string]interface{} `json:"arguments"`
	Status    string                 `json:"status"` // pending, running, completed, failed
	Result    *ToolResult            `json:"result,omitempty"`
	Error     string                 `json:"error,omitempty"`
	StartTime time.Time              `json:"startTime"`
	EndTime   *time.Time             `json:"endTime,omitempty"`
	Duration  time.Duration          `json:"duration,omitempty"`
}

// ToolStats informations statistiques d'outil
type ToolStats struct {
	ToolName     string     `json:"toolName"`
	TotalCalls   int        `json:"totalCalls"`
	SuccessCalls int        `json:"successCalls"`
	FailedCalls  int        `json:"failedCalls"`
	LastCallTime *time.Time `json:"lastCallTime,omitempty"`
}

// Prompt modèle de prompt
type Prompt struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Arguments   []PromptArgument `json:"arguments,omitempty"`
}

// PromptArgument argument de prompt
type PromptArgument struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

// GetPromptRequest requête d'obtention de prompt
type GetPromptRequest struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments,omitempty"`
}

// GetPromptResponse réponse d'obtention de prompt
type GetPromptResponse struct {
	Messages []PromptMessage `json:"messages"`
}

// PromptMessage message de prompt
type PromptMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Resource ressource
type Resource struct {
	URI         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
}

// ReadResourceRequest requête de lecture de ressource
type ReadResourceRequest struct {
	URI string `json:"uri"`
}

// ReadResourceResponse réponse de lecture de ressource
type ReadResourceResponse struct {
	Contents []ResourceContent `json:"contents"`
}

// ResourceContent contenu de ressource
type ResourceContent struct {
	URI      string `json:"uri"`
	MimeType string `json:"mimeType,omitempty"`
	Text     string `json:"text,omitempty"`
	Blob     string `json:"blob,omitempty"`
}

// SamplingRequest requête d'échantillonnage
type SamplingRequest struct {
	Messages    []SamplingMessage `json:"messages"`
	Model       string            `json:"model,omitempty"`
	MaxTokens   int               `json:"maxTokens,omitempty"`
	Temperature float64           `json:"temperature,omitempty"`
	TopP        float64           `json:"topP,omitempty"`
}

// SamplingMessage message d'échantillonnage
type SamplingMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// SamplingResponse réponse d'échantillonnage
type SamplingResponse struct {
	Content    []SamplingContent `json:"content"`
	Model      string            `json:"model,omitempty"`
	StopReason string            `json:"stopReason,omitempty"`
}

// SamplingContent contenu d'échantillonnage
type SamplingContent struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}
