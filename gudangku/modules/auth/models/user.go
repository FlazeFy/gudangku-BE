package models

type (
	UserLogin struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	UserRegister struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Timezone string `json:"timezone" binding:"required"`
	}
)
