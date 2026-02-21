package gogetssl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/elfincafe/annette"
)

type (
	Auth struct {
		LiveAPI string
	}
	AuthResponse struct {
		Key string `json:"key,omitempty"`
	}
)

func (api *Auth) Auth(username, apiPassword string) (*AuthResponse, error) {
	endpoint, err := url.Parse(fmt.Sprintf("%s/auth", api.LiveAPI))
	if err != nil {
		return nil, err
	}
	body := url.Values{}
	body.Set("user", username)
	body.Set("pass", apiPassword)

	var e Error
	var r AuthResponse
	req := annette.New(endpoint)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := req.Post(strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}
	if !res.IsStatus200s() {
		return nil, errors.New(res.Body())
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
