// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/craiggwilson/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFooBar() FooBar {
	foo := provideFoo()
	bar := provideBar()
	fooBar := FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

func injectPartFooBar() FooBar {
	foo := provideFoo()
	fooBar := FooBar{
		Foo: foo,
	}
	return fooBar
}
