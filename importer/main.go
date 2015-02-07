package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/kiasaki/yelp-dataset-api/data"
	"labix.org/v2/mgo"
)

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

var fileLocation = flag.String("file", "", "Location on Yelp json dump to import")
var importType = flag.String("type", "user", "Import type, options are: user, business, review")
var dbUrl = flag.String("mongo-url", "mongodb://localhost:27017/yelp-dataset-api", "MongoDB url")

func dialMongo(url string) *mgo.Session {
	dbSession, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	data.Index(dbSession.DB(""))
	return dbSession
}

func acquireFileHandle(location string) *os.File {
	if location == "" {
		fmt.Println("File to import location is required")
		os.Exit(1)
	} else {
		fmt.Println("Importing: " + location)
	}

	fileHandle, err := os.Open(location)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}

	return fileHandle
}

func main() {
	flag.Parse()

	dbSession := dialMongo(*dbUrl)
	fileHandle := acquireFileHandle(*fileLocation)

	i := 0
	reader := bufio.NewReader(fileHandle)
	line, err := Readln(reader)
	for err == nil {
		if *importType == "user" {
			var user data.YelpUser
			if uErr := json.Unmarshal([]byte(line), &user); uErr != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// Save to mongo
			if err := user.Save(dbSession.DB("")); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		line, err = Readln(reader)

		i = i + 1
		if i%1000 == 0 {
			fmt.Printf("Processed %d\n", i)
		}
		if i > 2 {
			break
		}
	}
	fmt.Println("\nDone!")
}
