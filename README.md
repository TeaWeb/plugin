## TeaWeb插件接口
### 如何实现一个自己的插件
1. 新建一个项目，项目结构为：
    ~~~
    demo-plugin/
      src/
        github.com/
            TeaWeb/
                plugin/
                    [https://github.com/TeaWeb/plugin源码]
        main/
            demo.go - 你的插件源文件  
            build.sh - 构建脚本                              
    ~~~
2. 在`main/`目录下建一个插件的Go文件，比如命名为`demo.go`；
3. 在`demo.go`中实现
    ~~~go
    package main
    
    import (
        "github.com/TeaWeb/plugin/charts"
        "github.com/TeaWeb/plugin/loader"
        "github.com/TeaWeb/plugin/plugins"
    )
    
    func main() {
        demoPlugin := plugins.NewPlugin()
        demoPlugin.Name = "Demo Plugin"
        demoPlugin.Code = "com.example.demo"
        demoPlugin.Developer = "Liu xiangchao"
        demoPlugin.Version = "1.0.0"
        demoPlugin.Date = "2018-10-15"
        demoPlugin.Site = "https://github.com/TeaWeb/build"
        demoPlugin.Description = "这是一个Demo插件"
        
        loader.Start(demoPlugin)
    }	
    ~~~
4. 可以修改`demoPlugin`，以提供插件的名称、描述等信息，或者实现其他功能；
5. 使用`go build -o demo.tea demo.go`编译插件；
6. 将编译成功后的`demo.tea`放到`TeaWeb`的`plugins/`目录下，重启`TeaWeb`后生效。

### 构建脚本
*build.sh*
~~~bash
#!/usr/bin/env bash

export GOPATH=`pwd`/../../
export CGO_ENABLED=1

# msgpack
if [ ! -d "${GOPATH}/src/github.com/vmihailenco/msgpack" ]
then
    go get "github.com/vmihailenco/msgpack"
fi

# TeaWeb
if [ ! -d "${GOPATH}/src/github.com/TeaWeb/plugin" ]
then
    go get "github.com/TeaWeb/plugin"
fi

go build -o demo.tea demo.go
~~~

### 代码示例
请见 [main/demo.go ](main/demo.go)。

### 安装插件
将`TeaWeb`插件拷贝到`plugins/`目录下，即可完成安装：
~~~
bin/
plugins/
  demo.tea
  ...
~~~

安装后，请重启`TeaWeb`后插件生效，然后在Web界面中的"插件"菜单中可以看到加载成功的插件。
    