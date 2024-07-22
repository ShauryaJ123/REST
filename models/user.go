package models

import (
	"errors"

	"abc.com/calc/db"
	"abc.com/calc/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO USERS (EMAIL,PASSWORD)
	VALUES (?,?)`
	smt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer smt.Close()

	hashedPwd,err:=utils.HashPassword(u.Password)
	if err!=nil{
		return err
	}
	result,err:=smt.Exec(u.Email,hashedPwd)
	if err!=nil{
		return err
	}
	userId,err:=result.LastInsertId()
	u.ID=userId
	return err
}


func(u *User)CheckCredentials()error {
	query:="SELECT id, PASSWORD FROM USERS WHERE EMAIL=?"
	row:=db.DB.QueryRow(query,u.Email)
	//this does not return an error 
	//errors can be found using scan
	var retreviedPassword string

	err:=row.Scan(&u.ID,&retreviedPassword)

	if err!=nil{
		return errors.New("invalid credentials")
	}

	validPassword:=utils.CheckPassword(u.Password,retreviedPassword)

	if !validPassword{
		return errors.New("invalid credentials")
	}

	return nil

	

}