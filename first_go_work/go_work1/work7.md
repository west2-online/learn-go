# Bonus2

## 切片与数组的区别

数组是值类型，切片是数组的一个引用，切片是引用类型，数组作为参数传递时是传拷贝，切片是传引用地址；

数组长度不可变，切片可改变长度；

切片本质上是一个结构体，由指针、长度、容量构成；

由数组创建的切片指针指向数组，但扩容后就不指向原数组

## 切片的创建方式

```go
var slice1 []type
var slice2 []type = []type{}
slice3:=[]type{}
```

```go
arr:=[10]int{0,1,2,3,4}
slice4:=arr[0:5]
slice5:=append(slice4,5)
slice6:=make([]int,10)
```

## map的创建方式

```go
var m1 map[keytype]valuetype
var m2 map[keytype]valuetype = map[keytype]valuetyle{}
var m3 map[keytype]valuetype = make(map[keytype]valuetype)
```

