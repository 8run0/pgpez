package messages

import (
	"github.com/kllpw/pgpez/src/pkg/contacts"
	"github.com/kllpw/pgpez/src/pkg/keys"
)

type EncryptDecrypt string

type PageData struct {
	PageTitle      string
	Keys      	[]*keys.ArmouredKeyPair
	Contacts    []*contacts.Contact
	Message        string
}

func (pd *PageData) GetData() interface{} {
	return pd
}
