package main

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/XPFarmers/invoice-generator/config"
	"github.com/XPFarmers/invoice-generator/database"
	"github.com/XPFarmers/invoice-generator/http"
	"github.com/XPFarmers/invoice-generator/invoices"
)

//go:embed html/*
var staticFiles embed.FS

var (
	Version = "dev"
	Build   = "local"
)

func main() {

	fmt.Printf("Version: %s, Build: %s\n", Version, Build)

	err := config.LoadConfig("config.yaml")

	if err != nil {
		panic(err)
	}

	htmlFS, err := fs.Sub(staticFiles, "html")

	if err != nil {
		panic(err)
	}

	db := database.NewDBConnection()
	repo := invoices.NewRepository(db)

	http := http.NewServer(repo, config.Global, htmlFS)

	http.Listen(":8080")
}
