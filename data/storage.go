package data

import (
	"database/sql"
	"my-app-tx/utils/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Inicializa la conexión a PostgreSQL
func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	// Verificar que la conexión es válida
	if err = db.Ping(); err != nil {
		db.Close()
		return err
	}
	return nil
}

// Agrega una transacción
func AddTrx(t models.Transaction, tenant string) (models.Transaction, error) {
	query := `INSERT INTO transaccion (cuenta,cuenta_destino,monto, tipo, fecha, descripcion,estado,empresa) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := db.QueryRow(query, t.Cuenta, t.Monto, t.Tipo, t.Fecha, t.Descripcion, t.Estado, t.Empresa).Scan(&t.ID)
	if err != nil {
		return models.Transaction{}, err
	}
	return t, nil
}

// Obtiene todas las Transacciones
func GetTrx(tenant string) ([]models.Transaction, error) {
	query := `SELECT id, cuenta,cuenta_destino, monto, tipo, fecha, descripcion,estado,empresa 
			  FROM transaccion WHERE empresa = $1 `
	rows, err := db.Query(query, tenant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lstTrx []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.Cuenta, &t.Monto, &t.Tipo, &t.Fecha, &t.Descripcion, &t.Estado, &t.Empresa); err != nil {
			return nil, err
		}
		lstTrx = append(lstTrx, t)
	}
	return lstTrx, nil
}

// Obtiene una transacción por ID
func GetTrxByID(id int8, tenant string) (models.Transaction, bool, error) {
	query := `SELECT id, cuenta, monto, tipo, fecha, descripcion,estado,empresa
	          FROM transaccion WHERE empresa = $1 AND id = $2 `
	var t models.Transaction
	err := db.QueryRow(query, tenant, id).Scan(&t.ID, &t.Cuenta, &t.Monto, &t.Tipo, &t.Fecha, &t.Descripcion, &t.Estado, &t.Empresa)
	if err == sql.ErrNoRows {
		return models.Transaction{}, false, nil
	}
	if err != nil {
		return models.Transaction{}, false, err
	}
	return t, true, nil
}

// Elimina una transacción por ID
func DeleteTrx(id int, tenant string) (bool, error) {
	query := `DELETE FROM transaccion WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}
