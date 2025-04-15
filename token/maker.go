package token

type Maker interface {
	CreateToken(email string, role string) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
