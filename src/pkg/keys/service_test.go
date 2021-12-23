package keys

import (
	"testing"
)

func TestKeyServiceImpl_DecryptMsg(t *testing.T) {
	test := "test message"
	akp := DefaultKeyService.GenerateKeyPair(test, test)
	encryptedMsg, _ := DefaultKeyService.EncryptMsg(akp.PubKey, test )
	decryptedMsg, err  := DefaultKeyService.DecryptMsg(akp.PrivKey, []byte(test), encryptedMsg)
	if err != nil {
		t.Fail()
	}
	if decryptedMsg != test {
		t.Fail()
	}
}
