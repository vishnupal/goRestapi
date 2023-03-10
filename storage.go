package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "goRestapi"
	dbname   = "postgres"
)

// we create interface so any type have this methods that can Satisfied our Storage Interface
// so single interface can handle any database
type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable", host, port, user, password, dbname)
	// connStr := "user=postgres dbname=postgres password=goRestapi sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil

}

func (s *PostgresStore) Init() error {
	return s.CreatedAccountTable()
}

func (s *PostgresStore) CreatedAccountTable() error {
	query := `CREATE TABLE if not exists account (
		id serial PRIMARY KEY,
		first_name VARCHAR ( 50 ),
		last_name VARCHAR ( 50 ),
		number serial,
		balance serial,
		created_at TIMESTAMP
	);`

	_, err := s.db.Exec(query)

	return err
}

// we created methods so it can Satisfied our Storage interface
func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO account (first_name, last_name, number, balance, created_at) VALUES ($1,$2,$3,$4,$5)`

	resp, err := s.db.Query(query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
	)
	if err != nil {
		return err
	}

	fmt.Printf("%+v \n", resp)
	return nil
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query(`select * from account`)
	if err != nil {
		return nil, err
	}
	// defer rows.Close()
	accounts := make([]*Account, 0)
	for rows.Next() {
		account := new(Account)

		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
