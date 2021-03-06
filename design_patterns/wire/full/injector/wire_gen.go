// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"gosrc/design_patterns/wire/full/provider"
)

// Injectors from wire.go:

func UserLoader() (func(int) *provider.User, error) {
	connectionOpt := provider.DefaultConnectionOpt()
	db, err := provider.NewBb(connectionOpt)
	if err != nil {
		return nil, err
	}
	v, err := provider.NewUserLoadFunc(db)
	if err != nil {
		return nil, err
	}
	return v, nil
}
