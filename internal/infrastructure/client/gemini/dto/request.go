package dto

type Request struct {
	Contents []RequestContent `json:"contents"`
}

type RequestContent struct {
	Parts []RequestPart `json:"parts"`
}

type RequestPart struct {
	Text string `json:"text"`
}
