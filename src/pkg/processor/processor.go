package processor

import (
	"encoding/base64"
	"github.com/kllpw/pgpez/src/pkg/contacts"
	"github.com/kllpw/pgpez/src/pkg/db"
	dbContacts "github.com/kllpw/pgpez/src/pkg/db/contacts"
	dbKeys "github.com/kllpw/pgpez/src/pkg/db/keys"
	"github.com/kllpw/pgpez/src/pkg/keys"
)

type Processor struct {
	keyService keys.KeyService
	keysDb     db.KeysDatabase
	contactsDb db.ContactsDatabase
}

type ProcReqs interface {
	GetAllKeys() ([]*keys.ArmouredKeyPair, error)
	GenerateAndStoreNewKeyPair(name string, passphrase string) (*keys.ArmouredKeyPair, error)
	GetKeyById(id string) (*keys.ArmouredKeyPair, error)
	GetKeyByName(name string) (*keys.ArmouredKeyPair, error)
	EncryptMessage(id string, message string, useContacts bool) (string, error)
	DecryptMessage(id string, passphrase string, message string) (string, error)
	EncryptMessageToBase64(id string, message string, useContacts bool) (string, string, error)
	DecryptMessageFromBase64(id string, passphrase string, message string) (string, error)
	DeleteKey(id string, passphrase string) error

	GetAllContacts() ([]*contacts.Contact, error)
	GetContactByName(name string) (*contacts.Contact, error)
	GetContactById(id string) (*contacts.Contact, error)
	AddContact(name string, base64pubkey string) (*contacts.Contact, error)
	DeleteContact(id string) error
	AuthKey(id string, passphrase string) error
}

func NewProcessor(path string) (ProcReqs, error) {
	p := &Processor{
		keysDb:     dbKeys.NewKeysDatabase(path),
		contactsDb: dbContacts.NewcontactsDatabase(path),
		keyService: keys.DefaultKeyService,
	}
	err := p.keysDb.InitDatabase()
	if err != nil {
		return nil, err
	}
	err = p.contactsDb.InitDatabase()
	return p, nil
}

func (p *Processor) GetAllKeys() ([]*keys.ArmouredKeyPair, error) {
	ks, err := p.keysDb.GetAllKeys()
	if err != nil {
		return nil, err
	}
	return ks, nil
}

func (p *Processor) DeleteKey(id string, passphrase string) error {
	key, err := p.keysDb.GetKeyById(id)
	if err != nil {
		return err
	}
	varEnc, _ := p.keyService.EncryptMsg(key.PubKey, "verification message")
	_, err = p.keyService.DecryptMsg(key.PrivKey, []byte(passphrase), varEnc)
	if err != nil {
		return err
	}
	err = p.keysDb.DeleteKeyById(id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Processor) GenerateAndStoreNewKeyPair(name string, passphrase string) (*keys.ArmouredKeyPair, error) {
	akp := p.keyService.GenerateKeyPair(name, passphrase)
	err := p.keysDb.StoreKeyToDb(akp)
	if err != nil {
		return nil, err
	}
	return akp, nil
}
func (p *Processor) GetKeyById(id string) (*keys.ArmouredKeyPair, error) {
	k, err := p.keysDb.GetKeyById(id)
	if err != nil {
		return nil, err
	}
	return k, err
}
func (p *Processor) GetKeyByName(name string) (*keys.ArmouredKeyPair, error) {
	k, err := p.keysDb.GetKeyByName(name)
	if err != nil {
		return nil, err
	}
	return k, err
}
func (p *Processor) EncryptMessage(id string, message string, useContacts bool) (string, error) {
	if useContacts {
		k, err := p.contactsDb.GetContactById(id)
		if err != nil {
			return "", err
		}
		return p.keyService.EncryptMsg(k.PubKey, message)
	}
	k, err := p.keysDb.GetKeyById(id)
	if err != nil {
		return "", err
	}
	return p.keyService.EncryptMsg(k.PubKey, message)
}

func (p *Processor) EncryptMessageToBase64(id string, message string, useContacts bool) (string, string, error) {
	if useContacts {
		k, err := p.contactsDb.GetContactById(id)
		if err != nil {
			return "", "", err
		}
		msg, err := p.keyService.EncryptMsg(k.PubKey, message)
		enc := base64.URLEncoding.EncodeToString([]byte(msg))
		return msg, enc, nil
	}
	k, err := p.keysDb.GetKeyById(id)
	msg, err := p.keyService.EncryptMsg(k.PubKey, message)
	if err != nil {
		return "", "", err
	}
	enc := base64.URLEncoding.EncodeToString([]byte(msg))
	return msg, enc, nil
}

func (p *Processor) DecryptMessage(id string, passphrase string, message string) (string, error) {
	k, err := p.keysDb.GetKeyById(id)
	if err != nil {
		return "", err
	}
	de, err := p.keyService.DecryptMsg(k.PrivKey, []byte(passphrase), message)
	if err != nil {
		return "", err
	}
	return de, nil
}

func (p *Processor) DecryptMessageFromBase64(id string, passphrase string, message string) (string, error) {
	decoded, err := base64.URLEncoding.DecodeString(message)
	if err != nil {
		return "", err
	}
	k, err := p.keysDb.GetKeyById(id)
	if err != nil {
		return "", err
	}
	de, err := p.keyService.DecryptMsg(k.PrivKey, []byte(passphrase), string(decoded))
	if err != nil {
		return "", err
	}
	return de, nil
}

func (p *Processor) AddContact(name string, pubkey string) (*contacts.Contact, error) {
	c := &contacts.Contact{
		Name:         name,
		PubKey:       pubkey,
		Base64PubKey: pubkey,
	}
	p.contactsDb.StoreContactToDb(c)
	return c, nil
}

func (p *Processor) GetAllContacts() ([]*contacts.Contact, error) {
	c, err := p.contactsDb.GetAllContacts()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (p *Processor)AuthKey(id string, passphrase string) error  {
	key, err := p.keysDb.GetKeyById(id)
	if err != nil {
		return err
	}
	varEnc, _ := p.keyService.EncryptMsg(key.PubKey, "verification message")
	_, err = p.keyService.DecryptMsg(key.PrivKey, []byte(passphrase), varEnc)
	if err != nil {
		return err
	}
	return nil
}

func (p *Processor) GetContactByName(name string) (*contacts.Contact, error) {
	c, err := p.contactsDb.GetContactByName(name)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (p *Processor) GetContactById(id string) (*contacts.Contact, error) {
	c, err := p.contactsDb.GetContactByName(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (p *Processor) DeleteContact(id string) error {
	err := p.contactsDb.DeleteContactById(id)
	if err != nil {
		return err
	}
	return nil
}

