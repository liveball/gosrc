// +build wireinject
package injector

import (
	"github.com/google/wire"
	"gosrc/design_patterns/wire/full/provider"
)

func UserLoader() (func(int) *provider.User, error) {
	panic(wire.Build(provider.Provider, provider.NewUserLoadFunc))
}
