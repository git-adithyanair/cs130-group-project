package token

type Maker interface {
	CreateToken(userID string, email string) (string, error)
	VerifyToken(token string) (*Payload, error)
}
