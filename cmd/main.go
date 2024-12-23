package main

import (
	"fmt"

	"github.com/KrittayotToin/ecommerce-platform/config"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Access the loaded database configuration
	fmt.Println("Database Host:", config.AppConfig.Database.Host)
	fmt.Println("Database User:", config.AppConfig.Database.User)
}
