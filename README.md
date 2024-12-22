Calc - это проект веб-сервер калькулятора.
При помощи этого проекта можно производить математические рассчеты.

Для запуска проекта, необходимо выполнить команду
```
go run cmd/main.go
```

Веб сервер доступен на порте 8080
Ручка для запросов: /api/v1/calculate

На вход ручка ожидает данные в формате json:
```
{
  "expression": "2+2*2"
}
```

***Примеры использования***
Неуспешное выполнение(500):
```
curl --location --request POST 'http://localhost:8080/api/v1/calculate'
```

Успешное выполнение(200):
```
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
В ответе приходит:
```
{"result":6}
```

Входные данные не соответствуют требованиям приложения(422):
```
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "a+2*2"
}'
```
