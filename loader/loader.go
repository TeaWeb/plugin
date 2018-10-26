package loader

import (
	"encoding/binary"
	"errors"
	"github.com/TeaWeb/plugin/charts"
	"github.com/TeaWeb/plugin/messages"
	"github.com/TeaWeb/plugin/plugins"
	"log"
	"net/http/httputil"
	"os"
	"reflect"
)

type Loader struct {
	plugin *plugins.Plugin

	methods   map[string]reflect.Method
	thisValue reflect.Value

	reader *os.File
	writer *os.File

	debug bool
}

func NewLoader(plugin *plugins.Plugin) *Loader {
	loader := &Loader{
		plugin:  plugin,
		methods: map[string]reflect.Method{},
		reader:  os.NewFile(uintptr(3), "parentReader"),
		writer:  os.NewFile(uintptr(4), "parentWriter"),
	}

	// 当前methods
	t := reflect.TypeOf(loader)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		loader.methods[method.Name] = method
	}

	loader.thisValue = reflect.ValueOf(loader)

	plugin.OnReloadApps(func() {
		loader.Write(&messages.ReloadAppsAction{
			Apps: plugin.Apps,
		})
	})

	return loader
}

func (this *Loader) Debug() {
	this.debug = true
}

func (this *Loader) Load() {
	// 注册插件
	this.Write(&messages.RegisterPluginAction{
		Plugin: this.plugin,
	})

	buf := make([]byte, 1024)
	msgData := []byte{}
	for {
		if this.debug {
			log.Println("[plugin]try to read buf")
		}

		n, err := this.reader.Read(buf)

		if n > 0 {
			msgData = append(msgData, buf[:n]...)

			if this.debug {
				log.Println("[plugin]read msg data:", string(msgData))
			}

			msgLen := uint32(len(msgData))
			h := uint32(24) // header length

			if msgLen > h { // 数据组成方式： id[8] | actionLen[8] | dataLen[8] | action | data[len-8]
				id := binary.BigEndian.Uint32(msgData[:8])
				l1 := binary.BigEndian.Uint32(msgData[8:16])
				l2 := binary.BigEndian.Uint32(msgData[16:24])

				if msgLen >= h+l1+l2 { // 数据已经完整了
					action := string(msgData[h : h+l1])
					valueData := msgData[h+l1 : h+l1+l2]

					msgData = msgData[h+l1+l2:]

					ptr, err := messages.Unmarshal(action, valueData)
					if err != nil {
						log.Println("[plugin]unmarshal message error:", err.Error())
						continue
					}

					err = this.CallAction(ptr, id)
					if err != nil {
						log.Println("[plugin]", err.Error())
						continue
					}
				}
			}
		}

		if err != nil {
			break
		}
	}
}

func (this *Loader) CallAction(ptr interface{}, messageId uint32) error {
	action, ok := ptr.(messages.ActionInterface)
	if !ok {
		return errors.New("ptr should be an action")
	}
	action.SetMessageId(messageId)

	method, found := this.methods["Action"+action.Name()]
	if !found {
		return errors.New("[plugin]handler for '" + action.Name() + "' not found")

	}
	method.Func.Call([]reflect.Value{this.thisValue, reflect.ValueOf(action)})
	return nil
}

func (this *Loader) ActionStart(action *messages.StartAction) {
	this.plugin.Start()
}

func (this *Loader) ActionReload(action *messages.ReloadAction) {
	this.plugin.Reload()

	for _, widget := range this.plugin.Widgets {
		widget.Reload()
	}
}

func (this *Loader) ActionStop(action *messages.StartAction) {
	this.plugin.Stop()
}

func (this *Loader) ActionReloadWidget(action *messages.ReloadWidgetAction) {
	widget := this.plugin.WidgetWithId(action.WidgetId)
	if widget != nil {
		widget.Reload()

		// 检查更新
		for _, c := range widget.Charts {
			chart, ok := c.(charts.ChartInterface)
			if !ok {
				continue
			}
			if chart.IsChanged() {
				this.Write(&messages.ReloadChartAction{
					Chart: chart,
				})
			}
		}
	}
}

// 刷新单个App
func (this *Loader) ActionReloadApp(action *messages.ReloadAppAction) {
	if action.App == nil {
		return
	}
	app := this.plugin.AppWithId(action.App.Id)
	if app != nil {
		err := app.Reload()
		if err != nil {
			log.Println("[plugin]reload app:", err.Error())
		} else {
			this.Write(&messages.ReloadAppAction{
				App: app,
			})
		}
	}
}

// 对Request进行过滤
func (this *Loader) ActionFilterRequest(action *messages.FilterRequestAction) {
	req, err := action.Request()
	if err != nil {
		log.Println("[plugin]", err.Error())
		return
	}

	b := this.plugin.FilterRequest(req)
	data, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Println("[plugin]", err.Error())
		return
	}

	// 修改后的Req
	respAction := &messages.FilterRequestAction{
		Continue: b,
		Data:     data,
	}
	respAction.SetMessageId(action.MessageId())
	this.Write(respAction)
}

// 对Response进行过滤
func (this *Loader) ActionFilterResponse(action *messages.FilterResponseAction) {
	resp, err := action.Response()
	if err != nil {
		log.Println("[plugin]", err.Error())
		return
	}

	b := this.plugin.FilterResponse(resp)
	data, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Println("[plugin]", err.Error())
		return
	}

	// 修改后的Resp
	respAction := &messages.FilterResponseAction{
		Continue: b,
		Data:     data,
	}
	respAction.SetMessageId(action.MessageId())
	this.Write(respAction)
}

func (this *Loader) Write(action messages.ActionInterface) error {
	msg := messages.NewActionMessage(action)
	msg.Id = action.MessageId()
	data, err := msg.Marshal()
	if err != nil {
		return err
	}
	_, err = this.writer.Write(data)
	return err
}
