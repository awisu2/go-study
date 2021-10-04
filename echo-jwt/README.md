# echo-jwt

[JWT Recipe | Echo - High performance, minimalist Go web framework](https://echo.labstack.com/cookbook/jwt/)

- 問い合わせサンプル(後述)
- このコードでは DB でのログイン処理も実装したが、中核は変わらない
  - token を生成する際の確認に DB にアクセスするか、ローカルデータにアクセスするかの違いのみ
- Header に kid を付与したい場合は、NewWithClaims で生成したインスタンスにセット(`token.Header["kid"] = "abc"`)
- JWTConfig(利用しそうな部分だけ)
  - KeyFunc, SigningKeys, SigningKey: token の Unsign 処理で利用される signKey の設定
    - 運用考えると KeyFunc, SigningKeys のどちらかを利用したほうがいい気がする。記載順に nil チェックされ利用される
  - ContextKey: jwt 認証を通った claim を取得する際のキー名(default: "user")

jwt の token を生成

```go
	// Set custom claims(データ)
	claims := &JwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
```

## 問い合わせサンプル

```bash
$ http -f POST localhost:1323/login username=jon password=shhh!

HTTP/1.1 200 OK
Content-Length: 160
Content-Type: application/json; charset=UTF-8
Date: Thu, 30 Sep 2021 00:26:30 GMT

{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9uIFNub3ciLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjMzMjIwNzkwfQ.2z4QGzSrsnQ0p828jiJ6IgLIOrnfC20pX_W9vym4r3w"
}

$ http localhost:1323/restricted

HTTP/1.1 400 Bad Request
Content-Length: 39
Content-Type: application/json; charset=UTF-8
Date: Thu, 30 Sep 2021 00:27:49 GMT

{
    "message": "missing or malformed jwt"
}

$ http localhost:1323/restricted Authorization:"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9uIFNub3ciLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjMzMjIwNzkwfQ.2z4QGzSrsnQ0p828jiJ6IgLIOrnfC20pX_W9vym4r3w"
HTTP/1.1 200 OK
Content-Length: 17
Content-Type: text/plain; charset=UTF-8
Date: Thu, 30 Sep 2021 00:29:09 GMT

Welcome Jon Snow!
```
