package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/zeroidentidad/fiber-hex-api/errors"
	"github.com/zeroidentidad/fiber-hex-api/logger"
)

type StorageDbCuenta struct {
	client *sqlx.DB
}

func (d StorageDbCuenta) Save(c Cuenta) (*Cuenta, *errors.AppError) {
	insertSql := "INSERT INTO cuentas(cliente_id, fecha_apertura, tipo_cuenta, cantidad, estatus) values(?,?,?,?,?)"

	result, err := d.client.Exec(insertSql, c.ClienteID, c.FechaApertura, c.TipoCuenta, c.Cantidad, c.Estatus)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	c.ID = strconv.FormatInt(id, 10)
	return &c, nil
}

/* transaccion = hacer entrada en tabla de transacciones + actualizar saldo en tabla de cuentas */
func (d StorageDbCuenta) SaveTransaccion(t Transaccion) (*Transaccion, *errors.AppError) {
	// iniciar bloque de transacciones a base de datos
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting new transaction for bank account transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	// insertar transacción de cuenta
	result, _ := tx.Exec(`INSERT INTO transacciones (cuenta_id, cantidad, tipo_transaccion, fecha_transaccion) 
	values (?, ?, ?, ?)`, t.CuentaID, t.Cantidad, t.TipoTransaccion, t.FechaTransaccion)

	// actualizar saldo de cuenta
	if t.EsRetiro() {
		_, err = tx.Exec(`UPDATE cuentas SET cantidad = cantidad - ? where cuenta_id = ?`, t.Cantidad, t.CuentaID)
	} else {
		_, err = tx.Exec(`UPDATE cuentas SET cantidad = cantidad + ? where cuenta_id = ?`, t.Cantidad, t.CuentaID)
	}

	// en caso de error Rollback, revertir cambios
	if err != nil {
		_ = tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	// confirmar transacción si todo bien
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	// obtener último ID de transacción de la tabla
	transaccionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	// obtener información de cuenta actualizada de tabla cuentas
	cuenta, appErr := d.FindBy(t.CuentaID)
	if appErr != nil {
		return nil, appErr
	}
	t.ID = strconv.FormatInt(transaccionId, 10)

	// actualizar estructura de transacción con último saldo
	t.Cantidad = cuenta.Cantidad
	return &t, nil
}

func (d StorageDbCuenta) FindBy(cuentaId string) (*Cuenta, *errors.AppError) {
	findByIdSql := "SELECT cuenta_id, cliente_id, fecha_apertura, tipo_cuenta, cantidad from cuentas where cuenta_id = ?"

	var cuenta Cuenta
	err := d.client.Get(&cuenta, findByIdSql, cuentaId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return &cuenta, nil
}

func NewStorageDbCuenta(dbClient *sqlx.DB) StorageDbCuenta {
	return StorageDbCuenta{dbClient}
}
