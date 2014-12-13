package request

type Users struct {
	Users_Id     uint64 // id пользователя
	Email        string // Email
	Login        string // Логин пользователя
	Password     string // Пароль пользователя (SHA256)
	CaptchaValue string
	CaptchaHash  string
	Language     string // Язык интерфейса
}

type Position struct {
	Id       uint64 // id сортируемого элемента
	TargetId uint64 // id элемента после которого встанет сортируемый (0 – если в начало)
}

type DbResult struct {
	Count    int64
	Position int64
	Max      int64
	Min      int64
	Avg      int64
	Id       uint64
	IdTarget uint64
	IdGrid   uint64
	IdParent uint64
	IdChild  uint64
}
