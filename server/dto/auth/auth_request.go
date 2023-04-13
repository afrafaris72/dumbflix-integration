package authdto

type AuthRequest struct {
	Fullname  string `json:"fullname" form:"fullname" validate:"required"`
	Email     string `json:"email" binding:"required, email" gorm:"unique, not null" `
	Password  string `json:"password" form:"password" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	Phone     string `json:"phone" form:"phone" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Subscribe bool   `json:"subscribe" form:"subscribe" `
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
