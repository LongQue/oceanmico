package service

import (
	"oceanmico/common"
	"oceanmico/proto/user"
	"oceanmico/user-service-2/model"
)

func RegisterUser(username, password, email string) user.UserLoginResponse {
	userModel := model.User{
		Username: username,
		Password: password,
		Email:    email,
		Channel: 2,
	}

	message := "Success"
	token := ""

	//获取db
	db := common.GetDB()
	/**
	获取 gorm.DB类型
	create := db.Create(&userModel)
	err := create.Error
	*/
	create := db.Create(&userModel)

	if err:=create.Error; err != nil {
		message = err.Error()
	} else {
		token = "Test Token"
	}
	return user.UserLoginResponse{
		Message: message,
		Token:   token,
	}
}
