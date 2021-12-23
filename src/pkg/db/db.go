package db

import (
	"github.com/kllpw/pgpez/src/pkg/contacts"
	"github.com/kllpw/pgpez/src/pkg/keys"
	_ "github.com/mattn/go-sqlite3"
)

const DRIVER = "sqlite3"

type Database interface {
	InitDatabase() error
}

type KeysDatabase interface {
	InitDatabase() error
	GetAllKeys() ([]*keys.ArmouredKeyPair, error)
	StoreKeyToDb(akp *keys.ArmouredKeyPair) error
	GetKeyById(id string) (*keys.ArmouredKeyPair, error)
	GetKeyByName(name string) (*keys.ArmouredKeyPair, error)
	DeleteKeyById(id string) error
}

type ContactsDatabase interface {
	InitDatabase() error
	GetAllContacts() ([]*contacts.Contact, error)
	StoreContactToDb(contact *contacts.Contact) error
	GetContactById(id string) (*contacts.Contact, error)
	GetContactByName(name string) (*contacts.Contact, error)
	DeleteContactById(id string) error
}


