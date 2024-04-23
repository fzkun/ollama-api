package types

import "time"

type OllamaGenerateResp struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Response  string    `json:"response"`
	Done      bool      `json:"done"`

	Metrics
}
type Message struct {
	Role    string      `json:"role"` // one of ["system", "user", "assistant"]
	Content string      `json:"content"`
	Images  []ImageData `json:"images,omitempty"`
}

type ImageData []byte

type Metrics struct {
	TotalDuration      time.Duration `json:"total_duration,omitempty"`
	LoadDuration       time.Duration `json:"load_duration,omitempty"`
	PromptEvalCount    int           `json:"prompt_eval_count,omitempty"`
	PromptEvalDuration time.Duration `json:"prompt_eval_duration,omitempty"`
	EvalCount          int           `json:"eval_count,omitempty"`
	EvalDuration       time.Duration `json:"eval_duration,omitempty"`
}
