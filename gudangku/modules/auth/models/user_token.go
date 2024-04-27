package models

type (
	UserToken struct {
		ContextType string `json:"context_type" binding:"required"`
		ContextId   string `json:"context_id" binding:"required"`
		Token       string `json:"token" binding:"required"`
	}
)
