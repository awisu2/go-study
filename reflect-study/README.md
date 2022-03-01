# reflect study

[reflect package \- reflect \- pkg\.go\.dev](https://pkg.go.dev/reflect)

- Type: [Type](https://pkg.go.dev/reflect#Type)
- Kind: [Kind](https://pkg.go.dev/reflect#Kind)

## samples

```go
v := []int{1, 2, 3}
i := interface{}(v)

typ := reflect.TypeOf(i)
typ.Kind() // reflect.Slice
typ.String() // "[]int"
typ.Name() // ""

// get element type
elmType := typ.Elem()
elmType.Kind() // reflect.Int
```

**struct**

```go
type User struct {
    Name   string
    Length float32
}
v := User{
    Name:   "jon",
    Length: 156.78,
}
i := interface{}(v)
rf := reflect.TypeOf(i)

names := []string{}
types := []reflect.Type{}
for i := 0; i < rf.NumField(); i++ {
    f := rf.Field(i)
    names = append(names, f.Name)
    types = append(types, f.Type)
}
```

**Value**

```go
v := reflect.ValueOf("a")
v.Kind() //reflect.String
v.String() // a
```
