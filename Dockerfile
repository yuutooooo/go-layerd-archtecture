FROM golang:1.21-alpine

WORKDIR /app

# 必要なパッケージをインストール
RUN apk update && apk add --no-cache git

# airの特定バージョンをインストール
RUN go install github.com/cosmtrek/air@v1.49.0

# アプリケーションの依存関係をコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

EXPOSE 8080

# airを使用してアプリケーションを実行
CMD ["air", "-c", ".air.toml"] 