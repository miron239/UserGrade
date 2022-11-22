package initializers

import userModel "dbAuth/model"

func SyncDatabase() {
	DB.AutoMigrate(&userModel.User{})
	DB.AutoMigrate(&userModel.UserGrade{})
}
