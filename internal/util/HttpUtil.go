package util

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	minRead = 16 * 1024 // 16kb
)

// 请求方法枚举
const (
	POST string = "POST"
	GET  string = "GET"
	PUT  string = "PUT"
)
const (
	Json      string = "application/json"
	Multipart string = "multipart/form-data"
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

var Client = &HttpConfig{
	// 跳过证书校验
	HttpClient: &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}},
}

func NewResponse() *Response {
	return &Response{}
}

// 请求结构体  内部暂时只有HttpClient客户端后续可对其增加重试等属性
type HttpConfig struct {
	HttpClient *http.Client
}

// post请求
func (c HttpConfig) Post(url string, headers http.Header, body, res interface{}) (statusCode int, err error) {
	return c.request(POST, url, Json, reqBody(Json, body), headers, res)
}

// post请求
func (c HttpConfig) PostReader(url string, headers http.Header, body io.Reader, res interface{}) (statusCode int, err error) {
	return c.request(POST, url, Json, body, headers, res)
}

// get请求
func (c HttpConfig) Get(url string, headers http.Header, body, res interface{}) (statusCode int, err error) {
	return c.request(GET, url, Json, reqBody(Json, body), headers, res)
}

// get请求  请求的 response 为非 json 格式时，用这个方法
func (c HttpConfig) GetString(url string, headers http.Header, body interface{}) (statusCode int, res string, err error) {
	return c.requestString(GET, url, Json, reqBody(Json, body), headers)
}

// put请求
func (c HttpConfig) Put(url string, headers http.Header, contentType string, body, res interface{}) (statusCode int, err error) {
	return c.request(PUT, url, contentType, reqBody(contentType, body), headers, res)
}

// put请求
func (c HttpConfig) PutReader(url string, headers http.Header, contentType string, body io.Reader, res interface{}) (statusCode int, err error) {
	return c.request(PUT, url, contentType, body, headers, res)
}

// put请求，返回response
func (c HttpConfig) PutReturnByteBody(url string, headers http.Header, contentType string, body, res interface{}) ([]byte, error) {
	var (
		bodyBytes  []byte
		err        error
		bodyReader io.Reader
	)
	// Multipart 需要单独处理，contentType 里要有 boundary
	if contentType == Multipart {
		bodyReader, contentType = reqBodyMultipart(body)
		_, bodyBytes, err = c.requestReturnBodyAndStatusCode(PUT, url, contentType, bodyReader, headers, res)
	} else {
		_, bodyBytes, err = c.requestReturnBodyAndStatusCode(PUT, url, contentType, reqBody(contentType, body), headers, res)
	}
	return bodyBytes, err
}

//基础请求方法  内部无重试等机制实现 后续可在 HttpConfig结构体中增加
func (c HttpConfig) requestReturnBodyAndStatusCode(method string, url string, contextType string, body io.Reader, headers http.Header, res interface{}) (int, []byte, error) {
	var (
		response *http.Response
		bs       []byte
	)
	newRequest, err := http.NewRequest(method, url, body)
	if err != nil {
		return 0, bs, errors.Wrap(err, method+" - request creation failed")
	}
	if headers == nil {
		headers = http.Header{}
	}
	headers.Set("Content-Type", contextType)
	newRequest.Header = headers
	response, err = c.HttpClient.Do(newRequest)
	if err != nil {
		return 0, bs, errors.Wrap(err, method+" - do request failed")
	}
	defer response.Body.Close()

	if bs, err = ReadAll(response.Body, minRead); err != nil {
		return response.StatusCode, bs, errors.Wrap(err, "readAll - readAll failed")
	}

	//if response.StatusCode >= http.StatusInternalServerError {
	//	err := errors.New(string(bs))
	//	return response.StatusCode, bs, err
	//}
	if response.StatusCode >= http.StatusBadRequest {
		err := errors.New(string(bs))
		return response.StatusCode, bs, err
	}

	if res != nil && json.Valid(bs) {
		if err = json.Unmarshal(bs, res); err != nil {
			err = errors.Wrap(errors.Wrap(err, "Unmarshal failed"), string(bs))
		}
	}
	return response.StatusCode, bs, err
}

// 只有在请求的 response 为 json 时，会自动将结果解析到 res
func (c HttpConfig) request(method string, url string, contextType string, body io.Reader, headers http.Header, res interface{}) (statusCode int, err error) {
	statusCode, _, err = c.requestReturnBodyAndStatusCode(method, url, contextType, body, headers, res)
	if err != nil {
		return statusCode, err
	}
	return statusCode, err
}

// 请求的 response 为非 json 格式时，用这个方法
func (c HttpConfig) requestString(method string, url string, contextType string, body io.Reader, headers http.Header) (statusCode int, res string, err error) {
	statusCode, bs, err := c.requestReturnBodyAndStatusCode(method, url, contextType, body, headers, nil)
	if err != nil {
		return 0, "", err
	}
	return statusCode, string(bs), err
}

func reqBody(contentType string, param interface{}) (body io.Reader) {
	var err error
	if contentType == Json {
		buff := new(bytes.Buffer)
		err = json.NewEncoder(buff).Encode(param)
		if err != nil {
			return
		}
		body = buff
	}
	return
}

func reqBodyMultipart(param interface{}) (body io.Reader, contentType string) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile := os.Open(param.(string))
	defer file.Close()
	part1, errFile := writer.CreateFormFile("file.text", filepath.Base(param.(string)))
	_, errFile = io.Copy(part1, file)
	if errFile != nil {
		log.Println(errFile)
		return
	}
	contentType = writer.FormDataContentType()
	err := writer.Close()
	if err != nil {
		log.Println(err)
		return
	}
	body = payload
	return
}

func ReadAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}
