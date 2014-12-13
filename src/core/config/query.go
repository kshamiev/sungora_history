package config

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"core"
	"lib/database/mysql"
)

var queryData = make(map[string][]string)

// compileQuery Прекопиляция sql запросов в текстовых файлах в исходные файлы на GO
func compileQuery() (err error) {
	// Вычисляем путь до проекта
	var pathRoot = core.Config.Main.WorkDir
	var path string

	// MYSQL
	// читаем все запросы
	path = pathRoot + `/src/core`
	if _, err = os.Stat(path); err != nil {
		return nil
	}
	if err = compileQueryMysql(path); err != nil {
		return
	}
	path = pathRoot + `/src/app`
	if _, err = os.Stat(path); err != nil {
		return nil
	}
	if err = compileQueryMysql(path); err != nil {
		return
	}

	// формируем результат
	var sql = "// mysql запросы приложения к БД\npackage mysql\n\n"
	sql += "var Query = make(map[string][]string)\n\n"
	sql += "func init() {\n\tQuery = map[string][]string{\n"
	for i := range queryData {
		sql += "\t\t`" + i + "`: []string{\n"
		for j := range queryData[i] {
			sql += "\t\t\t\"" + queryData[i][j] + "\",\n"
		}
		sql += "\t\t},\n"
	}
	sql += "\t}\n}"

	// сохраняем
	if err = ioutil.WriteFile(pathRoot+`/src/lib/database/mysql/query.go`, []byte(sql), 0777); err != nil {
		return
	}
	mysql.Query = queryData
	queryData = make(map[string][]string)

	// CASANDRA
	// ...

	return
}

// compileQueryMysql
func compileQueryMysql(path string) (err error) {
	var fileInfo, fileInfoSql []os.FileInfo
	if fileInfo, err = ioutil.ReadDir(path); err != nil {
		//fmt.Println("error sql dir read: ", err.Error())
		return
	}

	for i, _ := range fileInfo {
		if fileInfo[i].IsDir() == false {
			continue
		}
		var moduleName = fileInfo[i].Name()
		var pathModule = path + `/` + moduleName + `/query/mysql`
		if fileInfoSql, err = ioutil.ReadDir(pathModule); err != nil {
			//fmt.Println("error sql dir read: ", err.Error())
			err = nil
			continue
		}
		for i, _ := range fileInfoSql {
			// Только sql файлы
			if 0 < strings.LastIndex(fileInfoSql[i].Name(), `.sql`) {
				var data []byte
				if data, err = ioutil.ReadFile(pathModule + `/` + fileInfoSql[i].Name()); err != nil {
					//fmt.Println("error sql file read: ", err.Error())
					return
				}
				var index = strings.ToLower(fileInfoSql[i].Name())
				index = moduleName + `/` + strings.Replace(index, `.sql`, ``, 1)

				queryData[index] = []string{}
				var dataString = string(data)
				// избавляемся от комментариев
				var regdataString = regexp.MustCompile("-- .*")
				l := regdataString.FindStringSubmatch(dataString)
				for len(l) > 0 {
					dataString = strings.Replace(dataString, l[0], "", 1)
					l = regdataString.FindStringSubmatch(dataString)
				}
				// формируем и минимизируем данные в однострочные запросы
				l = strings.Split(dataString, ";")
				for i := range l {
					l[i] = strings.Replace(l[i], "\r", "", -1)
					l[i] = strings.Replace(l[i], "\n", " ", -1)
					l[i] = strings.Replace(l[i], "\t", " ", -1)
					l[i] = strings.Replace(l[i], "  ", " ", -1)
					l[i] = strings.Replace(l[i], "  ", " ", -1)
					l[i] = strings.Replace(l[i], "  ", " ", -1)
					l[i] = strings.Trim(l[i], " ")
					//fmt.Println(l[i])
					if l[i] != "" {
						queryData[index] = append(queryData[index], l[i]+";")
					}
				}
			}
		}
	}
	return
}
