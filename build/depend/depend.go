package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

var pathExclusion = map[string]bool{
	`github.com`:      true,
	`code.google.com`: true,
	`bitbucket.org`:   true,
}

func main() {
	var err error
	if err = app(); err != nil {
		debug(err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func app() (err error) {
	// Вычисляем путь до проекта
	var path, _ = os.Getwd()
	path += `/src`

	// читаем все пакеты
	var fileInfo []os.FileInfo
	if fileInfo, err = ioutil.ReadDir(path); err != nil {
		return
	}
	for i := range fileInfo {

		if fileInfo[i].IsDir() == true {
			if _, ok := pathExclusion[fileInfo[i].Name()]; ok == true {
				continue
			}
			if err = pkgAnalize(path+`/`+fileInfo[i].Name(), fileInfo[i].Name()); err != nil {
				return
			}
		}
	}

	var pkgList []string
	// обратная зависимость пакетов
	for pkg := range pkgData {
		pkgDataParent[pkg] = make(map[string]bool)
		pkgDataChild[pkg] = make(map[string]bool)
		pkgList = append(pkgList, pkg)
	}
	for pkgParent := range pkgData {
		for pkg := range pkgData[pkgParent] {
			if _, ok := pkgData[pkg]; ok == true {
				pkgDataParent[pkg][pkgParent] = true
				pkgDataChild[pkgParent][pkg] = true
			}
		}
	}
	sort.Strings(pkgList)

	// формируем результат
	var str string
	for _, pkg := range pkgList {
		if len(pkgDataChild[pkg]) == 0 && len(pkgDataParent[pkg]) == 0 {
			continue
		}
		str += pkg + "\n\tchild: "
		for i := range pkgDataChild[pkg] {
			str += i + `, `
		}
		str += "\n\tparent: "
		for i := range pkgDataParent[pkg] {
			str += i + `, `
		}
		str += "\n"
	}

	// сохраняем
	//if err = os.MkdirAll(path+`/sql`, 0777); err != nil {
	//	return
	//}
	if err = ioutil.WriteFile(path+`/../pkgInfo.txt`, []byte(str), 0777); err != nil {
		return
	}
	return
}

var pkgData = make(map[string]map[string]bool)
var pkgDataChild = make(map[string]map[string]bool)
var pkgDataParent = make(map[string]map[string]bool)

// pkgAnalize
func pkgAnalize(path, pkg string) (err error) {
	if _, ok := pkgData[pkg]; ok == true {
		return
	}
	pkgData[pkg] = make(map[string]bool)
	var fileInfo []os.FileInfo
	if fileInfo, err = ioutil.ReadDir(path); err != nil {
		return
	}
	for i, _ := range fileInfo {
		if fileInfo[i].IsDir() == true {
			pkgAnalize(path+`/`+fileInfo[i].Name(), pkg+`/`+fileInfo[i].Name())
		}
		// Только go файлы
		if -1 != strings.LastIndex(fileInfo[i].Name(), `_test.go`) || -1 == strings.LastIndex(fileInfo[i].Name(), `.go`) {
			continue
		}

		var data []byte
		if data, err = ioutil.ReadFile(path + `/` + fileInfo[i].Name()); err != nil {
			return
		}
		var dataString = string(data)
		dataString = strings.Replace(dataString, "\r", "", -1)
		dataString = strings.Replace(dataString, "\n", "", -1)
		dataString = strings.Replace(dataString, "\t", "", -1)

		var regdataString = regexp.MustCompile(`import \(([^)]*)`)
		match := regdataString.FindAllString(dataString, -1)
		for i := range match {
			regdataString1 := regexp.MustCompile(`"([^"]*)"`)
			m1 := regdataString1.FindAllStringSubmatch(match[i], -1)
			for j := range m1 {
				pkgData[pkg][m1[j][1]] = true
			}
			regdataString1 = regexp.MustCompile("`([^`]*)`")
			m1 = regdataString1.FindAllStringSubmatch(match[i], -1)
			for j := range m1 {
				pkgData[pkg][m1[j][1]] = true
			}
		}
		//debug(pkgData)
		//os.Exit(0)
	}
	return
}

// debug
func debug(idl ...interface{}) {
	var buf bytes.Buffer
	var wr io.Writer

	wr = io.MultiWriter(&buf)
	for _, field := range idl {
		fset := token.NewFileSet()
		ast.Fprint(wr, fset, field, ast.NotNilFilter)
	}
	fmt.Print(buf.String())
}
