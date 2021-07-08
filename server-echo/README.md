# server-echo

- `go run server.go` 時に windows のセキュリティアラートが出る: e.Start にドメインを付与する(e.Start(localhost:1323))
- db: `docker-compose up` (docker で動作せています。基本 default)
  - `PGPASSWORD=postgres psql -h localhost -U postgres -d postgres`

## links

- [Echo - High performance, minimalist Go web framework](https://echo.labstack.com/)
- [Guide | Echo - High performance, minimalist Go web framework](https://echo.labstack.com/guide/)

## getstart

```bash
go mod init myapp
go get github.com/labstack/echo/v4

# after created sever.go
go run server.go
```

## access test

- upload image: `http -f POST :1323/images name=test.png image@cycle.png`
