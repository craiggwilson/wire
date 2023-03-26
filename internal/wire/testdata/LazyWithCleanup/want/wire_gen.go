// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFooBar() (FooBar, func()) {
	foo := provideFoo()
	barFunc, cleanup := func() (func() Bar, func()) {
		var instance Bar
		var wasSet bool
		var cleanup func()
		return func() Bar {
				if !wasSet {
					wasSet = true
					instance, cleanup = provideBar()
				}
				return instance
			}, func() {
				if wasSet {
					cleanup()
				}
			}
	}()
	fooBar := provideFooBar(foo, barFunc)
	return fooBar, func() {
		cleanup()
	}
}
