package keys

import (
	"database/sql"
	"errors"
	dbp "github.com/kllpw/pgpez/src/pkg/db"
	"github.com/kllpw/pgpez/src/pkg/keys"
)

const (
	SQLCreateTable     = "create table if not exists keys(name text, privkey text, pubkey text, base64pubkey text);"
	SQLSelectAllKeys   = "select ROWID, name, privkey, pubkey, base64pubkey from keys order by name;"
	SQLSelectKeyByName = "select ROWID, name, privkey, pubkey, base64pubkey from keys where name like ?;"
	SQLSelectKeyById   = "select ROWID, name, privkey, pubkey, base64pubkey from keys where ROWID = ?;"
	SQLInsertKey       = "insert into keys (name, privkey, pubkey, base64pubkey) VALUES (?, ?, ?, ?);"
	SQLDeleteKey       = "delete from keys where ROWID = ?;"
)

type databaseKeysImpl struct {
	fileLocation string
}

func NewKeysDatabase(fileLocation string) dbp.KeysDatabase {
	dbimpl := &databaseKeysImpl{
		fileLocation: fileLocation,
	}
	return dbimpl
}

func (db *databaseKeysImpl) StoreKeyToDb(akp *keys.ArmouredKeyPair) error {
	k, err := db.GetKeyByName(akp.Name)
	if err != sql.ErrNoRows && err != nil {
		return err
	}
	if k != nil {
		return errors.New("name is unavailable")
	}
	database, err := sql.Open(dbp.DRIVER, db.fileLocation)
	if err != nil {
		return err
	}
	statement, err := database.Prepare(SQLInsertKey)
	if err != nil {
		return err
	}
	_, err = statement.Exec(akp.Name, akp.PrivKey, akp.PubKey, akp.Base64PubKey)
	if err != nil {
		return err
	}
	return err
}
func (db *databaseKeysImpl) InitDatabase() error {
	database, err := sql.Open(dbp.DRIVER, db.fileLocation)
	if err != nil {
		return err
	}
	defer database.Close()
	statement, err := database.Prepare(SQLCreateTable)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}

func (db *databaseKeysImpl) DeleteKeyById(id string) error {
	database, err := sql.Open(dbp.DRIVER, db.fileLocation)
	defer database.Close()
	if err != nil {
		return err
	}
	_, err = database.Exec(SQLDeleteKey, id)
	if err != nil {
		return err
	}
	return nil
}
func (db *databaseKeysImpl) GetKeyById(id string) (*keys.ArmouredKeyPair, error) {
	database, err := sql.Open(dbp.DRIVER, db.fileLocation)
	defer database.Close()
	if err != nil {
		return nil, err
	}
	row := database.QueryRow(SQLSelectKeyById, id)
	kp := keys.ArmouredKeyPair{}
	err = row.Scan(&kp.ID, &kp.Name, &kp.PrivKey, &kp.PubKey, &kp.Base64PubKey)
	if err != nil {
		return nil, err
	}
	return &kp, nil
}

func (db *databaseKeysImpl) GetKeyByName(name string) (*keys.ArmouredKeyPair, error) {
	database, err := sql.Open(dbp.DRIVER, db.fileLocation)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	row := database.QueryRow(SQLSelectKeyByName, name)
	kp := keys.ArmouredKeyPair{}
	err = row.Scan(&kp.ID, &kp.Name, &kp.PrivKey, &kp.PubKey, &kp.Base64PubKey)
	if err != nil {
		return nil, err
	}
	return &kp, nil
}
func (db *databaseKeysImpl) GetAllKeys() ([]*keys.ArmouredKeyPair, error) {
	database, err := sql.Open(dbp.DRIVER, db.fileLocation)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	rows, err := database.Query(SQLSelectAllKeys)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ks []*keys.ArmouredKeyPair
	for rows.Next() {
		akpOut := keys.ArmouredKeyPair{}
		rows.Scan(&akpOut.ID, &akpOut.Name, &akpOut.PrivKey, &akpOut.PubKey, &akpOut.Base64PubKey)
		ks = append(ks, &akpOut)
	}
	return ks, nil
}
