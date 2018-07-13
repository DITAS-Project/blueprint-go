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

func checkMethodList(t *testing.T, methodList MethodListType, name string) {

	if len(methodList.Methods) == 0 {
		t.Fatalf("No methods found in %s section", name)
	}

	if *methodList.Methods[0].Name != "patient-details" {
		t.Fatalf("Unexpected name for method in %s. Found %s", name, *methodList.Methods[0].Name)
	}

}

func TestReader(t *testing.T) {
	blueprint := ReadBlueprint("resources/vdc_blueprint_example_1.json")

	checkMethodList(t, blueprint.DataManagement, "Data Management")
	checkMethodList(t, blueprint.AbstractProperties, "Abstract Properties")
}
