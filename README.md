## `TeaWeb`插件接口
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
    func New() teainterfaces.PluginInterface {
        return &DemoPlugin{}
    }
    
    type DemoPlugin struct {
        plugins.Plugin
    }
    ~~~
4. 可以覆盖`DemoPlguin`中的方法，以提供插件的名称、描述等信息，或者实现其他功能；
5. 使用`go build -o demo.so -buildmode=plugin demo.go`编译插件；
6. 将编译成功后的`demo.so`放到`TeaWeb`的`plugins/`目录下，重启`TeaWeb`后生效。

### 构建脚本
*build.sh*
~~~
#!/usr/bin/env bash

export GOPATH=`pwd`/../../

go build -o demo.so -buildmode=plugin demo.go
~~~

### 代码示例
请见 `main/demo.go` 。

### 安装插件
`TeaWeb`插件安装在`plugins/`目录下：
~~~
bin/
plugins/
  demo.so
  ...
~~~

安装后，请重启`TeaWeb`后插件生效，然后在Web界面中的"插件"菜单中可以看到加载成功的插件。
    