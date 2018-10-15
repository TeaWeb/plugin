# `TeaWeb`插件接口
## 如何实现一个自己的插件
1. 建一个插件的Go文件，比如命名为`demo.go`；
2. 在`demo.go`中实现
    ~~~go
    func New() teainterfaces.PluginInterface {
        return &DemoPlugin{}
    }
    
    type DemoPlugin struct {
        plugins.Plugin
    }
    ~~~
3. 可以覆盖`DemoPlguin`中的方法，以提供插件的名称、描述等信息，或者实现其他功能；
4. 使用`go build -o demo.so -buildmode=plugin demo.go`编译插件；
5. 将编译成功后的`demo.so`放到`TeaWeb`的`plugins/`目录下，重启`TeaWeb`后生效。

## 示例
请见 `main/demo.go` 。
    