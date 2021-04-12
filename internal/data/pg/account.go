package pg

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/one-click-platform/deployer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const accountsTableName = "accounts"

func NewAccountsQ(db *pgdb.DB) data.AccountsQ {
	return &accountsQ{
		db:  db.Clone(),
		sql: sq.Select("accounts.*").From(accountsTableName),
	}
}

type accountsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q accountsQ) New() data.AccountsQ {
	return NewAccountsQ(q.db)
}

func (q accountsQ) Get() (*data.Account, error) {
	var result data.Account

	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q accountsQ) Insert(item data.Account) (data.Account, error) {
	clauses := structs.Map(item)
	var result data.Account
	stmt := sq.Insert(accountsTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q accountsQ) FilterByEmail(email string) data.AccountsQ {
	stmt := sq.Eq{"accounts.email": email}
	q.sql = q.sql.Where(stmt)
	return &q
}
