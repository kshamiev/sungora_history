-- 0
UPDATE `Controllers` SET `Position` =  `Position` - 1 WHERE `Position` >  ?;
-- 1 (0.1)
UPDATE `Controllers` SET `Position` =  `Position` + 1 WHERE `Position` >  ?;
-- 2 (0.2)
UPDATE `Controllers` SET `Position` =  ? WHERE `Id` = ?;
-- 3 удаление связи контроллера и группы
DELETE FROM GroupsUri WHERE Controllers_Id = ? AND Groups_Id = ?;
-- 4 получение максимальной позиции
SELECT MAX(Position) as Max FROM `Controllers`
