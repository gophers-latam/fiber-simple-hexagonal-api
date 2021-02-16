package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/zeroidentidad/fiber-hex-api/errors"
	"github.com/zeroidentidad/fiber-hex-api/logger"
)

type StorageDbCliente struct {
	client *sqlx.DB
}

func (d StorageDbCliente) FindAll(estatus string) ([]Cliente, *errors.AppError) {
	var findAllSql string
	clientes := make([]Cliente, 0)
	if estatus == "" {
		findAllSql = "SELECT cliente_id, nombre, ciudad, codigo_postal, fecha_nacimiento, estatus FROM clientes"
	} else {
		findAllSql = "SELECT cliente_id, nombre, ciudad, codigo_postal, fecha_nacimiento, estatus FROM clientes WHERE estatus = " + estatus
	}

	err := d.client.Select(&clientes, findAllSql)
	if err != nil {
		logger.Error("Error while querying customers table: " + err.Error())
		return nil, errors.NewNotFoundError("Unexpected database error")
	}

	return clientes, nil
}

func (d StorageDbCliente) ById(id string) (*Cliente, *errors.AppError) {
	findByIdSql := "SELECT cliente_id, nombre, ciudad, codigo_postal, fecha_nacimiento, estatus FROM clientes WHERE cliente_id = ?"

	var c Cliente
	err := d.client.Get(&c, findByIdSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while querying customers table " + err.Error())
			return nil, errors.NewNotFoundError("Customer not found")
		} else {
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewStorageDbCliente(dbClient *sqlx.DB) StorageDbCliente {
	return StorageDbCliente{dbClient}
}
