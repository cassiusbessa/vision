package ports

type TokenService interface {
	GetPayload(tokenString string) (string, error)
}
