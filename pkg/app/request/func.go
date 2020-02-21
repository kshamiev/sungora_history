package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Подготовка данных для отправки запроса
func (r *Request) requestSendData(method, query string, requestBody interface{}) (string, *bytes.Buffer, error) {
	var (
		err  error
		body = new(bytes.Buffer)
		data = []byte("")
	)

	if method == http.MethodPost || method == http.MethodPut {
		switch strings.Split(r.Header.Get("Content-Type"), ";")[0] {
		case headerTypeJSON:
			data, err = json.Marshal(requestBody)
		case headerTypeXML:
			data, err = xml.Marshal(requestBody)
		case headerTypeFormURLEncoded:
			if p, ok := requestBody.(map[string]interface{}); ok {
				data = []byte(uriParamsCompile(p))
			}
		}

		if err != nil {
			return "", nil, err
		}

		if _, err = body.Write(data); err != nil {
			return "", nil, err
		}
	}

	if p, ok := requestBody.(map[string]interface{}); ok {
		query += "?" + uriParamsCompile(p)
	}

	return query, body, nil
}

// Разбор данных ответа на запрос
func (r *Request) requestResiveData(response *http.Response, responseBody interface{}) (err error) {
	if responseBody != nil {
		switch strings.Split(response.Header.Get("Content-Type"), ";")[0] {
		case headerTypeJSON:
			err = json.Unmarshal(r.ResponseBody, responseBody)
		case headerTypeXML:
			err = xml.Unmarshal(r.ResponseBody, responseBody)
		}
	}

	return
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
