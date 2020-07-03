package libs

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"pb/global"
	"pb/model"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		//log.Println(err)
		return false
	}
	return true
}

func VerifyCode(pwd string) bool {
	var _code []model.AuthenCode

	global.GVA_DB.Find(&_code)
	for _, c := range _code {
		if ComparePasswords(c.Code, []byte(pwd)) {
			return true
		}
	}
	return false
}

func CreateANewCode(pwd string) {
	var hashpwd model.AuthenCode
	hashpwd.Code = HashAndSalt([]byte(pwd))
	if err := global.GVA_DB.Create(&hashpwd).Error; err != nil {
		log.Println("创建code错误")
	} else {
		log.Printf("创建密码：%v 成功", pwd)
	}
}
