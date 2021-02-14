package models

import "database/sql"

//SuccessResponseCreate structure for response of type success api
type SuccessResponseCreate struct {
	Days         int    `json:"Days"`
	DroughtDays  int    `json:"DroughtDays"`
	OptimumDays  int    `json:"OptimumDays"`
	RainDays     int    `json:"RainDays"`
	RainPeaksDay []int  `json:"RainPeaksDay"`
	RegularDays  int    `json:"RegularDays"`
	Token        string `json:"token"`
}

//QueryStringParams structure for get params on query string
type QueryStringParams struct {
	Token    string
	RealLife string
}

//GetDayWeaterSQL struct for get values of sql
type GetDayWeaterSQL struct {
	Weather sql.NullString
}

//SuccessResponseGetDay struct for response day weather
type SuccessResponseGetDay struct {
	Day     int    `json:"day"`
	Weather string `json:"weather"`
}

//ErrorResponse struct for response request error
type ErrorResponse struct {
	Message string `json:"message"`
}
