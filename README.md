# CMF sungora
(in language go and framework angularJS)

[www.sungora.ru](http://sungora.ru)

Version: 0.9.1


### Install

Инсталируем язык программирования Го (golang) с официального сайта.

Клонируем репозиторий в папку наших проектов под нужным нам именем.

git clone https://github.com/kshamiev/sungora projectName

Создаем папки `pkg`, `bin` в корневой папке проекта

Устанавливаем переменные окружения:

- Windows

		находясь в корневой папке проекта (репозитория):
		SET PATH=путь до компилятора Го (пример: C:\Go\bin)
		SET GOROOT=путь до корневой папки куда установлен Го (пример: C:\Go)
		Эти переменные обычно прописываются автомтически после установки Го

		SET GOPATH=путь до проекта (пример: C:\Projects\ProjectName)

		SET GOBIN=путь до бинарников проекта (пример: C:\Projects\ProjectName\bin)


- Unix

		находясь в корневой папке проекта (репозитория):
		export PATH=путь до компилятора Го (пример: /home/go/bin)
		export GOROOT=путь до корневой папки куда установлен Го (пример: /home/go)
		Эти переменные обычно прописываются автомтически после установки Го

		export GOPATH=путь до проекта (пример: /home/projects/projectName)
			sample: export GOPATH=/home/domains/www.sungora.ru

		export GOBIN=путь до бинарников проекта (пример: /home/projects/projectName/bin)
			sample: export GOBIN=/home/domains/www.sungora.ru/bin


Настриваем конфигурацию взяв за основу шаблон конфигурационного файла:<br>
`application.conf.sample` -> `bin/application.conf`

Инсталируем сторонние библотеки:

	находясь в корневой папке проекта (репозитория):
	go get bitbucket.org/kardianos/service
	go get github.com/dchest/captcha
	go get github.com/robfig/config
	go get github.com/ziutek/mymysql/mysql
	go get gopkg.in/fsnotify.v1
	go get code.google.com/p/winsvc/winapi

Возможно после инсталяции будет выведено сообщение о невозможности собрать инсталируемую либу (не обращаем внимания).

### Development & Run
Производим необходимкю нам работу над проектом.

IDE:

- Eclipse
- LiteIde

Собираем наш проект

###### install
находясь в папке проекта (репозитория):
`go install -a src/application.go src/modules.go`

или

###### build (этот способ не требует переменной GOBIN)
находясь в папке src проекта (репозитория):
`go build -o ../bin/application[.exe]`

--

Получаем исполняемый бинарный файл. Переходим в папку `bim`

application -h справка по запуску программы (проекта)

Наслаждаемся результатом.

##### Плюшки

Анализ зависмостей пакетов: `go run build/depend/depend.go`<br>
Формирует текстовой файл с информацией по зависимостям пакетов

Инженеринг моделей и типов: `go run build/engine/engine.go`<br>
Формирует типы и модели на основе указанной в программе БД. `src/typesEngine`


### Docs

[Дополнительная документация по рпоекту](http://sungora.ru)

##### Документация кода.
В корневой папке проекта запустить: `godoc -http=:6060 -goroot=путь до папки проекта`

Запустится сервер доступный по ссылке: [http://lolcahost:6060](http://lolcahost:6060)

### Test

Находясь в корневой папке проекта (репозитория):

`(SET или export) GOPATH=путь до проекта`

Системная библиотека логирования:
	
	go test -v lib/logs






