# 学习Go语言

#### 介绍
本仓库是从极客大学Go进阶训练营讲课中抽取出来的练习题。通过做题，可巩固所学知识。

#### 使用方法
1. 每个目录一个问题，选择要复习的问题，进入相应的目录。
2. 打开不带"_test"的".go"源码文件，阅读注释中的题目要求。
3. 根据要求完成代码编写。
4. 运行"go test"来验证你的解答是否通过测试。看到"PASS"字样，即为成功。
5. 若没有思路，则可直接查看answers目录下的解答。
6. 完成练习后，用git checkout命令重置源文件，方便下次使用。

#### 使用示例
```
$ cd 001_errors
$ vim errors.go
$ go test
PASS
ok      goerrors        0.001s
$ git checkout goerrors.go
```

#### 运行环境
1. Go 1.13
2. 部分题目需要配置代理，可参考：https://goproxy.cn/

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


