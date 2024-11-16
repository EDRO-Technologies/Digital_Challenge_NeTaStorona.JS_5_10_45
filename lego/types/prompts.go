package types

type UserRequest struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
	Context string `json:"context"`
}

type ModelRequest struct {
	Model       string  `json:"model"`
	Temperature float64 `json:"temperature"`
	TopP        int     `json:"top_p"`
	MaxTokens   int     `json:"max_tokens"`
	Stream      bool    `json:"stream"`
	Stop        string  `json:"stop"`
	RandomSeed  int     `json:"random_seed"`
	Messages    []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	ResponseFormat struct {
		Type string `json:"type"`
	} `json:"response_format"`
	Tools []struct {
		Type     string `json:"type"`
		Function struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Parameters  struct {
			} `json:"parameters"`
		} `json:"function"`
	} `json:"tools"`
	ToolChoice       string `json:"tool_choice"`
	PresencePenalty  int    `json:"presence_penalty"`
	FrequencyPenalty int    `json:"frequency_penalty"`
	N                int    `json:"n"`
	SafePrompt       bool   `json:"safe_prompt"`
}

type ModelResponse struct {
	Id     string `json:"id"`
	Object string `json:"object"`
	Model  string `json:"model"`
	Usage  struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Created int `json:"created"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role      string `json:"role"`
			Content   string `json:"content"`
			ToolCalls struct {
			} `json:"tool_calls"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}
