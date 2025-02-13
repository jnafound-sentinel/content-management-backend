package schemas

type UserUpdate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"user_role" binding:"required,oneof=student parent teacher admin"`
}

type CreateAccountDetails struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
	UserRole string `json:"user_role" binding:"required,oneof=student parent teacher admin"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type VerificationDetails struct {
	Email string `json:"email" binding:"required,email"`
}
