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
	"errors"
	"fmt"
	"strings"

	"github.com/craiggwilson/wire"
)

func main() {
	foo, err := injectFoo()
	fmt.Println(foo) // should be zero, the injector should ignore provideFoo's return value.
	if err == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println(strings.Contains(err.Error(), "there is no Foo"))
	}
}

type Foo int

func provideFoo() (Foo, error) {
	return 42, errors.New("there is no Foo")
}

var Set = wire.NewSet(provideFoo)
