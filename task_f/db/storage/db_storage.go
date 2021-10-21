package storage

import (
	"database/sql"
	"fmt"
)

type Account struct {
	Name        string
	Description string
}

type Storage struct {
	DBConnect *sql.DB
}

func NewDBStorage(dbConnection *sql.DB) Storage {
	out := Storage{DBConnect: dbConnection}

	return out
}
func (s Storage) GetAccounts() []Account {
	rows, err := s.DBConnect.Query(`SELECT "Name", "Description" FROM "Accounts"`)
	CheckError(err)

	defer rows.Close()
	var out []Account
	for rows.Next() {
		var name string
		var desc string

		err = rows.Scan(&name, &desc)
		CheckError(err)

		out = append(out, Account{
			Name:        name,
			Description: desc,
		})
	}
	return out
}

func (s Storage) StoreAccount(account *Account) {
	insertStmt := fmt.Sprintf(`insert into "Accounts"("Name", "Description") values('%s', '%s')`,
		account.Name, account.Description)
	_, e := s.DBConnect.Exec(insertStmt)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
