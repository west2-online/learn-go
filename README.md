# west2-online golang考核指南

这里是西二在线工作室go方向的考核指南，旨在为初学者提供一个循序渐进的golang学习路线

## 版权

本项目遵循GPL-3.0 License，转载请注明本项目仓库地址

## 如何开始

请点开`docs`, 里面还有一个介绍~

我们的学习希望是以一个文档引导 + 个人自学的方式, 我们会告诉你应当如何快速的上手这门语言, 但是你想要学好、学懂, 除了我们的引导外你还需要自己的个人提升

从功利的角度出发, 如果你想进west2-online工作室, 你需要在完成基础内容的基础上适当的完成**一定量的Bonus内容**, 我们的答辩会考察你一定的知识储备

## 问题解答

我们罗列了常见的问题解答位于这个仓库的[Discussions](https://github.com/west2-online/learn-go/discussions), 你也可以直接在这个页面的上方找到这个按钮进入

## 概览

### 基础
| 阶段                          | 学习内容                                        | 预期时长     | 是否需要答辩 |
|-----------------------------|---------------------------------------------|----------| ------------- |
| [基础](docs/1-基础语法.md)        | 基本语法，数组，切片，map，chan，Github入门                | 30天（1个月） | x |
| [爬虫](docs/2-爬虫.md)          | http请求/响应，http包，并发                          | 30天（1个月） | x |
| [备忘录](docs/3-备忘录.md)        | 命名/结构规范，基础框架（hertz、gorm）的使用                 | 30天（1个月） | &#10004; |
| [大作业](docs/4-大作品.md)        | 项目结构设计、三层架构、Docker、Redis等的学习与使用，实现一个入门级Demo | 60天（2个月） | &#10004; |
| [简单提升](docs/5(2025)-微服务.md) | 微服务架构、服务注册                                  | 45天（1个半月） | &#10004; |
| [微服务](docs/6-微服务.md)        | 待定                                          | 30天（1个月） | &#10004; |
| [6.824](docs/7-6.824.md)    | 掌握分布式系统设计/MapReduce/Raft算法                  | 45天（1个半月） | &#10004; |
| [合作](docs/8-合作.md)          | 与前端/客户端进行合作开发第一款相对成熟的产品，了解项目的对接、开发、测试       | 60天（2个月） | &#10004; |

预期时长以一名零基础全日制大学生为参考，如果是全身心投入学习语言，或者已经对其他语言有一定的了解，每一阶段所需的时间会比预期时长来的短

更多内容可以访问`docs`文件夹内的`README.md`文件, 会有关于学习建议、答辩内容以及一些更细节的安排

### 进阶

我们希望可以结合先进课程的实验、理念来提升个人使用go的综合能力,同时尽量避免陷入盲目的开发(我们不鼓励当API工程师)

目前我们**正在推进**的有如下课程作业的转换:

1. MIT 6.824 分布式系统 (2023 Spring)
2. MIT 6.031 软件工程

## 时间安排

考虑到学期的期末等因素，以学期为单位，安排学习内容如下

| 时间     | 完成内容                 |
| -------- | ------------------------ |
| 第一学期 | 基础、爬虫、备忘录       |
| 寒假     | 大作业                   |
| 第二学期 | 聊天室、微服务、底层源码 |
| 暑假     | 合作/开始进阶学习                     |

## 考核设计

对于每一轮考核，通常分为如下部分

| 名称  | 解释                                           |
| ----- | ---------------------------------------------- |
| 目的  | 本轮需要学习的内容                             |
| 背景  | （部分阶段有）增加部分趣味性的故事             |
| 任务  | 任务的具体描述                                 |
| Bonus | 在完成任务的基础上实现更加深入的功能/特性      |
| 要求  | （部分阶段有）对任务的具体细节要求             |
| 参考  | 提供的部分参考资料                             |
| 提示  | （部分阶段有）对考核，或者对语言学习的一些提示 |

## 考核目标

我们的目标是快速为初学者构建一套**相对广的知识体系**。也就是说，我们希望按照每一阶段的考核完成的同学可以熟悉当前golang的基础业务开发与基础工程项目能力。



但是很明显，**只有广而不深的知识体系并不能在就业/升学中形成有力竞争**，因此我们在2023级的考核内容中增加了`Bonus`，这部分内容以额外奖励的形式，引导同学们去学习一些更加深入的内容，而这些更加深入的内容将会在未来的读研中给你提供一定的底层思维能力，同时也会在就业中让你的面试更加得心应手。



如果你有意以golang作为你的主力语言，我们强烈建议认真负责的完成每一轮的`全部内容`。而不是为了考核而考核

## 项目结构

```
.
├── LICENSE
├── README.md
├── advance             // 进阶学习
├── etc                 // 推荐阅读的文章/资料
└── docs                // 考核内容
    ├── 0-推荐资料.md
    ├── 1-基础语法.md
    ├── 2-爬虫.md
    ├── 3-备忘录.md
    ├── 4-大作品.md
    ├── 5-简单提升.md
    ├── 6-微服务.md
    ├── 7-6.824.md
    └── 8-合作.md
```
