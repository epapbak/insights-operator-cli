/*
Copyright © 2019, 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

// Generated documentation is available at:
// https://godoc.org/github.com/RedHatInsights/insights-operator-cli/commands
//
// Documentation in literate-programming-style is available at:
// https://redhatinsights.github.io/insights-operator-cli/packages/commands/authors.html

import (
	"fmt"
)

// PrintAuthors function displays list of authors to standard output. It is
// possible to enable or disable colorized output of header via command line
// options.
func PrintAuthors() {
	fmt.Println(colorizer.Magenta("Authors"))
	fmt.Println(`
Angelina Nikiforova
Bohdan Iakymets
Pavel Tisnovsky`)
}
