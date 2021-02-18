package controllers

import (
	"planets/db"
	"planets/models"
	"planets/utils"
)

//Variables set structs
var (
	FGrade models.PlanetGrades
	BGrade models.PlanetGrades
	VGrade models.PlanetGrades
	DB     *db.DB
)

//PlanetsStructure struct local fot generate model
type PlanetsStructure struct {
	models.PlanetsStructure
}

//GeneratePlanets Functions for generate all planets
func GeneratePlanets() *PlanetsStructure {
	Planets := new(PlanetsStructure)
	Planets.Ferengi = models.PlanetModel{
		PlanetName:      "Ferengi",
		Distance:        500,
		Grades:          0,
		AngularVelocity: 1,
		Coordinates: models.CoordinatesPlanet{
			X: 0,
			Y: 0,
		},
	}

	Planets.Betasoide = models.PlanetModel{
		PlanetName:      "Betasoide",
		Distance:        2000,
		Grades:          0,
		AngularVelocity: 3,
		Coordinates: models.CoordinatesPlanet{
			X: 0,
			Y: 0,
		},
	}

	Planets.Vulcano = models.PlanetModel{
		PlanetName:      "Vulcano",
		Distance:        1000,
		Grades:          0,
		AngularVelocity: 5,
		Coordinates: models.CoordinatesPlanet{
			X: 0,
			Y: 0,
		},
	}

	return Planets
}

//RunDays funtion for run days on orbit planet
func (P *PlanetsStructure) RunDays(years int, realLife bool) {
	days := 0

	if realLife == true {
		days = utils.YearsSinceToday(years)
	} else {
		days = utils.YearBaseDayOfPlanets(years)
	}

	DB = db.NewDatabase()
	P.Token = DB.CreateSolarSystem()
	P.Days = days

	for i := 0; i < days; i++ {
		P.CurrentDay = i
		P.GenerateGrades(i)
	}

	P.IntensityPeak()
}

//GenerateGrades function for generate Grades to planets for day
func (P *PlanetsStructure) GenerateGrades(day int) {
	P.Ferengi.Grades = (-P.Ferengi.AngularVelocity * day)
	P.Betasoide.Grades = (-P.Betasoide.AngularVelocity * day)
	P.Vulcano.Grades = (P.Vulcano.AngularVelocity * day)

	P.SetCondition()
}

//SetCondition function for set weather condition for day and position of planets
func (P *PlanetsStructure) SetCondition() {
	var ConditionDay models.ConditionForDay
	ConditionDay.Day = P.CurrentDay

	Sun := models.CoordinatesPlanet{X: 0, Y: 0}
	FX, FY := utils.XYPositionPlanet(P.Ferengi.Grades, P.Ferengi.Distance)
	BX, BY := utils.XYPositionPlanet(P.Betasoide.Grades, P.Betasoide.Distance)
	VX, VY := utils.XYPositionPlanet(P.Vulcano.Grades, P.Vulcano.Distance)

	P.Ferengi.Coordinates = models.CoordinatesPlanet{
		X: FX,
		Y: FY,
	}
	P.Betasoide.Coordinates = models.CoordinatesPlanet{
		X: BX,
		Y: BY,
	}
	P.Vulcano.Coordinates = models.CoordinatesPlanet{
		X: VX,
		Y: VY,
	}

	FerengiSlope := utils.Slope(P.Betasoide.Coordinates, P.Ferengi.Coordinates)
	VulcanoSlope := utils.Slope(P.Betasoide.Coordinates, P.Vulcano.Coordinates)
	SunSlope := utils.Slope(P.Betasoide.Coordinates, Sun)

	if utils.ValidateGrade(P.Ferengi.Grades) && utils.ValidateGrade(P.Betasoide.Grades) && utils.ValidateGrade(P.Vulcano.Grades) {
		P.Drought += 1
		ConditionDay.Condition = "Drought"

		P.WeatherCondition = append(P.WeatherCondition, ConditionDay)
	}

	if FerengiSlope == VulcanoSlope && FerengiSlope != SunSlope {
		P.Optimum += 1
		ConditionDay.Condition = "Optimum"
		P.WeatherCondition = append(P.WeatherCondition, ConditionDay)
	} else {
		FerengiBetasoide := utils.DistanceBetweenPlanets(P.Betasoide.Coordinates, P.Ferengi.Coordinates)
		FerengiVulcano := utils.DistanceBetweenPlanets(P.Vulcano.Coordinates, P.Ferengi.Coordinates)
		BetasoideVulcano := utils.DistanceBetweenPlanets(P.Betasoide.Coordinates, P.Vulcano.Coordinates)

		Perimeter := FerengiBetasoide + FerengiVulcano + BetasoideVulcano
		ConditionDay.Perimeter = Perimeter

		if utils.CenterSunTriangle(P.Betasoide.Coordinates, P.Ferengi.Coordinates, P.Vulcano.Coordinates, Sun) {
			P.Rain += 1
			ConditionDay.Condition = "Rain"
			P.WeatherCondition = append(P.WeatherCondition, ConditionDay)
		} else {
			P.Regular += 1
			ConditionDay.Condition = "Regular"
			P.WeatherCondition = append(P.WeatherCondition, ConditionDay)
		}
	}
}

//IntensityPeak function for know peak of rain intensity
func (P *PlanetsStructure) IntensityPeak() {

	Rain := []models.ConditionForDay{}

	for _, RainCondition := range P.WeatherCondition {
		if RainCondition.Condition == "Rain" {
			Rain = append(Rain, RainCondition)
		}
	}

	MaxPeak := Rain[0].Perimeter

	for _, Condition := range Rain {
		if Condition.Perimeter > MaxPeak {
			MaxPeak = Condition.Perimeter
		}
	}

	for _, Condition := range Rain {
		if Condition.Perimeter == MaxPeak {
			Condition.Condition = "Peak"
			P.RainPeaksDay = append(P.RainPeaksDay, Condition.Day)
		}
	}
}

//End Function for get final structure of predition weather
func (P *PlanetsStructure) End() (Days, Drought, Optimum, Rain int, PK []int, Regular int, Token string) {
	DB.SaveWeatherSystem(P.Drought, P.Optimum, P.Rain, P.Regular)
	DB.SaveWheatherDay(P.WeatherCondition, P.RainPeaksDay)
	DB.CloseConnection()
	return P.Days, P.Drought, P.Optimum, P.Rain, P.RainPeaksDay, P.Regular, P.Token
}
