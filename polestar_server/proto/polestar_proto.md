
# 下载protoc编译器和go插件(放到$GOPATH/bin目录)

* windows

https://github.com/civet148/protoc-plugins/tree/master/windows-x64

* linux

https://github.com/civet148/protoc-plugins/tree/master/linux-amd64

# 创建gogo目录并克隆相关代码

```shell script
$ mkdir -p $GOPATH/src/github.com/gogo
$ cd $GOPATH/src/github.com/gogo
$ git clone https://github.com/gogo/protobuf.git
```