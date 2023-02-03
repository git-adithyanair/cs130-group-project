package token

type Maker interface {
	CreateToken(userID int64, email string) (string, error)
	VerifyToken(token string) (*Payload, error)
}
