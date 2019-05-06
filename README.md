# README #

这是一个golang版本的fromework-config的实现， 目前只实现了mysql

### 使用目标 ###

目标是为了使用golang的老师便利使用我们现有的配置中心架构

### 亮点 ###

本程序手动实现了Aes128Ecb go版本的算法特性，因为在新的版本中此算法系统的库以及不再支持，如果有小伙伴需要此算法支持，这个就是答案

### 使用方法 ###

由于golang使用最先进的模块依赖策略，整个系统使用的包可以直接依赖网上的版本，但是如果不能上网，需要go get 相关软件

### Contribution guidelines ###

loan.go是主要测试入口，不过整个package为main,后续变更package名词

### 有问题请联系 ###

有问题请联系vincent@6estates.com，或者自行解决
