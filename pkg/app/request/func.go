package request

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/kshamiev/sungora/pkg/app/response"
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
		case strings.Split(headerTypeJSON, ";")[0]:
			data, err = json.Marshal(requestBody)
		case strings.Split(headerTypeXML, ";")[0]:
			data, err = xml.Marshal(requestBody)
		case strings.Split(headerTypeFormURLEncoded, ";")[0]:
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
func (r *Request) requestResiveData(resp *http.Response, responseBody interface{}) (err error) {
	if responseBody != nil {
		switch strings.Split(resp.Header.Get("Content-Type"), ";")[0] {
		case strings.Split(headerTypeJSON, ";")[0]:
			err = json.Unmarshal(r.ResponseBody, responseBody)
		case strings.Split(headerTypeXML, ";")[0]:
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

// ContextGRPC передача данных (метаданных) в контексте по grpc
func ContextGRPC(ctx context.Context, m map[string]string) context.Context {
	if m == nil {
		m = make(map[string]string)
	}
	m[string(response.CtxUUID)] = ctx.Value(response.CtxUUID).(string)
	m[string(response.CtxAPI)] = ctx.Value(response.CtxAPI).(string)

	return metadata.NewOutgoingContext(ctx, metadata.New(m))
}
