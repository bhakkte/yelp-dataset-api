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
	fan := make(chan bool)
	errChannel := make(chan error)

	lineRequestChan := make(chan bool)
	lineFeedChan := make(chan string)

	go func() {
		for {
			select {
			case <-lineRequestChan:
				if line, err := Readln(reader); err != nil {
					errChannel <- err
					break
				} else {
					lineFeedChan <- line
				}
			}
		}
	}()

	for count := 0; count < 30; count++ {
		go func() {
			workerSession := dbSession.Copy()
			for {
				var err error
				lineRequestChan <- true
				line := <-lineFeedChan

				var model data.Model
				if *importType == "user" {
					parsed := data.YelpUser{}
					err = json.Unmarshal([]byte(line), &parsed)
					model = parsed
				} else if *importType == "business" {
					parsed := data.YelpBusiness{}
					err = json.Unmarshal([]byte(line), &parsed)
					parsed.Loc = []float32{parsed.Longitude, parsed.Latitude}
					model = parsed
				} else if *importType == "review" {
					parsed := data.YelpReview{}
					err = json.Unmarshal([]byte(line), &parsed)
					model = parsed
				}

				if err != nil {
					errChannel <- err
					break
				}
				if err = data.Save(workerSession.DB(""), model); err != nil {
					errChannel <- err
					break
				}

				fan <- true
			}
		}()
	}

	for {
		select {
		case <-fan:
			i++
			if i%1000 == 0 {
				fmt.Printf("Processed %d\n", i)
			}
		case err := <-errChannel:
			if err.Error() == "EOF" {
				fmt.Println("\nDone!")
				os.Exit(0)
				break
			} else {
				panic(err)
				break
			}
		}
	}
}
