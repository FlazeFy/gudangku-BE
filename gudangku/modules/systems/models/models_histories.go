package models

type (
	GetHistory struct {
		ID             string `json:"id"`
		HistoryType    string `json:"history_type"`
		HistoryContext string `json:"history_context"`

		// Props
		CreatedAt string `json:"created_at"`
	}
)
