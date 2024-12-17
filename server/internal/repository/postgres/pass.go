package postgres

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"server/internal/models"
)

func PassGetByID(db *sqlx.DB, id int) (*models.Pass, error) {
	pass := models.Pass{}
	query := `SELECT * FROM passes WHERE id = $1`

	err := db.Get(&pass, query, id)
	if err != nil {
		return nil, err
	}
	return &pass, nil
}

func PassCreate(db *sqlx.DB, pass *models.Pass) (*models.Pass, error) {
	query := `INSERT INTO passes (name, phone, type, duration) VALUES ($1, $2, $3, $4) RETURNING id`

	err := db.QueryRow(query, pass.Name, pass.Phone, pass.Type, pass.Duration).Scan(&pass.ID)
	if err != nil {
		return nil, err
	}
	return pass, nil
}

func PassDelete(db *sqlx.DB, id int) error {
	query := `DELETE FROM passes WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func PassGetAll(db *sqlx.DB) (*[]models.Pass, error) {
	var pass []models.Pass
	query := `SELECT * FROM passes`
	err := db.Select(&pass, query)
	if err != nil {
		return nil, err
	}
	return &pass, nil
}

func PassExistsID(db *sqlx.DB, id int) (bool, error) {
	exists := 0
	query := "SELECT 1 FROM passes WHERE id = $1 LIMIT 1"
	err := db.QueryRow(query, id).Scan(&exists)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}
