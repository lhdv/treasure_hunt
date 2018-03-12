package core

import "testing"

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

}
