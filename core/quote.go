package core

import (
	"time"
)

// Quote group a series of bonds on a certain date/time (future implementation)
type Quote struct {
	Date  time.Time
	Kind  int
	Bonds []Bond
}

/*
// Retrieve quotes from database
func (q *Quote) GetQuotes(kind int) error {

	sqlQuery := "SELECT * FROM quotes WHERE type = $1 ORDER BY type, bond_name, bond_last_update DESC, date"

	// log.Println("### Getquotes - INIT")

	db := database.Open()

	rows, err := db.Query(sqlQuery, kind)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bonds := make([]Bond, 0)

	var date string
	var typeBond int

	for rows.Next() {

		var id int
		var name string
		var index string
		var oldName string
		var dueDate string
		var rate string
		var minPrice string
		var unitPrice string
		var lastUpdate string

		b := Bond{}

		err := rows.Scan(&id,
			&date,
			&typeBond,
			&name,
			&index,
			&oldName,
			&dueDate,
			&rate,
			&minPrice,
			&unitPrice,
			&lastUpdate)
		if err != nil {
			log.Fatal(err)
		}

		b.SetBond(name, dueDate, rate, minPrice, unitPrice, lastUpdate)

		// log.Println("### Quote.GetQuotes: ", rate, b.Rate, unitPrice, b.UnitPrice)

		bonds = append(bonds, b)
	}

	q.Date = FormatDate(date)

	q.Kind = typeBond
	q.Bonds = bonds

	database.Close(db)

	return err
}
*/
