package domain

type SignupRequest struct {
	Username string `json:"username" binding:"required,min=3,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type CollectKYCRequest struct {
	FullName           string `json:"full_name" binding:"required"`
	PhoneNumber        string `json:"phone_number" binding:"required"`
	NationalID         string `json:"national_id" binding:"required"`
	NationalIDDocument string `json:"national_id_document" binding:"required"`
	Selfie             string `json:"selfie" binding:"required"`
	Location           string `json:"location" binding:"required"`
}
