package db

type DBEntitie interface {
	ToEntitie(rowValues []interface{})
}
