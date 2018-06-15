package fakedb

import (
	"database/sql"
	"database/sql/driver"
	"strings"
)

var d *Driver

func init() {
	d = &Driver{}
	sql.Register("fakedb", d)
}

type Driver struct {
}

type Conn struct {
}

type Stmt struct {
	conn  *Conn
	input int
}

type Result struct {
	conn *Conn
}

type Tx struct {
	conn *Conn
}

type DataRows struct {
	conn *Conn
	line int
}

func (d *Driver) Open(string) (driver.Conn, error) {
	return &Conn{}, nil
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return &Stmt{conn: c, input: strings.Count(query, "?")}, nil
}

func (c *Conn) Close() error {
	return nil
}

func (c *Conn) Begin() (driver.Tx, error) {
	return &Tx{conn: c}, nil
}

func (r *Result) LastInsertId() (int64, error) {
	return 0, nil
}

func (r *Result) RowsAffected() (int64, error) {
	return 0, nil
}

func (s *Stmt) Close() error {
	return nil
}

func (s *Stmt) NumInput() int {
	return s.input
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return &Result{conn: s.conn}, nil
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	return &DataRows{}, nil
}

func (r *DataRows) Columns() []string {
	return nil
}

func (r *DataRows) Close() error {
	return nil
}

func (r *DataRows) Next(dest []driver.Value) error {
	return nil
}

func (tx *Tx) Commit() error {
	return nil
}

func (tx *Tx) Rollback() error {
	return nil
}
