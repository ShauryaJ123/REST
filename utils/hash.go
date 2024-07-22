package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string,error) {
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes),err
}

func CheckPassword(password ,hashPass string) bool{



	// hashedPwd,err:=HashPassword(password)
	// if err!=nil{
	// 	return false
	// }

	// if hashPass==hashedPwd{
	// 	return true
	// }else{
	// 	return false
	// }
	err:=bcrypt.CompareHashAndPassword([]byte(hashPass),[]byte(password))
	return err==nil
}