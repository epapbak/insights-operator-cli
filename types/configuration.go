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

package types

// Generated documentation is available at:
// https://godoc.org/github.com/RedHatInsights/insights-operator-cli/types
//
// Documentation in literate-programming-style is available at:
// https://redhatinsights.github.io/insights-operator-cli/packages/types/configuration.html

// ConfigurationResponse structure represents response of controller service to
// configuration request.
//     Status: status of response
//     Configuration: JSON string of single configuration
type ConfigurationResponse struct {
	Status        string `json:"status"`
	Configuration string `json:"configuration"`
}
