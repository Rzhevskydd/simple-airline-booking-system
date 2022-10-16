package app

import (
	"database/sql"
	"errors"
	"fmt"
	//"github.com/Rzhevskydd/techno-db-forum/project/app/units"
	//forumDelivery "github.com/Rzhevskydd/techno-db-forum/project/app/units/forum/delivery"
	//postDelivery "github.com/Rzhevskydd/techno-db-forum/project/app/units/post/delivery"
	//serviceDelivery "github.com/Rzhevskydd/techno-db-forum/project/app/units/service/delivery"
	//threadDelivery "github.com/Rzhevskydd/techno-db-forum/project/app/units/thread/delivery"
	//userDelivery "github.com/Rzhevskydd/techno-db-forum/project/app/units/user/delivery"
	"booking-system/src/repository"
	usecase "booking-system/src/use_cases"
	delivery "booking-system/src/controllers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const ApiPrefix = "/api"

type Config struct {
	Port string
	Addr string
	DbHost string
	DbName string
	DbPort string
	DbUser string
	DbPwd string
}

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Run(addr string) {
	defer func() {
		if err := a.DB.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) Initialize(cfg Config) {
	var err error
	err = a.initializeDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter().PathPrefix(ApiPrefix).Subrouter()

	a.initializeApplication()
}

func (a *App) initializeDatabase(cfg Config) (err error) {
	if a.DB != nil {
		return errors.New("db already initialized")
	}

	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ",
			cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPwd, cfg.DbName)
	a.DB, err = sql.Open("postgres", connectionString)

	a.DB.SetMaxOpenConns(100)
	a.DB.SetMaxIdleConns(30)
	a.DB.SetConnMaxLifetime(time.Hour)

	if query, err := ioutil.ReadFile("sql/db_init_tables.sql"); err != nil {
		return err
	} else {
		_, err = a.DB.Exec(string(query))
		return err
	}

}


func (a *App) initializeApplication() {
	repos := repository.CreateRepositories(a.DB)
	useCases := usecase.NewUseCase(repos)

	flightsController := a.Router.PathPrefix("/flights").Subrouter()
	delivery.HandleFlightsRoutes(flightsController, useCases)

}
