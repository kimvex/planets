package utils

import (
	"math"
	"planets/models"
)

//XYPositionPlanet function for get position of planet in the cartesian plane
func XYPositionPlanet(Grades int, Radius int) (x, y float64) {
	Angle := float64(Grades) * math.Pi / 180

	PositionX := float64(Radius) * math.Cos(Angle)
	PositionY := float64(Radius) * math.Sin(Angle)

	return PositionX, PositionY
}

//Slope function of get slope since a planet to other planet
func Slope(From, To models.CoordinatesPlanet) float64 {
	SlopeCalculate := (To.Y - From.Y) / (To.X - From.X)
	ToFloat := float64(10)
	return math.Round(SlopeCalculate*ToFloat) / ToFloat
}

//DistanceBetweenPlanets function for get distance between two planets
func DistanceBetweenPlanets(From, To models.CoordinatesPlanet) float64 {
	RaisedBaseX := math.Pow(To.X-From.X, 2)
	RaisedBaseY := math.Pow(To.Y-From.Y, 2)

	return math.Sqrt(RaisedBaseY + RaisedBaseX)
}

//CenterSunTriangle function to know position of sun into triangle of planets
func CenterSunTriangle(FirsPlanet, SecondPlanet, ThirdPlanet, Sun models.CoordinatesPlanet) bool {
	FirstX, FirstY := SecondPlanet.X-FirsPlanet.X, SecondPlanet.Y-FirsPlanet.Y
	SecondX, SecondY := ThirdPlanet.X-FirsPlanet.X, ThirdPlanet.Y-FirsPlanet.Y

	side1 := (SecondX*(FirsPlanet.Y-Sun.Y) + SecondY*(Sun.X-FirsPlanet.X)) / (FirstX*SecondY - FirstY*SecondX)
	side2 := (Sun.Y - FirsPlanet.Y - side1*FirstY) / SecondY

	return (side1 >= 0.0) && (side2 >= 0.0) && ((side1 + side2) <= 1.0)
}

//ValidateGrade function for validate grades
func ValidateGrade(Grades int) bool {
	if Grades < 0 {
		// fmt.Println(Grades, "Grados dentro", math.Abs(float64(Grades)))
		Grades = int(math.Abs(float64(Grades)))
	}

	switch Grades % 360 {
	case 0:
		return true
	case 90:
		return true
	case 180:
		return true
	case 270:
		return true
	case 360:
		return true
	default:
		return false
	}

}
