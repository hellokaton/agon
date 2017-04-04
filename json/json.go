package json

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"strconv"
	"errors"
)

type JSON struct {
	data interface{}
}

func NewJson(data string) *JSON {
	j := new(JSON)
	var f interface{}
	err := json.Unmarshal([]byte(data), &f)
	if nil != err {
		return j
	}
	j.data = f
	return j
}

// return map in Go
func (j *JSON) GetMapData() map[string]interface{} {
	if m, ok := (j.data).(map[string]interface{}); ok {
		return m
	}
	return nil
}

// Acoording to the key of the returned data information , return js.data
// if you know json is an object
func (j *JSON) Get(key string) *JSON {
	m := j.GetMapData()
	if v, ok := m[key]; ok {
		return &JSON{data:v}
	}
	j.data = nil
	return j
}

// GetIndex get []interface or map in Go
func (j *JSON) GetIndex(i int) *JSON {
	num := i - 1
	if m, ok := (j.data).([]interface{}); ok {
		if num <= len(m) - 1 {
			v := m[num]
			j.data = v
		} else {
			j.data = nil
		}
		return j
	}
	if m, ok := (j.data).(map[string]interface{}); ok {
		var n = 0
		var data = make(map[string]interface{})
		for i, v := range m {
			if n == num {
				switch vv := v.(type) {
				case float64:
					data[i] = strconv.FormatFloat(vv, 'f', -1, 64)
					j.data = data
					return j
				case string:
					data[i] = vv
					j.data = data
					return j
				case []interface{}:
					j.data = vv
					return j
				}
			}
			n++
		}
	}
	j.data = nil
	return j
}

// The data must be []interface{}, According to your custom number to return key adn array data
func (j *JSON) GetKey(key string, i int) (*JSON, error) {
	num := i - 1
	if i > len((j.data).([]interface{})) {
		return nil, errors.New("index out of range list")
	}
	if m, ok := (j.data).([]interface{}); ok {
		v := m[num].(map[string]interface{})
		if h, ok := v[key]; ok {
			j.data = h
			return j, nil
		}
	}
	j.data = nil
	return j, nil
}

// According to the custom of the PATH to fing element
// You can use function this to find recursive map
func (j *JSON) GetPath(args ...string) *JSON {
	d := j
	for i := range args {
		m := d.GetMapData()
		if val, ok := m[args[i]]; ok {
			d.data = val
		} else {
			d.data = nil
			return d
		}
	}
	return d
}

// String return string
func (j *JSON) String() string {
	if m, ok := j.data.(string); ok {
		return m
	}
	if m, ok := j.data.(float64); ok {
		return strconv.FormatFloat(m, 'f', -1, 64)
	}
	return ""
}

func (j *JSON) ToArray() (k, d []string) {
	var key, data []string
	if m, ok := (j.data).([]interface{}); ok {
		for _, value := range m {
			for index, v := range value.(map[string]interface{}) {
				switch vv := v.(type) {
				case float64:
					data = append(data, strconv.FormatFloat(vv, 'f', -1, 64))
					key = append(key, index)
				case string:
					data = append(data, vv)
					key = append(key, index)
				}
			}
		}
		return key, data
	}
	if m, ok := (j.data).(map[string]interface{}); ok {
		for index, v := range m {
			switch vv := v.(type) {
			case float64:
				data = append(data, strconv.FormatFloat(vv, 'f', -1, 64))
				key = append(key, index)
			case string:
				data = append(data, vv)
				key = append(key, index)
			}
		}
		return key, data
	}
	return nil, nil
}

// Array return array
func (j *JSON) Array() ([]string, error) {
	if a, ok := (j.data).([]interface{}); ok {
		array := make([]string, 0)
		for _, v := range a {
			switch vv := v.(type) {
			case float64:
				array = append(array, strconv.FormatFloat(vv, 'f', -1, 64))
			case string:
				array = append(array, vv)
			}
		}
		return array, nil
	}
	return nil, errors.New("type assertion to []interface{} failed")
}

// type to json string
func (j *JSON) Stringify(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("err was %v", err)
		return ""
	} else {
		return string(b)
	}
}

func (j *JSON) ToString() string {
	b, err := json.Marshal(j.data)
	if err != nil {
		fmt.Printf("err was %v", err)
		return ""
	} else {
		return string(b)
	}
}
// parse string to some type
func (j *JSON) Parse(v string, _type interface{}) {
	json.Unmarshal([]byte(v), &_type)
}

// load config
func (j *JSON) LoadJson(filename string) (map[string]string, error) {
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
