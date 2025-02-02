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
// https://redhatinsights.github.io/insights-operator-cli/packages/commands/commands_test.html

import (
	"github.com/redhatinsighs/insights-operator-cli/commands"
	"github.com/tisnik/go-capture"
	"regexp"
	"strings"
	"testing"
)

func tryToFindTrigger(t *testing.T, captured string, trigger string) {
	if !strings.Contains(captured, trigger) {
		t.Fatal("Can not find trigger:", trigger)
	}
}

// TestListOfTriggers function checks whether the non-empty list of triggers
// read via REST API is displayed correctly.
func TestListOfTriggers(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMock{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.ListOfTriggers(restAPIMock)
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "List of triggers for all clusters") {
		t.Fatal("Unexpected output:\n", captured)
	}

	// we expect six lines - title, column headers and four triggers
	numlines := strings.Count(captured, "\n")
	if numlines < 6 {
		t.Fatal("Not all triggers are listed in the output:\n", captured)
	}

	expectedTriggers := []string{
		"must-gather",
		"must-gather",
		"must-gather",
		"different-trigger",
	}
	for _, expectedTrigger := range expectedTriggers {
		tryToFindTrigger(t, captured, expectedTrigger)
	}
}

// TestListOfTriggers function checks whether the empty list of triggers read
// via REST API is displayed correctly.
func TestListOfTriggersEmptyList(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMockEmpty{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.ListOfTriggers(restAPIMock)
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "List of triggers for all clusters") {
		t.Fatal("Unexpected output:\n", captured)
	}

	// we expect two lines - title and column headers
	numlines := strings.Count(captured, "\n")
	if numlines > 2 {
		t.Fatal("Unexpected output:\n", captured)
	}
}

// TestListOfTriggersErrorHandling function checks whether error returned by
// REST API is handled correctly.
func TestListOfTriggersErrorHandling(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMockErrors{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.ListOfTriggers(restAPIMock)
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Error reading list of triggers") {
		t.Fatal("Unexpected output:\n", captured)
	}
}

// TestDescribeActivatedTrigger function checks whether it is possible to read
// and displays information about activated trigger.
func TestDescribeActivatedTrigger(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMock{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.DescribeTrigger(restAPIMock, "0")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Trigger info") {
		t.Fatal("Unexpected output:\n", captured)
	}
	if !strings.Contains(captured, "ffffffff-ffff-ffff-ffff-ffffffffffff") {
		t.Fatal("Can not find cluster ID:\n", captured)
	}
	if !strings.Contains(captured, "tester") {
		t.Fatal("Can not find name of user two triggered the trigger:\n", captured)
	}
	match, err := regexp.MatchString(`Active:.*yes`, captured)
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatal("Trigger is not activated as expected:\n", captured)
	}
}

// TestDescribeInactivatedTrigger function checks whether it is possible to
// read and displays information about inactivated trigger.
func TestDescribeInactivatedTrigger(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMock{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.DescribeTrigger(restAPIMock, "1")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Trigger info") {
		t.Fatal("Unexpected output:\n", captured)
	}
	if !strings.Contains(captured, "ffffffff-ffff-ffff-ffff-ffffffffffff") {
		t.Fatal("Can not find cluster ID:\n", captured)
	}
	if !strings.Contains(captured, "tester") {
		t.Fatal("Can not find name of user two triggered the trigger:\n", captured)
	}
	match, err := regexp.MatchString(`Active:.*no`, captured)
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatal("Trigger is not deactivated as expected:\n", captured)
	}
}

// TestDescribeNonMustGatherTrigger function checks whether it is possible to
// read and displays information about other type of trigger.
func TestDescribeNonMustGatherTrigger(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMock{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.DescribeTrigger(restAPIMock, "2")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Trigger info") {
		t.Fatal("Unexpected output:\n", captured)
	}
	if !strings.Contains(captured, "00000000-0000-0000-0000-000000000000") {
		t.Fatal("Can not find cluster ID:\n", captured)
	}
	if !strings.Contains(captured, "tester") {
		t.Fatal("Can not find name of user two triggered the trigger:\n", captured)
	}
	match, err := regexp.MatchString(`Active:.*no`, captured)
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatal("Trigger is not deactivated as expected:\n", captured)
	}
}

// TestDescribeTriggerErrorHandling function checks how REST API-related issues
// are reported and handled.
func TestDescribeTriggerErrorHandling(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMockErrors{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.DescribeTrigger(restAPIMock, "1")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Error reading selected trigger") {
		t.Fatal("Unexpected output:\n", captured)
	}
}

// TestAddTriggerImpl function checks the ability to add a new trigger via REST
// API.
func TestAddTriggerImpl(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMock{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.AddTriggerImpl(restAPIMock, "tester", "cluster", "reason", "link")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Trigger has been created") {
		t.Fatal("Unexpected output:\n", captured)
	}
}

// TestAddTriggerImplError function checks error handling during new trigger
// registration.
func TestAddTriggerImplErrorHandling(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMockErrors{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.AddTriggerImpl(restAPIMock, "tester", "cluster", "reason", "link")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Error communicating with the service") {
		t.Fatal("Unexpected output:\n", captured)
	}
}

// TestDeleteTrigger function checks. the command 'delete trigger'.
func TestDeleteTrigger(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMock{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.DeleteTrigger(restAPIMock, "0")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Trigger 0 has been deleted") {
		t.Fatal("Unexpected output:\n", captured)
	}
}

// TestDeleteTriggerErrorHandling function check error handling for the command
// 'delete trigger'.
func TestDeleteTriggerErrorHandling(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMockErrors{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.DeleteTrigger(restAPIMock, "0")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Error communicating with the service") {
		t.Fatal("Unexpected output:\n", captured)
	}
}

// TestActivateTrigger function checks the command 'activate trigger'.
func TestActivateTrigger(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMock{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.ActivateTrigger(restAPIMock, "0")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Trigger 0 has been activated") {
		t.Fatal("Unexpected output:\n", captured)
	}
}

// TestActivateTriggerErrorHandling function checks the error handling for
// command 'activate trigger'.
func TestActivateTriggerErrorHandling(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMockErrors{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.ActivateTrigger(restAPIMock, "0")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Error communicating with the service") {
		t.Fatal("Unexpected output:\n", captured)
	}
}

// TestDeactivateTrigger function checks the command 'deactivate trigger'.
func TestDeactivateTrigger(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMock{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.DeactivateTrigger(restAPIMock, "0")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Trigger 0 has been deactivated") {
		t.Fatal("Unexpected output:\n", captured)
	}
}

// TestDeactivateTriggerErrorHandling function checks the error handling for
// command 'deactivate trigger'.
func TestDeactivateTriggerErrorHandling(t *testing.T) {
	// turn off any colorization on standard output
	configureColorizer()

	// use mocked REST API instead of the real one
	restAPIMock := RestAPIMockErrors{}

	// use go-capture package to capture all writes to standard output
	captured, err := capture.StandardOutput(func() {
		commands.DeactivateTrigger(restAPIMock, "0")
	})

	// check if capture was done correctly
	checkCapturedOutput(t, captured, err)

	// test the captured output
	if !strings.HasPrefix(captured, "Error communicating with the service") {
		t.Fatal("Unexpected output:\n", captured)
	}
}
