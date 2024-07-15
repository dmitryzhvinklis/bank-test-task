# bank-test-task

Этот проект представляет собой REST API, который имитирует работу банкомата. Приложение поддерживает операции по созданию аккаунтов, пополнению баланса, снятию средств и проверке баланса. Все операции логируются в консоль и обрабатываются с использованием горутин для обеспечения асинхронности.

## Функциональные возможности

- Создание нового аккаунта
- Пополнение баланса
- Снятие средств
- Проверка баланса

## Запуск

Запустите сервер:
```sh
go run main.go
```

## Создание нового аккаунта

- Endpoint: 
- POST /accounts
- Пример запроса:

```sh
curl -X POST http://localhost:8080/accounts
```
- Ответ:
```json
{
    "ID": 1,
    "Balance": 0
}
```
## Пополнение баланса
- Endpoint: 
- POST /accounts/{id}/deposit
- Пример запроса:

```sh
curl -X POST -d "amount=100.0" http://localhost:8080/accounts/1/deposit
```
- Ответ:
```json
{
    "status": "success"
}
```
## Снятие средств
- Endpoint: 
- POST /accounts/{id}/withdraw
- Пример запроса:
```sh
curl -X POST -d "amount=50.0" http://localhost:8080/accounts/1/withdraw
```
- Ответ:

```json
{
    "status": "success"
}
```
## Проверка баланса

- Endpoint:

- GET /accounts/{id}/balance

- Пример запроса:

```sh
curl -X GET http://localhost:8080/accounts/1/balance
```
- Ответ:
```json
{
    "balance": 50.0
}
```
## Логирование

Все операции логируются в консоль с указанием времени выполнения операции и идентификатора аккаунта. 

- Пример логов:

```
INFO[0000] Created new account with ID 1
INFO[0001] Deposited 100.0 to account 1
INFO[0002] Withdrew 50.0 from account 1
INFO[0003] Checked balance of account 1: 50.0
```