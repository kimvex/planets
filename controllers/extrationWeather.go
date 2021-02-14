package controllers

import (
	"planets/db"
)

var (
	database *db.DB
)

//GetWeatherDay controller for get weather day
func GetWeatherDay(day int, token string) (string, *string) {
	database = db.NewDatabase()
	wheather, errorGetDay := database.GetDay(day, token)
	database.CloseConnection()

	if errorGetDay != nil {
		errorDB := "Error with get weather"
		return "", &errorDB
	}

	return wheather, nil
}
