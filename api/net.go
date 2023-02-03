package api

import (
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/redirect"
	"gopkg.in/h2non/gentleman.v2/plugins/timeout"
	"math/rand"
	"time"
)

// NewGetClient returns a new Gentleman (HTTP) client with a GET request
func NewGetClient() *gentleman.Client {
	cli := gentleman.New()
	cli.Use(redirect.Limit(3))
	cli.Use(timeout.Request(20 * time.Second))
	cli.Use(timeout.Dial(5*time.Second, 10*time.Second))
	req := cli.Request()
	req.Method("GET")
	return cli
}

// APIGet performs a GET request on the given URL and returns the status code, body and error
func APIGet(cli *gentleman.Client, url string) (int, []byte, error) {
	var err error
	RandomSleep(2)
	cli.URL(url)
	resp, err := cli.Request().Send()
	if err != nil {
		logger.Warn(err.Error())
		return resp.StatusCode, resp.Bytes(), err

	}
	if resp.Error != nil {
		logger.Warn(err.Error())
		return resp.StatusCode, resp.Bytes(), err
	}
	return resp.StatusCode, resp.Bytes(), err
}

// RandomSleep sleeps for a random number of seconds up to the given number
func RandomSleep(longest int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(longest) // n will be between 0 and longest
	time.Sleep(time.Duration(n) * time.Second)
}
