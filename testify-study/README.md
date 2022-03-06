# testify study

my stretchr/testify study

- [stretchr/testify: A toolkit with common assertions and mocks that plays nicely with the standard library](https://github.com/stretchr/testify)

NOTE

- testing
  - [assert](https://github.com/stretchr/testify#assert-package): it can change each test to oneline and return ok(bool) it is success or not
  - [require](https://github.com/stretchr/testify#require-package): require like assert but terminate at test is missing.
    - This means that if the test fails, no further tests will be done.
- test support
  - [mock](https://github.com/stretchr/testify#mock-package): Support mock function. It can changebable behavior. details later.
    - note: Why our function use interface?
      - As far as I know, it's good to know architectures such as clean architecture and DDD
  - [suite](https://github.com/stretchr/testify#suite-package): testable struct. exsample write later.
    - what can this, show here. [suite package \- github\.com/stretchr/testify/suite \- pkg\.go\.dev](https://pkg.go.dev/github.com/stretchr/testify/suite)

## assert & require

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ValueObject = struct {
	Value string
}

func TestAssert(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "want 123")
	assert.Equal(t, "hello", "hello", "want hello")

	// assert for nil (good for errors)
	obj := &ValueObject{Value: "Samething"}
	assert.Nil(t, obj)

	// assert for not nil (good when you expect something)
	if assert.NotNil(t, obj) {
		// now we know that object isn't nil
		assert.Equal(t, "Something", obj.Value)
	}
}

func TestRequire(t *testing.T) {
	// require
	require.Equal(t, 123, 234, "want 123")

	// not run because before test missing
	require.Equal(t, 123, 999, "want 123")
}
```

## mock

create mock

## 1: our funtion

```go
package main

import (
	"github.com/stretchr/testify/mock"
)

// our faunction
type Something = interface {
	Hello(name string) (string, error)
}

func Hello(something Something, name string) (string, error) {
	return something.Hello(name)
}
```

- note: Why our function use interface?
  - As far as I know, it's good to know architectures such as clean architecture and DDD


## 2: create mock

```go
// mock (tipycaly write under mock package)
type SomethingMock struct {
	mock.Mock
}

func (m *SomethingMock) Hello(name string) (string, error) {
	args := m.Called(name)
	return args.String(0), args.Error(1)
}
```

## 3: use on test

```go
package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMock(t *testing.T) {
	somethingMock := new(SomethingMock)

	// set function behavior
	name := "World"
	somethingMock.On("Hello", name).Return(fmt.Sprintf("Hello %v", name), nil)

	// run with mock
	//
	// note: The argument must be equivalent to the set value
	res, err := Hello(somethingMock, name)

	// test
	require.Nil(t, err, "not error happen")
	assert.Equal(t, res, "Hello World", "want Hello world")
}
```

## suite

### 1: create suite and run

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
}

func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// normal test. run this when exec `go test`
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
```

### 2: add test

```go
// update: add VariableThatShouldStartAtFive
type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Setup ! it is valuable part of suite
func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// suite test
func (suite *ExampleTestSuite) TestExample() {
	suite.Equal(suite.VariableThatShouldStartAtFive, 5)
}
```

