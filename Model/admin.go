package model

import postgres "main.go/dataStore"

type Admin struct {
	FirstName string `json:"fname"`
	LastName string `json:"lname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (a *Admin) CreateAdmin()error{
	const query = "insert into admin(FirstName,LastName,Email,Password) values($1,$2,$3,$4);"
	_,err := postgres.Db.Exec(query,a.FirstName,a.LastName,a.Email,a.Password)
	return err
}
func (a *Admin)Check(email string)error{
	const query = "select * from admin where Email = $1;"
	return postgres.Db.QueryRow(query,email).Scan(&a.FirstName,&a.LastName,&a.Email,&a.Password)
}