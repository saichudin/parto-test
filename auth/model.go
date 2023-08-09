package auth

type Account struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
