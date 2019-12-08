package response

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kshamiev/sungora/pkg/logger"
)

const cookiePath = "/"

// Структура для работы с входящим запросом
type Response struct {
	Request  Request
	response http.ResponseWriter
	lg       logger.Logger
}

// New Функционал по работе с входящим запросом для формирования ответа
func New(r *http.Request, w http.ResponseWriter) *Response {
	_ = r.ParseForm()
	var rw = &Response{
		response: w,
		lg:       logger.GetLogger(r.Context()),
		Request:  Request{request: r},
	}
	return rw
}

// CookieGet Получение куки.
func (rw *Response) CookieGet(name string) (c string, err error) {
	sessionID, err := rw.Request.request.Cookie(name)
	if err != nil {
		return "", err
	}
	lg := logger.GetLogger(rw.Request.request.Context())
	lg.WithField("COOKIE", "GET").Infof("%s = %s", name, sessionID.Value)
	return sessionID.Value, nil
}

// CookieSet Установка куки. Если время не указано кука сессионная (пока открыт браузер).
func (rw *Response) CookieSet(name, value string, t ...time.Time) {
	var cookie = new(http.Cookie)
	cookie.HttpOnly = true
	cookie.Name = name
	cookie.Value = value
	cookie.Domain = strings.Split(rw.Request.request.Host, ":")[0]
	cookie.Path = cookiePath
	if len(t) > 0 {
		cookie.Expires = t[0]
	}
	lg := logger.GetLogger(rw.Request.request.Context())
	lg.WithField("COOKIE", "SET").Infof("%s = %s", name, value)
	http.SetCookie(rw.response, cookie)
}

// CookieRem Удаление куков.
func (rw *Response) CookieRem(name string) {
	var cookie = new(http.Cookie)
	cookie.HttpOnly = true
	cookie.Name = name
	cookie.Domain = strings.Split(rw.Request.request.Host, ":")[0]
	cookie.Path = cookiePath
	cookie.Expires = time.Now()
	lg := logger.GetLogger(rw.Request.request.Context())
	lg.WithField("COOKIE", "REM").Infof("%s", name)
	http.SetCookie(rw.response, cookie)
}

// interface for responses with an error
type Error interface {
	Error() string
	Trace() []string
	Response() string
	HTTPCode() int
}

// JsonError ответ с ошибкой в формате json
func (rw *Response) JSONError(err error) {
	if e, ok := err.(Error); ok {
		rw.lg.Error(e.Error())
		for _, t := range e.Trace() {
			rw.lg.Trace(t)
		}
		rw.JSON(e.Response(), e.HTTPCode())
	} else {
		rw.lg.WithError(err).Error("Other (unexpected) error")
		rw.JSON(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// JSON ответ в формате json
func (rw *Response) JSON(object interface{}, status ...int) {
	data, err := json.Marshal(object)
	if err != nil {
		rw.lg.WithField("app", "response").Error(err.Error())
		// Заголовки
		rw.generalHeaderSet("application/json; charset=utf-8", int64(len(data)), status[0])
		// Тело документа
		_, _ = rw.response.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}
	// Статус ответа
	if len(status) == 0 {
		status = append(status, http.StatusOK)
	}
	if status[0] < http.StatusBadRequest {
		rw.lg.Infof("%d:%s:%s", status[0], rw.Request.request.Method, rw.Request.request.URL.Path)
	} else {
		rw.lg.Errorf("%d:%s:%s", status[0], rw.Request.request.Method, rw.Request.request.URL.Path)
	}
	// Заголовки
	rw.generalHeaderSet("application/json; charset=utf-8", int64(len(data)), status[0])
	// Тело документа
	_, _ = rw.response.Write(data)
}

// Static ответ - отдача статических данных
func (rw *Response) Static(pathFile string) {
	var fi os.FileInfo
	var err error
	if fi, err = os.Stat(pathFile); err != nil {
		d := []byte(http.StatusText(http.StatusNotFound))
		rw.Bytes(d, filepath.Base(pathFile), "text/html; charset=utf-8", http.StatusNotFound)
		return
	}
	if fi.IsDir() {
		if rw.Request.request.URL.Path != "/" {
			pathFile += string(os.PathSeparator)
		}
		pathFile += "index.html"
		if _, err = os.Stat(pathFile); err != nil {
			d := []byte(http.StatusText(http.StatusNotFound))
			rw.Bytes(d, filepath.Base(pathFile), "text/html; charset=utf-8", http.StatusNotFound)
			return
		}
	}
	// content
	var data []byte
	if data, err = ioutil.ReadFile(pathFile); err != nil {
		d := []byte(http.StatusText(http.StatusInternalServerError))
		rw.Bytes(d, filepath.Base(pathFile), "text/html; charset=utf-8", http.StatusInternalServerError)
		return
	}
	// type
	var typ = `application/octet-stream`
	l := strings.Split(pathFile, ".")
	fileExt := `.` + l[len(l)-1]
	if mimeType := mime.TypeByExtension(fileExt); mimeType != `` {
		typ = mimeType
	}
	// Аттач если документ не картинка и не текстововой
	if strings.LastIndex(typ, `image`) == -1 && strings.LastIndex(typ, `text`) == -1 {
		rw.response.Header().Set("Content-Disposition", "attachment; filename = "+filepath.Base(pathFile))
	}
	// Заголовки
	rw.generalHeaderSet(typ, int64(len(data)), http.StatusOK)
	// Тело документа
	_, _ = rw.response.Write(data)
}

// Reader ответ
func (rw *Response) Reader(data io.Reader, dataLen int64, fileName, mimeType string, status int) {
	// Аттач если документ не картинка и не текстововой
	if strings.LastIndex(mimeType, `image`) == -1 && strings.LastIndex(mimeType, `text`) == -1 {
		rw.response.Header().Set("Content-Disposition", "attachment; filename = "+fileName)
	}
	// Заголовки
	rw.generalHeaderSet(mimeType, dataLen, status)
	// Тело документа
	_, _ = io.Copy(rw.response, data)
}

// Bytes ответ
func (rw *Response) Bytes(data []byte, fileName, mimeType string, status int) {
	// Аттач если документ не картинка и не текстововой
	if strings.LastIndex(mimeType, `image`) == -1 && strings.LastIndex(mimeType, `text`) == -1 {
		rw.response.Header().Set("Content-Disposition", "attachment; filename = "+fileName)
	}
	// Заголовки
	rw.generalHeaderSet(mimeType, int64(len(data)), status)
	// Тело документа
	_, _ = rw.response.Write(data)
}

// Redirect 301
func (rw *Response) Redirect301(redirectURL string) {
	rw.response.Header().Set("Location", redirectURL)
	rw.response.WriteHeader(http.StatusMovedPermanently)
}

// Redirect 302
func (rw *Response) Redirect302(redirectURL string) {
	rw.response.Header().Set("Location", redirectURL)
	rw.response.WriteHeader(http.StatusFound)
}

// generalHeaderSet общие заголовки любого ответа
func (rw *Response) generalHeaderSet(contentTyp string, l int64, status int) {
	t := time.Now()
	// запрет кеширования
	rw.response.Header().Set("Cache-Control", "no-cache, must-revalidate")
	rw.response.Header().Set("Pragma", "no-cache")
	rw.response.Header().Set("Date", t.Format(time.RFC3339))
	rw.response.Header().Set("Last-Modified", t.Format(time.RFC3339))
	// размер и тип контента
	rw.response.Header().Set("Content-Type", contentTyp)
	if l > 0 {
		rw.response.Header().Set("Content-Length", fmt.Sprintf("%d", l))
	}
	// status
	rw.response.WriteHeader(status)
}
