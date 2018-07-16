package blueprint

import (
	"testing"
)

func TestGetMethodMap(t *testing.T) {
	blueprint := ReadBlueprint("resources/concrete_blueprint_doctor.json")
	data := blueprint.GetMethodMap()

	if val, ok := data["getAllValuesForBloodTestComponent"]; !ok {
		t.Fatalf("exptexted getAllValuesForBloodTestComponent in %v", data)
	} else {
		if val.Method != "GET" {
			t.Fatalf("Wront Method for getAllValuesForBloodTestComponent")
		}
		if val.Path != "/patient/{SSN}/blood-test/component/{component}" {
			t.Fatalf("Wront Path for getAllValuesForBloodTestComponent")
		}
	}
}
