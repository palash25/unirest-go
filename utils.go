package unirest

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pborman/uuid"
)

func iterateParams(params map[string]interface{}) url.Values {
	var valType string
	values := url.Values{}

	for k, v := range params {
		valType = reflect.TypeOf(v).Kind().String()
		if valType != "string" {
			str := toString(reflect.ValueOf(v), valType)
			values.Add(k, str)
		} else {
			values.Add(k, v.(string))
		}
	}
	return values
}

func (r *Request) rawEncode() *Request {
	var _url *url.URL
	_url, _ = url.Parse(r.url)

	values := iterateParams(r.body.(map[string]interface{}))
	_url.RawQuery = values.Encode()

	r.url = _url.String()
	return r
}

func (r *Request) formEncode() (*Request, error) {
	form := iterateParams(r.body.(map[string]interface{}))
	r.headers.Set("Content-Type", "application/x-www-form-urlencoded")
	req, err := http.NewRequest(r.method, r.url, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	r.HTTPRequest = req
	return r, nil
}

func (r *Request) multiPartFormEncode(paramName, path string, params map[string]interface{}) (*Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val.(string))
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(r.method, r.url, body)
	if err != nil {
		return nil, err
	}
	r.HTTPRequest = req
	r.headers.Set("Content-Type", writer.FormDataContentType())

	return r, nil
}

func toString(value reflect.Value, valType string) string {
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
