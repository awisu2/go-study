# time

- now: `t := time.Now()`
- add(dec): `t.Add({duration})`
- check
  - t > t2: `t.After(t2)`
  - t < t2: `t.Before(t2)`
  - t = t2: `t.Equal(t2)`
- durations: 加算処理に利用できる
  - `time.Second, time.Hour, time.Microsecond, time.Millisecond, time.Minute, time.Nanosecond`
- convert:
  - str > time : `time.Parse({layout}, {str})`
    - zone がない場合は強制で UTC になる
    - 文字列と pattern が一致しないときエラーになる
  - str > time (with location):`time.ParseInLocation({layout}, {str}, {location})`
    - zone がない場合設定した location になる
  - time > str: `t.Format({layout})`
  - layouts:
    - https://golang.org/src/time/format.go
    - "yyyy-mm-ddThh:mm:dd:ssZ": `time.RFC3339` or `2006-01-02T15:04:05Z07:00`
