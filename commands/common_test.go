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

package commands_test

// Documentation in literate-programming-style is available at:
// https://redhatinsights.github.io/insights-operator-cli/packages/commands/common_test.html

import (
	"github.com/logrusorgru/aurora"
	"github.com/redhatinsighs/insights-operator-cli/commands"
)

// configureColorizer function configures the Aurora colorizer. For tests (i.e.
// unit tests and functional tests as well) it is preferred to turn-off
// the colorization.
func configureColorizer() {
	colorizer := aurora.NewAurora(false)
	commands.SetColorizer(colorizer)
}
