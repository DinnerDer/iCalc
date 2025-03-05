# iCalc

Это мой проект по заданиям курса YandexLyceum.
Проект носит название iCalc, и представляет из себя сервис подсчёта арифметических выражений.
Лаконичное решение, без добаления логирования, обработки конкретных ошибок и прочего. 
Этот код решает конкретное ТЗ.

## Установка

 - Для установки нужно выбрать директорию проекта:
```bash
cd <your_dir>
```
 - Потом необходимо выполнить эту команду:
```bash
git clone https://github.com/romanSPB15/Calculator_Service
```
 - В выбранной папке появится папка ```Calculator_Service``` c проектом.

## Работа с API

### Конфигурация
#### Переменные среды
Сначала необходимо открыть файл ```./config/.env``` и установить параметры:

 - **TIME_ADDITION_MS** - время вычисления сложения(в миллисекундах);

 - **TIME_SUBTRACTION_MS** - время вычисления вычитания;

 - **TIME_MULTIPLICATIONS_MS** - время вычисления умножения;

 - **TIME_DIVISIONS_MS** - время вычисления деления;

 - **COMPUTING_POWER** - максмальное количество *worker*'ов, которые параллельно выполняют арифметические действия.

#### Другие параметры

Потом необходимо открыть файл ```config.json``` в той же папке и установить следущие параметры(**true** - включено, **false** - выключено):

 - ```debug``` - отладка(вывод событий в лог)

 - ```web``` - веб-интерфейс(об его использовании читайте дальше в **Веб-интерфейс**)

По умолчанию и то и другое выключено.

### Запуск
 - Для запуска API необходимо выбрать директорию проекта:
```
cd <путь к папке Calculator_Service>
```
 - Далее надо запустить файл ```cmd/main.go```:
```
go run cmd/main.go
```

### Управление

Для работы через терминал на *Curl* настоятельно рекомендую использовать **Git Bash**.<br> Но можно использовать и **Postman** - приложение для отправки *HTTP-запросов*, по-моему в нём работать проще.

#### Добавление вычисления арифметического выражения

Добавление выражения для вычисления на **API**.

##### Curl
```
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": <выражение>
}'
```
##### Postman
 - **URL** localhost/api/v1/calculate;
 - Запрос **POST**;
 - **Body** **RAW** {"expression": <выражение>};
 - Нажать на ****SEND****.

##### Коды ответа: 
 - 201 - выражение принято для вычисления
 - 422 - невалидные данные
 - 500 - что-то пошло не так

##### Тело ответа

```
{
    "id": <уникальный идентификатор выражения> // его ID
}
```
#### Получение списка выражений
##### Curl
```
curl --location 'localhost/api/v1/expressions'
```
##### Postman
 - **URL** localhost/api/v1/expressions;
 - Запрос **GET**;
 - **Body** **NONE**;
 - Нажать на ****SEND****.

##### Тело ответа:

Получение всех сохранённых выражений(**ID** не нужен).

```
{
    "expressions": [
        {
            "id": 8251431,
            "status": "OK",
            "result": 3>
        },
        {
            "id": 34942763,
            "status": "Wait",
            "result": 0
        }
    ]
}
```
##### Коды ответа:
 - 200 - успешно получен список выражений
 - 500 - что-то пошло не так

#### Получение выражения по его идентификатору

Получение выражения по его идентификатору.

*Примечание:* Для того, чтобы получить выражение по его ID, необходимо сохранить полученный при **Добавление вычисления арифметического выражения** индитефикатор.

##### Curl

```
curl --location 'localhost/api/v1/expressions/<id выражения>'
```

##### Postman:
 - **URL** localhost/api/v1/expressions/<id выражения>;
 - Запрос **GET**;
 - Тело **NONE**;
 - Нажать на ****SEND****.

##### Тело ответа:

```
{
    "expression":
        {
            "id": <идентификатор выражения>,
            "status": <статус вычисления выражения>,
            "result": <результат выражения>
        }
}
```

##### Коды ответа:
 - 200 - успешно получено выражение
 - 404 - нет такого выражения
 - 500 - что-то пошло не так

## Примеры работы с API

### Простой пример

#### Делаем запрос на вычисление выражения

##### Curl
```
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2/2"
}'
```
##### Postman
 - **URL** localhost/api/v1/calculate;
 - Запрос **POST**;
 - **Body** **RAW** {"expression": "2+2/2"};
 - Нажать на ****SEND****.

##### Ответ
Статус 201(успешно создано);
```
{
    "id": 12345 // пример
}
```

#### Получаем наше выражение
##### Curl
```
curl --location 'localhost/api/v1/expressions/12345' // 12345 - это ID выше.
```
##### Postman
 - **URL** localhost/api/v1/expressions/12345;
 - Запрос **GET**;
 - Тело **NONE**;
 - Нажать на ****SEND****

##### Ответ
Статус 200(успешно получено);
```
{
    "expression":
        {
            "id": 12345,
            "status": "OK",
            "result": 321
        }
}
```

#### Получаем все выражения
##### Curl
```
curl --location 'localhost/api/v1/expressions'
```
##### Postman
 - **URL** localhost/api/v1/expressions;
 - Запрос **GET**;
 - **Body** **NONE**;
 - Нажать на ****SEND****.

##### Ответ
Статус 200(успешно получены);
```
{
    "expressions": [
        {
            "id": 12345,
            "status": "OK",
            "result": 321
        },
    ]
}
```

### Пример с ошибкой в запросе №1

#### Делаем **неправильный** запрос на вычисление выражения

##### Curl

```
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "radhgsags": "2+2/2"
}'
```

##### Postman
 - **URL** localhost/api/v1/calculate;
 - Запрос **POST**;
 - **Body** **RAW** {"radhgsags": "2+2/2"};
 - Нажать на ****SEND****.

##### Ответ
Статус 422(**неправильный** запрос);


### Пример с ошибкой в запросе №2

#### Делаем **правильный** запрос на вычисление выражения

##### Curl
```
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2/2"
}'
```


##### Postman
 - **URL** localhost/api/v1/calculate;
 - Запрос **POST**;
 - **Body** **RAW** {"radhgsags": "2+2/2"};
 - Нажать на ****SEND****.

##### Ответ
Статус 201(успешно создано);
```
{
    "id": 12345 // пример
}
```
#### Далее получаем наше выражение(**неправильный** ID)
##### Curl
```
curl --location 'localhost/api/v1/expressions/45362'
```

##### Postman:
 - **URL** localhost/api/v1/expressions/45362;
 - Запрос **GET**;
 - Тело **NONE**;
 - Нажать на ****SEND****.

##### Ответ
Статус 404(не найдено);


### Пример с ошибкой в запросе №3

#### Делаем запрос с **некорректным** URL на вычисление выражения

##### Curl
```
curl --location 'localhost/api/v1/abc' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "121+2"
}'
```
##### Postman
 - **URL** localhost/api/v1/abc;
 - Запрос **POST**;
 - **Body** **RAW** {"expression": "121+2"};
 - Нажать на ****SEND****.

##### Ответ
Статус 404(**NOT FOUND**);


### Веб-интерфейс

Вот ссылки на веб-страницы:

 - [Главная страница](http://localhost:8080/api/v1/web)

 - [Вычисление выражения](http://localhost:8080/api/v1/web/calculate)

 - [Просмотр выражений](http://localhost:8080/api/v1/web/expressions)

****ВАЖНО:**** По умолчанию веб-интерфейс выключен. Чтобы его включить, нужно изменить параметр *Веб интерфейс* в **Конфигурация/Другие Параметры**.
### КОНЕЦ README.md
