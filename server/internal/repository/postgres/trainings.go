package postgres

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"server/internal/models"
)

//type Training struct {
//	ID           int    `json:"id" db:"id" example:"137"`
//	Name         string `json:"name" db:"name" example:"Анна"`
//	Phone        string `json:"phone" db:"phone" example:"+79164043522"`
//	Confirmation string `json:"confirmation" db:"confirmation" example:"0"`
//}

func TrainingExists(db *sqlx.DB, name, phone string) (bool, error) {
	exists := 0
	query := "SELECT 1 FROM trainings WHERE name = $1 AND phone = $2 LIMIT 1"
	err := db.QueryRow(query, name, phone).Scan(&exists)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

func TrainingExistsID(db *sqlx.DB, id int) (bool, error) {
	exists := 0
	query := "SELECT 1 FROM trainings WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

func TrainingGetByID(db *sqlx.DB, id int) (*models.Training, error) {
	training := models.Training{}
	query := "SELECT * FROM trainings WHERE id = $1"
	err := db.Get(&training, query, id)
	if err != nil {
		return nil, err
	}
	return &training, nil
}

func TrainingCreate(db *sqlx.DB, training *models.Training) (*models.Training, error) {
	query := "INSERT INTO trainings (name, phone) VALUES (:name, :phone) RETURNING id"
	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	err = stmt.Get(&training.ID, *training)
	if err != nil {
		return nil, err
	}
	return training, nil
}

func TrainingUpdate(db *sqlx.DB, training *models.Training) (*models.Training, error) {
	query := `UPDATE trainings SET name = $1, phone = $2 WHERE id = $3`
	_, err := db.Exec(query, training.Name, training.Phone, training.ID)
	if err != nil {
		return nil, err
	}
	return training, nil
}

func TrainingDelete(db *sqlx.DB, id int) error {
	query := "DELETE FROM trainings WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func ConfirmationUpdate(db *sqlx.DB, status string, id int) error {
	query := `UPDATE trainings SET confirmation = $1 WHERE id = $2`
	_, err := db.Exec(query, status, id)
	if err != nil {
		return err
	}
	return nil
}

func TrainingGetAll(db *sqlx.DB) (*[]models.Training, error) {
	trainings := []models.Training{}
	query := "SELECT * FROM trainings"
	err := db.Select(&trainings, query)
	if err != nil {
		return nil, err
	}
	return &trainings, nil
}
