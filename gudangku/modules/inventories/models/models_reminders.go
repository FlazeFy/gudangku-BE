package models

type (
	GetReminderModel struct {
		ID              string `json:"id"`
		ReminderDesc    string `json:"reminder_desc"`
		ReminderType    string `json:"reminder_type"`
		ReminderContext string `json:"reminder_context"`
		CreatedAt       string `json:"created_at"`
	}
)
