package paseto

import (
	"github.com/o1egl/paseto"
	"mygrpcp_project/config"
	auth "mygrpcp_project/gRPC/proto"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg *config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(cfg.Paseto.Key),
	}
}

func (maker *PasetoMaker) CreateNewToken(auth *auth.AuthData) (string, error) {
	//randomBytes := make([]byte, 16)
	//rand.Read(randomBytes)

	return maker.Pt.Encrypt(maker.Key, auth, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) error {

	var a *auth.AuthData

	return maker.Pt.Decrypt(token, maker.Key, &a, nil)

}
