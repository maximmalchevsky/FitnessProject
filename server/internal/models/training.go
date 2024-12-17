package models

//const (
//	createTableTrainings = `
//		CREATE TABLE trainings (
//		    id SERIAL PRIMARY KEY,
//		    name VARCHAR NOT NULL,
//		    phone VARCHAR NOT NULL,
//		    confirmation VARCHAR NOT NULL DEFAULT '-'
//		);
//`
//	createTablePasses = `
//		CREATE TABLE passes (
//		    id SERIAL PRIMARY KEY,
//		    name VARCHAR NOT NULL,
//		    phone VARCHAR NOT NULL,
//		    type VARCHAR NOT NULL,
//		    Duration VARCHAR NOT NULL
//		)
//`
//)

type Training struct {
	ID           int    `json:"id" db:"id" example:"137"`
	Name         string `json:"name" db:"name" example:"Анна"`
	Phone        string `json:"phone" db:"phone" example:"+79164043522"`
	Confirmation string `json:"confirmation" db:"confirmation" example:"0"`
}

type CreateTraining struct {
	Name  string `json:"name" db:"name" example:"Анна"`
	Phone string `json:"phone" db:"phone" example:"+79164043522"`
}

type UpdateTraining struct {
	ID    int    `json:"id" db:"id" example:"137"`
	Name  string `json:"name" db:"name" example:"Павел"`
	Phone string `json:"phone" db:"phone" example:"+79164041337"`
}
