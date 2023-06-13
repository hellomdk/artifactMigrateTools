package util

import (
	"encoding/base64"
	"net/http"
)

type HttpClient struct {
	BaseURL  string
	Header   http.Header
	Username string
	Password string
}

//func NewClient(config *config.Config) HttpClient {
//	c := &HttpClient{
//		BaseURL:  config.SourceRepo.URL,
//		Username: config.SourceRepo.Username,
//		Password: config.SourceRepo.Password,
//		Header:   http.Header{},
//	}
//	basic := base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))
//	c.Header.Add("Authorization", "Basic "+basic)
//	return *c
//}

func Auth(c HttpClient) HttpClient {
	basic := base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))
	c.Header.Add("Authorization", "Basic "+basic)
	return c
}
