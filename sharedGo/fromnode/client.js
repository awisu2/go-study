var ref = require('ref-napi')
var ffi = require('ffi-napi')
var Struct = require('ref-struct-napi')
var ArrayType = require('ref-array-napi')
var LongArray = ArrayType(ref.types.longlong)

var GoSlice = Struct({
  data: LongArray,
  len: 'longlong',
  cap: 'longlong'
})
var GoString = Struct({
  p: 'string',
  n: 'longlong'
})

// var file = './awesome.dll'
var file = 'D:\\develop\\study\\go-study\\sharedGo\\fromnode\\awesome.dll'
var fs = require('fs')
if (!fs.existsSync(file)) {
  console.log('not exists file. ' + file)
  process.exit(1)
}

// TOOD: Dynamic Linking Error: Win32 error 126 と出て通らない
var awesome = ffi.Library(file, {
  Add: ['longlong', ['longlong', 'longlong']],
  Cosine: ['double', ['double']],
  Sort: ['void', [GoSlice]],
  Log: ['longlong', [GoString]]
})
console.log('awesome.Add(12, 99) = ', awesome.Add(12, 99))
console.log('awesome.Cosine(1) = ', awesome.Cosine(1))
nums = LongArray([12, 54, 0, 423, 9])

var slice = new GoSlice()
slice['data'] = nums
slice['len'] = 5
slice['cap'] = 5
awesome.Sort(slice)
str = new GoString()
str['p'] = 'Hello Node!'
str['n'] = 11
awesome.Log(str)
