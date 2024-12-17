package models

//CREATE TABLE IF NOT EXISTS passes (
//		    id SERIAL PRIMARY KEY,
//		    name VARCHAR NOT NULL,
//		    phone VARCHAR NOT NULL,
//		    type VARCHAR NOT NULL,
//		    Duration VARCHAR NOT NULL
//		)

type Pass struct {
	ID       int    `json:"id" db:"id" example:"1"`
	Name     string `json:"name" db:"name" example:"Анна"`
	Phone    string `json:"phone" db:"phone" example:"+79164043522"`
	Type     string `json:"type" db:"type" example:"Персональный"`
	Duration int    `json:"duration" db:"duration" example:"6"`
}

type CreatePass struct {
	Name     string `json:"name" db:"name" example:"Анна"`
	Phone    string `json:"phone" db:"phone" example:"+79164043522"`
	Type     string `json:"type" db:"type" example:"Персональный"`
	Duration int    `json:"duration" db:"duration" example:"6"`
}
