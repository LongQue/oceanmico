package service

import (
	"oceanmico/common"
	"oceanmico/proto/user"
	"oceanmico/user-service/model"
)

func RegisterUser(username, password, email string) user.UserLoginResponse {
	userModel := model.User{
		Username: username,
		Password: password,
		Email:    email,
		Channel:  1,
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

	if err := create.Error; err != nil {
		message = err.Error()
	} else {
		token = "Test Token"
	}
	return user.UserLoginResponse{
		Message: message,
		Token:   token,
	}
}

func UserLogin(username string, password string) user.UserLoginResponse {
	db := common.GetDB()
	userModel := &model.User{}
	err := db.Where("username = ? And password = ?", username, password).Find(userModel).Error
	message := "Success"
	token := ""
	if err != nil {
		message = "Failed"
	} else {
		token = common.GenToken(userModel.Username, userModel.ID)
	}
	return user.UserLoginResponse{
		Message: message,
		Token:   token,
	}
}
