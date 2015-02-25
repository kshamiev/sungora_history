package rw

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Скачивание методом GET больших данных
//	+ url string http ссылка на скачиваемый файл
//	+ pathFile string полный путь до файла в который будут записаны скаченные данные
//	- error ошибка операции
func (self *RW) DownloadLongData(url, pathFile string) (err error) {
	err = errors.New("Неизвестная ошибка")
	var timeout time.Duration = time.Hour * 10

	// Подготовка клиента
	var respReq *http.Response
	client := http.Client{
		Transport: &http.Transport{
			Dial:              timeoutDialler(timeout),
			DisableKeepAlives: true,
		},
	}
	respReq, err = client.Get(url)
	if err != nil {
		// errors.New(fmt.Sprintf("Ошибка получения файла %s: %s", url, err.Error()))
		return
	}
	defer respReq.Body.Close()

	// Создание временного файла
	var file *os.File
	file, err = os.Create(pathFile)
	if err != nil {
		//errors.New(fmt.Sprintf("Ошибка создания временного файла %s: %s", pathFile, err.Error()))
		return
	} else {
		defer file.Close()
	}

	// Копирование
	_, err = io.Copy(file, respReq.Body)
	if err != nil {
		os.Remove(pathFile)
		pathFile = ""
		//errors.New(fmt.Sprintf("Ошибка загрузки файла %s: %s", url, err.Error()))
		return
	}

	//logs.Dumper(fmt.Sprintf("Загружен файл %s размером %v во временную папку %s", url, nbytes, pathFile))

	return
}

// Контроль таймаута загрузки данных
func timeoutDialler(timeout time.Duration) func(net, addr string) (client net.Conn, err error) {
	return func(netw, addr string) (client net.Conn, err error) {
		client, err = net.DialTimeout(netw, addr, time.Duration(timeout))
		if err != nil {
			//errors.New(fmt.Sprintf("Ошибка таймаута при подключении к %s: %s", addr, err.Error()))
			return
		}
		client.SetDeadline(time.Now().Add(timeout))
		return
	}
}

////

// Запрос к сторннему ресурсу в формате json
// Для реализации запрос к сторонним API
//	+ method string Метод запроса в парадигме REST
//	+ url string Полная ссылка запроса
//	+ requestData ...[]byte Данные запроса (POST)
//	- []byte Данные полученные в ответ на запрос
//	- error Ошибка операции
func (self *RW) RequestJson(method, url string, requestData ...[]byte) (responseData []byte, err error) {
	return request(method, url, "application/json; charset=utf-8", requestData...)
}

// Низкоуровневой запрос к сторннему ресурсу
func request(method, url, typ string, requestData ...[]byte) (responseData []byte, err error) {
	var request *http.Request
	var response *http.Response
	// Make request
	if len(requestData) == 0 {
		requestData = append(requestData, []byte(``))
	}
	body := bytes.NewBuffer(requestData[0])
	request, err = http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	// Set user agent as web browser
	request.Header.Set("User-Agent", `Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.72 Safari/537.36`)
	// тип документа
	request.Header.Set("Content-Type", typ)
	// размер контента
	request.Header.Set("Content-Length", strconv.Itoa(len(requestData[0])))
	// получаем результат
	client := http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return
	}
	if response.StatusCode != 200 {
		err = errors.New(strconv.Itoa(response.StatusCode))
		return
	}
	responseData, err = ioutil.ReadAll(response.Body)
	response.Body.Close()
	return
}
