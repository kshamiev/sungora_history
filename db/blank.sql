/*
SQLyog Ultimate v11.33 (64 bit)
MySQL - 5.5.25a-log : Database - zegota
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
/*Table structure for table `Content` */

DROP TABLE IF EXISTS `Content`;

CREATE TABLE `Content` (
  `Uri_Id` bigint(20) DEFAULT NULL COMMENT 'Uri',
  `Lang` varchar(10) NOT NULL COMMENT 'Язык',
  `Title` varchar(100) DEFAULT NULL COMMENT 'Заголовок',
  `Keywords` varchar(100) DEFAULT NULL COMMENT 'Ключи',
  `Description` varchar(300) DEFAULT NULL COMMENT 'Описание',
  `Content` longblob COMMENT 'Контент',
  `Block` varchar(50) NOT NULL COMMENT 'Блок',
  KEY `Uri_Id` (`Uri_Id`),
  KEY `Lang` (`Lang`),
  CONSTRAINT `Content_ibfk_1` FOREIGN KEY (`Uri_Id`) REFERENCES `Uri` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Контент';

/*Table structure for table `Controllers` */

DROP TABLE IF EXISTS `Controllers`;

CREATE TABLE `Controllers` (
  `Id` bigint(64) NOT NULL AUTO_INCREMENT COMMENT 'Id',
  `Name` varchar(128) DEFAULT NULL COMMENT 'Человеческое название контроллера',
  `Path` varchar(255) NOT NULL COMMENT 'Контроллер (модуль/контроллер/метод)',
  `IsBefore` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Порядок выполнения контроллеров по умолчанию (до или после)',
  `IsInternal` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Внутренний контроллер',
  `IsDefault` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Контроллер по умолчанию',
  `IsHidden` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Скрытый контроллер',
  `Position` int(11) NOT NULL DEFAULT '1' COMMENT 'Сортировка (приоритет выполнения)',
  `Date` datetime DEFAULT NULL COMMENT 'Дата регистрации контроллера',
  `Domain` varchar(255) DEFAULT NULL COMMENT 'Домен или regexp описывающий домен',
  `Content` text COMMENT 'Контент контроллера (текстовые или бинарные данные)',
  `ContentTime` datetime DEFAULT NULL COMMENT 'Дата и время копии файла контента в файловой системе (синхронизация БД с файловой системой)',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Контроллеры';

/*Table structure for table `Groups` */

DROP TABLE IF EXISTS `Groups`;

CREATE TABLE `Groups` (
  `Id` bigint(64) NOT NULL AUTO_INCREMENT COMMENT 'Id',
  `Name` varchar(50) DEFAULT NULL COMMENT 'Наименование',
  `Description` text COMMENT 'Описание группы',
  `IsDefault` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Группа по умолчанию',
  `Del` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Запись удалена',
  `Hash` varchar(64) DEFAULT NULL COMMENT 'Контрольная сумма для синхронизации (SHA256)',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Группы';

/*Table structure for table `GroupsUri` */

DROP TABLE IF EXISTS `GroupsUri`;

CREATE TABLE `GroupsUri` (
  `Groups_Id` bigint(64) DEFAULT NULL COMMENT 'Группа',
  `Uri_Id` bigint(64) DEFAULT NULL COMMENT 'Роутинг',
  `Controllers_Id` bigint(64) DEFAULT NULL COMMENT 'Контроллер',
  `Get` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Запрос на получение данных',
  `Post` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Запрос на добавление данных',
  `Put` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Запрос на изменение данных',
  `Delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Запрос на удаление данных',
  `Options` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Запрос на получение опций',
  `Disable` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Запрещен запуск контроллера (для группы)',
  KEY `Groups_Id` (`Groups_Id`),
  KEY `Uri_Id` (`Uri_Id`),
  KEY `Controllers_Id` (`Controllers_Id`),
  CONSTRAINT `GroupsUri_ibfk_1` FOREIGN KEY (`Groups_Id`) REFERENCES `Groups` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `GroupsUri_ibfk_2` FOREIGN KEY (`Uri_Id`) REFERENCES `Uri` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `GroupsUri_ibfk_3` FOREIGN KEY (`Controllers_Id`) REFERENCES `Controllers` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Права';

/*Table structure for table `Uri` */

DROP TABLE IF EXISTS `Uri`;

CREATE TABLE `Uri` (
  `Id` bigint(64) NOT NULL AUTO_INCREMENT COMMENT 'Id',
  `Method` set('GET','OPTIONS','HEAD','POST','PUT','PATCH','DELETE','TRACE','LINK','UNLINK','CONNECT','WS') DEFAULT 'GET' COMMENT 'Метод доступа, разрешается перечисление через запятую без пробелов. ALL,POST,GET,DELETE,UPDATE и т.п.',
  `Domain` varchar(255) DEFAULT NULL COMMENT 'Домен или regexp описывающий домен',
  `Uri` varchar(255) NOT NULL DEFAULT '/' COMMENT 'URI без указания домена и протокола до необязательных параметров',
  `Name` varchar(255) DEFAULT NULL COMMENT 'Название раздела',
  `Redirect` varchar(512) DEFAULT NULL COMMENT 'Если не пусто, то содержит адрес безусловной переадресации',
  `Layout` varchar(150) DEFAULT NULL COMMENT 'Макет шаблон',
  `IsAuthorized` tinyint(1) NOT NULL DEFAULT '0' COMMENT '=1 - доступ к разделу разрешен только авторизованным пользователям',
  `IsMenuVisible` tinyint(1) NOT NULL DEFAULT '0' COMMENT '=1 - раздел отображается в стандартном меню',
  `IsDisable` tinyint(1) NOT NULL DEFAULT '0' COMMENT '=1 - Раздел отключен, при попытке доступа выдается 404',
  `Content` mediumblob COMMENT 'Шаблон или контент раздела (текстовые или бинарные данные)',
  `ContentTime` datetime DEFAULT NULL COMMENT 'Дата и время копии файла контента в файловой системе (синхронизация БД с файловой системой)',
  `ContentType` varchar(50) NOT NULL DEFAULT 'text/html' COMMENT 'Mime type контента раздела, может быть переопределён контроллером',
  `ContentEncode` varchar(50) NOT NULL DEFAULT 'utf-8' COMMENT 'Кодировка контента (по умолчанию utf-8) для заголовка',
  `Position` int(11) NOT NULL DEFAULT '1' COMMENT 'Сортировка, приоритет в роутинге',
  `Title` varchar(255) DEFAULT NULL COMMENT 'Заголовок раздела - title',
  `KeyWords` varchar(255) DEFAULT NULL COMMENT 'Ключевые слова - keywords',
  `Description` text COMMENT 'Описание раздела - description',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Uri` (`Uri`,`Domain`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Роутинг';




/*Table structure for table `Users` */

DROP TABLE IF EXISTS `Users`;

CREATE TABLE `Users` (
  `Id` bigint(64) NOT NULL AUTO_INCREMENT COMMENT 'Id',
  `Users_Id` bigint(20) DEFAULT NULL COMMENT 'Пользователь',
  `Login` varchar(128) NOT NULL COMMENT 'Логин пользователя',
  `Password` varchar(64) DEFAULT NULL COMMENT 'Пароль пользователя (SHA256)',
  `Email` varchar(50) NOT NULL COMMENT 'Email',
  `LastName` varchar(255) DEFAULT NULL COMMENT 'Фамилия',
  `Name` varchar(255) DEFAULT NULL COMMENT 'Имя',
  `MiddleName` varchar(255) DEFAULT NULL COMMENT 'Отчество',
  `IsAccess` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Доступ разрешен',
  `IsCondition` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Условия пользователя',
  `IsActivated` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Активированный',
  `DateOnline` datetime DEFAULT NULL COMMENT 'Дата последнего посещения',
  `Date` datetime DEFAULT NULL COMMENT 'Дата регистрации',
  `Del` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Запись удалена',
  `Hash` varchar(64) DEFAULT NULL COMMENT 'Контрольная сумма для синхронизации (SHA256)',
  `Token` varchar(64) DEFAULT NULL COMMENT 'Токен идентификации пользователя',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Login` (`Login`),
  KEY `Users_Id` (`Users_Id`),
  CONSTRAINT `Users_ibfk_1` FOREIGN KEY (`Users_Id`) REFERENCES `Users` (`Id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Пользователи';

/*Table structure for table `UsersGroups` */

DROP TABLE IF EXISTS `UsersGroups`;

CREATE TABLE `UsersGroups` (
  `Users_Id` bigint(64) NOT NULL COMMENT 'Id пользователя',
  `Groups_Id` bigint(64) NOT NULL COMMENT 'Id группы',
  KEY `Users_Id` (`Users_Id`),
  KEY `Groups_Id` (`Groups_Id`),
  CONSTRAINT `UsersGroups_ibfk_1` FOREIGN KEY (`Users_Id`) REFERENCES `Users` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `UsersGroups_ibfk_2` FOREIGN KEY (`Groups_Id`) REFERENCES `Groups` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Связь пользователя с группой';

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
