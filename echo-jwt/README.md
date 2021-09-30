# echo-jwt

[JWT Recipe | Echo - High performance, minimalist Go web framework](https://echo.labstack.com/cookbook/jwt/)

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
