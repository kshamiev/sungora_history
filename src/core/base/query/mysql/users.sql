-- 0 создание связи пользователя и группы
INSERT UsersGroups SET `Users_Id` = ?, `Groups_Id` = ?;

-- 1 удаление связи пользователя и группы
DELETE FROM UsersGroups WHERE Users_Id = ? AND Groups_Id = ?;

-- 2 обновление онлайн статуса пользователя
UPDATE Users SET `DateOnline` = ?, `Token` = ? WHERE Id = ?;
