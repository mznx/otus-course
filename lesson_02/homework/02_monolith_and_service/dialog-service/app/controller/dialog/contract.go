package dialog

type SendMessageRequest struct {
	Text string `json:"text"`
}

type GetListMessagesResponse struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}
