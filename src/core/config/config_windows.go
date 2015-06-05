package config

import (
	"os"
	"path/filepath"
	"strings"
)

// searchConfigPaths Инициализация списка путей расположения конфигурационного файла
func searchConfigPaths() (pathConfig []string) {

	arr := strings.Split(filepath.Base(os.Args[0]), `.`)
	arr[len(arr)-1] = `conf`
	var f = `/` + strings.Join(arr, `.`)
	p, _ := filepath.Abs(os.Args[0])
	p = filepath.Dir(p)

	var path0 = strings.Replace(p, `\`, `/`, -1)
	var path1 = strings.Replace(filepath.Dir(path0), `\`, `/`, -1)
	var path2 = strings.Replace(filepath.Dir(path1), `\`, `/`, -1)
	var path3 = strings.Replace(os.Getenv("SYSTEMROOT"), `\`, `/`, -1)
	var path4 = strings.Replace(os.Getenv("PROGRAMFILES"), `\`, `/`, -1)
	var path5 = strings.Replace(os.Getenv("COMMONPROGRAMFILES"), `\`, `/`, -1)
	var path6 = strings.Replace(os.Getenv("PROGRAMDATA"), `\`, `/`, -1)
	var path7 = strings.Replace(os.Getenv("LOCALAPPDATA"), `\`, `/`, -1)

	// в текщей директории
	pathConfig = append(pathConfig, path0+f)

	// в родительской директории
	pathConfig = append(pathConfig, path1+f)

	// в вышестоящей директории
	pathConfig = append(pathConfig, path2+f)

	// в системной папке ОС
	pathConfig = append(pathConfig, path3+f)

	// в папке прикладных программ
	pathConfig = append(pathConfig, path4+f)
	pathConfig = append(pathConfig, path5+f)
	pathConfig = append(pathConfig, path6+f)

	// в папке пользователя
	pathConfig = append(pathConfig, path7+f)

	return
}
