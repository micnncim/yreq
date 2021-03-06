package yreq

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

const (
	contentType = "application/json"
)

func SendReq(url string, reqBody []byte) (interface{}, error) {
	r := bytes.NewReader(reqBody)
	resp, err := http.Post(url, contentType, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var dst interface{}
	if err := json.Unmarshal(respBody, &dst); err != nil {
		return nil, err
	}
	return dst, nil
}

func ToYamls(srcs []interface{}) ([]byte, error) {
	if reflect.TypeOf(srcs).Kind() != reflect.Slice {
		return nil, errors.New("request yaml should be array of hash")
	}
	yamlByte, err := yaml.Marshal(srcs)
	if err != nil {
		return nil, err
	}
	return yamlByte, nil
}

func FromYaml(buf []byte) (reqs [][]byte, err error) {
	dst := make([]map[string]interface{}, 0)
	if err = yaml.Unmarshal(buf, &dst); err != nil {
		return
	}
	if reflect.TypeOf(dst).Kind() != reflect.Slice {
		return nil, errors.New("request yaml should be array of hash")
	}

	for _, v := range dst {
		var y []byte
		y, err = yaml.Marshal(v)
		if err != nil {
			return
		}
		var j []byte
		j, err = yaml.YAMLToJSON(y)
		if err != nil {
			return
		}
		reqs = append(reqs, j)
	}
	return
}
