package contacts

import (
	"github.com/kllpw/pgpez/src/pkg/contacts"
)

type PageData struct {
	PageTitle string
	ContactCount  int
	Contacts      []*contacts.Contact
}

func (pd *PageData) GetData() interface{} {
	return pd
}
