# [Template] ai-interview-golang
ai-interviewのバックエンド
---

## 1. 動かすために必要なこと

### 1.1 .envファイルの作成
```shell
cp .env.sample .env
```
envファイルの中身を特定のものに書き換えてください

---

## 2. 起動・停止方法
### 2.1 起動方法
imageの作成
```shell
docker compose build
```

imageからコンテナの起動
```shell
docker compose up -d
```

---

### 2.2 停止方法
```shell
docker compose down
```

---

## 3. 必要最低限のエンドポイント

### <span style="color:green">GET</span> 動画をS3にアップロードする
```text
http://localhost:8080/
```

#### file: (binary)
| FormData Keys | FormData Value |
|---------------|----------------|
| file          | (binaryデータ)    |

