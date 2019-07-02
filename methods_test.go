package blueprint

import (
	"testing"
)

func TestGetMethodMap(t *testing.T) {
	blueprint, err := ReadBlueprint("resources/test_blueprint.json")

	if err != nil {
		t.Fatalf("could not read the test blueprint: %+v", err)
	}

	data := blueprint.GetMethodMap()

	if val, ok := data["getAllValuesForBloodTestComponent"]; !ok {
		t.Fatalf("exptexted getAllValuesForBloodTestComponent in %v", data)
	} else {
		if val.HTTPMethod != "GET" {
			t.Fatalf("Wrong HTTPMethod for getAllValuesForBloodTestComponent")
		}
		if val.Path != "/patient/{SSN}/blood-test/component/{component}" {
			t.Fatalf("Wrong Path for getAllValuesForBloodTestComponent")
		}
	}

}
