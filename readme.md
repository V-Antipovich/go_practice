# Практика в НИУ ВШЭ - Мини-курс "Введение в Go"
- Задачки на 0.5 баллов в папке "prep_taks", по уровню соответственно (решено всё из блоков 1 и 3)
- ДЗ2 - Банковское приложение (до вторника) - в fiber_app/
- ДЗ3 - То же приложение с grpc - в grpc_app/
- ДЗ4 - Приложение с grpc и БД Postgresql - в psql/
База данных postgres запускалась командой
```
docker run --name goprac_psql  -p 5432:5432 -e POSTGRES_USER=vitalii -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=godb -d postgres
```
(эти же параметры передаются в connectString в серверной части приложения)
База состоит из одной таблицы:
```
CREATE TABLE accounts(
    name VARCHAR(256) PRIMARY KEY
    amount INT 
)
```