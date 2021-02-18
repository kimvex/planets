package main

import (
	"planets/utils"
	"testing"
)

func TestPositionXY500(t *testing.T) {
	x, y := utils.XYPositionPlanet(90, 500)

	if float64(x) != 3.061616997868379e-14 && float64(y) != 500 {
		t.Error("Problem with get position of planet")
	}
}

func TestPositionXY1000(t *testing.T) {
	x, y := utils.XYPositionPlanet(90, 1000)

	if float64(x) != 6.123233995736757e-14 && float64(y) != 1000 {
		t.Error("Problem with get position of planet")
	}
}
func TestPositionXY2000(t *testing.T) {
	x, y := utils.XYPositionPlanet(90, 2000)

	if float64(x) != 1.2246467991473515e-13 && float64(y) != 2000 {
		t.Error("Problem with get position of planet")
	}
}
