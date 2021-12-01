# UserBalanceService

## Запуск приложения
```
make run
```

## Описание API и примеры использования

#### Пополнение баланса
```
curl --location --request POST 'localhost:8000/balance-service/refill' 
--header 'Content-Type: application/json' 
--data-raw '{
"id": 3,
"value": 1000
}'
```
#### Снятие с баланса
```
curl --location --request POST 'localhost:8000/balance-service/debit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 3,
    "value": 100
}'
```
#### Получение баланса
```
curl --location --request GET 'localhost:8000/balance-service/balance/1' \
--data-raw ''
```
#### Перевод
```
curl --location --request POST 'localhost:8000/balance-service/transfer' \
--header 'Content-Type: application/json' \
--data-raw '{
"sender_id": 1,
"receiver_id": 2,
"value": 500
}'
```