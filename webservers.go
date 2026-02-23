package gogetssl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/elfincafe/annette"
)

type (
	GetWebServersResponse struct {
		WebServer []struct {
			Id       int    `json:"id"`
			Software string `json:"software"`
		} `json:"webservers"`
		Success bool `json:"success"`
	}
)

func (api *Api) GetWebServers(supplierId int) (*GetWebServersResponse, error) {
	if supplierId < 1 || supplierId > 2 {
		return nil, errors.New("wrong supplier ID")
	}
	endpoint, err := url.Parse(fmt.Sprintf("%s/tools/webservers/%d", api.LiveAPI, supplierId))
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Set("auth_key", api.Key)
	endpoint.RawQuery = q.Encode()

	var e Error
	var r GetWebServersResponse
	req := annette.New(endpoint)
	res, err := req.Get()
	if err != nil {
		return nil, err
	}
	if !res.IsStatus200s() {
		err = json.Unmarshal(res.Binary(), &e)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s. %s", e.Message, e.Description)
	}
	resBody := res.Binary()

	err = json.Unmarshal(resBody, &e)
	if err == nil && e.Error {
		return nil, fmt.Errorf("%s. %s", e.Message, e.Description)
	}
	err = json.Unmarshal(resBody, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
