package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jszwec/csvutil"
)

type Address struct {
	City    string
	Country string
}

type User struct {
	Name string
	Address
	Age       int `csv:"age,omitempty"`
	CreatedAt time.Time
}

func main() {
	users := []User{
		{
			Name:      "John",
			Address:   Address{"Boston", "USA"},
			Age:       26,
			CreatedAt: time.Date(2010, 6, 2, 12, 0, 0, 0, time.UTC),
		},
		{
			Name:    "Alice",
			Address: Address{"SF", "USA"},
		},
	}

	buf := generateCSVDataWithoutHeader(users)
	fmt.Println(buf.String())

	saveCSV(buf, "output-no-header.csv")
}

func generateCSVDataWithoutHeader(users []User) bytes.Buffer {
	var buf bytes.Buffer

	w := csv.NewWriter(&buf)
	w.Comma = '\t'

	enc := csvutil.NewEncoder(w)
	enc.AutoHeader = false
	for _, user := range users {
		if err := enc.Encode(user); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	return buf
}

func saveCSV(buf bytes.Buffer, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	file.Write(buf.Bytes())
	defer file.Close()
}
