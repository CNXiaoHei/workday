package workday

import "testing"

func TestSpecialDay(t *testing.T) {
	dayMap := GetSpecialDayMap()
	if len(dayMap) == 0 {
		t.Errorf("map is empty")
	}
}
