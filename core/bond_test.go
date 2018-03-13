package core

import (
	"testing"
	"time"
)

func TestListDistinctBonds(t *testing.T) {

	buyBonds := ListDistinctBonds(BondBuy)
	if len(buyBonds) <= 0 {
		t.Error("It should returns an array with bonds")
	}

	sellBonds := ListDistinctBonds(BondSell)
	if len(sellBonds) <= 0 {
		t.Error("It should returns an array with bonds")
	}

	noBonds := ListDistinctBonds(999)
	if noBonds != nil {
		t.Error("It should returns nil when a wrong kind was set")
	}
}

func TestGetBondsByNameType(t *testing.T) {
	var bondName string
	var bonds []Bond

	bondName = "Tesouro Prefixado 2023"
	bonds = GetBondsByNameType(BondBuy, bondName)
	if bonds == nil {
		t.Error("It should return a slice of Bonds when a bond name was found")
	}
	bonds = GetBondsByNameType(BondSell, bondName)
	if bonds == nil {
		t.Error("It should return a slice of Bonds when a bond name was found")
	}

	bondName = "FooBar"
	bonds = GetBondsByNameType(BondBuy, bondName)
	if bonds != nil {
		t.Error("It should return nil when a bond name was not found")
	}
	bonds = GetBondsByNameType(BondSell, bondName)
	if bonds != nil {
		t.Error("It should return nil when a bond name was not found")
	}
}

func TestSetBond(t *testing.T) {
	var kind int
	var fetchDate, name, dueDate, rate, minPrice, unitPrice, lastUpdate string

	var b Bond

	//
	// Nonsense fetchDate data
	//
	kind = 0
	fetchDate = "9999-99-99"
	name = ""
	dueDate = ""
	rate = ""
	minPrice = ""
	unitPrice = ""
	lastUpdate = ""

	b.SetBond(kind, fetchDate, name, dueDate, rate, minPrice, unitPrice, lastUpdate)

	if b.FetchDate.Hour() != time.Now().Hour() || b.FetchDate.Minute() != time.Now().Minute() {
		t.Error("Bond FetchDate should be time.Now(). Expected:", time.Now().String(), "Got:", b.FetchDate.String())
	}

	//
	// Nonsense name data
	//
	kind = 0
	fetchDate = ""
	name = "FooBar (XPTO)"
	dueDate = ""
	rate = ""
	minPrice = ""
	unitPrice = ""
	lastUpdate = ""

	b.SetBond(kind, fetchDate, name, dueDate, rate, minPrice, unitPrice, lastUpdate)

	if b.Name != "FooBar (XPTO)" {
		t.Error("Bond Name have the same string. Expected:", name, "Got:", b.Name)
	}

	//
	// Nonsense dueDate data
	//
	kind = 0
	fetchDate = ""
	name = ""
	dueDate = "9XY9-99-99"
	rate = ""
	minPrice = ""
	unitPrice = ""
	lastUpdate = ""

	b.SetBond(kind, fetchDate, name, dueDate, rate, minPrice, unitPrice, lastUpdate)

	if b.DueDate.Hour() != time.Now().Hour() || b.DueDate.Minute() != time.Now().Minute() {
		t.Error("Bond DueDate should be time.Now(). Expected:", time.Now().String(), "Got:", b.DueDate.String())
	}

	//
	// Nonsense rate, minPrice, unitPrice data
	//
	kind = 0
	fetchDate = ""
	name = ""
	dueDate = ""
	rate = "123.XXX"
	minPrice = "zzz123.XXX"
	unitPrice = "0.zz1"
	lastUpdate = ""

	b.SetBond(kind, fetchDate, name, dueDate, rate, minPrice, unitPrice, lastUpdate)

	if b.Rate != 0 {
		t.Error("Bond Rate should be 0. Expected:", 0, "Got:", b.Rate)
	}

	if b.MinPrice != 0 {
		t.Error("Bond MinPrice should be 0. Expected:", 0, "Got:", b.MinPrice)
	}

	if b.UnitPrice != 0 {
		t.Error("Bond UnitPrice should be 0. Expected:", 0, "Got:", b.UnitPrice)
	}

	//
	// Nonsense lastUpdate data
	//
	kind = 0
	fetchDate = ""
	name = ""
	dueDate = ""
	rate = ""
	minPrice = ""
	unitPrice = ""
	lastUpdate = "90-90-90"

	b.SetBond(kind, fetchDate, name, dueDate, rate, minPrice, unitPrice, lastUpdate)

	if b.LastUpdate.Hour() != time.Now().Hour() || b.LastUpdate.Minute() != time.Now().Minute() {
		t.Error("Bond LastUpdate should be time.Now(). Expected:", time.Now().String(), "Got:", b.LastUpdate.String())
	}
}
