
<p align="center">
  <img alt="GoReleaser Logo" src="https://avatars2.githubusercontent.com/u/24697112?v=3&s=200" height="140" />
  <h3 align="center">API-Scheduler</h3>
  <p align="center">
    <a href="https://github.com/imyuliz/api-scheduler/releases/latest"><img alt="Codecov" src="https://img.shields.io/github/v/release/imyuliz/api-scheduler.svg?logo=github&style=flat-square"></a>
    <a href="https://codecov.io/gh/imyuliz/api-scheduler"><img alt="Codecov" src="https://img.shields.io/codecov/c/github/imyuliz/api-scheduler?logo=codecov&style=flat-square"></a>
    <a href="https://github.com/imyuliz/api-scheduler/actions?query=workflow%3A%22Lint+Test+Build%22"><img alt="GitHub marketplace" src="https://github.com/imyuliz/api-scheduler/workflows/Lint%20Test%20Build/badge.svg"></a>
    <a href="https://github.com/imyuliz/api-scheduler/actions?query=workflow%3ARelease"><img alt="GitHub release" src="https://github.com/imyuliz/api-scheduler/workflows/Release/badge.svg"></a>
    
  
  </p>
</p>

---


Proposal
---

### Ø Introduction

伴随着业务的发展和低代码的迅速推进，可见的问题随之暴露出来。从业务发展的角度来看，由于业务的复杂度越来越高，导致API的调用链越来越长，依赖的上下游越来越复杂，

* 功能


### Ø Proposal

### Ø Plan of Action

### Ø Will it work?

### Ø Desired outcomes

### Ø Necessary Resources

### Ø Preparations Made

### Ø Conclusion

<h5 align="right">Created in ChengDU rainy night</h5>

### 常见命令

1. 手动测试镜像构建 `docker build -t template-go -f build/docker/Dockerfile .`
2. 发布新版本 `git tag v1.0.0 && git push --tags`
### Git 提交规范

1. 在提交 Commit message 信息时, 应该保证 message 能够简要说明此次变更的性质和作用域。
比如:
```
fix #issue: math/big: fix typo in documentation for Int.Exp
```
2. 当代码变更需要引入第三方包时, 单独单独提交一个commit, 来说明对第三方包的变更, 通常, 会使用: `vendor: add package name` 的方式来提交 commit 信息。 这样有利于 code review。 
3. 提交代码时, 一个功能特性或者bug fix 应该独占一个 commit, 而不能把多个特性揉合成一个 commit 提交代码, 因为这要非常不利于 code review 和 代码追溯。

常用的一些提交术语:
```
feat:       新特性
fix:        bug 被完全修复
update:     issue/bug 只有部分被修复, 还需要提高和改善
docs:       文档变更
vendor:     代码的依赖包整理
style:      改变代码的风格但是未影响代码的功能和意义(比如对 空格, 格式化等的修改)
refactor:   既不修正错误也不添加功能的代码更改
perf:       提高代码的性能.
test:       对单元测试进行添加/删除/修改
build:      构建工具更新
conf:       运行时的配置变更(注: 通常应该配置分离)
cve:        对安全漏洞的修复
chore:      更改构建过程或辅助工具和库，例如生成文档
```
 



###  代码编写规范

0. 使用 goreporter 或者相关系统做代码审查和质量评估 
1. 编程风格遵守vscode代码提示规范,尽可能的避免避免绿色波浪线
2. 非内部方法一定有func注释
3. 代码质量按照Goreport质量等级规范
4. 使用sonar深度分析代码质量
5. 使用火焰图性能调优
6. 熟悉常见的代码性能分析知识
7. 编译时, 把编译环境和`commitid`写入程序中, 并通过日志打印出来,有利于问题排查



### 代码优化习惯

1. 尽可能少的在循环体内部创建切片或者map,或者复杂对象,因为这样会造成性能不必要的损失。
2. 当既可以使用slice又可以使用map的时候,如果不是非要使用map,那么我们可以多使用slice,因为map在在赋值的时候会信息hash判重,这是因为golang的map是又hashmap实现决定的,slice相当于map来说写入块，但是也比map更占用内存,所以,通常两者都可以的时候,取决于是时间换空间,还是空间换时间.




### 相关推荐
1. Go 官方 CodeReview 规范标准 [英文版](https://github.com/golang/go/wiki/CodeReviewComments), [中文版](https://github.com/panchengtao/articles/issues/8)
2. Go 语言源码贡献官方指导文档 [英文版](https://golang.org/doc/contribute.html), [中文版](https://gocn.vip/topics/10185)
3. Uber Go 编码规范 ( 多厂借鉴 ) [英文版](https://github.com/uber-go/guide/blob/master/style.md), [中文版](https://github.com/xxjwxc/uber_go_guide_cn)
