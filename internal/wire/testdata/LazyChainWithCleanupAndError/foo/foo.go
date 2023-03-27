// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/craiggwilson/wire"
)

func main() {
	s, cleanup, _ := injectFooBar()
	defer cleanup()
	fmt.Println(s)
}

type Foo int
type Bar int
type FooBar int

func provideFoo() (Foo, func(), error) {
	return 40, func() { fmt.Println("foo cleanup") }, nil
}

func provideBar(fooFn func() (Foo, error)) (Bar, func(), error) {
	foo, err := fooFn()
	if err != nil {
		return Bar(0), nil, err
	}

	return Bar(foo) + Bar(2), func() { fmt.Println("bar cleanup") }, nil
}

func provideFooBar(barFn func() (Bar, error)) (FooBar, error) {
	bar, err := barFn()
	if err != nil {
		return FooBar(0), err
	}
	return FooBar(bar), nil
}

var Set = wire.NewSet(
	wire.Lazy(provideFoo),
	wire.Lazy(provideBar),
	provideFooBar)
