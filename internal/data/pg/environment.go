package pg

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/one-click-platform/deployer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const envsTableName = "environments"

func NewEnvsQ(db *pgdb.DB) data.EnvsQ {
	return &envsQ{
		db:  db.Clone(),
		sql: sq.Select("environments.*").From(envsTableName),
	}
}

type envsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q envsQ) New() data.EnvsQ {
	return NewEnvsQ(q.db)
}

func (q envsQ) Get() (*data.Env, error) {
	var result data.Env

	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q envsQ) FilterByName(name string) data.EnvsQ {
	stmt := sq.Eq{"environment.name": name}
	q.sql = q.sql.Where(stmt)
	return &q
}

func (q envsQ) FilterByAccountID(accountID int64) data.EnvsQ {
	stmt := sq.Eq{"environment.account_id": accountID}
	q.sql = q.sql.Where(stmt)
	return &q
}

func (q envsQ) Insert(item data.Env) (data.Env, error) {
	clauses := structs.Map(item)
	var result data.Env
	stmt := sq.Insert(envsTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}
