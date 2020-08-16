package service

import (
	"webGo/model"
)

func UserLogin(user model.UserLogin) (map[string]interface{},error) {
	if user.Name=="sdg" && user.Password =="pwd"{
		return (map[string]interface{}{"msg":"success"}),nil
	}else {
		return (map[string]interface{}{"msg":"failed"}),nil
	}

}
