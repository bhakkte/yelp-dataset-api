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

func clearTypeTable(dbSession *mgo.Session, importType string) {
	var collection string
	if importType == "user" {
		collection = "users"
	} else if importType == "business" {
		collection = "businesses"
	} else if importType == "review" {
		collection = "reviews"
	} else {
		fmt.Println("Import type didn't match user, business or review")
		os.Exit(1)
	}
	// Empty db
	if _, err := dbSession.DB("").C(collection).RemoveAll(nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleFatalError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	dbSession := dialMongo(*dbUrl)
	fileHandle := acquireFileHandle(*fileLocation)
	clearTypeTable(dbSession, *importType)

	i := 0
	reader := bufio.NewReader(fileHandle)
	line, err := Readln(reader)
	for err == nil {
		if *importType == "user" {
			model := data.YelpUser{}
			err = json.Unmarshal([]byte(line), &model)
			handleFatalError(err)
			err = data.Save(dbSession.DB(""), model)
			handleFatalError(err)
		} else if *importType == "business" {
			model := data.YelpBusiness{}
			err = json.Unmarshal([]byte(line), &model)
			handleFatalError(err)
			err = data.Save(dbSession.DB(""), model)
			handleFatalError(err)
		} else if *importType == "review" {
			model := data.YelpReview{}
			err = json.Unmarshal([]byte(line), &model)
			handleFatalError(err)
			err = data.Save(dbSession.DB(""), model)
			handleFatalError(err)
		}
		// Read next
		line, err = Readln(reader)

		i = i + 1
		if i%1000 == 0 {
			fmt.Printf("Processed %d\n", i)
		}
		// TODO: Temporary limit to not bloat local db
		if i >= 100 {
			break
		}
	}
	fmt.Println("\nDone!")
}
