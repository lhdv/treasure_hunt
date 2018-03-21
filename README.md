TREASURE HUNT
=============

[![Go Report Card](https://goreportcard.com/badge/github.com/lhdv/treasure_hunt)](https://goreportcard.com/report/github.com/lhdv/treasure_hunt)

Webapp build in Go to keep track of Brazilian T-Bonds(Tesouro Direto) rates and values.

Official Bond page: http://www.tesouro.fazenda.gov.br/tesouro-direto

Support Libraries
-----------------

- SQLLite driver: https://github.com/mattn/go-sqlite3
- datatables: https://datatables.net/
- gorm: https://github.com/jinzhu/gorm
- chart.js: http://www.chartjs.org/ 

Process Flow
------------

1. Get page from http://tesouro.fazenda.gov.br/web/stn/tesouro-direto-precos-e-taxas-dos-titulos
2. Look for the last date/time page update
3. Retrieve all bonds for buying, add them to Bond struct and ask for save it
4. Retrieve all bonds for selling, add them to Bond type and ask for save it
5. A quote is just a representation of a group of bonds on a specific date/time (future implementation)

Packages
--------

- assets: files used by weapp package when serving http
- core: main types and logic
- core/scrap: contains the logic from scrapping the site and retrieve data
- database: connect/disconnect and perform database requests
- util: general purpose functions
- webapp: contains all web interface and http handling

Package Core
============

type Bond
---------

- **__Fields__**
    - name(string): Bond name, e.g.: Tesouro Selic 2023 (LFT)
    - kind(string): Generaly the bond index, e.g.: Selic
    - index(string): The financial indicator which interest     is build
    - old_name(string): The previous bond naming, e.g.: LTF
    - due_date(datetime): When the bond expires, e.g.:  01/03/2023 
    - rate(float64): Interests that you will earn, e.g.: 10,18
    - min_price(float64): The minimum amount which you can  buy a bond, based on its unit price, e.g.: 92.59
    - unit_price(float64): How much the bond costs, e.g.:   9.258,94
    - last_update(float64): the last datetime when it was   updated by Tesouro Nacional
    - fetch_datetime(datetime): When the happened the fetch of Tesouro Direto site

- **__Methods__**
    - SetBond: set struct fields
    - Add: save a bond to database.
    - String: Stringer implementation

type Quote (TBD)
----------------

It aggregates buy or sell bonds in a specific fetch_date and fech_time bonds table.

- **__Fields__**
    - datetime(datetime): when a quote was retrieved
    - type(string): if it's a quote for buying or selling a bond
    - bonds(Bond): all the bonds on an specific date and type

- **__Methods__**
    - GetQuote: method to receive type and date or time(or both) and get all bonds

Package Functions
-----------------

- ListDistinctBonds: get a list of all bons already loaded, by its type
- GetBondsByNameType: get a list of bonds by its name and type

Database Definition
===================

Table - BONDS
-------------
```
id          INTEGER,     - unique identifier          
fetch_date  TEXT,        - when(date) the bond quote happened
fetch_time  TEXT,        - when(time) the bond quote happened
buy_sell    INTEGER,     - identify a bond value for buying(0) or selling(1)
name        TEXT,        - bond name, e.g.: Tesouro Selic 2023 (LFT)
index       TEXT,        - which index the bond is based, e.g.: Selic
old_name    TEXT,        - (deprecated) the old name of a bond, e.g.: LTF
due_date    TEXT,        - bond's due date, e.g.: 01/03/2023
rate        REAL,        - interest rate, e.g.: 10,18
min_price   REAL,        - a buying bond has a minimum amount(a fraction of a bond), e.g.: 92.59
unit_price  REAL,        - bond's full price, e.g.: 9.258,94
last_update TEXT,        - last date/time government has updated bond's page
PRIMARY KEY(id ASC)
```
