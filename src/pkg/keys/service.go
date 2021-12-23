package keys

import (
	"encoding/base64"
	"errors"
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
)

type ArmouredKeyPair struct {
	ID int
	Name         string
	PrivKey      string
	PubKey       string
	Base64PubKey string
}

type ArmouredMessage struct {
	Message string
	Base64Message string
}

type KeyService interface {
	GenerateKeyPair(name string, passphrase string) *ArmouredKeyPair
	EncryptMsg(pubKey string, message string) (string, error)
	DecryptMsg(privKey string, passphrase []byte, message string) (string, error)
}

type keyServiceImpl struct {
	keyType    string
	namePrefix string
	keySize    int
}

type KeyType = string
const (
	x25 KeyType = "x25519"
	rsa KeyType= "rsa"

	DefaultSize       = 2048
	DefaultNamePrefix = "pki-"
	DefaultKeyType    = rsa
)

var (
	errInvalidKey = errors.New("invalid keys type")
)

var DefaultKeyService = NewKeyService(DefaultKeyType, DefaultNamePrefix, DefaultSize)

func NewKeyService(keyType KeyType, namePrefix string, keySize int) KeyService {
	return &keyServiceImpl{keyType: keyType, namePrefix: namePrefix, keySize: keySize}
}

func (ks *keyServiceImpl) GenerateKeyPair(name string, passphrase string) *ArmouredKeyPair {
	privkey, pubkey, base64pubkey := ks.generatePKI(name, passphrase)
	akp := &ArmouredKeyPair{Name: name, PrivKey: privkey, PubKey: pubkey, Base64PubKey: base64pubkey}
	return akp
}

func (ks *keyServiceImpl) generatePKI(name string, passphrase string) (privkey string, pubkey string, base64pubkey string) {
	privkey, _ = helper.GenerateKey(name + ks.namePrefix, "email", []byte(passphrase), ks.keyType, ks.keySize)
	key, _ := crypto.NewKeyFromArmored(privkey)
	pubkey, _ = key.GetArmoredPublicKey()
	base64pubkey = base64.StdEncoding.EncodeToString([]byte(pubkey))
	return privkey, pubkey, base64pubkey
}

func (ks *keyServiceImpl) DecryptMsg(privkey string, passphrase []byte, armor string) (string, error) {
	return helper.DecryptMessageArmored(privkey, passphrase, armor)
}

func (ks *keyServiceImpl) EncryptMsg(pubKey string, message string) (string, error) {
	return helper.EncryptMessageArmored(pubKey, message)
}


