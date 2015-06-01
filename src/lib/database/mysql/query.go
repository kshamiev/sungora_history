package mysql

// Запросы к БД по индексам
var Query = make(map[string][]string)

// Инициализация базы запросов к БД
func init() {
	Query = map[string][]string{
		`base/users`: []string{
			"INSERT UsersGroups SET `Users_Id` = ?, `Groups_Id` = ?;",
			"DELETE FROM UsersGroups WHERE Users_Id = ? AND Groups_Id = ?;",
			"UPDATE Users SET `DateOnline` = ?, `Token` = ? WHERE Id = ?;",
		},
		`base/controllers`: []string{
			"UPDATE `Controllers` SET `Position` = `Position` - 1 WHERE `Position` > ?;",
			"UPDATE `Controllers` SET `Position` = `Position` + 1 WHERE `Position` > ?;",
			"UPDATE `Controllers` SET `Position` = ? WHERE `Id` = ?;",
			"DELETE FROM GroupsUri WHERE Controllers_Id = ? AND Groups_Id = ?;",
			"SELECT MAX(Position) as Max FROM `Controllers`;",
		},
		`base/uri`: []string{
			"UPDATE `Uri` SET `Position` = `Position` - 1 WHERE `Position` > ?;",
			"UPDATE `Uri` SET `Position` = `Position` + 1 WHERE `Position` > ?;",
			"UPDATE `Uri` SET `Position` = ? WHERE `Id` = ?;",
			"DELETE FROM GroupsUri WHERE Uri_Id = ? AND Controllers_Id = ?;",
			"DELETE FROM GroupsUri WHERE Uri_Id = ? AND Groups_Id = ?;",
			"SELECT MAX(Position) as Max FROM `Uri`;",
		},
	}
}