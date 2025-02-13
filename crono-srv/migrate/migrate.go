package main

import (
	"github.com/krgleon/crono/initializers"
	"github.com/krgleon/crono/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}

func main() {

	initializers.DB.AutoMigrate(&models.User{})
}
