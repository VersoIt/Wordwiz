package user_repo

import (
	trmngr "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	ctxGetter *trmngr.CtxGetter
	db        *sqlx.DB
}

func New(db *sqlx.DB, ctxGetter *trmngr.CtxGetter) *Repository {
	return &Repository{
		ctxGetter: ctxGetter,
		db:        db,
	}
}
