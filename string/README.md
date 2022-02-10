# strings

- split each word: `strings.Split(input, "")`
- replace(all): `strings.Replace("aaaa", "a", "b", 2)` = bbaa
  - replaceAll:
    - `strings.Replace("aaaa", "a", "b", -1)` = bbbb
    - `strings.ReplaceAll("aaaa", "a", "b")` = bbbb
- substring: `"abc"[1:3]` = bc
- convert:
  - string to code: `"b"[0]` = 98, `"abc"[1]` = 98
  - code to string: `string(98)` = b
  - str > code > str: `string("abc"[1])` = b

## code

### range return code

```go
s := "xaybx"
for i, code := range "abc" {
  fmt.Printf("%d: %c\n", i, code)
  // 0: 97
  // 1: 98
  // 2: 99

  s = strings.ReplaceAll(s, string(code), "=")
}
fmt.Printf("%s\n", s) // x=y=z
```
