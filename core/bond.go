// Package core groups elementary data types and its methods
package core

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/lhdv/treasure_hunt/database"
	"github.com/lhdv/treasure_hunt/util"
)

// Bond group all relevant info scrapped from a bond.
// A typical bond info will look like: Tesouro Selic 2023 01/03/2023      0,00    R$92,63 R$9.263,84
//
//  Name: Bond name, e.g.: Tesouro Selic 2023 (LFT)
//  Kind: If it's amounts for buy(0) or sell(1) bond
//  Index: The index used by the bond to fix a rate, e.g.: Selic
//  OldName: The previous bond naming, e.g.: LTF. - Deprecated
//  DueDate: When the bond expires, e.g.: 01/03/2023
//  Rate: Interests that you will earn, e.g.: 10,18
//  MinPrice: The minimum amount which you can buy a bond, based on its unit price, e.g.: 92.59
//  UnitPrice: How much one bond costs, e.g.: 9.258,94
//  LastUpdate: the last datetime when it was updated by Tesouro Nacional
//  FetchDate: When a bond amounts was fetched
type Bond struct {
	Name       string
	Kind       int
	Index      string
	OldName    string
	DueDate    time.Time
	Rate       float64
	MinPrice   float64
	UnitPrice  float64
	LastUpdate time.Time
	FetchDate  time.Time
}

const (
	// BondBuy const defines the value 0 for a bond which is a buy type
	BondBuy = 0
	// BondSell const defines the value 1 for a bond which is a sell type
	BondSell = 1
)

// ListDistinctBonds gets a distinct list of bond names, its index and rate
func ListDistinctBonds(kind int) []Bond {

	db, err := database.Open("")
	defer database.Close(db)
	if err != nil {
		util.LogError("ListDistinctBonds - database.Open", err)
		return nil
	}

	rows, err := db.Query(database.ListDistinctBonds, kind)
	if err != nil {
		util.LogError("ListDistinctBonds - Query", err)
		return nil
	}

	util.LogMessage("ListDistinctBonds - " + strconv.Itoa(kind))

	defer rows.Close()

	bonds := make([]Bond, 0)

	for rows.Next() {

		var name string
		var index string
		var dueDate string

		b := Bond{}

		err := rows.Scan(&name, &index, &dueDate)
		if err != nil {
			log.Fatal(err)
		}

		b.SetBond(kind, "", name, dueDate, "0", "0", "0", "")

		bonds = append(bonds, b)
	}

	if len(bonds) <= 0 {
		return nil
	}

	return bonds
}

// GetBondsByNameType return an array of bonds for a given name and type(kind)
func GetBondsByNameType(kind int, name string) []Bond {

	db, err := database.Open("")
	defer database.Close(db)
	if err != nil {
		util.LogError("GetBondsByNameType - database.Open", err)
		return nil
	}

	rows, err := db.Query(database.GetBondsByNameType, name, kind)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	util.LogMessage("GetBondsByNameType - " + strconv.Itoa(kind) + " / " + name)

	bonds := make([]Bond, 0)

	for rows.Next() {

		var id int
		var fetchDate string
		var buySell string
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
			&fetchDate,
			&buySell,
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

		b.SetBond(kind, fetchDate, name, dueDate, rate, minPrice, unitPrice, lastUpdate)

		bonds = append(bonds, b)
	}

	if len(bonds) <= 0 {
		return nil
	}

	return bonds
}

// SetBond type fields, extract more additional info from name parameter
func (b *Bond) SetBond(kind int, fetchDate, name, dueDate, rate, minPrice, unitPrice,
	lastUpdate string) {

	b.Kind = kind
	b.Name = name

	bondInfo := util.ExtractBondInfo(b.Name)

	if len(bondInfo) > 0 {
		b.Index = bondInfo[1]
		if len(bondInfo) > 2 {
			b.OldName = bondInfo[2]
		}
	}

	if dueDate != "" {
		b.DueDate = util.FormatDate(dueDate)
	}

	rate = util.FormatStringAmount(rate)
	b.Rate, _ = strconv.ParseFloat(rate, 64)

	minPrice = util.FormatStringAmount(minPrice)
	b.MinPrice, _ = strconv.ParseFloat(minPrice, 64)

	unitPrice = util.FormatStringAmount(unitPrice)
	b.UnitPrice, _ = strconv.ParseFloat(unitPrice, 64)

	if lastUpdate != "" {
		b.LastUpdate = util.FormatDateTime(lastUpdate)
	}

	if fetchDate != "" {
		b.FetchDate = util.FormatDate(fetchDate)
	} else {
		b.FetchDate = util.FormatDateTime(util.FormatTimeNowToStr())
	}

}

// Add a bond to database
func (b *Bond) Add() error {

	var err error

	db, err := database.Open("")
	defer database.Close(db)
	if err != nil {
		util.LogError("Add - database.Open", err)
		return nil
	}

	if checkForSameBondQuote(b) <= 0 {
		_, err = db.Exec(database.InsertBondsTable,
			b.FetchDate.String(),
			b.Kind,
			b.Name,
			b.Index,
			b.OldName,
			b.DueDate.String(),
			b.Rate,
			b.MinPrice,
			b.UnitPrice,
			b.LastUpdate.String())
	}

	return err
}

func (b Bond) String() string {

	var kind string

	if b.Kind == BondBuy {
		kind = "BUY"
	} else {
		kind = "SELL"
	}

	return kind + " - " + b.Name + " / R$ " + strconv.FormatFloat(b.UnitPrice, 'f', 2, 64) + " / " + util.FormatDateTimeToStr(b.LastUpdate)
}

/****************************************************************************
 *
 * Functions not exported
 *
 ****************************************************************************/

func checkForSameBondQuote(b *Bond) int {

	result := 0

	db, err := database.Open("")
	defer database.Close(db)
	if err != nil {
		util.LogError("Add - database.Open", err)
		return result
	}

	err = db.QueryRow(database.CheckBondExist, b.LastUpdate.String(), b.Name, b.Kind).Scan(&result)
	switch {
	case err == sql.ErrNoRows:
		result = 0
	case err != nil:
		log.Fatal(err)
	}

	return result
}
