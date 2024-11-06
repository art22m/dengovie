package store

//
//import (
//	"context"
//	"fmt"
//
//	"github.com/jackc/pgx/v5/pgxpool"
//)
//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "test"
//	password = "test"
//	dbname   = "dengovie"
//)
//
//// create user test with password 'test';
//// grant all privileges on database "dengovie" to test;
//
//func CreateDatabase(ctx context.Context) (*Database, error) {
//	pool, err := pgxpool.New(ctx, generateDsn())
//	if err != nil {
//		return nil, err
//	}
//	return NewDatabase(pool), nil
//}
//
//func generateDsn() string {
//	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
//}
