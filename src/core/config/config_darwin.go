package config

import (
	"os"
	"path/filepath"
)

// searchConfigPaths Инициализация списка путей расположения конфигурационного файла
func searchConfigPaths() []string {
	var err error
	var pt, filename, currentFolder string
	var parent, home string
	var ps = string(os.PathSeparator)
	var pattern []string

	pt, err = filepath.Abs(os.Args[0])
	filename = filepath.Base(os.Args[0])
	parent = filepath.Clean(filepath.Dir(pt) + ps + `..`)
	home = os.Getenv(`HOME`)

	// Конфигурационный файл в папке с исполняемым файлом
	pattern = append(pattern, pt+`.conf`)

	// Конфигурационный файл в вышестоящей папке от исполняемого файла
	pattern = append(pattern, parent+ps+filename+`.conf`)

	// Конфигурационный файл в папке conf расположенной в вышестоящей папке от исполняемого файла
	pattern = append(pattern, parent+ps+`conf`+ps+filename+`.conf`)

	// Конфигурационный файл в папке /etc
	pattern = append(pattern, `/etc`+ps+filename+`.conf`)

	// Конфигурационный файл в папке [filename] расположенной в /etc
	pattern = append(pattern, `/etc`+ps+filename+ps+filename+`.conf`)

	// Конфигурационный файл в папке [filename] расположенной в /opt
	pattern = append(pattern, `/opt`+ps+filename+ps+filename+`.conf`)

	// Конфигурационный файл в /usr/local/etc
	pattern = append(pattern, `/usr/local/etc`+ps+filename+`.conf`)

	// Конфигурационный файл в папке [filename] расположенной в /usr/local/etc
	pattern = append(pattern, `/usr/local/etc`+ps+filename+ps+filename+`.conf`)

	// Конфигурационный файл в домашней папке пользователя
	pattern = append(pattern, home+ps+filename+`.conf`)

	// Конфигурационный файл в папке etc расположенной в домашней папке
	pattern = append(pattern, home+ps+`etc`+ps+filename+`.conf`)

	// Конфигурационный файл в папке conf расположенной в домашней папке
	pattern = append(pattern, home+ps+`conf`+ps+filename+`.conf`)

	// Конфигурационный файл в папке [filename] расположенной в домашней папке
	pattern = append(pattern, home+ps+filename+ps+filename+`.conf`)

	// Конфигурационный файл в текущей папке
	currentFolder, err = os.Getwd()
	if err == nil {
		pattern = append(pattern, currentFolder+ps+filename+`.conf`)
	}

	return pattern
}
