package data

type AccountsQ interface {
	New() AccountsQ
	Get() (*Account, error)
	Insert(data Account) (Account, error)
	FilterByEmail(login string) AccountsQ
}

type Account struct {
	ID       int64  `db:"id" structs:"-"`
	Email    string `db:"email" structs:"email"`
	Password []byte `db:"password" structs:"password"`
}
