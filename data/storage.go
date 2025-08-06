package data

import (
	"database/sql"
	"my-app-tx/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Inicializa la conexión a PostgreSQL
func InitDB(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return db.Ping()
}

// Agrega una transacción
func AddTransaccion(t models.Transaccion) (models.Transaccion, error) {
	query := `INSERT INTO transacciones (monto, tipo, fecha, descripcion) 
              VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(query, t.Monto, t.Tipo, t.Fecha, t.Descripcion).Scan(&t.ID)
	if err != nil {
		return models.Transaccion{}, err
	}
	return t, nil
}

// Obtiene todas las transacciones
func GetTransacciones() ([]models.Transaccion, error) {
	query := `SELECT id, monto, tipo, fecha, descripcion FROM transacciones`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transacciones []models.Transaccion
	for rows.Next() {
		var t models.Transaccion
		if err := rows.Scan(&t.ID, &t.Monto, &t.Tipo, &t.Fecha, &t.Descripcion); err != nil {
			return nil, err
		}
		transacciones = append(transacciones, t)
	}
	return transacciones, nil
}

// Obtiene una transacción por ID
func GetTransaccionByID(id int) (models.Transaccion, bool, error) {
	query := `SELECT id, monto, tipo, fecha, descripcion FROM transacciones WHERE id = $1`
	var t models.Transaccion
	err := db.QueryRow(query, id).Scan(&t.ID, &t.Monto, &t.Tipo, &t.Fecha, &t.Descripcion)
	if err == sql.ErrNoRows {
		return models.Transaccion{}, false, nil
	}
	if err != nil {
		return models.Transaccion{}, false, err
	}
	return t, true, nil
}

// Elimina una transacción por ID
func DeleteTransaccion(id int) (bool, error) {
	query := `DELETE FROM transacciones WHERE id = $1`
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
