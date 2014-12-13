-- 0
UPDATE `Uri` SET `Position` =  `Position` - 1 WHERE `Position` >  ?;
-- 1 (0.1)
UPDATE `Uri` SET `Position` =  `Position` + 1 WHERE `Position` >  ?;
-- 2 (0.2)
UPDATE `Uri` SET `Position` =  ? WHERE `Id` = ?;
-- 3 удаление связи ури и контроллера
DELETE FROM GroupsUri WHERE Uri_Id = ? AND Controllers_Id = ?;
-- 4 удаление связи ури и группы
DELETE FROM GroupsUri WHERE Uri_Id = ? AND Groups_Id = ?;
-- 5 получение максимальной позиции
SELECT MAX(Position) as Max FROM `Uri`
