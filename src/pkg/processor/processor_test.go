package processor

import (
	"strconv"
	"testing"
)

func TestDecryptMessage(t *testing.T) {
	testStr := "test"
	testMsg := "test - message"
	pro, err := NewProcessor("./test/test.db")
	if err != nil {
		t.Errorf("err %s", err)
	}
	akp, err := pro.GetKeyByName(testStr)
	if err != nil {
		t.Errorf("err %s", err)
	}
	id := strconv.Itoa(akp.ID)
	encMsg, err := pro.EncryptMessage(id, testMsg, false)
	if err != nil {
		t.Errorf("err %s", err)
	}
	decrypted, err := pro.DecryptMessage(id, testStr, encMsg)
	if err != nil {
		t.Errorf("err %s", err)
	}
	if decrypted != testMsg {
		t.Errorf("err %s", "decrypted does not match")
	}
}

func TestDecryptMessageFromBase64(t *testing.T) {
	testStr := "test"
	testMsg := "test - message"
	pro, err := NewProcessor("./test/test.db")
	if err != nil {
		t.Errorf("err %s", err)
	}
	akp, err := pro.GetKeyByName(testStr)
	if err != nil {
		t.Errorf("err %s", err)
	}
	id := strconv.Itoa(akp.ID)
	_, encryptedBase64, err := pro.EncryptMessageToBase64(id, testMsg, false)
	if err != nil {
		t.Errorf("err %s", err)
	}
	decrypted, err := pro.DecryptMessageFromBase64(id, testStr, encryptedBase64)
	if err != nil {
		t.Errorf("err %s", err)
	}
	if decrypted != testMsg {
		t.Errorf("err %s", "decrypted does not match")
	}
}
