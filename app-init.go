package main

//go:generate go run app-init.go

import instanceGen "github.com/skeletonkey/lib-instance-gen-go/app"

func main() {
	app := instanceGen.NewApp("tv-tracker", "app")
	app.SetupApp(
		app.WithDependencies("github.com/pioz/tvdb"),
		app.WithGoVersion("1.23"),
		app.WithMakefile(),
		app.WithCodeOwners("* shiny.gift6738@fastmail.com"),
	).Generate()
}
