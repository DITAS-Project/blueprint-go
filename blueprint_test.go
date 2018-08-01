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

func findAttribute(attr string, constraints []ConstraintType) bool {
	for _, constraint := range constraints {
		if *constraint.ID == attr {
			return true
		}
	}
	return false
}

func verifyTree(t *testing.T, tree TreeStructureType, constraints []ConstraintType) {

	if len(tree.Children) == 0 && len(tree.Leaves) == 0 {
		t.Fatalf("Invalid tree without leaves or children")
	}

	if tree.Type == nil {
		t.Fatalf("Invalid tree without type")
	}

	for _, child := range tree.Children {
		verifyTree(t, child, constraints)
	}

	for _, leaf := range tree.Leaves {
		for _, attr := range leaf.Attributes {
			if !findAttribute(attr, constraints) {
				t.Fatalf("Can't find attribute %s in list of constraints", attr)
			}
		}
	}
}

func verifyConstraints(t *testing.T, property AbstractPropertiesMethodType, method DataManagementMethodType) {
	verifyTree(t, property.GoalTrees.DataUtility, method.Attributes.DataUtility)
}

func findMethod(methodId string, methods []DataManagementMethodType) *DataManagementMethodType {
	for _, method := range methods {
		if *method.MethodId == methodId {
			return &method
		}
	}
	return nil
}

func checkRules(t *testing.T, methods []DataManagementMethodType, abstractProperties []AbstractPropertiesMethodType) {
	if len(abstractProperties) == 0 {
		t.Fatalf("List of abstract properties is empty")
	}

	for _, property := range abstractProperties {
		if property.MethodId == nil {
			t.Fatalf("Found abstract property with empty method id")
		}
		method := findMethod(*property.MethodId, methods)

		if method == nil {
			t.Fatalf("Can't find method information of %s", *property.MethodId)
		}

		verifyConstraints(t, property, *method)
	}
}

func checkAppendix(t *testing.T, appendix CookbookAppendix) {
	if appendix.Name == "" {
		t.Errorf("Mandatory name not found in appendix")
	}

	if len(appendix.Infrastructure) == 0 {
		t.Errorf("Empty infrastructure in blueprint")
	}

	for _, inf := range appendix.Infrastructure {
		name := inf.Name
		if inf.Type == "" {
			t.Errorf("Can't find type in infrastructure %s", name)
		}

		if len(inf.Resources) == 0 {
			t.Errorf("Resources empty for infrastructure %s", name)
		}

		for _, res := range inf.Resources {
			resName := res.Name
			if resName == "" {
				t.Errorf("Can't find name for resource in infrastructure %s", name)
			}

			if res.Type == "" {
				t.Errorf("Empty type for resource %s", resName)
			}

			if res.CPUs == "" {
				t.Errorf("Empty CPUs value for resource %s", resName)
			}

			if res.BaseImage == "" {
				t.Errorf("Empty base image for resource %s", resName)
			}

			if res.OS == "" {
				t.Errorf("Empty OS for resource %s", resName)
			}

			if res.RAM == "" {
				t.Errorf("Empty RAM for resource %s", resName)
			}

			if res.Role == "" {
				t.Errorf("Empty role for resource %s", resName)
			}
		}
	}
}
func TestReader(t *testing.T) {

	blueprint, err := ReadBlueprint("resources/concrete_blueprint_doctor.json")

	if err != nil {
		t.Fatalf("could not read the test blueprint: %+v", err)
	}

	checkRules(t, blueprint.DataManagement, blueprint.AbstractProperties)
	checkAppendix(t, blueprint.CookbookAppendix)
}
