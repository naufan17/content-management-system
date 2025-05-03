package main

import (
	"github.com/naufan17/content-management-system/config"
)

func main() {
	db := config.ConnectDB()

	config.SeedAll(db)
}
