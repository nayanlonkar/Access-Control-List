package repository

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/rest-api/model"
)

type IRepository interface {
	GetByID(context.Context, int64) (interface{}, error)
	Create(context.Context, interface{}) (interface{}, error)
	Update(context.Context, interface{}) (interface{}, error)
	Delete(context.Context, int64) error
	GetAll(context.Context) ([]interface{}, error)
}

type JRepository interface {
	Login(context.Context, int64, string) (interface{}, error)
}

type UGRepository interface {
	GetGroupByID(context.Context, int64) (interface{}, error)
	GetUserByGID(context.Context, int64) (interface{}, error)
}

type FRepository interface {
	GetFilesByPID(context.Context, int64) (interface{}, error)
}

type Repository struct {
}

func (repo *Repository) GetByID(cntx context.Context, id int64) (obj interface{}, err error) {
	return
}

func (repo *Repository) Create(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}

func (repo *Repository) Update(cntx context.Context, obj interface{}) (uobj interface{}, err error) {
	return
}

func (repo *Repository) Delete(cntx context.Context, id int64) (deleted bool, err error) {
	return
}

func (repo *Repository) GetAll(cntx context.Context) (obj []interface{}, err error) {
	return
}

func (repo *Repository) Login(cntx context.Context, id int64, password string) (obj interface{}, err error) {
	return
}

func (repo *Repository) GetGroupByID(conn *sql.DB, object model.IModel, id int64) (obj model.IModel, err error) {
	return
}

func (repo *Repository) GetUserByGID(conn *sql.DB, object model.IModel, id int64) (obj model.IModel, err error) {
	return
}

func (repo *Repository) GetFilesByPID(conn *sql.DB, object model.IModel, id int64) (obj model.IModel, err error) {
	return
}
