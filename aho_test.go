package go_aho_corasic

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-aho-corasic/aho"
	"io/ioutil"
	"os"
	"testing"
)

type TestData []struct {
	Title  string `json:"title"`
	Type   string `json:"type"`
	Text   string `json:"text"`
	Result []aho.Result
}

func getTestData() TestData {
	jsonFile, err := os.Open("./test-data/test-data.json")
	if err != nil {
		fmt.Println(err)
	}
	var tests TestData
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &tests)
	return tests
}

func getSearches() map[string][]string {
	jsonFile, err := os.Open("./test-data/searches.json")
	if err != nil {
		fmt.Println(err)
	}
	jsonMap := make(map[string][]string)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &jsonMap)
	return jsonMap
}

func TestSearches(t *testing.T) {
	t.Parallel()
	tests := getTestData()
	searches := getSearches()
	for i := 0; i < len(tests); i++ {
		test := tests[i]
		t.Run(test.Title, func(t *testing.T) {
			phrases := searches[test.Type]
			search := aho.NewSearch(phrases)
			search.Build()
			result := search.Exec(test.Text)
			assert.EqualValues(t, test.Result, result, "they should be equal")
		})
	}
}
