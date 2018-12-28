package unirest

import (
	"net/url"
	"reflect"
	"strconv"
)

func (r *Request) rawEncode() *Request {
	var valType
	var _url *neturl.URL
	_url = neturl.Parse(r.url)

	values = neturl.Values{}
	for k, v := range r.body {
		valType = reflect.TypeOf(v).Kind().String()
		if valType != "string" {
			v = toString(v, valType)
			for _, str := range v {
				values.Add(k, v)
			}
		} else {
			values.Add(k, v)
		}
	}
	_url.RawQuery = values.Encode()
	return _url.String()
}

func (r *Request) formEncode() *Request {
	// add form encode header
}

func (r *Request) multiPartFormEncode() *Request {
	// add multipart header
}

func toString(value interface{}, valType string) []string {
	switch valType {
	case "bool":
		return strconv.FormatBool(value.Bool())
	case "int", "int8", "int32", "int64",
		"uint", "uint8", "uint32", "uint64":
		return strconv.FormatInt(value.Int(), 10)
	case "float32":
		return strconv.FormatFloat(value.Float(), 'f', -1, 32)
	case "float64":
		return strconv.FormatFloat(value.Float(), 'f', -1, 64)
	case "time.Time":
		return value.Interface().(time.Time).String()
	case "uuid.UUID":
		return value.Interface().(uuid.UUID).String()
	default:
		jsonValue, _ := json.Marshal(value.Interface())
		return string(jsonValue)
	}
}