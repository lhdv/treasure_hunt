package util

import (
	"testing"
	"time"
)

func TestExtractBondInfo(t *testing.T) {

	resultOK := ExtractBondInfo("Tesouro IPCA+ com Juros Semestrais 2020")
	if len(resultOK) <= 0 {
		t.Error("ExtractBondInfo should return an array size > 0.")
	}

	resultErr := ExtractBondInfo("FooBarXpto")
	if len(resultErr) > 0 {
		t.Error("ExtractBondInfo should return an empty array.")
	}
}

func TestFormatStringAmount(t *testing.T) {

	var input, expected, got string

	input = "ABC"
	expected = "0.00"
	got = FormatStringAmount(input)
	if got != expected {
		t.Error("FormatStringAmount return different string. Expected:", expected, "Got:", got)
	}

	input = "0.000"
	expected = "0.000"
	got = FormatStringAmount(input)
	if got != expected {
		t.Error("FormatStringAmount return different string. Expected:", expected, "Got:", got)
	}

	input = "R$ 1.000,00"
	expected = "1000.00"
	got = FormatStringAmount(input)
	if got != expected {
		t.Error("FormatStringAmount return different string. Expected:", expected, "Got:", got)
	}

	input = "R$ 1"
	expected = "1"
	got = FormatStringAmount(input)
	if got != expected {
		t.Error("FormatStringAmount return different string. Expected:", expected, "Got:", got)
	}

	input = "R$ 0,25"
	expected = "0.25"
	got = FormatStringAmount(input)
	if got != expected {
		t.Error("FormatStringAmount return different string. Expected:", expected, "Got:", got)
	}

	input = "0,25"
	expected = "0.25"
	got = FormatStringAmount(input)
	if got != expected {
		t.Error("FormatStringAmount return different string. Expected:", expected, "Got:", got)
	}

}

func TestFormatDate(t *testing.T) {
	var input string
	var expected, got time.Time

	loc, _ := time.LoadLocation(dateLocale)

	input = "01/07/2017"
	expected = time.Date(2017, 7, 1, 0, 0, 0, 0, loc)

	got = FormatDate(input)
	if !got.Equal(expected) {
		t.Error("FormatDate return different values. Expected:", expected, "Got:", got)
	}

	input = "01/90/2017"
	expected = time.Now()

	got = FormatDate(input)
	if got.Hour() != expected.Hour() || got.Minute() != expected.Minute() {
		t.Error("FormatDate return different values. Expected:", expected, "Got:", got)
	}

	input = "2017-07-01 00:01:01 -0300 -03"
	expected = time.Date(2017, 7, 1, 0, 1, 1, 0, loc)

	got = FormatDate(input)
	if !got.Equal(expected) {
		t.Error("FormatDate return different values. Expected:", expected, "Got:", got)
	}

	input = "2017-90-01 00:01:01 -0300 -03"
	expected = time.Now()

	got = FormatDate(input)
	if got.Hour() != expected.Hour() || got.Minute() != expected.Minute() {
		t.Error("FormatDate return different values. Expected:", expected, "Got:", got)
	}
}

func TestFormatDateTime(t *testing.T) {
	var input string
	var expected, got time.Time

	loc, _ := time.LoadLocation(dateLocale)

	input = "01/07/2017"
	expected = time.Date(2017, 7, 1, 0, 0, 0, 0, loc)

	got = FormatDate(input)
	if !got.Equal(expected) {
		t.Error("FormatDate return different values. Expected:", expected, "Got:", got)
	}

	input = "01/90/2017"
	expected = time.Now()

	got = FormatDate(input)
	if got.Hour() != expected.Hour() || got.Minute() != expected.Minute() {
		t.Error("FormatDate return different values. Expected:", expected, "Got:", got)
	}

	input = "2017-07-01 00:01:01 -0300 -03"
	expected = time.Date(2017, 7, 1, 0, 1, 1, 0, loc)

	got = FormatDate(input)
	if !got.Equal(expected) {
		t.Error("FormatDate return different values. Expected:", expected, "Got:", got)
	}

	input = "2017-90-01 00:01:01 -0300 -03"
	expected = time.Now()

	got = FormatDate(input)
	if got.Hour() != expected.Hour() || got.Minute() != expected.Minute() {
		t.Error("FormatDate return different values. Expected:", expected, "Got:", got)
	}
}

func TestFormatDateToStr(t *testing.T) {
	var input time.Time
	var expected, got string

	loc, _ := time.LoadLocation(dateLocale)

	input = time.Date(2017, 7, 1, 0, 0, 0, 0, loc)
	expected = "01/07/2017"

	got = FormatDateToStr(input)
	if got != expected {
		t.Error("FormatDateToStr return different values. Expected:", expected, "Got:", got)
	}
}

func TestFormatDateTimeToStr(t *testing.T) {
	var input time.Time
	var expected, got string

	loc, _ := time.LoadLocation(dateLocale)

	input = time.Date(2017, 7, 1, 10, 10, 10, 0, loc)
	expected = "01/07/2017 10:10:10"

	got = FormatDateTimeToStr(input)
	if got != expected {
		t.Error("FormatDateToStr return different values. Expected:", expected, "Got:", got)
	}
}
