package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"

	"lib"
	"lib/i18n"
	"lib/logs"
	typConfig "types/config"
	typDb "types/db"
)

// Потоковые данные и управление вводом и выводом. Прерывания. (Веб сокет сюда войдет так понимаю?)
type RW struct {
	Server         *typConfig.Server   // Конфигурация сервера обрабатывающего запрос
	Writer         http.ResponseWriter // Writer контента
	Request        *http.Request       // Reader контента
	UriSegment     map[string]string   // Сегментированные параметры Uri
	UriParams      map[string][]string // Параметры текущего Uri запроса после знака вопроса
	Content        *content            // Результирующий контент
	Interrupt      bool                // Прерывание выполнения других контрллеров. "Я последний контроллер" (по умолчанию false)
	InterruptHard  bool                // Прерывание выполнения всего включая управляющий контроллер, управляющий контроллер делает return Status (по умолчанию false)
	Lang           string              // Префикс языка
	Log            *logs.Log           // Лог
	Token          string              // Кука (хеш-строка). Для идентификации пользователя
	DocumentRoot   string              // Корневой путь до сайта
	DocumentFolder string              // Папка сайта (host)
}

// NewRW Инициализация потоков ввода и вывода
func NewRW(w http.ResponseWriter, r *http.Request, sc *typConfig.Server) (*RW, bool) {
	var self = new(RW)
	self.Server = sc
	self.Writer = w
	self.Request = r
	self.Lang = Config.Main.Lang
	self.Content = newContent(nil)
	// папка и путь до статичных документов
	dl := strings.Split(sc.Domain, `,`)
	self.DocumentFolder = strings.Split(r.Host, `:`)[0]
	l := strings.Split(self.DocumentFolder, `.`)
	if len(l) > 3 {
		self.DocumentFolder = dl[0]
	} else if len(l) > 2 {
		self.DocumentFolder = l[0] + `.` + l[1] + `.` + l[2]
	} else if self.DocumentFolder != `localhost` {
		self.DocumentFolder = `www.` + self.DocumentFolder
	}
	self.DocumentRoot = Config.View.Path + `/` + self.DocumentFolder
	// проверка допустимости домена
	var ok = false
	for i := range dl {
		if strings.LastIndex(dl[i], self.DocumentFolder) != -1 {
			ok = true
			break
		}
	}
	return self, ok
}

// NewRW Инициализация потоков ввода и вывода
func (self *RW) InitParams(uri *typDb.Uri, uriSegment map[string]string, uriParams map[string][]string) {
	self.UriParams = uriParams
	self.UriSegment = uriSegment
	self.Content = newContent(uri)
	// токен
	if _, ok := uriSegment[`token`]; ok == true {
		self.Token = uriSegment[`token`]
	}
	// язык запроса
	if _, ok := uriSegment[`lang`]; ok == true {
		self.Lang = uriSegment[`lang`]
	}
}

// NewRW Инициализация потоков ввода и вывода
func (self *RW) InitLog(moduleName string) {
	self.Log = logs.NewLog(self.Lang, moduleName)
}

// Lation Перевод по ключевому слову
func (self *RW) Translation(key string, messages ...interface{}) (translation string) {
	return i18n.Translation(self.Lang, key, messages...)
}

// RequestParse Разбор входящего запроса в формате JSON и сохранение в переданный объект
func (self *RW) RequestJsonParse(object interface{}) (err error) {
	var buf []byte
	//var count int
	if buf, err = ioutil.ReadAll(self.Request.Body); err != nil {
		return logs.Error(103, self.Request.Method, self.Request.URL.Path).Error
	}
	if err = json.Unmarshal(buf, object); err != nil {
		return logs.Error(104, self.Request.URL.Path, err).Error
	}
	//model := reflect.TypeOf(object)
	//if model.Kind() == reflect.Ptr {
	//	model = model.Elem()
	//}
	//num := model.NumField()
	//dataString := strings.ToLower(string(buf))
	//for i := 0; i < num; i++ {
	//	field := model.Field(i)
	//	fieldString := strings.ToLower(field.Name)
	//	if 1 < len(strings.Split(dataString, fieldString)) {
	//		count++
	//	}
	//}
	//aa := float64(num)
	//bb := float64(count)
	//ab := aa / 100
	//dd = int(bb / ab)
	return
}

func (self *RW) Redirect(url string) {
	logs.Info(301, url)
	// запрет кеширования
	self.Writer.Header().Set("Cache-Control", "no-cache, must-revalidate")
	self.Writer.Header().Set("Pragma", "no-cache")
	self.Writer.Header().Set("Date", lib.Time.Now().Format(time.RFC1123))
	self.Writer.Header().Set("Last-Modified", lib.Time.Now().Format(time.RFC1123))
	//self.Writer.Header().Set("Location", url)
	//self.Writer.WriteHeader(http.StatusMovedPermanently)
	http.Redirect(self.Writer, self.Request, url, http.StatusMovedPermanently)
	self.InterruptHard = true
	return
}

// Структура отдаваемых данных в формате JSON
type response struct {
	ErrorCode    int         `json:"errorCode"`         // Внешний код сообщения ( 0 < ошибка )
	ErrorMessage string      `json:"errorMessage"`      // Сообщение ( описание кода)
	Content      interface{} `json:"content,omitempty"` // Тело ответа (данные)
}

// ResponseJson Выдача браузеру страницы в формате JSON
func (self *RW) ResponseJson(data interface{}, status int, codeLocal int, messages ...interface{}) (err error) {
	self.Content.Type = `application/json`
	self.Content.Encode = `utf-8`
	self.Content.Status = status
	var code int
	var message string
	_, path, _, ok := runtime.Caller(1)
	if 0 < codeLocal && ok == true {
		moduleName := strings.Split(strings.Split(path, `src/`)[1], `/`)[1]
		code, message = i18n.Message(moduleName, self.Lang, codeLocal, messages...)
	}
	con := new(response)
	con.ErrorCode = code
	con.ErrorMessage = message
	con.Content = data
	if self.Content.Content, err = json.Marshal(con); err != nil {
		lg := logs.Error(105, err)
		con := new(response)
		con.ErrorCode = lg.Code
		con.ErrorMessage = lg.Message
		self.Content.Content, _ = json.Marshal(con)
		self.Content.Status = 500
		self.Response()
		return
	}
	self.Response()
	return
}

// Redirect Переадресация на указанную страницу (301, только для html)
func (self *RW) ResponseFile(filePath string) {
	self.Content.Content = nil
	self.Content.Encode = ``
	self.Content.Type = `application/octet-stream`
	l := strings.Split(filePath, ".")
	fileExt := `.` + l[len(l)-1]
	mimeType := mime.TypeByExtension(fileExt)
	if mimeType != `` {
		self.Content.Type = mimeType
	}
	self.Content.File = filePath
	self.Content.Status = 200
	self.Response()
}

// Redirect Переадресация на указанную страницу (301, только для html)
func (self *RW) ResponseError(status int) {
	self.Content.Type = `text/html`
	self.Content.Encode = `utf-8`
	self.Content.Content = nil
	var path = self.DocumentRoot + `/` + strconv.Itoa(status) + `.html`
	if _, err := os.Stat(path); err != nil {
		path = Config.View.Tpl + `/` + strconv.Itoa(status) + `.html`
	}
	self.Content.Variables[`Message`] = `Доработать передачу ошибок`
	self.Content.ExecuteFile(path)
	self.Content.Status = status
	self.Response()
}

// Отдаем статику (css, images, js, download ...)
func (self *RW) ResponseStatic() bool {
	var pathAbs = self.DocumentRoot + strings.TrimRight(self.Request.URL.Path, `/`)
	if fi, e := os.Stat(pathAbs); e == nil {
		if fi.IsDir() == false {
			self.ResponseFile(pathAbs)
			return true
		} else {
			pathAbs += `/index.html`
			if _, e := os.Stat(pathAbs); e == nil {
				self.ResponseFile(pathAbs)
				return true
			}
		}
	}
	return false
}

// ResponseCustom Выдача браузеру правильных данных в абстрактном формате
// Возможные типы документов
// application/json, text/html, image/*
func (self *RW) Response() {
	// Тип и Кодировка документа
	t := self.Content.Type
	if self.Content.Encode != "" {
		self.Writer.Header().Set("Content-Type", t+"; charset="+self.Content.Encode)
	} else {
		self.Writer.Header().Set("Content-Type", t)
	}

	// Аттач если документ не картинка и не текстововой
	if strings.LastIndex(t, `image`) == -1 && strings.LastIndex(t, `text`) == -1 {
		self.Writer.Header().Set("Content-Disposition", "attachment; filename = "+filepath.Base(self.Content.File))
	}

	// Если контент пустой и указан файл
	if len(self.Content.Content) == 0 && self.Content.File != `` {
		var err error
		if self.Content.Content, err = ioutil.ReadFile(self.Content.File); err != nil {
			self.Content.Content = []byte(err.Error())
			self.Content.Status = 500
		}
	}
	if self.Content.Status == 0 {
		self.Content.Status = 200
	}

	// запрет кеширования
	self.Writer.Header().Set("Cache-Control", "no-cache, must-revalidate")
	self.Writer.Header().Set("Pragma", "no-cache")
	self.Writer.Header().Set("Date", lib.Time.Now().Format(time.RFC1123))
	self.Writer.Header().Set("Last-Modified", lib.Time.Now().Format(time.RFC1123))
	// размер контента
	self.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(self.Content.Content)))
	// Статус ответа
	self.Writer.WriteHeader(self.Content.Status)
	// Тело документа
	self.Writer.Write(self.Content.Content)
	//
	self.InterruptHard = true
	logs.Info(106, self.Content.Status, self.Request.URL.Path)
	return
}

// SetCookie Установка куки. Если время не указано кука сессионная (пока открыт браузер).
func (self *RW) SetCookie(name, value string, t ...time.Time) {
	var cookie = new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Domain = self.Request.URL.Host
	cookie.Path = `/`
	if 0 < len(t) {
		cookie.Expires = t[0]
		logs.Info(101, name, value)
	} else {
		logs.Info(100, name, value)
	}
	http.SetCookie(self.Writer, cookie)
}

// RemCookie Удаление куков.
func (self *RW) RemCookie(name string) {
	var cookie = new(http.Cookie)
	cookie.Name = name
	cookie.Domain = self.Request.URL.Host
	cookie.Path = `/`
	cookie.Expires = lib.Time.Now()
	http.SetCookie(self.Writer, cookie)
	logs.Info(175, name)
}

///////////////////////////////////////////////////////////////////////////////
//
//type uri struct {
//	SegmentAll map[string]string   // Сегментированные параметры Uri
//	Segment    segment             // Зарезервированные сегментированные параметры запроса
//	Params     map[string][]string // Параметры текущего Uri запроса после знака вопроса
//	Path       string              //
//}

// Сегменты запроса
//type segment struct {
//	Parent string // Относительные данные
//	Pid    uint64 // Идентификатор рабочего объекта
//	Self   string // Относительные данные
//	Sid    uint64 // Идентификатор рабочего объекта
//	Child  string // Относительные данные
//	Cid    uint64 // Идентификатор рабочего объекта
//}

//func newUri(r *http.Request) *uri {
//	var self = new(uri)
//	self.SegmentAll = make(map[string]string)
//	self.Params, _ = url.ParseQuery(r.URL.Query().Encode())
//	self.Path = r.URL.Path
//	return self
//}

func (self *RW) GetSegmentUriInt(paramName string) (value int64, ok bool) {
	var err error
	if p, ok := self.UriSegment[paramName]; ok == true {
		if value, err = strconv.ParseInt(p, 0, 64); err != nil {
			logs.Error(135, paramName, self.Request.URL.Path, err)
			return 0, false
		}
		return value, true
	}
	logs.Warning(139, paramName, self.Request.URL.Path)
	return 0, false
}
func (self *RW) GetSegmentUriUint(paramName string) (value uint64, ok bool) {
	var err error
	if p, ok := self.UriSegment[paramName]; ok == true {
		if value, err = strconv.ParseUint(p, 0, 64); err != nil {
			logs.Error(135, paramName, self.Request.URL.Path, err)
			return 0, false
		}
		return value, true
	}
	logs.Warning(139, paramName, self.Request.URL.Path)
	return 0, false
}
func (self *RW) GetSegmentUriFloat(paramName string) (value float64, ok bool) {
	var err error
	if p, ok := self.UriSegment[paramName]; ok == true {
		if value, err = strconv.ParseFloat(p, 64); err != nil {
			logs.Error(135, paramName, self.Request.URL.Path, err)
			return 0, false
		}
		return value, true
	}
	logs.Warning(139, paramName, self.Request.URL.Path)
	return 0, false
}
func (self *RW) GetSegmentUriString(paramName string) (value string, ok bool) {
	value, ok = self.UriSegment[paramName]
	return
}
func (self *RW) GetSegmentUri() map[string]string {
	return self.UriSegment
}

func (self *RW) GetParamUriInt(paramName string) (value int64, ok bool) {
	var err error
	if p, ok := self.UriParams[paramName]; ok == true {
		if value, err = strconv.ParseInt(p[0], 0, 64); err != nil {
			logs.Error(135, paramName, self.Request.URL.Path, err)
			return 0, false
		}
		return value, true
	}
	logs.Warning(139, paramName, self.Request.URL.Path)
	return 0, false
}
func (self *RW) GetParamUriUint(paramName string) (value uint64, ok bool) {
	var err error
	if p, ok := self.UriParams[paramName]; ok == true {
		if value, err = strconv.ParseUint(p[0], 0, 64); err != nil {
			logs.Error(135, paramName, self.Request.URL.Path, err)
			return 0, false
		}
		return value, true
	}
	logs.Warning(139, paramName, self.Request.URL.Path)
	return 0, false
}
func (self *RW) GetParamUriFloat(paramName string) (value float64, ok bool) {
	var err error
	if p, ok := self.UriParams[paramName]; ok == true {
		if value, err = strconv.ParseFloat(p[0], 64); err != nil {
			logs.Error(135, paramName, self.Request.URL.Path, err)
			return 0, false
		}
		return value, true
	}
	logs.Warning(139, paramName, self.Request.URL.Path)
	return 0, false
}
func (self *RW) GetParamUriString(paramName string) (value string, ok bool) {
	if p, o := self.UriParams[paramName]; o == true {
		return p[0], true
	}
	return ``, false
}
func (self *RW) GetParamsUri() map[string][]string {
	return self.UriParams
}

///////////////////////////////////////////////////////////////////////////////

//type log struct {
//	rw         *RW
//	ModuleName string // Имя модуля, контроллер которого работает в текущий момент
//}

//func newLog(rw *RW) *log {
//	var self = new(log)
//	self.rw = rw
//	return self
//}

/*
func (self *log) Info(codeLocal int16, messages ...interface{}) (err error) {
	code, message := i18n.Message(self.ModuleName, self.rw.Lang, codeLocal, messages...)
	logs.Info(code, message)
	return
}

func (self *log) Error(codeLocal int16, messages ...interface{}) (err error) {
	code, message := i18n.Message(self.ModuleName, self.rw.Lang, codeLocal, messages...)
	return logs.Error(code, message).Error
}
*/
///////////////////////////////////////////////////////////////////////////////

// Сессионные данные результирующего контенте
type content struct {
	Content   []byte                 // Шаблон контента - контент (по умолчанию пусто)
	Type      string                 // Тип контента (по умолчанию application/octet-stream)
	Encode    string                 // Кодировка контента для заголовка (по умолчанию пусто)
	File      string                 // Абсолютный путь до файла (или имя файла если он уже прочитан)
	Variables map[string]interface{} // Variable (по умолчанию пустой)
	Functions map[string]interface{} // html/template.FuncMap (по умолчанию пустой)
	Status    int                    // Стандартный код ответа для браузера
}

func newContent(uri *typDb.Uri) *content {
	var self = new(content)
	self.Functions = make(map[string]interface{})
	self.Variables = make(map[string]interface{})
	self.Variables[`BuildInfo`] = `(` + VERSION + `) ` + Config.Main.VersionBuild
	if uri != nil {
		self.File = uri.Uri
		self.Content = uri.Content
		self.Encode = uri.ContentEncode
		self.Type = uri.ContentType
		self.Variables[`Title`] = uri.Title
		self.Variables[`Keywords`] = uri.KeyWords
		self.Variables[`Description`] = uri.Description
	}
	self.Status = 200
	return self
}

// ExecuteFile Выполнение шаблона (вставка данных в шаблон)
// tplPath - абсолютный путь до шаблона
func (self *content) ExecuteFile(tplPath string) (err error) {
	//self.Variables[`Content`] = string(self.Content)
	if _, ok := self.Variables[`Content`]; ok == false {
		self.Variables[`Content`] = string(self.Content)
	}
	tpl, err := template.New(filepath.Base(tplPath)).Funcs(self.Functions).ParseFiles(tplPath)
	if err != nil {
		return err
	}
	return self.execute(tpl)
}

// ExecuteFile Выполнение шаблона (вставка данных в шаблон)
// tplPath - абсолютный путь до шаблона
func (self *content) ExecuteBytes(data []byte) (err error) {
	//var data1 = self.Content
	//self.Variables[`Content`] = string(data1)
	//self.Variables[`Content`] = string(self.Content)
	if _, ok := self.Variables[`Content`]; ok == false {
		self.Variables[`Content`] = string(self.Content)
	}
	tpl, err := template.New(`ExecuteStringTemplate`).Funcs(self.Functions).Parse(string(data))
	if err != nil {
		return err
	}
	return self.execute(tpl)
	//logs.Dumper(self)
	//return
}

func (self *content) ExecuteString(str string) (err error) {
	//var data1 = self.Content
	//self.Variables[`Content`] = string(data1)
	if _, ok := self.Variables[`Content`]; ok == false {
		self.Variables[`Content`] = string(self.Content)
	}
	tpl, err := template.New(`ExecuteStringTemplate`).Funcs(self.Functions).Parse(str)
	if err != nil {
		return err
	}
	return self.execute(tpl)
	//logs.Dumper(self)
	//return
}

// ExecuteBytes Выполнение шаблона (вставка данных в шаблон)
func (self *content) Execute() error {
	//var data = make([]byte, len(self.Content))
	//copy(data, self.Content)
	//var data = self.Content
	tpl, err := template.New(`ExecuteStringTemplate`).Funcs(self.Functions).Parse(string(self.Content))
	if err != nil {
		return err
	}
	return self.execute(tpl)
}

// ExecuteBytes Выполнение шаблона (вставка данных в шаблон)
func (self *content) execute(tpl *template.Template) (err error) {
	var ret bytes.Buffer
	if err = tpl.Execute(&ret, self.Variables); err != nil {
		return
	}
	//self.Variables = make(map[string]interface{})
	self.Content = ret.Bytes()
	return
}

// ResponseJsonContent Выдача браузеру правильной страницы в указанном ранее формате
//func (self *RW) response(data []byte, status int) {
//	// запрет кеширования
//	self.Writer.Header().Set("Cache-Control", "no-cache, must-revalidate")
//	self.Writer.Header().Set("Pragma", "no-cache")
//	self.Writer.Header().Set("Date", lib.Time.Now().Format(time.RFC1123))
//	self.Writer.Header().Set("Last-Modified", lib.Time.Now().Format(time.RFC1123))
//	// размер контента
//	self.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
//	// Статус ответа
//	self.Writer.WriteHeader(status)
//	// Тело документа
//	self.Writer.Write(data)
//	//
//	self.InterruptHard = true
//	logs.Info(106, status, self.Request.URL.Path)
//}

// ResponseCustom Выдача браузеру правильных данных в абстрактном формате
// Возможные типы документов
// application/json, text/html, image/*
//func (self *RW) ResponseCustom(con Content) {
//	var fileExt, contentType, mimeType string
//	var err error

//	// Тип документа
//	contentType = "application/octet-stream"
//	if con.Type != "" {
//		contentType = con.Type
//	} else if con.File != "" {
//		l := strings.Split(con.File, ".")
//		fileExt = `.` + l[len(l)-1]
//		mimeType = mime.TypeByExtension(fileExt)
//		if mimeType != `` {
//			contentType = mimeType
//		}
//	}

//	// Кодировка
//	if con.Encode != "" {
//		contentType += "; charset=" + con.Encode
//	}

//	self.Writer.Header().Set("Content-Type", contentType)

//	// Аттач если документ бинарный файл и не картинка
//	if mimeType != `` && strings.LastIndex(mimeType, `image`) == -1 && strings.LastIndex(mimeType, `text`) == -1 {
//		self.Writer.Header().Set("Content-Disposition", "attachment; filename = "+filepath.Base(con.File))
//	}

//	// Если контент пустой
//	if len(con.Content) == 0 && con.File != `` {
//		if con.Content, err = ioutil.ReadFile(con.File); err != nil {
//			con.Content = []byte(err.Error())
//		}
//	}
//	if con.Status == 0 {
//		con.Status = 200
//	}

//	self.response(con.Content, con.Status)
//}

// ResponseJson Выдача браузеру страницы в формате JSON
//func (self *RW) ResponseJson(data interface{}, status, codeLocal int16, messages ...interface{}) {

//	self.Mode = ModeApi
//	content, err := json.Marshal(data)
//	if err != nil {
//		lg := logs.Error(105, err)
//		content, _ = json.Marshal(lg.Message)
//		self.Content.Content = []byte(`,"value":` + string(content))
//		self.Content.Status = 500
//		return
//	}
//	t := reflect.TypeOf(data)
//	if t.Kind() == reflect.Slice { // срезы
//		self.Content.Content = []byte(`,"slice":` + string(content))
//	} else if t.Kind() == reflect.Map { // хеши
//		self.Content.Content = []byte(`,"map":` + string(content))
//	} else if t.Kind() == reflect.Struct || t.Kind() == reflect.Ptr { // объекты
//		self.Content.Content = []byte(`,` + string(content[1:len(content)-1]))
//	} else if `""` != string(content) { // скалярные значения
//		self.Content.Content = []byte(`,"value":` + string(content))
//	} else { // скалярные значения
//		self.Content.Content = []byte(``)
//	}
//}

// ResponseHtmlError Выдача браузеру правильнойошибочной страницы текстового формата
//func (self *RW) ResponseError(status int) {
//	lg := logs.Error(status)
//	if self.Mode == ModeHtml {
//		self.Content.Type = `text/html`
//		self.Content.Encode = `utf-8`
//		self.Content.Content = nil
//		self.Content.File = Config.Path.Www + tplError[status]
//	} else if self.Mode == ModeApi {
//		self.Content.Type = `application/json`
//		self.Content.Encode = `utf-8`
//		data := new(Response)
//		data.ErrorCode = lg.Code
//		data.ErrorMessage = lg.Message
//		self.Content.Content, _ = json.Marshal(data)
//	}
//	self.Content.Status = status
//	self.Response()
//}

// ResponseHtmlError Выдача браузеру правильнойошибочной страницы текстового формата
//func (self *RW) ResponseError(status int) {
//	// Расположение служебных шаблонов для не стандартных (ошибочных) запросов работы сервера
//	var tplError = map[int]string{
//		403: "/403.html",
//		404: "/404.html",
//		500: "/500.html",
//	}
//	lg := logs.Error(status)
//	if self.Mode == ModeHtml {
//		self.Content.Type = `text/html`
//		self.Content.Encode = `utf-8`
//		self.Content.Content = nil
//		self.Content.File = Config.Path.Www + tplError[status]
//	} else if self.Mode == ModeApi {
//		self.Content.Type = `application/json`
//		self.Content.Encode = `utf-8`
//		data := new(Response)
//		data.ErrorCode = lg.Code
//		data.ErrorMessage = lg.Message
//		self.Content.Content, _ = json.Marshal(data)
//	}
//	self.Content.Status = status
//	self.Response()
//}
