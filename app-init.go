package main

//go:generate go run app-init.go

import instanceGen "github.com/skeletonkey/lib-instance-gen-go/app"

func main() {
	app := instanceGen.NewApp("tv-tracker", "app")
	app.SetupApp(
		app.WithCodeOwners("* shiny.gift6738@fastmail.com"),
		app.WithDependencies(
			"github.com/go-playground/validator/v10",
			"github.com/google/uuid",
			"github.com/labstack/echo/v4",
			"github.com/mattn/go-sqlite3",
			"github.com/patrickmn/go-cache",
			"golang.org/x/crypto",
		),
		app.WithGoVersion("1.24.1"),
		app.WithMakefile("db"),
		app.WithPackages("db", "server", "tvdb"),
	).Generate()
}
