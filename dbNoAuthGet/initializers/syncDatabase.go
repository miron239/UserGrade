package initializers

import userModel "dbAuth/model"

func SyncDatabase() {
	DB.AutoMigrate(&userModel.UserGrade{})
}
