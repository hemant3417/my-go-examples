package dog

import "testing"

func TestYears(t *testing.T) {

	y := Years(7)
	if y != 70 {
		t.Error("Expected 70, got ", y)
	}

}
