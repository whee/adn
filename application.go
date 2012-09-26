// Copyright (c) 2012 Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package adn

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Application struct {
	Id          string
	Secret      string
	RedirectURI string
	Scopes      []Scope
}

type UserAccess struct {
	Token  string
	Scopes []Scope
}

type Scope string

const (
	ScopeBasic     Scope = "basic"
	ScopeStream          = "stream"
	ScopeEmail           = "email"
	ScopeWritePost       = "write_post"
	ScopeFollow          = "follow"
	ScopeMessages        = "messages"
	ScopeExport          = "export"
)

var DefaultApplication = &Application{}
var apiHttpClient = &http.Client{}

func (c *Application) Do(ua *UserAccess, name string, args EpArgs) (body io.ReadCloser, err error) {
	var path bytes.Buffer
	err = epTemplates.ExecuteTemplate(&path, name, args)
	if err != nil {
		return
	}

	method := apiEndpoints[name].Method
	url := apiHost + path.String()
	req, err := http.NewRequest(string(method), url, nil)
	if err != nil {
		return
	}
	req.Header.Add("X-ADN-Migration-Overrides", "response_envelope=1")

	if ua != nil && ua.Token != "" {
		req.Header.Add("Authorization", "Bearer " + ua.Token)
	}

	resp, err := apiHttpClient.Do(req)
	if err != nil {
		return
	}
	body = resp.Body

	return
}

func (c *Application) Get(ua *UserAccess, name string, args EpArgs, v interface{}) error {
	body, err := c.Do(ua, name, args)
	if err != nil {
		return err
	}
	defer body.Close()

	resp, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	re := &responseEnvelope{Data: v}
	err = json.Unmarshal(resp, re)
	if err != nil {
		return err
	}

	if re.Meta.ErrorId != "" {
		return APIError(re.Meta)
	}

	return err
}
