package scrap

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
)

const wrongURLPage = "FOO*BAR://www."
const urlPage = "http://www.example.com"

const htmlPage = `<!DOCTYPE html> 
<html class="ltr" dir="ltr" lang="pt-BR">
<head>
        <title>Tesouro Direto - Preços e Taxas dos Títulos - STN </title>
</head>
<body class=" yui3-skin-sam controls-visible guest-site signed-out public-page site $add_css">
        <table class="mercadostatus mercadoaberto"><tbody> <tr><td class="tittuloTabelaTesouroDireto informacaotabelaTesouroDireto"> <div> Investir </div> <span>Mercado Aberto</span><br> <span class="dataprecos">9h30min às 18h</span><br> <tr class="informacoes"><td><center>Preços e taxas dos títulos públicos disponíveis para <b>investir</b></center></td></tr> </tbody></table> <table cellspacing="0" cellpadding="3" border="0" class="tabelaPrecoseTaxas"> <tbody><tr class="tabelaTitulo"> <th align="center" class="tittuloTabelaTesouroDireto"><center><b> Título</b></center></th> <th align="center" class="tittuloTabelaTesouroDireto"><center><b> Vencimento</b></center></th> <!-- TD class="tabelaTitulo" rowspan=2 align=center><b>Indexador</b></TD --> <th class="tittuloTabelaTesouroDireto"> <center><b>Taxa de Rendimento (% a.a.)</b></center></th> <th class="tittuloTabelaTesouroDireto"> <center><b>Valor Mínimo</b></center></th> <th class="tittuloTabelaTesouroDireto"> <center><b>Preço Unitário</b></center></th> </tr> <tr class="tituloprefixado"><td bgcolor="FFFF9C" class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" ><b>Indexados ao IPCA</b></td> <td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" colspan="6"> </tr> <tr class="camposTesouroDireto "> <td class="listing0">Tesouro IPCA+ 2024</td> <td align="center" class="listing">15/08/2024</td> <td align="center" class="listing ">4,65</td> <td align="center" class="listing ">R$45,73</td> <td align="center" class="listing ">R$2.286,93</td> </tr> <tr class="camposTesouroDireto "> <td class="listing0">Tesouro IPCA+ 2035</td> <td align="center" class="listing">15/05/2035</td> <td align="center" class="listing ">5,11</td> <td align="center" class="listing ">R$39,11</td> <td align="center" class="listing ">R$1.303,81</td> </tr> <tr class="camposTesouroDireto "> <td class="listing0">Tesouro IPCA+ 2045</td> <td align="center" class="listing">15/05/2045</td> <td align="center" class="listing ">5,11</td> <td align="center" class="listing ">R$31,72</td> <td align="center" class="listing ">R$793,18</td> </tr> <tr class="camposTesouroDireto "> <td class="listing0">Tesouro IPCA+ com Juros Semestrais 2026</td> <td align="center" class="listing">15/08/2026</td> <td align="center" class="listing ">4,68</td> <td align="center" class="listing ">R$33,47</td> <td align="center" class="listing ">R$3.347,08</td> </tr> <tr class="camposTesouroDireto "> <td class="listing0">Tesouro IPCA+ com Juros Semestrais 2035</td> <td align="center" class="listing">15/05/2035</td> <td align="center" class="listing ">4,98</td> <td align="center" class="listing ">R$34,73</td> <td align="center" class="listing ">R$3.473,92</td> </tr> <tr class="camposTesouroDireto "> <td class="listing0">Tesouro IPCA+ com Juros Semestrais 2050</td> <td align="center" class="listing">15/08/2050</td> <td align="center" class="listing ">5,06</td> <td align="center" class="listing ">R$35,28</td> <td align="center" class="listing ">R$3.528,07</td> </tr> <tr class="tituloprefixado"><td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preprefixado mercadoaberto" ><b>Prefixados</b></td> <td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" colspan="6"> </tr> <tr class="camposTesouroDireto "> <td class="listing0">Tesouro Prefixado 2021</td> <td align="center" class="listing">01/01/2021</td> <td align="center" class="listing ">8,26</td> <td align="center" class="listing ">R$31,99</td> <td align="center" class="listing ">R$799,87</td> </tr> <tr class="camposTesouroDireto "> <td class="listing0">Tesouro Prefixado 2025</td> <td align="center" class="listing">01/01/2025</td> <td align="center" class="listing ">9,57</td> <td align="center" class="listing ">R$32,22</td> <td align="center" class="listing ">R$537,07</td> </tr> <tr class="camposTesouroDireto "> <td class="listing0">Tesouro Prefixado com Juros Semestrais 2029</td> <td align="center" class="listing">01/01/2029</td> <td align="center" class="listing ">9,69</td> <td align="center" class="listing ">R$31,18</td> <td align="center" class="listing ">R$1.039,49</td> </tr> <tr class="tituloprefixado"><td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" ><b>Indexados à Taxa Selic</b></td> <td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" colspan="6"> </tr> <tr class="camposTesouroDireto "> <td class="listing0">Tesouro Selic 2023</td> <td align="center" class="listing">01/03/2023</td> <td align="center" class="listing ">0,01</td> <td align="center" class="listing ">R$93,91</td> <td align="center" class="listing ">R$9.391,01</td> </tr> </tbody></table> <br> 

        <!-- Resgaste --> <table height="40px" class="mercadostatus mercadoaberto "><tbody> <tr><td class="tittuloTabelaTesouroDireto informacaotabelaTesouroDireto"> <div class="resgatar"> Resgatar </div> <span>Mercado Aberto</span><br> <span class="dataprecos">9h30min às 18h</span><br> </td> </tr> <tr class="informacoes clicksanfonado"><td><center>Preços e taxas dos títulos públicos disponíveis para <b>resgatar</b></center></td></tr> </tbody></table> 

        <div class="sanfonado"> <table cellspacing="0" width="100%" cellpadding="3" border="0" class="tabelaPrecoseTaxas sanfonado"> <tbody><tr class="tabelaTitulo"> <th align="center" class="tittuloTabelaTesouroDireto"><center><b> Título</b></center></th> <th align="center" class="tittuloTabelaTesouroDireto"><center><b> Vencimento</b></center></th> <!-- TD class="tabelaTitulo" rowspan=2 align=center><b>Indexador</b></TD --> <th class="tittuloTabelaTesouroDireto"> <center><b>Taxa de Rendimento (% a.a.)</b></center></th> <th class="tittuloTabelaTesouroDireto"> <center><b>Preço Unitário</b></center></th> </tr> <tr class="tituloprefixado"><td bgcolor="FFFF9C" class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" ><b>Indexados ao IPCA</b></td> <td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" colspan="3"> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IPCA+ 2019</td> <td align="center" class="listing">15/05/2019</td> <td align="center" class="listing ">2,64</td> <td align="center" class="listing ">R$2.969,06</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IPCA+ 2024</td> <td align="center" class="listing">15/08/2024</td> <td align="center" class="listing ">4,77</td> <td align="center" class="listing ">R$2.270,17</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IPCA+ 2035</td> <td align="center" class="listing">15/05/2035</td> <td align="center" class="listing ">5,23</td> <td align="center" class="listing ">R$1.278,58</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IPCA+ 2045</td> <td align="center" class="listing">15/05/2045</td> <td align="center" class="listing ">5,23</td> <td align="center" class="listing ">R$769,03</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IPCA+ com Juros Semestrais 2020</td> <td align="center" class="listing">15/08/2020</td> <td align="center" class="listing ">3,68</td> <td align="center" class="listing ">R$3.233,04</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IPCA+ com Juros Semestrais 2024</td> <td align="center" class="listing">15/08/2024</td> <td align="center" class="listing ">4,69</td> <td align="center" class="listing ">R$3.290,35</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IPCA+ com Juros Semestrais 2026</td> <td align="center" class="listing">15/08/2026</td> <td align="center" class="listing ">4,80</td> <td align="center" class="listing ">R$3.321,08</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IPCA+ com Juros Semestrais 2035</td> <td align="center" class="listing">15/05/2035</td> <td align="center" class="listing ">5,10</td> <td align="center" class="listing ">R$3.430,01</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IPCA+ com Juros Semestrais 2045</td> <td align="center" class="listing">15/05/2045</td> <td align="center" class="listing ">5,17</td> <td align="center" class="listing ">R$3.486,97</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IPCA+ com Juros Semestrais 2050</td> <td align="center" class="listing">15/08/2050</td> <td align="center" class="listing ">5,18</td> <td align="center" class="listing ">R$3.465,36</td> </tr> <tr class="tituloprefixado"><td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preprefixado mercadoaberto" ><b>Prefixados</b></td> <td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" colspan="3"></td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Prefixado 2019</td> <td align="center" class="listing">01/01/2019</td> <td align="center" class="listing ">6,51</td> <td align="center" class="listing ">R$949,98</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Prefixado 2020</td> <td align="center" class="listing">01/01/2020</td> <td align="center" class="listing ">7,42</td> <td align="center" class="listing ">R$878,01</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Prefixado 2021</td> <td align="center" class="listing">01/01/2021</td> <td align="center" class="listing ">8,38</td> <td align="center" class="listing ">R$797,39</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Prefixado 2023</td> <td align="center" class="listing">01/01/2023</td> <td align="center" class="listing ">9,24</td> <td align="center" class="listing ">R$653,96</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Prefixado 2025</td> <td align="center" class="listing">01/01/2025</td> <td align="center" class="listing ">9,69</td> <td align="center" class="listing ">R$533,08</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Prefixado com Juros Semestrais 2021</td> <td align="center" class="listing">01/01/2021</td> <td align="center" class="listing ">8,26</td> <td align="center" class="listing ">R$1.059,20</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Prefixado com Juros Semestrais 2023</td> <td align="center" class="listing">01/01/2023</td> <td align="center" class="listing ">9,03</td> <td align="center" class="listing ">R$1.054,29</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Prefixado com Juros Semestrais 2025</td> <td align="center" class="listing">01/01/2025</td> <td align="center" class="listing ">9,47</td> <td align="center" class="listing ">R$1.044,04</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Prefixado com Juros Semestrais 2027</td> <td align="center" class="listing">01/01/2027</td> <td align="center" class="listing ">9,65</td> <td align="center" class="listing ">R$1.038,81</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Prefixado com Juros Semestrais 2029</td> <td align="center" class="listing">01/01/2029</td> <td align="center" class="listing ">9,81</td> <td align="center" class="listing ">R$1.031,78</td> </tr> <tr class="tituloprefixado"><td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" ><b>Indexados à Taxa Selic</b></td> <td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" colspan="3"></tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Selic 2021</td> <td align="center" class="listing">01/03/2021</td> <td align="center" class="listing ">0,03</td> <td align="center" class="listing ">R$9.387,32</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro Selic 2023</td> <td align="center" class="listing">01/03/2023</td> <td align="center" class="listing ">0,05</td> <td align="center" class="listing ">R$9.372,38</td> </tr> <tr class="tituloprefixado"><td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" ><b>Indexados ao IGP-M</b></td> <td class="tittuloTabelaTesouroDireto prefixadoTesouroDireto preipca mercadoaberto" colspan="3"></tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IGPM+ com Juros Semestrais 2021</td> <td align="center" class="listing">01/04/2021</td> <td align="center" class="listing ">3,96</td> <td align="center" class="listing ">R$3.912,44</td> </tr> <tr class="camposTesouroDireto"> <td class="listing0">Tesouro IGPM+ com Juros Semestrais 2031</td> <td align="center" class="listing">01/01/2031</td> <td align="center" class="listing ">4,93</td> <td align="center" class="listing ">R$6.014,64</td> </tr> </tbody></table> </div> <br> Atualizado em: <b>07/03/2018 15:24</b>
</html> 
`

func TestGetBonds(t *testing.T) {

	bonds := GetBonds(htmlPage)
	if len(bonds) == 33 {
		t.Error("It should returns an array sized on 33")
	}

	noBonds := GetBonds(urlPage)
	if noBonds != nil {
		t.Error("It should returns nil when no bonds were found in URL")
	}

	wrongURL := GetBonds(wrongURLPage)
	if wrongURL != nil {
		t.Error("It should returns nil when a wrong URL was set")
	}

}
func TestGetDocToScrap(t *testing.T) {

	var doc *goquery.Document

	doc = getDocToScrap(urlPage)
	if doc == nil {
		t.Error("A URL was not parsed")
	}

	doc = getDocToScrap(wrongURLPage)
	if doc == nil {
		t.Error("A wrong URL should be parsed as a io.Reader")
	}

	doc = getDocToScrap(htmlPage)
	if doc == nil {
		t.Error("HTML page content was not parsed")
	}
}

func TestBuildBonds(t *testing.T) {

	doc := getDocToScrap(htmlPage)

	b := buildBonds(doc)
	if b == nil {
		t.Error("Bond array not built")
	}
	if len(b) <= 0 {
		t.Error("Bond array zero len")
	}
}

func TestBondGetLastUpdate(t *testing.T) {

	var result string
	var doc *goquery.Document

	doc = getDocToScrap(htmlPage)
	result = bondGetLastUpdate(doc)
	if result == "" {
		t.Error("It should return a string date and time")
	}

	doc = getDocToScrap(urlPage)
	result = bondGetLastUpdate(doc)
	if result != "" {
		t.Error("It should not return a string date and time")
	}
}
