package services

import (
	"errors"
	"fmt"
	"userapi/daos"
	"userapi/utils"
)

func CreateUser(account, password string) (err error) {
	if account == "" {
		return errors.New("account can't empty")
	}
	if password == "" {
		return errors.New("password can't empty")
	}
	user := queryUser(account)
	if user.Id != 0 {
		return errors.New("user are exist.")
	}
	user = &daos.User{
		Account:  account,
		Password: utils.ToMd5(password),
	}
	return utils.GetDb().Insert(user)
}

func Login(account, password string) (bool, error) {
	if account == "" {
		return false, errors.New("account can't empty")
	}
	if password == "" {
		return false, errors.New("password can't empty")
	}
	user := queryUser(account)
	if user.Id != 0 && user.Password == utils.ToMd5(password) {
		return true, nil
	}
	return false, errors.New("login Failed!")
}

func DeleteUser(account string) bool {
	db := utils.GetDb()
	user := queryUser(account)
	if user.Id == 0 {
		return false
	}
	db.DeleteById(&daos.User{}, "id", user.Id)
	return true
}

func ChangePwdByUser(account, password string) error {
	if account == "" {
		return errors.New("account can't empty")
	}
	if password == "" {
		return errors.New("password can't empty")
	}
	user := queryUser(account)
	if user.Id == 0 {
		return errors.New("account is not exist")
	}
	user.Password = utils.ToMd5(password)
	utils.GetDb().Update(&user)
	return nil
}

func queryUser(account string) *daos.User {
	user := &daos.User{}
	utils.GetDb().QueryByAny(user, "account", account)
	fmt.Println(user)
	return user
}

func IsOk(ok bool) struct{ IsOk bool } {
	return struct {
		IsOk bool
	}{
		IsOk: ok,
	}
}
