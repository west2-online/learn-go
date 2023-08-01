# Golang 第一轮考核

## 目的

Go语言基本语法

- 条件，选择
- 循环
- 键值对
- 切片，集合
- 函数
- 通道 Channel
- Go协程 Goroutine

## 任务

## Task1.基础语法
请使用golang完成下列任务

1. 洛谷P1001：https://www.luogu.com.cn/problem/P1001
2. 洛谷P1046：https://www.luogu.com.cn/problem/P1046
3. 洛谷P5737：https://www.luogu.com.cn/problem/P5737
4. AtCoder ARC017A：https://www.luogu.com.cn/problem/AT_arc017_1
   - 对于这道题，请编写一个判断质数的函数`isPrime(x int) bool` ，并且在主函数中调用它

5. 创建一个**切片(slice)** 使其元素为数字`1-50`，从切⽚删掉数字为`3`的倍数的数，并且在末尾再增加⼀个数`114514`，输出切⽚。

**输出示例**

```go
[1 2 4 5 7 8 10 11 13 14 16 17 19 20 22 23 25 26 28 29 31 32 34 35 37 38 40 41 43 44 46 47 49 50 666]
```

### Bonus

1. 写一个99乘法表，并且把结果保存到同⽬录下ninenine.txt，⽂件保存命名为"6.go"。

2. 回答问题：Go语言中的切片和数组的区别有哪些？答案越详细越好。Go中创建切片有几种方式？创建map
   呢？

3. 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那
   两个 整数，并返回它们的数组下标。

   你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

   你可以按任意顺序返回答案。

   **示例 1：**

   > 输入：nums = [2,7,11,15], target = 9
   > 输出：[0,1]
   > 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 

   **示例2**

   > 输入：nums = [3,2,4], target = 6
   > 输出：[1,2]

* 是否有复杂度`O(n)`的算法？

3. 运行下面代码，在你认为重要的地方写好注释，同时回答下面这些问题
   - 这个代码实现了什么功能？
   - 这个代码利用了golang的什么特性？
   - 这个代码相较于普通写法，是否有性能上的提升？（性能提升：求解速度更快了）


```go
package main

import (
	"fmt"
)

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in chan int, out chan int, prime int) {
	for {
		num := <-in
		if num%prime != 0 {
			out <- num
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)
	for i := 0; i < 6; i++ {
		prime := <-ch 
		fmt.Printf("prime:%d\n", prime)
		out := make(chan int)
		go filter(ch, out, prime)
		ch = out
	}
}
```

## Task2. Git与Github

现在我们来讨论一下Git和Github, 这是一个计算机学生绕不开的话题.不论你计划是升学还是就业,掌握Git和Github对你的帮助都是莫大的.

目前西二在线编写了如下的文档: [Git与Github的超容易入门](https://west2-online.feishu.cn/wiki/Lsz9w3CiGinXzgkevtmceHZknrf),这个文档简单的介绍了如何使用git和github,但是更多的功能仍然需要你自己去探索,同时,这个文档并没有编写完毕

我们希望可以通过这个文档让你快速上手git, 但光看文档肯定是没有用的,你还需要完成下列任务

- 如果你没有自己的Github账号,请创建一个自己的Github账号
- 为你的Github账号添加一个头像
- 为你的Github账号开启2FA(多因素账号登录)
- 为你的Github账号写一个README(请去网上查阅如何美化自己的Github主页)
- 访问这个仓库[[Github-Introduction](https://github.com/west2-online-reserve/Github-Introduction)],在这个仓库中添加一个新的issue,模板选择[Bug Report],按照模板内的要求填写内容(关于bug的复原部分可以随便写,比如click xxx),发布这个issue
- fork上面这个仓库,在fork的仓库中新增一个READDME.md文件,在这个文件中填写你的Github ID
- 将上一步的修改提交到你fork的仓库,并且给主仓库提交一个关于这个修改的pr

**注意:fork仓库并添加文件这个步骤,请在你的电脑本地操作,之后push到你fork的仓库.不要试图直接在Github上面操作,是可以区分出在Github上面操作和本地操作后push的**

### Bonus

1. 请注意你的commit message,你可以自行查找一些git commit规范,我们希望你可以先建立起一定的规范操作
2. 请创建一个仓库,这个仓库存放着你本次考核的代码.仓库名和其他内容不做要求,但是我们希望你的一切操作都看起来是**具有一定规范**的

在未来的工程项目中,规范是非常重要但很难掌握的一个内容,你未来势必不会单打独斗,你会和其他不同的人一起创造、修缮你们的伟大工程.同时,一个工程可能会持续很长时间,可能最初的人因为各种原因被新来的人接替,这时候规范问题更重要了:你该如何使用一定的规范,让新来的人可以很快的上手你们的伟大工程?这是一个值得思考的问题!

请务必注意你的规范问题.以及,请务必学好git和github.

## 要求

1.  不要抄袭 
2.  不要抄袭
3.  不要抄袭
4.  遇到不会的地⽅时，⾸先尝试⾃⼰去解决，可以去百度、⾕歌上寻求帮助，能⽤`搜索引擎`解决**⾃⼰**的问题是⼀项⾮常⾮常重要的能⼒。

## 参考

- 菜鸟教程Go语言 https://www.runoob.com/go/go-tutorial.html
- B站老男孩Go语言入门视频 https://www.bilibili.com/video/BV1fz4y1m7Pm
- 七天入门Go语言 https://blog.csdn.net/weixin_45304503/category_11253946.html
- go语言中文社区 https://learnku.com/go

