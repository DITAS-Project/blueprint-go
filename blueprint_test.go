/*
Copyright 2017 Atos

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
package blueprint

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func checkConstraints(t *testing.T, name string, constraints []ConstraintType) {

}

func checkDataManagementAttributes(t *testing.T, attributes DataManagementAttributesType) {
	checkConstraints(t, "Data Utility", attributes.DataUtility)
	checkConstraints(t, "Security", attributes.Security)
	checkConstraints(t, "Privacy", attributes.Privacy)
}

func checkDataManagement(t *testing.T, methodList []DataManagementMethodType) {

	if len(methodList) == 0 {
		t.Fatalf("No methods found in section Data Management")
	}

	for _, method := range methodList {
		if method.MethodId == nil {
			t.Fatalf("Found method in data management without a method id")
		}
		checkDataManagementAttributes(t, method.Attributes)
	}

}

func TestReader(t *testing.T) {
	blueprint := ReadBlueprint("resources/concrete_blueprint_doctor.json")

	checkDataManagement(t, blueprint.DataManagement)
}
