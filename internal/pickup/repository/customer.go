package repository

import "github.com/agadilkhan/pickup-point-service/internal/pickup/database/postgres"

type CustomerRepo struct {
	Main    *postgres.Db
	Replica *postgres.Db
}

func NewCustomerRepo(main *postgres.Db, replica *postgres.Db) *CustomerRepo {
	return &CustomerRepo{
		main,
		replica,
	}
}
