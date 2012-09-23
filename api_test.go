package adn

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestGet(t *testing.T) {
	body, err := epExecute("retrieve post", epArgs{Post: "511604"})
	if err != nil {
		t.Error("retrieve post", err)
	}
	defer body.Close()

	resp, err := ioutil.ReadAll(body)
	if err != nil {
		t.Error("reading post", err)
	}

	p := &Post{}
	err = json.Unmarshal(resp, p)
	if err != nil {
		t.Error("unmarshal post", err)
	}

	if p.User.Username != "whee" {
		t.Error("Post username is incorrect")
	}
}
