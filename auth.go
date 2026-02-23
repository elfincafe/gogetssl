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
	AuthResponse struct {
		Key string `json:"key"`
	}
)

func (api *Api) Auth(username, apiPassword string) error {
	endpoint, err := url.Parse(fmt.Sprintf("%s/auth", api.LiveAPI))
	if err != nil {
		return err
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
		return err
	}
	if !res.IsStatus200s() {
		return errors.New(res.Body())
	}
	resBody := res.Binary()

	err = json.Unmarshal(resBody, &e)
	if err == nil && e.Error {
		return fmt.Errorf("%s. %s", e.Message, e.Description)
	}
	err = json.Unmarshal(resBody, &r)
	if err != nil {
		return err
	}
	api.Key = r.Key
	return nil
}
