package handlers

type OkPayload struct {
	Health int    `json:"health,omitempty"`
	Test   string `json:"test,omitempty"`
}
