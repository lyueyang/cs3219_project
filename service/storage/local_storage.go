package storage

var acctStorage = map[string]Account{}

type Account struct {
	Name        string
	Description string
}

type Storage struct{}

func (s Storage) GetAccounts() []Account {
	var out []Account

	for _, acct := range acctStorage {
		out = append(out, acct)
	}

	return out
}

func (s Storage) StoreAccount(a Account) {
	if _, exists := acctStorage[a.Name]; !exists {
		// only write if the entry doesn't exist so we don't overwrite
		acctStorage[a.Name] = a
	}
}

func (s Storage) DeleteAccount(acctName string) {
	if _, exists := acctStorage[acctName]; exists {
		delete(acctStorage, acctName)
	}
}

func (s Storage) UpdateAccount(a Account) {
	if _, exists := acctStorage[a.Name]; exists {
		acctStorage[a.Name] = a
	}
}
