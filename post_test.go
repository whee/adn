// Copyright (c) 2012 Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package adn

import (
	"encoding/json"
	"testing"
)

var testPostData = []byte(`
{
    "created_at":"2012-09-21T00:13:05Z",
    "entities":{
      "hashtags":[

      ],
      "links":[
        {
          "len":21,
          "pos":51,
          "text":"http:\/\/tcrn.ch\/UlUQZv",
          "url":"http:\/\/tcrn.ch\/UlUQZv"
        },
        {
          "len":20,
          "pos":74,
          "text":"http:\/\/bit.ly\/SG83ej",
          "url":"http:\/\/bit.ly\/SG83ej"
        }
      ],
      "mentions":[

      ]
    },
    "html":"<span itemscope=\"https:\/\/app.net\/schemas\/Post\">Nope, nothing ironic about this rolling out today: <a href=\"http:\/\/tcrn.ch\/UlUQZv\">http:\/\/tcrn.ch\/UlUQZv<\/a>\n\n<a href=\"http:\/\/bit.ly\/SG83ej\">http:\/\/bit.ly\/SG83ej<\/a><\/span>",
    "id":"489932",
    "machine_only":false,
    "num_replies":7,
    "num_reposts":4,
    "num_stars":1,
    "source":{
      "client_id":"caYWDBvjwt2e9HWMm6qyKS6KcATHUkzQ",
      "link":"https:\/\/alpha.app.net",
      "name":"Alpha"
    },
    "text":"Nope, nothing ironic about this rolling out today: http:\/\/tcrn.ch\/UlUQZv\n\nhttp:\/\/bit.ly\/SG83ej",
    "thread_id":"489932",
    "user":{
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
    },
    "you_reposted":false,
    "you_starred":false
}`)

func TestPostDecode(t *testing.T) {
	p := &Post{}
	if err := json.Unmarshal(testPostData, p); err != nil {
		t.Error("json.Unmarshal(testPostData)", err)
	}
}

func TestPostGet(t *testing.T) {
	if _, err := GetPost("", "511604"); err != nil {
		t.Error("getting post", err)
	}
}
