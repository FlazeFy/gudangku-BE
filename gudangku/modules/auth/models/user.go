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
	UserProfileModel struct {
		Username         string `json:"username"`
		Email            string `json:"email"`
		TelegramUserId   string `json:"telegram_user_id"`
		TelegramIsValid  int    `json:"telegram_is_valid"`
		FirebaseFCMToken string `json:"firebase_fcm_token"`
		LineUserId       string `json:"line_user_id"`
		Phone            string `json:"phone"`
		CreatedAt        string `json:"created_at"`
		UpdatedAt        string `json:"updated_at"`
		Timezone         string `json:"timezone"`
	}
)
