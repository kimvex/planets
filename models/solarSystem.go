package models

//PlanetsStructure struct for save planest
type PlanetsStructure struct {
	Ferengi          PlanetModel
	Betasoide        PlanetModel
	Vulcano          PlanetModel
	FerengiGrades    []PlanetGrades
	BetasoideGrades  []PlanetGrades
	VulcanoGrades    []PlanetGrades
	WeatherCondition []ConditionForDay
	Drought          int
	Optimum          int
	Rain             int
	Regular          int
	CurrentDay       int
	RainPeaksDay     []int
	Token            string
	Days             int
}

//PlanetModel struct for generate planets
type PlanetModel struct {
	PlanetName      string
	Distance        int
	Grades          int
	AngularVelocity int
	Coordinates     CoordinatesPlanet
	Clockwise       bool
}

//CoordinatesPlanet struct for position of planet
type CoordinatesPlanet struct {
	X float64
	Y float64
}

//PlanetGrades struc
type PlanetGrades struct {
	Grades int
}

//ConditionForDay struct
type ConditionForDay struct {
	Condition string
	Day       int
	Perimeter float64
}
