package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-backend/db"
	"go-backend/handler"
	"go-backend/model"
	"go-backend/repository"
	"go-backend/router"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

func main() {
	var cfg model.Config
	loadConfig(&cfg)
	setupEnv(&cfg)

	var sql = new(db.SQL)
	sql.Connect(cfg)
	defer sql.Close()

	e := echo.New()

	userHandler := handler.UserHandler{
		UserRepo: repository.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupAdminRouter()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":8000")))
}

func setupEnv(cfg *model.Config) {
	jwtExpires := strconv.Itoa(cfg.Server.JwrExpires)
	os.Setenv("jwtExpires", jwtExpires)
}

func loadConfig(cfg *model.Config) {
	f, err := os.Open("../../env.pro.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
}
