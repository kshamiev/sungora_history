LITEIDE (27.2.1)

1)
Добавить в конфигурационный файл сборки для Го следующие строчки
(корневая папка программы/share/liteide/litebuild/gosrc.xml)

		<action id="Install1" menu="Build" img="gray/install.png" key="Ctrl+1" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP1)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install2" menu="Build" img="gray/install.png" key="Ctrl+2" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP2)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install3" menu="Build" img="gray/install.png" key="Ctrl+3" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP3)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install4" menu="Build" img="gray/install.png" key="Ctrl+4" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP4)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install5" menu="Build" img="gray/install.png" key="Ctrl+5" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP5)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install6" menu="Build" img="gray/install.png" key="Ctrl+6" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP6)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install7" menu="Build" img="gray/install.png" key="Ctrl+7" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP7)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install8" menu="Build" img="gray/install.png" key="Ctrl+8" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP8)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install9" menu="Build" img="gray/install.png" key="Ctrl+9" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP9)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>

Для продвинутых можно поиграться с атрибутом args
Собственно это указание параметров компиляции конкретного приложения
"Ctrl+9" Это горячая клавиша по кторой будет собираться определенное приложение
		
2) 
Добавить в конфигурационный файл переменных окружения следующие строчки
(корневая папка программы/share/liteide/liteenv/win64.env)

PATHAPP1=application.go
PATHAPP2=application.go
PATHAPP3=application.go
PATHAPP4=application.go
PATHAPP5=application.go
PATHAPP6=application.go
PATHAPP7=application.go
PATHAPP8=application.go
PATHAPP9=application.go

Первые два пункта нужны для настройки удобной сборки (компиляции) ваших многочисленных приложений проекта по горячим клавишам из любого исходного файла.
"application.go" это собственно имя главного файла с точкой входа в конкретное приложение лежащий в корне исходниколв (src)
Задайте свои имена приложений которые вы реализуете.
Как видите их количество здесь предствалено 9. Более чем достаточно.
Но если нужно больше, можете добавлять конфигурационные строчки сколько угодно.
Главное чтобы они были синхронизированы между 1 и 2 пунктом

3)
Не забудьте раскоментировать переменную GOBIN и прописать путь до папки куда будут собираться бинарники. По умолчанию она закоментирована.
А также установить переменную GOPATH в визуальном менеджере IDE

4) 
Keys.kms
Это моя редакция горячих клавиш для удобства работы
Кому как нравиться...

5)
Все остальные параметры по вкусу
