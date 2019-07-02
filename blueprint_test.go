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

func findAttribute(attr string, constraints []DataUtility) bool {
	for _, constraint := range constraints {
		if *constraint.ID == attr {
			return true
		}
	}
	return false
}

func verifyTree(t *testing.T, tree TreeStructureType, constraints []DataUtility) {

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

func verifyConstraints(t *testing.T, property AbstractPropertiesMethodType, method DataManagement) {
	verifyTree(t, property.GoalTrees.DataUtility, method.Attributes.DataUtility)
}

func findMethod(methodId string, methods []DataManagement) *DataManagement {
	for _, method := range methods {
		if method.MethodID == methodId {
			return &method
		}
	}
	return nil
}

func checkRules(t *testing.T, methods []DataManagement, abstractProperties []AbstractPropertiesMethodType) {
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

func checkDeployment(t *testing.T, deployment DeploymentInfo) {
	if deployment.Name == "" {
		t.Errorf("Mandatory name not found in appendix")
	}

	if len(deployment.Infrastructures) == 0 {
		t.Errorf("Empty infrastructure in blueprint")
	}

	for _, inf := range deployment.Infrastructures {
		name := inf.Name
		if inf.Type == "" {
			t.Errorf("Can't find type in infrastructure %s", name)
		}

		if len(inf.Nodes) == 0 {
			t.Errorf("Resources empty for infrastructure %s", name)
		}

		for role, nodes := range inf.Nodes {
			if role != "master" && role != "worker" {
				t.Errorf("Invalid role found: %s. Only master and worker are supported", role)
			}
			for _, res := range nodes {
				resName := res.Hostname
				if resName == "" {
					t.Errorf("Can't find host name name for resource in infrastructure %s", name)
				}

				if res.CPU == 0 {
					t.Errorf("Empty CPUs value for resource %s", resName)
				}

				if res.RAM == 0 {
					t.Errorf("Empty RAM for resource %s", resName)
				}

				if res.Role == "" {
					t.Errorf("Empty role for resource %s", resName)
				}

				if res.Cores == 0 {
					t.Errorf("Empty number of cores for resource %s", resName)
				}

				if res.IP == "" {
					t.Errorf("Empty IP for resource %s", resName)
				}

				if res.DriveSize == 0 {
					t.Errorf("Empty boot drive size for resource %s", resName)
				}
			}
		}
		if inf.VDCs != nil && len(inf.VDCs) > 0 {
			for vdcID, vdc := range inf.VDCs {
				if vdc.Ports == nil || len(vdc.Ports) == 0 {
					t.Errorf("Expected ports information in vdc %s of infrastructure %s", vdcID, inf.ID)
				}
			}
		}
	}
}

func checkAppendix(t *testing.T, appendix CookbookAppendix) {
	checkDeployment(t, appendix.Deployment)
}
func TestReader(t *testing.T) {

	blueprint, err := ReadBlueprint("resources/test_blueprint.json")

	if err != nil {
		t.Fatalf("could not read the test blueprint: %+v", err)
	}

	checkRules(t, blueprint.DataManagement, blueprint.AbstractProperties)
	checkAppendix(t, blueprint.CookbookAppendix)
}
