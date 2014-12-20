package config

import (
	"core"
	"lib"
	"lib/database/mysql"
	"lib/logs"
	"lib/mailer"
	"lib/uploader"
	"os"
	"path/filepath"
)

// initLib Инициализация используемых в приложении библиотек
func library() (err error) {
	// Библиотека lib
	lib.Time.Location = core.Config.Main.TimeLocation

	// Библиотека lib/database/mysql
	mysql.InitMysql(core.Config.Mysql)

	// Служба ведения логов logs
	logs.Init(&core.Config.Logs)

	// Библиотека отправки почты
	mailer.Init(&core.Config.Mail)

	// Модуль загрузки файлов и получение их по идентификатору
	if err = uploader.Init(core.Config.Main.Upload, 30); err != nil {
		return
	}
	return
}

// initFolders Инициализация (создание если их нет) папок приложения
func folders(perm os.FileMode) (err error) {
	if err = os.MkdirAll(core.Config.Main.WorkDir, perm); err != nil {
		return
	}
	if err = os.MkdirAll(filepath.Dir(core.Config.Logs.File), perm); err != nil {
		return
	}
	if err = os.MkdirAll(core.Config.Main.Keys, perm); err != nil {
		return
	}
	if err = os.MkdirAll(core.Config.View.Path, perm); err != nil {
		return
	}
	if err = os.MkdirAll(filepath.Dir(core.Config.Main.Pid), perm); err != nil {
		return
	}
	return
}
