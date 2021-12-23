package contacts

import (
	"database/sql"
	"errors"
	"github.com/kllpw/pgpez/src/pkg/contacts"
	dbp "github.com/kllpw/pgpez/src/pkg/db"
)

const (
	SQLCreateTable       = "create table if not exists contacts(name text, pubkey text, base64pubkey text);"
	SQLSelectAllContacts = "select ROWID, name, pubkey, base64pubkey from contacts order by name;"
	SQLSelectKeyByName   = "select ROWID, name, pubkey, base64pubkey from contacts where name like ?;"
	SQLSelectKeyById     = "select ROWID, name, pubkey, base64pubkey from contacts where ROWID = ?;"
	SQLInsertKey         = "insert into contacts (name, pubkey, base64pubkey) VALUES (?, ?, ?);"
	SQLDeleteKey         = "delete from contacts where ROWID = ?;"
)

type databasecontactsImpl struct {
	fileLocation string
}

func (db *databasecontactsImpl) InitDatabase() error {
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

func (db *databasecontactsImpl) GetAllContacts() ([]*contacts.Contact, error) {
	database, err := sql.Open(dbp.DRIVER, db.fileLocation)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	rows, err := database.Query(SQLSelectAllContacts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ks []*contacts.Contact
	for rows.Next() {
		contactOut := contacts.Contact{}
		rows.Scan(&contactOut.ID, &contactOut.Name, &contactOut.PubKey, &contactOut.Base64PubKey)
		ks = append(ks, &contactOut)
	}
	return ks, nil
}

func (db *databasecontactsImpl) StoreContactToDb(contact *contacts.Contact) error {
	k, err := db.GetContactByName(contact.Name)
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
	_, err = statement.Exec(contact.Name, contact.PubKey, contact.Base64PubKey)
	if err != nil {
		return err
	}
	return err
}

func (db *databasecontactsImpl) GetContactById(id string) (*contacts.Contact, error) {
	database, err := sql.Open(dbp.DRIVER, db.fileLocation)
	defer database.Close()
	if err != nil {
		return nil, err
	}
	row := database.QueryRow(SQLSelectKeyById, id)
	contact := contacts.Contact{}
	err = row.Scan(&contact.ID, &contact.Name, &contact.PubKey, &contact.Base64PubKey)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (db *databasecontactsImpl) GetContactByName(name string) (*contacts.Contact, error) {
	database, err := sql.Open(dbp.DRIVER, db.fileLocation)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	row := database.QueryRow(SQLSelectKeyByName, name)
	contact := contacts.Contact{}
	err = row.Scan(&contact.ID, &contact.Name, &contact.PubKey, &contact.Base64PubKey)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (db *databasecontactsImpl) DeleteContactById(id string) error {
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

func NewcontactsDatabase(fileLocation string) dbp.ContactsDatabase {
	dbimpl := &databasecontactsImpl{
		fileLocation: fileLocation,
	}
	return dbimpl
}
