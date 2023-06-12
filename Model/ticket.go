package model

import (
	postgres "main.go/dataStore"
)

type Ticket struct {
	TikId     int64  `json:"tikid"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Phone     int64  `json:"phone"`
}


const queryInsertUser = "INSERT INTO ticket(tikid, firstname, lastname, phone) VALUES($1, $2, $3, $4);"

func (s *Ticket) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.TikId, s.FirstName, s.LastName, s.Phone)
	return err
}

const queryGetUser = "SELECT tikid, firstname, lastname, phone FROM ticket WHERE tikid=$1;"

func (s *Ticket) Read() error {
	return postgres.Db.QueryRow(queryGetUser, s.TikId).Scan(&s.TikId, &s.FirstName, &s.LastName, &s.Phone)
}

const queryDeleteUser = "DELETE FROM ticket WHERE tikid=$1;"

func (s *Ticket) Delete() error {
	if _, err := postgres.Db.Exec(queryDeleteUser, s.TikId); err != nil {
		return err
	}
	return nil
}

const queryUpdateUser = "UPDATE ticket SET tikid=$1, firstname=$2, lastname=$3, phone=$4 WHERE tikid=$5;"

func (s *Ticket) Update(oldID int64) error {
	_, err := postgres.Db.Exec(queryUpdateUser,
		s.TikId, s.FirstName, s.LastName, s.Phone, oldID)
	return err
}

func GetAllTickets() ([]Ticket, error) {
	rows, getErr := postgres.Db.Query("SELECT * from ticket;")
	if getErr != nil {
		return nil, getErr
	}
	tickets := []Ticket{}

	for rows.Next() {
		var s Ticket
		dbErr := rows.Scan(&s.TikId, &s.FirstName, &s.LastName, &s.Phone)
		if dbErr != nil {
			return nil, dbErr
		}
		tickets = append(tickets, s)
	}
	rows.Close()
	return tickets, nil
}
