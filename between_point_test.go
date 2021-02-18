package main

import (
	"planets/models"
	"planets/utils"
	"testing"
)

func TestDistanceBetWeenPlanets(t *testing.T) {
	var FirstPlanet models.CoordinatesPlanet
	var SecondPlanet models.CoordinatesPlanet
	FirstPlanet.X = 10
	FirstPlanet.Y = 20
	SecondPlanet.X = 30
	SecondPlanet.Y = 60

	if utils.DistanceBetweenPlanets(FirstPlanet, SecondPlanet) != 44.721359549995796 {
		t.Error("Test fail: Problem with get distance between two points")
	}
}
