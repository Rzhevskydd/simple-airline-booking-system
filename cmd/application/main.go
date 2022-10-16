package main

import (
	"booking-system/src/app"
)

func main() {
	// TODO аргументы командной строки
	cfg := app.Config{
		Port:   "8000",
		Addr:   "",
		DbHost: "localhost",
		DbPort: "5432",
		DbName: "booking",
		DbUser: "booking",
		DbPwd:  "booking",
	}

	a := app.App{}
	a.Initialize(cfg)

	a.Run(":" + cfg.Port)
}