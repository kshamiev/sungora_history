package response

type Captcha struct {
	CaptchaImage string // Капча - бинарные данные в base64
	CaptchaHash  string // Хеш капчи для ее проверки
}
