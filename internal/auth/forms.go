package auth

type LoginForm struct {
	Method   string `form:"login-method" validate:"required"`
	Email    string `form:"email"        validate:"email,omitempty"`
	Username string `form:"username"     validate:"omitempty"`
	Password string `form:"password"     validate:"required"`
}

type RegisterForm struct {
	FirstName       string `form:"first-name"       validate:"required"`
	LastName        string `form:"last-name"        validate:"required"`
	Email           string `form:"email"            validate:"email,required"`
	Username        string `form:"username"         validate:"required"`
	Password        string `form:"password"         validate:"required"`
	ConfirmPassword string `form:"confirm-password" validate:"required,eqfield=Password"`
}
