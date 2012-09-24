package adn

import (
	"encoding/json"
	"testing"
)

var testUserData = []byte(`
{
    "avatar_image":{
      "height":200,
      "url":"https:\/\/d2rfichhc2fb9n.cloudfront.net\/image\/4\/zBIezrirTTAVL1JQFy9mgA8RePruklFAwH9TNCqvR3fxoeGPC-JnYYN3yvkMT907ZvnrEA7HWt3Nk8kBBPmnwa78iRo3HArocVxx97_zjfL5nv-vVvjG-63zeQeSKH1iPdGuxPF70NBajBFy2qkCXGDX44A",
      "width":200
    },
    "counts":{
      "followers":5239,
      "following":455,
      "posts":3121,
      "stars":216
    },
    "cover_image":{
      "height":456,
      "url":"https:\/\/d2rfichhc2fb9n.cloudfront.net\/image\/4\/iV3sRFel7xEjDK7hCB9R0xgIFAH7fQEtmq_KmotfjaDfuOEPbfcy2go0g5_fees1lwanOhzqz-BNSyC9bF09pHQryBA-Gsi_3Z_Gz81IZMibsiFaZ4TajDFDhoBFxE6H4TGXJZxdW1ZGA0VmZQN0n6MsEdM",
      "width":1103
    },
    "created_at":"2012-08-03T01:17:14Z",
    "description":{
      "entities":{
        "hashtags":[
          
        ],
        "links":[
          {
            "len":7,
            "pos":12,
            "text":"App.net",
            "url":"http:\/\/App.net"
          },
          {
            "len":18,
            "pos":28,
            "text":"daltoncaldwell.com",
            "url":"http:\/\/daltoncaldwell.com"
          }
        ],
        "mentions":[
          
        ]
      },
      "html":"<span itemscope=\"https:\/\/app.net\/schemas\/Post\">Founder\/CEO <a href=\"http:\/\/App.net\">App.net<\/a> \r\nBlog: <a href=\"http:\/\/daltoncaldwell.com\">daltoncaldwell.com<\/a><\/span>",
      "text":"Founder\/CEO App.net \r\nBlog: daltoncaldwell.com"
    },
    "follows_you":false,
    "id":"1",
    "is_follower":true,
    "is_following":false,
    "is_muted":false,
    "locale":"en_US",
    "name":"Dalton Caldwell",
    "timezone":"America\/Los_Angeles",
    "type":"human",
    "username":"dalton",
    "you_follow":true,
    "you_muted":false
}`)

func TestUserDecode(t *testing.T) {
	u := &User{}
	if err := json.Unmarshal(testUserData, u); err != nil {
		t.Error("json.Unmarshal(testUserData)", err)
	}
}

func TestUserGet(t *testing.T) {
	u := &User{}
	// Not authenticated. This should fail.
	err := u.Get("@whee")
	if err == nil {
		t.Error("unauthenticated User.Get() should have failed", err)
	}
}
