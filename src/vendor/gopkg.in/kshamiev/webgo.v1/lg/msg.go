package lg

var messages = map[int]string{
	100: `Установка сессионной куки: [%s] = [%s]`,
	101: `Установка постоянной куки: [%s] = [%s]`,
	102: `Шаблон статуса ошибки [%d] не найден: %v`,
	103: `Неудалось прочитать JSON данные запроса [%s] [%s]`,
	104: `Ошибка разбора JSON: [%s] : %v`,
	105: `Ошибка формирования JSON: %v`,
	106: `+++ Status [%d] [%s]`,
	107: `+++ Выполнение редиректа на URL: [%s]`,
	108: `-- Некорректный формат свойства для БД: [%s].[%s]`,
	109: `Ошибка выполнения запроса [%s] : %v`,
	110: `Ошибка компиляции запроса [%s] : %v`,
	111: `Невозможно соединиться с БД Mysql конфиг № [%d] : %v`,
	112: `Отсутсвует конфигурация Mysql № [%d]`,
	113: `Ключевое поле (свойство) [%s] отсутствует в структуре [%s]`,
	114: `Ошибка выполнения пакетного запроса (QueryByte)`,
	115: `Нет прав на метод запроса [%s] [%s]`,
	116: `Нет прав на метод запроса [%s] [%s], переход на авторизацию`,
	117: `Запоминаем куда пользователю вернуться [%s]`,
	118: `Ошибка удаления временного файла контента по ключу [%s] : %v`,
	119: `Ошибка общей проверки для сценария [%s] %v`,
	120: `Ошибка копирования типа в модели: %v`,
	121: `Сценарий [%s] для источника [%s] не найден`,
	123: `Найден ранее загруженный файл контента [%s]`,
	124: `Объекты типа [%s] отсутсвуют в памяти`,
	125: `Объекты типа [%s] должны храниться в срезе`,
	126: `Запуск контроллера [%d] [%s]`,
	127: `Ошибка контроллера [%d] [%s]`,
	128: `Ошибка выполнения (парсинг) шаблона [%s] %v`,
	129: `Статика: [%s]`,
	130: `Метод [%s] не поодерживается Uri [%s]`,
	131: `Найден URI: [id:%d]:[%s]`,
	132: `Токен [%s]`,
	134: `Пользователь определен как [%s] : [%d]`,
	135: `Ошибка разбора сегмента Uri [%s][%s] %v`,
	136: `Ошибка разбора сегмента Uri 'relationid' [%s] %v`,
	137: `Не указан шаблон почтового сообщения`,
	138: `Отсутствует файл шаблона [%s] %v`,
	139: `Сегмент uri отсутствует [%s][%s]`,
	143: `Завершение работы приложения`,
	144: `Ошибка получения запроса по индексу [%s]`,
	145: `Ошибка обновления контента контроллера [%s] %v`,
	146: `Ошибка обновления контента uri [%s] %v`,
	147: `Ошибка обновления дефолтового контент-шаблона uri [%s] [%s]`,
	148: `Ошибка открытия порта для сервера [%s]`,
	149: `Сервер запущен по адресу [%s]`,
	150: `Сервер остановлен по адресу [%s]`,
	151: `Сервер [%s] не был запущен`,
	152: `Объект отдающий данные должен быть передан по ссылке: [%s]`,
	153: `Объект принимающий данные должен быть передан по ссылке 'CopyTyp': [%s]`,
	154: `Объект отдающий данные не инициализирован [%s]`,
	155: `Объект принимающий данные не инициализирован [%s]`,
	156: `Объект не найден в БД [%s]`,
	157: `Объект принимающий данные должен быть передан по ссылке 'mysql.LoadArray': [%s] [%s]`,
	158: `Неопределенный тип хранения данных 'mysql.LoadData' [%s]`,
	159: `Ошибка загрузки. Объект не найден в памяти. [%s] [%d]`,
	160: `Сценарий не найден: [%s] -> [%s]`,
	161: `Id пользователя разработчик инициализировано не верно`,
	162: `Email [%s] уже занят`,
	163: `Логин [%s] занят`,
	164: `Пароли [%s] != [%s] не совпадают`,
	165: `Id пользователя гость инициализировано не верно`,
	166: `Id группы разработчик инициализировано не верно`,
	167: `Id группы гость инициализировано не верно`,
	168: `Использование БД отключено либо не реализовано [core.Config.Main.TypeDb = %d]`,
	169: `Объект принимающий данные должен быть срезом 'mysql.LoadArray': [%s] [%s]`,
	170: `Ошибочное свойство для загрузки из БД 'mysql.LoadArray': [%s] [%s]`,
	171: `Ошибочное поле в запросе для загрузки из БД 'mysql.LoadArray': [%s] [%s]`,
	172: `Неверный путь до контроллера: [%s]`,
	173: `Контроллер [%s/%s] отсутсвует`,
	174: `Контроллер [%s/%s] не имеет метода [%s]`,
	175: `Удаление куки: [%s]`,
	176: `Объект принимающий данные должен быть передан по ссылке 'mysql.Load': [%s] [%s]`,
	177: `Ошибочное свойство для загрузки из БД 'mysql.Load': [%s] [%s]`,
	178: `Ошибочное поле в запросе для загрузки из БД 'mysql.Load': [%s] [%s]`,
	179: `Пароль изменен для [%s]`,
	180: `Объект принимающий данные должен быть инициализирован 'mysql.Load': [%s] [%s]`,
	181: `Ошибочное свойство для загрузки из БД 'mysql.LoadData': [%s]`,
	182: ``,
	183: ``,
	301: `Переадресация сюда [%s]`,
	404: `Документ не найден`,
	403: `Доступ запрещен`,
	500: `Ошибка сервера`,
}
