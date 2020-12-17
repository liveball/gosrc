package provider

import (
	"fmt"
	"github.com/google/wire"
)

var Provider = wire.NewSet(DefaultConnectionOpt, NewBb)

type ConnectionOpt struct{}

type DB struct {
	Name string
}

type User struct{ Age int }

func DefaultConnectionOpt() *ConnectionOpt {
	return &ConnectionOpt{}
}

func NewBb(opt *ConnectionOpt) (*DB, error) {
	return &DB{
		Name: "aaa",
	}, nil
}

func NewUserLoadFunc(db *DB) (func(int) *User, error) {
	fmt.Println(db.Name)

	return func(a int) *User {
		return &User{a}
	}, nil
}
