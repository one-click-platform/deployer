package data

type EnvsQ interface {
	New() EnvsQ
	Get() (*Env, error)
	FilterByAccountID(accountID int64) EnvsQ
	FilterByName(name string) EnvsQ
	Insert(data Env) (Env, error)
}

type Env struct {
	ID        int64  `db:"id" structs:"-"`
	Name      string `db:"name" structs:"name"`
	AccountID int64  `db:"account_id" structs:"account_id"`
}
