package initializers

import "github.com/binoymanoj/jwt-auth-go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
