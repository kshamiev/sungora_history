package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Структура для работы с исходящими запросами
type Request struct {
	url          string
	ResponseBody []byte
	Header       http.Header
}

// New Функционал по работе с исходящими запросами к внешним ресурсам
func New(link string) *Request {
	return &Request{
		url:    link,
		Header: http.Header{},
	}
}

// GET запрос
func (r *Request) GET(uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	return r.request(http.MethodGet, uri, requestBody, responseBody)
}

// POST запрос
func (r *Request) POST(uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	return r.request(http.MethodPost, uri, requestBody, responseBody)
}

// PUT запрос
func (r *Request) PUT(uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	return r.request(http.MethodPut, uri, requestBody, responseBody)
}

// DELETE запрос
func (r *Request) DELETE(uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	return r.request(http.MethodDelete, uri, requestBody, responseBody)
}

// OPTIONS запрос
func (r *Request) OPTIONS(uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	return r.request(http.MethodOptions, uri, requestBody, responseBody)
}

func (r *Request) request(method, uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	var (
		query   = r.url + uri
		request *http.Request
		body    = new(bytes.Buffer)
	)

	const headerTypeJSON = "application/json"

	// Данные исходящего запроса
	if (method == http.MethodPost || method == http.MethodPut) && strings.Split(r.Header.Get("Content-Type"), ";")[0] == headerTypeJSON {
		data, err := json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}

		if _, err = body.Write(data); err != nil {
			return nil, err
		}
	}

	if p, ok := requestBody.(map[string]interface{}); ok {
		query += "?" + uriParamsCompile(p)
	}

	// Запрос
	request, err := http.NewRequest(method, query, body)
	if err != nil {
		return nil, err
	}

	// Заголовки
	request.Header = r.Header
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false}, // ignore expired SSL certificates
	}
	c := http.Client{Transport: transCfg}

	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if r.ResponseBody, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if responseBody != nil && strings.Split(response.Header.Get("Content-Type"), ";")[0] == headerTypeJSON {
		_ = json.Unmarshal(r.ResponseBody, responseBody)
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("%s:[%d]:%s", method, response.StatusCode, query)
	}

	return response, err
}

// uriParamsCompile
func uriParamsCompile(postData map[string]interface{}) string {
	q := &url.Values{}

	for k, v := range postData {
		switch v1 := v.(type) {
		case uint64:
			q.Add(k, strconv.FormatUint(v1, 10))
		case int64:
			q.Add(k, strconv.FormatInt(v1, 10))
		case int:
			q.Add(k, strconv.Itoa(v1))
		case float64:
			q.Add(k, strconv.FormatFloat(v1, 'f', -1, 64))
		case bool:
			q.Add(k, strconv.FormatBool(v1))
		case string:
			q.Add(k, v1)
		}
	}

	return q.Encode()
}
