// Package util groups all general purpouse functions
package util

import (
	"log"
	"regexp"
	"strings"
	"time"
)

const dateFormat = "02/01/2006"
const dateTimeFormat = "02/01/2006 15:04"
const dateTimeFullFormat = "2006-01-02 15:04:05 -0700 -07"
const dateLocale = "America/Sao_Paulo"

// ExtractBondInfo receives the bond full name and extract its Index and Old Name
func ExtractBondInfo(name string) []string {

	// This expression was valid until 05/02.
	// Probably because on 02/02 there was an update on Tesouro Direto system
	// This regexp is valid on string: Tesouro IPCA+ com Juros Semestrais 2024 (NTNB)
	//exp := `Tesouro\\s(.*?\\s).*\\((.*?)\\)`

	// // This regexp is valid on string: Tesouro IPCA+ com Juros Semestrais 2020
	exp := `Tesouro\s(.*?\s)`

	re := regexp.MustCompile(exp)

	return re.FindStringSubmatch(name)
}

/****************************************************************************
 *
 * DATA FORMAT FUNCTIONS
 *
 ****************************************************************************/

// FormatStringAmount formats a string like "R$ 99.999,99" to "99999.99"
func FormatStringAmount(input string) string {

	var r string
	var re *regexp.Regexp

	r = strings.Trim(input, " ")

	if r != " " {

		r = strings.TrimPrefix(r, "R$")

		// Transfor numeric string only with matches the pattern 9.999,99
		re = regexp.MustCompile(`\..*\,`)
		if re.MatchString(input) {
			r = strings.Replace(r, ".", "", 1)
			r = strings.Replace(r, ",", ".", 1)
			return r
		}

		re = regexp.MustCompile(`\,.*\.`)
		if re.MatchString(input) {
			r = strings.Replace(r, ",", "", 1)
			return r
		}

		re = regexp.MustCompile(`\,`)
		if re.MatchString(input) {
			r = strings.Replace(r, ",", ".", 1)
			return r
		}

	}

	return r
}

// FormatDate return a date formated as dd/mm/yyyy
func FormatDate(d string) time.Time {

	loc, _ := time.LoadLocation(dateLocale)

	format := ""
	if len(d) == 10 {
		format = dateFormat
	} else {
		format = dateTimeFullFormat
	}

	t, _ := time.ParseInLocation(format, d, loc)

	return t
}

// FormatDateTime return a date and time formated as dd/mm/yyyy hh:mm
func FormatDateTime(d string) time.Time {

	loc, _ := time.LoadLocation(dateLocale)

	format := ""
	if len(d) == 16 {
		format = dateTimeFormat
	} else {
		format = dateTimeFullFormat
	}

	t, err := time.ParseInLocation(format, d, loc)
	if err != nil {
		LogError("", err)
	}

	return t
}

// FormatDateToStr converts time to DD/MM/YYYY
func FormatDateToStr(t time.Time) string {
	return t.Format("02/01/2006")
}

// FormatDateTimeToStr converts time to DD/MM/YYYY HH:MM:SS
func FormatDateTimeToStr(t time.Time) string {
	return t.Format("02/01/2006 15:04:05")
}

// FormatTimeNowToStr converts time.Now() to dateTimeFullFormat
func FormatTimeNowToStr() string {
	return time.Now().Format(dateTimeFullFormat)
}

/****************************************************************************
 *
 * LOGGING FUNCTIONS
 *
 ****************************************************************************/

// LogMessage format a log information message
func LogMessage(msg string) {
	log.Printf("[MSG]: %s", msg)
}

// LogError format a log error message
func LogError(msg string, err error) {
	log.Printf("[ERR]: %v (%s)", err, msg)
}
