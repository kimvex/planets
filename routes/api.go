package routes

import (
	"fmt"
	"planets/controllers"
	"planets/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//API function of management for routes
func API(app fiber.Router) {
	app.Get("/start_simulation/:years", StartSimulation)
	app.Get("/weather_day/:day", WheatherDay)
}

//StartSimulation Handler for start simulation
func StartSimulation(c *fiber.Ctx) error {
	years := c.Params("years")
	integerYears, err := strconv.Atoi(years)

	if err != nil {
		fmt.Println("Error to convert years")
	}

	QueryString := new(models.QueryStringParams)
	if err := c.QueryParser(QueryString); err != nil {
		fmt.Println(err, "Error to parsing real life value")
	}

	isRealLife, errRealLife := strconv.ParseBool(QueryString.RealLife)

	if errRealLife != nil {
		fmt.Println("Problem with convert string to bloean")
	}

	intYears, _ := strconv.Atoi(years)

	if intYears > 60 {
		return c.JSON(models.ErrorResponse{Message: "The limit of years to simulate is 60"})
	}

	var response models.SuccessResponseCreate

	Planets := controllers.GeneratePlanets()

	Planets.RunDays(integerYears, isRealLife)

	Days, Drought, Optimum, Rain, RainPeaksDay, Regular, Token := Planets.End()

	response.Days = Days
	response.DroughtDays = Drought
	response.OptimumDays = Optimum
	response.RainDays = Rain
	response.RainPeaksDay = RainPeaksDay
	response.RegularDays = Regular
	response.Token = Token

	return c.JSON(response)
}

//WheatherDay Handler for get weather of a day
func WheatherDay(c *fiber.Ctx) error {
	dayParam := c.Params("day")
	day, _ := strconv.Atoi(dayParam)

	QueryString := new(models.QueryStringParams)
	if err := c.QueryParser(QueryString); err != nil {
		fmt.Println(err, "Error to parsing token")
		return c.JSON(models.ErrorResponse{Message: "Error with token"})
	}

	token := QueryString.Token
	weather, errorGetDay := controllers.GetWeatherDay(day, token)

	if errorGetDay != nil {
		return c.JSON(models.ErrorResponse{Message: "out of range of days"})
	}

	var Response models.SuccessResponseGetDay
	Response.Day = day
	Response.Weather = weather

	return c.JSON(Response)
}
