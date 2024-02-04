// csvparser/parser.go

package csvparser

import (
	"encoding/csv"
	"io"
	"os"
)

// Recipient represents the recipient details extracted from the CSV.
type Recipient struct {
	To     string
	Params []string
}

// ParseCSV takes a filepath and returns a slice of Recipients with their details.
func ParseCSV(filepath string) ([]Recipient, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return _parseCSV(file)
}

// ParseCSV takes a filepath and returns a slice of Recipients with their details.
func _parseCSV(reader io.Reader) ([]Recipient, error) {

	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	recipients := map[string]*Recipient{}
	for _, record := range records[1:] {
		to, ok := recipients[record[0]]
		if ok {
			to.Params = append(to.Params, record[1])
		} else {
			recipients[record[0]] = &Recipient{
				To:     record[0],
				Params: []string{record[1]},
			}
		}
	}

	res := []Recipient{}
	for _, r := range recipients {
		res = append(res, *r)
	}

	return res, nil
}
