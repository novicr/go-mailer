package csvparser

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

type ByTo []Recipient

func (r ByTo) Len() int {
	return len(r)
}

func (r ByTo) Less(i, j int) bool {
	return r[i].To < r[j].To
}

func (r ByTo) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func Test_ParseCSV_success(t *testing.T) {
	csvContent := `
to,param
test1@example.com,123
test1@example.com,124
test2@example.com,125
test1@example.com,126
test3@example.com,127
`

	expected := []Recipient{
		{"test1@example.com", []string{"123", "124", "126"}},
		{"test2@example.com", []string{"125"}},
		{"test3@example.com", []string{"127"}},
	}
	reader := strings.NewReader(csvContent)

	recipients, err := _parseCSV(reader)
	if err != nil {
		t.Error(err)
	}
	sort.Sort(ByTo(recipients))

	if !reflect.DeepEqual(recipients, expected) {
		t.Errorf("\nExpected: %+v \ngot: %+v\n",
			expected, recipients)
	}

}
func Test_ParseCSV_fail(t *testing.T) {
	csvContent := `
to,param
test1@example.com,123,bad
test3@example.com,127
`
	reader := strings.NewReader(csvContent)

	_, err := _parseCSV(reader)
	if err == nil {
		t.Error(err)
	}

}
