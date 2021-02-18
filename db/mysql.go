package db

import (
	"database/sql"
	"fmt"
	"planets/models"
	"planets/utils"
	"strings"

	sq "github.com/Masterminds/squirrel"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	DBConnection *sql.DB
	Token        string
}

//MySQLConnect function for connect to mysql
func MySQLConnect() (dbs *sql.DB) {
	dbs, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/weather")

	if err != nil {
		fmt.Println(err, "Problem with connection to database")
	}

	return dbs
}

//NewDatabase function for connect to database
func NewDatabase() *DB {
	dbs := new(DB)
	dbs.DBConnection = MySQLConnect()

	solarSystemTable := `
		CREATE TABLE IF NOT EXISTS solarsystem(
			solarsystem_id INT AUTO_INCREMENT,
			token VARCHAR(15) NOT NULL,
    	drought INT(12) NULL,
    	optimum INT(12) NULL,
    	rain INT(12) NULL,
    	regular INT(12) NULL,
    	UNIQUE(token),
    	PRIMARY KEY (solarsystem_id, token)
		)ENGINE=INNODB;
	`

	_, errorCreateSolarSystem := dbs.DBConnection.Exec(solarSystemTable)

	if errorCreateSolarSystem != nil {
		fmt.Println(errorCreateSolarSystem)
	}

	weatherTable := `
		CREATE TABLE IF NOT EXISTS weather_day(
			weather_day_id INT AUTO_INCREMENT,
			day VARCHAR(12) NOT NULL,
    	weather varchar(12) NOT NULL,
    	token VARCHAR(15) NOT NULL,
    	FOREIGN KEY(token) REFERENCES solarsystem(token),
    	PRIMARY KEY (weather_day_id)
		)ENGINE=INNODB;
	`

	_, errorCreateWeather := dbs.DBConnection.Exec(weatherTable)

	if errorCreateWeather != nil {
		fmt.Println(errorCreateWeather)
	}

	return dbs
}

//CreateSolarSystem function for intialization of DB
func (db *DB) CreateSolarSystem() string {
	GenerateCode := utils.GenerateCode(3)
	_, errorInsertToken := sq.Insert("solarsystem").
		Columns(
			"token",
		).
		Values(
			GenerateCode,
		).
		RunWith(db.DBConnection).
		Exec()
	db.Token = GenerateCode

	if errorInsertToken != nil {
		fmt.Println(errorInsertToken, "Problem with insert token")
	}

	return GenerateCode
}

//SaveWeatherSystem save information of weather system
func (db *DB) SaveWeatherSystem(drought int, optimum int, rain int, regular int) {
	_, ErrorSolarSystem := sq.Update("solarsystem").
		Set("drought", drought).
		Set("optimum", optimum).
		Set("rain", rain).
		Set("regular", regular).
		Where("token = ?", db.Token).
		RunWith(db.DBConnection).
		Exec()

	if ErrorSolarSystem != nil {
		fmt.Println(ErrorSolarSystem, "Problem with update solar system")
	}
}

//SaveWheatherDay save wheather for day
func (db *DB) SaveWheatherDay(WhearerDays []models.ConditionForDay, PeakDays []int) {
	prepare := "INSERT INTO weather_day(`day`, `weather`, `token`) VALUES "
	modelValues := "(?,?,?)"
	var ColumsValues []string
	values := []interface{}{}

	for i := 0; i < len(WhearerDays); i++ {
		ColumsValues = append(ColumsValues, modelValues)
		if utils.FindArray(WhearerDays[i].Day, PeakDays) {
			WhearerDays[i].Condition = "PeakRain"
		}

		values = append(values, WhearerDays[i].Day, WhearerDays[i].Condition, db.Token)

		if (i % 360) == 0 {
			prepare = prepare + strings.Join(ColumsValues, ",")

			prepareBulk, errorPrepare := db.DBConnection.Prepare(prepare)

			if errorPrepare != nil {
				fmt.Println("Error to preparate bulk", errorPrepare)
			}

			_, errorBulk := prepareBulk.Exec(values...)

			if errorBulk != nil {
				fmt.Println("Error to save bulk", errorBulk)
			} else {
				prepare = "INSERT INTO weather_day(`day`, `weather`, `token`) VALUES "
				modelValues = "(?,?,?)"
				ColumsValues = []string{}
				values = []interface{}{}
			}
		}
	}
}

//GetDay function for get weather day
func (db *DB) GetDay(day int, token string) (string, error) {
	var GetDayWeather models.GetDayWeaterSQL
	ErrorDayWeather := sq.Select("weather").
		From("weather_day").
		Where("day = ? AND token = ?", day, token).
		RunWith(db.DBConnection).
		QueryRow().
		Scan(
			&GetDayWeather.Weather,
		)

	if ErrorDayWeather != nil {
		fmt.Println("Problem with get day weather", ErrorDayWeather)
		return "", ErrorDayWeather
	}

	return GetDayWeather.Weather.String, nil
}

//CloseConnection  function for close connecton to database
func (db *DB) CloseConnection() {
	db.DBConnection.Close()
}
