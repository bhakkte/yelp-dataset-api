package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/kiasaki/yelp-dataset-api/data"
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

func main() {
	flag.Parse()

	if *fileLocation == "" {
		fmt.Println("File to import location is required")
		os.Exit(1)
	} else {
		fmt.Println("Importing: " + *fileLocation)
	}

	fileHandle, err := os.Open(*fileLocation)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}

	i := 0
	reader := bufio.NewReader(fileHandle)
	line, err := Readln(reader)
	for err == nil {
		if importType == "user" {
			var user data.YelpUser
			if uErr := json.Unmarshal(line, &user); uErr != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if formatted, err := json.MarshalIndent(user, "", "  "); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(string(formatted))
		}
		line, err = Readln(reader)

		i = i + 1
		if i > 10 {
			break
		}
	}
	fmt.Println("Done")
}
