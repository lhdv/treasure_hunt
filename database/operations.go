package database

const (
	createBondsTable = `CREATE TABLE bonds (
                id              INTEGER,
                bond_fetch_datetime  TEXT,
                bond_buy_sell        INTEGER,
                bond_name            TEXT,
                bond_index           TEXT,
                bond_old_name        TEXT,
                bond_due_date        TEXT,
                bond_rate            REAL,
                bond_min_price       REAL,
                bond_unit_price      REAL,
                bond_last_update     TEXT,
                PRIMARY KEY(id ASC) );`

	// InsertBondsTable query to insert a new bond into the table
	InsertBondsTable = `INSERT INTO bonds(
                        id                 ,
                        bond_fetch_datetime,
                        bond_buy_sell      ,
                        bond_name          ,
                        bond_index         ,
                        bond_old_name      ,
                        bond_due_date      ,
                        bond_rate          ,
                        bond_min_price     ,
                        bond_unit_price    ,
                        bond_last_update
                ) VALUES (
                        NULL,
                        $1,
                        $2,
                        $3,
                        $4,
                        $5,
                        $6,
                        $7,
                        $8,
                        $9,
                        $10);
                        `
	// CheckBondExist query to check if a bond exists on a specific datetime
	CheckBondExist = `SELECT COUNT(*) FROM bonds WHERE bond_last_update = $1 AND bond_name = $2 AND bond_buy_sell = $3`

	// ListDistinctBonds query list all bonds names from a specific type
	ListDistinctBonds = `SELECT DISTINCT bond_name, bond_index, bond_due_date FROM bonds WHERE bond_buy_sell = $1 ORDER BY bond_name, bond_due_date;`

	// GetBondsByNameType query list all bonds of a specific name and type
	GetBondsByNameType = `SELECT * FROM bonds WHERE bond_name = $1 AND bond_buy_sell = $2 ORDER BY bond_fetch_datetime DESC, bond_due_date;`
)
