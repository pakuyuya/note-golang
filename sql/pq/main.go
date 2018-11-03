package main

/**
 * @see https://qiita.com/hiro9/items/e6e41ec822a7077c3568
 * @see @see https://stackoverflow.com/questions/20170275/how-to-find-a-type-of-an-object-in-go
 */

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Tbl struct {
	Id   int
	Name string
}

func main() {
	db, err := sql.Open("postgres", "host=192.168.1.13 port=5555 user=root password=password dbname=postgres sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	id := 3
	//	name := "test"
	//	_ = db.QueryRow("INSERT INTO tbl(id, name) VALUES($1,$2) RETURNING id", id, name)

	fmt.Println(id)

	rows, err := db.Query("SELECT id, name FROM tbl")

	if err != nil {
		fmt.Println(err)
	}

	columns, _ := rows.Columns()

	//
	fmt.Printf("%s", columns)
	for rows.Next() {
		row := make([]interface{}, len(columns))
		for idx := range columns {
			row[idx] = new(MetalScanner)
		}

		err := rows.Scan(row...)
		if err != nil {
			fmt.Println(err)
		}

		for idx, column := range columns {
			var scanner = row[idx].(*MetalScanner)
			fmt.Println(column, ":", scanner.value)
		}
	}
}

type MetalScanner struct {
	valid bool
	value interface{}
}

func (scanner *MetalScanner) getBytes(src interface{}) []byte {
	if a, ok := src.([]uint8); ok {
		return a
	}
	return nil
}

func (scanner *MetalScanner) Scan(src interface{}) error {
	switch src.(type) {
	case int64:
		if value, ok := src.(int64); ok {
			scanner.value = value
			scanner.valid = true
		}
	case float64:
		if value, ok := src.(float64); ok {
			scanner.value = value
			scanner.valid = true
		}
	case bool:
		if value, ok := src.(bool); ok {
			scanner.value = value
			scanner.valid = true
		}
	case string:
		//value := scanner.getBytes(src)
		scanner.value = src.(string)
		scanner.valid = true
	case []byte:
		value := scanner.getBytes(src)
		scanner.value = value
		scanner.valid = true
	case time.Time:
		if value, ok := src.(time.Time); ok {
			scanner.value = value
			scanner.valid = true
		}
	case nil:
		scanner.value = nil
		scanner.valid = true
	}
	return nil
}
