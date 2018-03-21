// Package webapp groups all web server functions
package webapp

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/lhdv/treasure_hunt/core"
	"github.com/lhdv/treasure_hunt/core/scrap"
	"github.com/lhdv/treasure_hunt/util"
)

const urlPage = "http://tesouro.fazenda.gov.br/web/stn/tesouro-direto-precos-e-taxas-dos-titulos"

// Server func bring http server up, for the sole pourpose to serve us
func Server(host string, port string) {

	assetsFS := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assetsFS))

	http.HandleFunc("/", rootPage)
	http.HandleFunc("/getBondList", getBondListAction)
	http.HandleFunc("/getBonds", getBondsAction)

	util.LogMessage("Init WebServer")
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}

/****************************************************************************
 *
 * Handler functions (not exported)
 *
 ****************************************************************************/

func rootPage(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Version   string
		Title     string
		BondsBuy  []core.Bond
		BondsSell []core.Bond
	}

	getBondsFromURL()

	bSell := core.ListDistinctBonds(core.BondSell)
	bBuy := core.ListDistinctBonds(core.BondBuy)

	data.Version = "2.1"
	data.Title = "TREASURE HUNT " + data.Version
	data.BondsBuy = bBuy
	data.BondsSell = bSell

	t := template.Must(template.ParseFiles("assets/index.html"))
	err := t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}

}

func getBondsAction(w http.ResponseWriter, r *http.Request) {

	var bondToFind core.Bond

	if r.Method == "POST" {
		util.LogMessage("getBondsAction - POST request")

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&bondToFind)
		if err != nil {
			util.LogError("getBondsAction - Error on json.Decode", err)
		}

		bondQuotes := core.GetBondsByNameType(bondToFind.Kind, bondToFind.Name)
		reponse, err := json.Marshal(bondQuotes)
		if err != nil {
			util.LogError("Error on json.Marshal", err)
		}

		w.Write(reponse)
	}
}

func getBondListAction(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		reqData := struct {
			Kind int
		}{
			Kind: 0,
		}

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&reqData)
		if err != nil {
			util.LogError("getBondListAction - Error on json.Decode", err)
		}

		util.LogMessage("getBondListAction - POST request - " + string(reqData.Kind))

		bonds := core.ListDistinctBonds(reqData.Kind)
		reponse, err := json.Marshal(bonds)
		if err != nil {
			util.LogError("Error on json.Marshal", err)
		}

		w.Write(reponse)
	}

}

func getAssetsAction(w http.ResponseWriter, r *http.Request) {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}

/****************************************************************************
 *
 * Functions not exported
 *
 ****************************************************************************/

func getBondsFromURL() {

	bonds := scrap.GetBonds(urlPage)
	for _, b := range bonds {
		err := b.Add()
		if err != nil {
			util.LogError("Can't save bond ", err)
		}
	}
}
