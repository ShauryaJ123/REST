package models

import (
	"abc.com/calc/db"
	"time"
)

type Event struct {
	ID          int64 
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time 
	UserId 		int
}

func(e *Event) Save() error{
	//later add to DB
	query:=`INSERT INTO EVENTS 
	(NAME,DESCRIPTION,LOCATION,DATETIME,USER_ID)
	VALUES
	(?,?,?,?,?)`
	smt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer smt.Close()
	//can also directly use exec
	result,err:=smt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserId)

	if err!=nil{
		return err
	}
	id,err:=result.LastInsertId()
	if err!=nil{
		return err
	}
	e.ID=id
	return nil
	
}
func GetAll() ([]Event,error){
	query:=`SELECT * FROM EVENTS`
	//PREPARE IS NOT USED ,ONLY USED WHEN WE NEED TO REUSE 
	//WITH DIFFERENT VALUES IN AN EFFICIENT WAY
	//EXEC IS USED WHEN WE NEED CHANGE DATA IN DB
	//QUERY IS USED WHEN WE GET MULTIPLE ROWS IN RETURN 
	rows,err:=db.DB.Query(query)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()

	var events []Event

	for rows.Next(){
		var e Event
		err:=rows.Scan(&e.ID,&e.Name,&e.Description,&e.Location,&e.DateTime,&e.UserId)
		if err!=nil{
			return nil,err
		}
		events=append(events,e )
	}
	return events,nil
	
}

func GetById(ident int64) (*Event,error){
	query:=`SELECT * FROM EVENTS WHERE ID =?`
	
	row:=db.DB.QueryRow(query,ident)
	
	

	var e Event
	err:=row.Scan(&e.ID,&e.Name,&e.Description,&e.Location,&e.DateTime,&e.UserId)
	if err!=nil{
		return nil,err
	}
	
	
	return &e,nil
	
}

func(event Event)Update() error{
	query:=`UPDATE EVENTS SET NAME=?,DESCRIPTION=?,LOCATION=?
	WHERE ID =?`
	smt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer smt.Close()
	_,err=smt.Exec(event.Name,event.Description,event.Location,event.ID)
	return err
}

func(e Event)DeleteIt ()(error){
	query:=`DELETE FROM EVENTS 
	WHERE ID=?`
	smt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer smt.Close()
	_,err=smt.Exec(e.ID)
	return err
}

func (e Event) Register (userId int64) error{
	query:=`insert into registrations values (?,?)`
	smt,err:=db.DB.Prepare(query)
	if err!=nil{
		return nil
	}

	defer smt.Close()

	_,err=smt.Exec(e.ID,userId)
	return err

}

func (e Event) Cancel (userId int64) error{
	query:=`delete from  registrations where event_id=? and user_id=?`
	smt,err:=db.DB.Prepare(query)
	if err!=nil{
		return nil
	}

	defer smt.Close()

	_,err=smt.Exec(e.ID,userId)
	return err

}