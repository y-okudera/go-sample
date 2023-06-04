# go-sample

## go get

```
go get xorm.io/xorm github.com/go-sql-driver/mysql
```

## docker-compose

起動

```
docker-compose up -d
```

停止

```
docker-compose down
```

## APIの動作確認

POST

```
curl -X POST "http://localhost:3000/v1/users" -H "Content-Type: application/json" -d '{"name": "鈴木", "address": "東京都"}'
```

GET

```
curl -X GET "http://localhost:3000/v1/users/1"
```

PUT

```
curl -X PUT -H "Content-Type: application/json" -d '{"name": "鈴木", "address": "神奈川県"}' "http://localhost:3000/v1/users/1"
```

DELETE

```
curl -X DELETE "http://localhost:3000/v1/users/1"
```
