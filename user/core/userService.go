package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"user/model"
	"user/services"
)

func BuildUser(item model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	err := model.DB.Where("user_name=?", req.UserName).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.Code = 400
			return nil
		}
		fmt.Println("UserLogin err = ", err)
		resp.Code = 500
		return nil
	}
	if user.CheckPassword(req.Password) == false {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	count := 0
	err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		err = errors.New("该用户名已存在")
		return err
	}
	user := model.User{
		UserName:       req.UserName,
		PasswordDigest: req.Password,
	}
	//加密密码
	err = user.SetPassword(req.Password)
	if err != nil {
		return err
	}
	//创建用户
	err = model.DB.Create(&user).Error
	if err != nil {
		return err
	}
	resp.UserDetail = BuildUser(user)
	return nil
}
