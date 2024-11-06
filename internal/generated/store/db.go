package store

//
//import (
//	"context"
//
//	"github.com/georgysavva/scany/v2/pgxscan"
//	"github.com/jackc/pgx/v5"
//	"github.com/jackc/pgx/v5/pgconn"
//	"github.com/jackc/pgx/v5/pgxpool"
//)
//
//type DatabaseOperations interface {
//	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
//	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
//	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
//	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
//	GetPool() *pgxpool.Pool
//}
//
//type Database struct {
//	cluster *pgxpool.Pool
//}
//
//func NewDatabase(cluster *pgxpool.Pool) *Database {
//	return &Database{cluster: cluster}
//}
//
//func (db *Database) GetPool() *pgxpool.Pool {
//	return db.cluster
//}
//
//func (db *Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
//	return pgxscan.Get(ctx, db.cluster, dest, query, args...)
//}
//
//func (db *Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
//	return pgxscan.Select(ctx, db.cluster, dest, query, args...)
//}
//
//func (db *Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
//	return db.cluster.Exec(ctx, query, args...)
//}
//
//func (db *Database) ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
//	return db.cluster.QueryRow(ctx, query, args...)
//}
