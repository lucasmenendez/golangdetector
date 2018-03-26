package golangdetector

import "testing"

func TestDetect(t *testing.T) {
	var input string = "apuesta apuro arado araña exhibit"
	var probs map[string]float64 = Detect(input)
	if len(probs) > 0 {
		t.Errorf("Expected len 0, got %d", len(probs))
	}

	input = "apuesta apuro arado araña exhibit apuesta apuro arado araña exhibit apuesta apuro arado araña exhibit apuesta apuro arado araña exhibit apuesta apuro arado araña exhibit"
	probs = Detect(input)
	if len(probs) == 0 {
		t.Errorf("Expected len > 0, got %d", len(probs))
	}

	var expected map[string]float64 = map[string]float64{"es": 0.8, "en": 0.2}
	for code, prob := range probs {
		if prob != expected[code] {
			t.Errorf("Expected %f, got %f", expected[code], prob)
		}
	}

	input = ""
}
