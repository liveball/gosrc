// +build wireinject

package impl

import "github.com/google/wire"

func Init() *App {

	panic(
		wire.Build(
			provideFoo,
			//wire.Struct(new(App), "Foo"),
			DefaultApp,
		),
	)
}
