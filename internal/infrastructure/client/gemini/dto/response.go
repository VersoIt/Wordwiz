package dto

type Response struct {
	Candidates    []Candidate `json:"candidates"`
	UsageMetadata Metadata    `json:"usageMetadata"`
	ModelVersion  string      `json:"modelVersion"`
	ResponseID    string      `json:"responseId"`
}

type Candidate struct {
	Content      ResponseContent `json:"content"`
	FinishReason string          `json:"finishReason"`
	AvgLogprobs  float64         `json:"avgLogprobs"`
}

type ResponseContent struct {
	Parts []ResponsePart `json:"parts"`
	Role  string         `json:"role"`
}

type ResponsePart struct {
	Text string `json:"text"`
}

type PromptTokensDetails struct {
	Modality   string `json:"modality"`
	TokenCount int    `json:"tokenCount"`
}

type CandidatesTokensDetails struct {
	Modality   string `json:"modality"`
	TokenCount int    `json:"tokenCount"`
}

type Metadata struct {
	PromptTokenCount        int                       `json:"promptTokenCount"`
	CandidatesTokenCount    int                       `json:"candidatesTokenCount"`
	TotalTokenCount         int                       `json:"totalTokenCount"`
	PromptTokensDetails     []PromptTokensDetails     `json:"promptTokensDetails"`
	CandidatesTokensDetails []CandidatesTokensDetails `json:"candidatesTokensDetails"`
}
