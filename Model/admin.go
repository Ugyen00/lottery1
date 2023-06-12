package model

import (
	postgres "main.go/dataStore"
)

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

// const queryDeleteAdmin = "DELETE FROM admin WHERE email=$1;"

// func (a *Admin) Delete() error {
// 	if _, err := postgres.Db.Exec(queryDeleteUser, a.Email); err != nil {
// 		return err
// 	}
// 	return nil
// }

// const queryUpdateAdmin = "UPDATE admin SET firstname=$1, lastname=$2, email=$3 WHERE Email=$4;"

// func (a *Admin) Update(oldID int64) error {
// 	_, err := postgres.Db.Exec(queryUpdateUser,
// 		a.FirstName, a.LastName, a.Email, oldID)
// 		fmt.Println("dkdkdmoel",err)
// 	return err
// }

// func GetAllAdmin() ([]Admin, error) {
// 	rows, getErr := postgres.Db.Query("SELECT * from admin;")
// 	if getErr != nil {
// 		return nil, getErr
// 	}
// 	Admins := []Admin{}

// 	for rows.Next() {
// 		var a Admin
// 		dbErr := rows.Scan(&a.FirstName, &a.LastName, &a.Email)
// 		if dbErr != nil {
// 			return nil, dbErr
// 		}
// 		Admins = append(Admins, a)
// 	}
// 	rows.Close()
// 	return Admins, nil
// }
