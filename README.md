# go-sample

## setup

### コンテナを起動状態で作成

8000でバインディング

```
docker run --name hello -d -p 8000:8080 go-sample-image
```

### コンテナの一覧を表示

停止中のコンテナも含めて表示

```
docker ps -a
```

### コンテナの停止

<CONTAINER ID>は、`docker ps -a` で確認可能

```
docker stop <CONTAINER ID>
```

### APIの動作確認

```
curl http://localhost:8000/hello
```
