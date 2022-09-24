package external

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

type IExternalCreateUser interface {
	// procedure random  call to third party
	CreateUser(ctx context.Context) error
}

type _externalCreateUser struct {
}

func (e *_externalCreateUser) CreateUser(ctx context.Context) error {
	selectedNum, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	switch selectedNum.Int64() % 6 {
	case 0:
	case 3:
	case 4:
	case 5:
		{
			// error
			return fmt.Errorf("send error")
		}

	case 1:
		{
			// success
			return nil
		}

	case 2:
		{
			// timeout and error
			time.Sleep(30 * time.Second)
			return fmt.Errorf("time out and error")
		}
	}
	return fmt.Errorf("alway return error")
}

func NewExternalCreateUser() IExternalCreateUser {
	b := make([]byte, 1024)
	rand.Reader.Read(b)
	return &_externalCreateUser{}
}
