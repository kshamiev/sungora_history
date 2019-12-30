### Работа с БД

Для работы с базой данных мы используем библиотеку: [sqlboiler](https://github.com/volatiletech/sqlboiler)

#### БД Миграция 
Для миграций мы используем библиотеку: [migrate](https://github.com/webnice/migrate)

* миграция пишется в обе стороны и строго на локальной БД.
* вспомогательный инструментарий в "Makefile.Sample"

**Важно:**

При разработке и выкладыванию своей работы в репу.
Ваша миграция должна быть самой последней. И стоять после всех полученных миграций от дургих разработчиков!
Это достигается путем перебивания на новую миграцию (то есть актуализации Вашей миграции по дате)

**дополнительно для справки:**
создание дампа БД
pg_dumpall -c -U postgres --database=dbName > data/dump_import.sql
pg_dump --file "/home/.../work/temp/back.sql" --host "localhost" --port "5433" --username "postgres" --no-password --verbose --format=p "dev"

востановление БД
createdb -U postgres $(DBNAME)
/usr/bin/psql -h "localhost" -p "5432" -U postgres -w -d $(DBNAME) -f "data/dump.sql"
