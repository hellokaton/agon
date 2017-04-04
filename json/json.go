package json

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

// load config
func LoadJson(filename string) (map[string]string, error) {
	var cmap map[string]string
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}
	if err := json.Unmarshal(bytes, &cmap); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}
	return cmap, nil
}
