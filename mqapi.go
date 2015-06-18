package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	. "github.com/eaciit/mq/client"
	. "github.com/eaciit/mq/helper"
	. "github.com/eaciit/mq/msg"
	"github.com/go-martini/martini"
)

type TokenData struct {
	Token string
	Valid time.Time
}

type PutData struct {
	Node        int
	Owner       string
	Valid, Size int64
}

const (
	tokenLength = 32
)

func main() {
	port := flag.Int("port", 8090, "Port of RCP call. Default is 1234")
	serverHost := flag.String("master", "127.0.0.1:7890", "Default master host")
	flag.Parse()

	client, _ := NewMqClient(*serverHost, time.Second*10)

	m := martini.Classic()
	m.Get("/api/gettoken/username=(?P<name>[a-zA-Z0-9]+)&password=(?P<password>[a-zA-Z0-9]+)", GetToken)
	m.Get("/api/get/token=(?P<token>[a-zA-Z0-9]+)&key=(?P<key>[a-zA-Z0-9]+)", func(w http.ResponseWriter, params martini.Params) {
		Get(w, params, client)
	})
	m.Post("/api/put/token=(?P<token>[a-zA-Z0-9]+)&key=(?P<key>[a-zA-Z0-9]+)", func(w http.ResponseWriter, r *http.Request, params martini.Params) {
		Put(w, r, params, client)
	})

	m.RunOnAddr(fmt.Sprint(":", *port))
}

func GetToken(w http.ResponseWriter, params martini.Params) string {
	var result string
	auth, e := Auth(params["name"], params["password"])
	if auth && e == nil {
		data := TokenData{}
		data.Token = GenerateRandomString(tokenLength)
		data.Valid = time.Now().Add(15 * time.Minute)
		PrintJSON(w, true, data, "")
	} else {
		PrintJSON(w, false, "", "wrong username and password combination")
	}
	return result
}

func Get(w http.ResponseWriter, params martini.Params, c *MqClient) {
	result, err := c.Call("Get", "public|"+params["key"])
	if err != nil {
		PrintJSON(w, false, "", err.Error())
	} else {
		PrintJSON(w, true, result, "")
	}
}

func Put(w http.ResponseWriter, r *http.Request, params martini.Params, c *MqClient) {
	key := BuildKey("", "", params["key"])
	arg := MqMsg{Key: key, Value: r.FormValue("value")}

	item, err := c.Call("Set", arg)

	if err != nil {
		PrintJSON(w, false, "", err.Error())
	} else {
		result := PutData{Owner: item.Owner, Size: item.Size, Valid: item.Duration}
		PrintJSON(w, true, result, "")
	}
}

func Auth(username, password string) (bool, error) {
	c, _ := NewMqClient("127.0.0.1:7890", time.Second*10)
	isLoggedIn := false
	msg := MqMsg{Key: username, Value: password}
	i, e := c.CallToLogin(msg)
	if e != nil {
		return false, e
	}

	if i.Value.(ClientInfo).IsLoggedIn {
		isLoggedIn = true
	}
	return isLoggedIn, nil
}

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(randInt(0, 100))
	}
	return b
}

func GenerateRandomString(s int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
