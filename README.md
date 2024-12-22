# iCalc

Это мой проект по финальному заданию Спринта №1 YandexLyceum.
Проект носит название iCalc, и представляет из себя сервис подсчёта арифметических выражений.
Лаконичное решение, без добаления логирования, обработки конкретных ошибок и прочего. 
Этот код решает конкретное ТЗ.

## Запуск проекта

1) **Установите Go.**
2) **Установите Git.**
Склонируйте проект с GitHub используя командную строку:
```
git clone https://github.com/DinnerDer/iCalc
```
Перейдите в папку проекта, выполните команду:
```
go mod tidy
```
Для запуска проекта необходимо в окне терминала ввести одну команду: 
```
go run ./main.go
```
Сервис будет доступен по адресу: [http://localhost:8080/api/v1/calculate](http://localhost:8080/api/v1/calculate)

## Ввод

У сервиса 1 endpoint с url-ом /api/v1/calculate. 
Пользователь отправляет на этот url POST-запрос с телом:
```{
"expression": "выражение, которое ввёл пользователь"
}
```
В ответ пользователь получает HTTP-ответ с телом:
```{
"result": "результат выражения"
}
```
и кодом 200, если выражение вычислено успешно, либо HTTP-ответ с телом:
```{
"error": "Expression is not valid"
}
```
и кодом 422, если входные данные не соответствуют требованиям приложения — например, кроме цифр и разрешённых операций пользователь ввёл символ английского алфавита.

Ещё один вариант HTTP-ответа:
```
{
"error": "Internal server error"
}
```
и код 500 в случае какой-либо иной ошибки («Что-то пошло не так»).

**Чтобы получить ответ, нужно отправить запрос через Postman или cURL:**

URL: http://localhost:8080/api/v1/calculate
Метод: POST
Заголовок: Content-Type: application/json
И сам пример, который будет приведен ниже.
Для этого нужен сайт Postman. 

Выберите метод POST, введите URL: http://localhost:8080/api/v1/calculate. 
Перейдите на вкладку Headers, и добавьте новый заголовок Content-Type: application/json. 
Потом перейдите во вкладку Body, выберите raw и выберите JSON данные.

{
"expression": "2 + 2 * 2"
}

В итоге мы получим следующее:

{"result":"6.000000"}

Если мы попробуем вставить:

{
"expression": "1 + "
}

То получим следующее:

{
"error": "Internal server error"
}

То есть ошибку 500: 
Internal Server Error

Или например если вставить вот такое выражение:

{
"expression": "1 / 0"
}

То выдаст ошибку 500: 
Internal server error

## Эндпоинты

***POST /api/v1/calculate***

Эндпоинт принимает JSON с математическим выражением.

Пример запроса с использованием curl
Пример для cmd:
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\\"expression\": \\"1\\"}"
```
Пример корректного запроса, код:200

git bash
```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
Пример запроса с пустым выражением, код: 422, ошибка:empty expression
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\\"expression\": \\"\\"}"
```
Пример запроса с делением на 0, код: 422, ошибка:division by zero
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\\"expression\": \\"1/0\\"}"
```
Пример запроса с неверным выражением, код: 422, ошибка:invalid expression
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\\"expression\\": \\"1++*2\\"}"
```
Для запросов можно использовать программу postman

###КОНЕЦ README.md
