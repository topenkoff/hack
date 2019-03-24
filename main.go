package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

var (
	db         *sql.DB
	config     producter.Config
	configPath = flag.String("c", "config.json", "-c")
)

func main() {
	flag.Parse()

	file, err := os.Open(*configPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db = connect(config.DBUrl)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.POST("/api/getAllMalls", GetAllMallsHandler)

	e.File("/getIndoorInfo", "DataPoints-2.json")

	e.Start(":" + config.Port)
}
