## Step to Run
- Clone this repository
- Run `go get -d ./...`
- Run `docker-compose up`
- Run `go run main.go`

## Api
All api using header token
```
- Header:
  Authorization: {{token}}
```
### Token
To acces this API, using JWT token. To get token you can use
```
POST http://localhost:3003/token/generate
```

```json
{
    "user_name": "tetsing_user",
    "password": "pasword-test"
}
```

#### respons
```json
{
    "code": 200,
    "status": true,
    "message": "success",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGV0c2luZ191c2VyIiwiaWQiOjMsImV4cCI6MTY4NDA3Njc3NH0.kR0gSfBWYF0cuaXIysqiymiyQx9jodO3H5eHuXIRL5w"
}
```

`data` is JWT token

### Customer 
#### Insert
```
POST http://localhost:3003/customer
```

```json
{
    "name": "testing apa",
    "email": "email@mail.com"
}
```

#### List and Search
```
GET http://localhost:3003/customer?search=nama&page=1
```

#### Detail
```
GET http://localhost:3003/customer/2
```
2 adalah id dari customer

#### Update
```
PUT http://localhost:3003/customer
```

```json
{
    "id": 2,
    "name": "testing nama",
    "email": "email@mail.com"
}
```

#### Delete
``` 
DELETE http://localhost:3003/customer
```

```json
{
    "id": 2
}
```

### Order 
#### Insert
```
POST http://localhost:3003/order
```

```json
{
    "customer_id": 4,
    "product_id": 1,
    "total_order": 20
}
```

#### List and Search
```
GET http://localhost:3003/order?page=1&searchCustomer=name&searchProduct=komix
```

#### Detail
```
GET http://localhost:3003/order/2
```
2 adalah id dari order

#### Update
```
PUT http://localhost:3003/order
```

```json
{
    "id": 2,
    "customer_id": 4,
    "product_id": 1,
    "total_order": 30
}
```

#### Update
```
DELETE http://localhost:3003/order
```

```json
{
    "id": 2
}
```

