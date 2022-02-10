# file

go's file controll module study

- [os package \- os \- pkg\.go\.dev](https://pkg.go.dev/os)

**NOTE**

- make directory: `os.Mkdir(name string, perm FileMode) error`
  - return error if duplicate
  - recurcive: `os.MkdirAll(path string, perm FileMode) error`
    - not return error if duplicate
- remove file and directry: `os.Remove(name string) error`
  - return error if not exists
  - recurcive: `os.RemoveAll(path string) error`
    - not return error if not exists
- rename: `os.Rename(oldpath, newpath string) error`
- isexists: below
- save file, read file: below

## code

### isExists

```go
// check file exists
//
// if ignore error hapen it's return false
// because if all error return this function is too deficult to use
//
func IsExists(filePath string) bool {
  if \_, err := os.Stat(filePath); err != nil {
    // if os.IsNotExist(err) {}
    return false
  }
  return true
}
```

### save file

```go
// save file
func Save(data []byte, savePath string) error {
  f, err := os.Create(savePath)
  if err != nil {
    return err
  }
  defer f.Close()

  f.Write(data)
  return nil
}
```

### read file

```go
// read file
func Read(readPath string) ([]byte, error) {
  if !IsExists(readPath) {
    return nil, os.ErrNotExist
  }

  data, err := ioutil.ReadFile(readPath)
  if err != nil {
    return nil, err
  }

  return data, nil
}
```
