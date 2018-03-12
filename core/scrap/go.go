// Package scrap groups all the functions that will get HTML page and scrap all required data
package scrap

import (
	"net/url"
	"strings"

	"github.com/lhdv/treasure_hunt/core"
	"github.com/lhdv/treasure_hunt/util"

	"github.com/PuerkitoBio/goquery"
)

// GetBonds scrap and build an array containing all scrapped bonds from url
func GetBonds(path string) []core.Bond {
	doc := getDocToScrap(path)
	return buildBonds(doc)
}

// Create the basic document to goquery be able to scrap all the data we want.
func getDocToScrap(path string) *goquery.Document {

	var doc *goquery.Document

	if _, err := url.ParseRequestURI(path); err == nil {
		doc, err = goquery.NewDocument(path)
		if err != nil {
			util.LogError("goquery.NewDocument", err)
		}
	} else {
		r := strings.NewReader(path)
		doc, err = goquery.NewDocumentFromReader(r)
		if err != nil {
			util.LogError("goquery.NewDocumentFromReader", err)
		}
	}

	return doc
}

// Find each html that contains bonds amounts
func buildBonds(doc *goquery.Document) []core.Bond {

	bonds := make([]core.Bond, 0)
	lastUpdate := bondGetLastUpdate(doc)

	// Find specific class to scrap info from all bonds
	cssToFind := ".tabelaPrecoseTaxas"
	doc.Find(cssToFind).Each(func(i int, s *goquery.Selection) {

		// For each bond found e retrieved, append them on master slice
		bb := buildBondFromSelect(i, s, lastUpdate)
		for _, b := range bb {
			bonds = append(bonds, b)
		}
	})

	if len(bonds) <= 0 && lastUpdate == "" {
		bonds = nil
	}

	return bonds

}

// Scrap all bonds from a specific kind and build a slice of them
func buildBondFromSelect(kind int, s *goquery.Selection, lastUpd string) []core.Bond {

	var bonds []core.Bond

	if kind == core.BondBuy || kind == core.BondSell {

		// The CSS class .camposTesouroDireto is applied on <tr>
		s.Find(".camposTesouroDireto").Each(func(i int, ss *goquery.Selection) {

			record := ss.Find("td")

			name := record.First().Text()
			record = record.Next()

			dueDate := record.First().Text()
			record = record.Next()

			rate := record.First().Text()
			record = record.Next()

			minValue := "0"
			if kind == core.BondBuy {
				minValue = record.First().Text()
				record = record.Next()
			}

			value := record.First().Text()

			// log.Printf("%d: %s - %s - %s - %s\n", i, name, rate, minValue, value)

			bond := core.Bond{}
			bond.SetBond(kind, "", name, dueDate, rate, minValue, value, lastUpd)
			bonds = append(bonds, bond)

		})

	}

	return bonds
}

// Retrieve last update date time
func bondGetLastUpdate(doc *goquery.Document) string {
	return doc.Find("div.sanfonado").First().Next().Next().Text()
}
