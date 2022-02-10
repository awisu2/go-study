# strings

- split each word: `strings.Split(input, "")`
- substring: `"abc"[1:3]` = bc
- convert:
  - string to code: `"b"[0]` = 98, `"abc"[1]` = 98
  - code to string: `string(98)` = b
  - str > code > str: `string("abc"[1])` = b

## code

### range return code

```go
for i, code := range "abc" {
  fmt.Printf("%d: %v\n", i, code)
}
// 0: 97
// 1: 98
// 2: 99
```
