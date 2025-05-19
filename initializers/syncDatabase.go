package initializers

import "jwt-auth-go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
