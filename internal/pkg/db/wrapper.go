// +build linux,amd64 darwin,amd64

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	ds "github.com/jimmyislive/gocve/internal/pkg/ds"

	"github.com/lib/pq"
	// make linting happy
	_ "github.com/mattn/go-sqlite3"
)

// RemoveDB xxx
func RemoveDB(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		// If the file did not exist, then this is a no-op
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	fmt.Println(fmt.Sprintf("%v has been removed", fileName))
	return err
}

// PopulateDB populates the DB with cve data from the recordsList
func PopulateDB(cfg *ds.Config, recordsList [][]string) error {
	fmt.Println("Inserting data into DB...")

	var (
		err error
		db  *sql.DB
	)

	if cfg.DBtype == "sqlite" {
		db, err = sql.Open("sqlite3", cfg.DBname)
	} else {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBhost, cfg.DBport, cfg.DBuser, cfg.Password, cfg.DBname)
		db, err = sql.Open("postgres", psqlInfo)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table if not exists cve (cveid text, status text, description text, reference text, phase text, category text);
	delete from cve;
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	//stmt, err := db.Prepare("insert into cve(cveid, status, description, reference, phase, category) values(?, ?, ?, ?, ?, ?)")
	stmt, err := db.Prepare("insert into cve(cveid, status, description, reference, phase, category) values($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < len(recordsList); i++ {
		_, err = stmt.Exec(recordsList[i][0], recordsList[i][1], recordsList[i][2], recordsList[i][3], recordsList[i][4], recordsList[i][5])
		if err != nil {
			//log.Fatal(err)
			fmt.Println(recordsList[i][0])
			fmt.Println(err)
		}
	}

	fmt.Println("DB Created")

	return nil
}

// ListCVE lists all available CVEs from the DB
func ListCVE(cfg *ds.Config) [][]string {
	var (
		err     error
		db      *sql.DB
		records [][]string
	)

	if cfg.DBtype == "sqlite" {
		db, err = sql.Open("sqlite3", cfg.DBname)
	} else {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBhost, cfg.DBport, cfg.DBuser, cfg.Password, cfg.DBname)
		db, err = sql.Open("postgres", psqlInfo)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT cveid, description FROM cve")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cveid, description string

		err = rows.Scan(&cveid, &description)
		if err != nil {
			log.Fatal(err)
		}
		record := []string{cveid, description}
		records = append(records, record)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return records

}

// SearchCVE searches for a pattern in the CVE DB
func SearchCVE(cfg *ds.Config, searchText string) [][]string {

	var (
		records [][]string
		err     error
		db      *sql.DB
	)

	if cfg.DBtype == "sqlite" {
		db, err = sql.Open("sqlite3", cfg.DBname)
	} else {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBhost, cfg.DBport, cfg.DBuser, cfg.Password, cfg.DBname)
		db, err = sql.Open("postgres", psqlInfo)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	searchTextLikeStr := fmt.Sprintf("%%%s%%", searchText)
	stmt := fmt.Sprintf("SELECT cveid, description FROM %s where description LIKE %s OR cveid LIKE %s ", pq.QuoteIdentifier(cfg.Tablename), pq.QuoteLiteral(searchTextLikeStr), pq.QuoteLiteral(searchTextLikeStr))

	rows, err := db.Query(stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cveid, description string

		err = rows.Scan(&cveid, &description)
		if err != nil {
			log.Fatal(err)
		}
		record := []string{cveid, description}
		records = append(records, record)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return records

}

// GetCVE returns details of a specific CVE
func GetCVE(cfg *ds.Config, cveid string) []string {

	var (
		record                                          []string
		status, description, reference, phase, category string
		err                                             error
		db                                              *sql.DB
	)

	if cfg.DBtype == "sqlite" {
		db, err = sql.Open("sqlite3", cfg.DBname)
	} else {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBhost, cfg.DBport, cfg.DBuser, cfg.Password, cfg.DBname)
		db, err = sql.Open("postgres", psqlInfo)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt := fmt.Sprintf("SELECT status, description, reference, phase, category FROM %s where cveid=%s", pq.QuoteIdentifier(cfg.Tablename), pq.QuoteLiteral(cveid))
	row := db.QueryRow(stmt)

	switch err := row.Scan(&status, &description, &reference, &phase, &category); err {
	case nil:
		record = append(record, cveid, status, description, reference, phase, category)
	default:
		fmt.Println(err)
	}

	return record

}
