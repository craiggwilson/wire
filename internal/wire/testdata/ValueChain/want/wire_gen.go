// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/craiggwilson/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFooBar() FooBar {
	foo := _wireFooValue
	fooBar := provideFooBar(foo)
	return fooBar
}

var (
	_wireFooValue = Foo(41)
)
