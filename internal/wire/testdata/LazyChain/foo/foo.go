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

	"github.com/google/wire"
)

func main() {
	fmt.Println(injectFooBar())
}

type Foo int
type Bar int
type FooBar int

func provideFoo() Foo {
	return 40
}

func provideBar(fooFn func() Foo) Bar {
	return Bar(fooFn()) + Bar(2)
}

func provideFooBar(barFn func() Bar) FooBar {
	return FooBar(barFn())
}

var Set = wire.NewSet(
	wire.Lazy(provideFoo),
	wire.Lazy(provideBar),
	provideFooBar)
