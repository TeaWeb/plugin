package loader

import (
	"encoding/binary"
	"errors"
	"github.com/TeaWeb/plugin/pkg/messages"
	"github.com/TeaWeb/plugin/pkg/plugins"
	"log"
	"net/http/httputil"
	"reflect"
)

type Loader struct {
	plugin *plugins.Plugin

	methods   map[string]reflect.Method
	thisValue reflect.Value

	reader PipeInterface
	writer PipeInterface

	debug bool
}

type PipeInterface interface {
	Read([]byte) (n int, err error)
	Write([]byte) (n int, err error)
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
			log.Println("[plugin]", err.Error())
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
		return errors.New("handler for '" + action.Name() + "' not found")

	}
	method.Func.Call([]reflect.Value{this.thisValue, reflect.ValueOf(action)})
	return nil
}

func (this *Loader) ActionStart(action *messages.StartAction) {
	this.plugin.Start()
}

func (this *Loader) ActionReload(action *messages.ReloadAction) {
	this.plugin.Reload()
}

func (this *Loader) ActionStop(action *messages.StartAction) {
	this.plugin.Stop()
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

// 刷新所有的Apps
func (this *Loader) ActionReloadApps(action *messages.ReloadAppsAction) {
	this.plugin.ReloadApps()
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
