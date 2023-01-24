package initializers

import "github.com/Loa212/todo2/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Todo{})
}