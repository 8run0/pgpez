package home

import (
	"github.com/kllpw/pgpez/src/pkg/keys"
)

type PageData struct {
	PageTitle string
	Name      string
	Key      *keys.ArmouredKeyPair
}

func (pd *PageData) GetData() interface{} {
	return pd
}

