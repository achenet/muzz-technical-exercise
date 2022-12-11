package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "os"
)

const dbName = "mysql"

var connString string

type Profile struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
}

type PartialProfile struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

type DB struct {
	*sql.DB
}

func Connect() (*DB, error) {
	// dbIP := os.Getenv("DB_IP")
	dbIP := "172.17.0.2"
	connString := fmt.Sprintf("root:toto@tcp(%s:3306)/muzz", dbIP)

	db, err := sql.Open(dbName, connString)
	return &DB{db}, err
}

func (db *DB) CreateProfile(p *Profile) error {
	cmdStr := "INSERT INTO profiles (email, password, name, gender, age)" +
		"VALUES (?, ?, ?, ?, ?)"

	_, err := db.Exec(cmdStr, p.Email, p.Password, p.Name, p.Gender, p.Age)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetProfiles(id string) ([]*PartialProfile, error) {
	out := []*PartialProfile{}
	rows, err := db.Query("SELECT id, name, gender, age FROM profiles WHERE id != ?", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		p := &PartialProfile{}
		err = rows.Scan(&p.ID, &p.Name, &p.Gender, &p.Age)
		if err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, nil
}
