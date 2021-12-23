package keys

import (
	"github.com/kllpw/pgpez/src/pkg/keys"
)

type PageData struct {
	PageTitle string
	Locked bool
	KeyCount  int
	Keys      []*keys.ArmouredKeyPair
	Key      *keys.ArmouredKeyPair
}

func (pd *PageData) GetData() interface{} {
	return pd
}