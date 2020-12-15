/*
	这个包是一个gin的中间件，用来简化入参和出参方式
*/
package wine

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

const (
	BindingErrorCode = "A0410"
	BindingErrorMsg  = "参数绑定异常"

	SystemErrorCode = "B0001"
	SystemErrorMsg  = "系统异常"

	DatabaseErrorCode = "C0300"
	DatabaseErrorMsg  = "数据库异常"
)

type ErrorResult struct {
	Code string
	Msg  string
}

var (
	BindingError  = ErrorResult{BindingErrorCode, BindingErrorMsg}
	SystemError   = ErrorResult{SystemErrorCode, SystemErrorMsg}
	DatabaseError = ErrorResult{DatabaseErrorCode, DatabaseErrorMsg}
)

var (
	ginContextType = reflect.TypeOf(&gin.Context{})
)

// Wine 中间件使用
func Wine(handlerFunc interface{}) gin.HandlerFunc {
	handlerFuncType := reflect.TypeOf(handlerFunc)
	if handlerFuncType.Kind() != reflect.Func {
		panic("handler need func but got " + handlerFuncType.Name())
	}
	return func(c *gin.Context) {
		handlerType := reflect.TypeOf(handlerFunc)
		if handlerType.Kind() == reflect.Ptr {
			handlerType = handlerType.Elem()
		}
		var param interface{}
		var args = make([]reflect.Value, 0, handlerType.NumIn())

		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"errno": SystemErrorCode, "errmsg": "red raw data err:" + err.Error()})
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		for i := 0; i < handlerType.NumIn(); i++ {
			reqType := handlerType.In(i)

			switch reqType {
			case ginContextType:
				// 是ginContext上下文
				args = append(args, reflect.ValueOf(c))

			default:
				// 默认是参数
				param = reflect.New(reqType).Interface()
				if err := c.ShouldBind(param); err != nil {
					log.Printf("bind err:%v", err)
					c.JSON(http.StatusOK, gin.H{"errno": BindingErrorCode, "errmsg": BindingErrorMsg})
					return
				}
				args = append(args, reflect.ValueOf(param).Elem())
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			}
		}

		handlerFuncVal := reflect.ValueOf(handlerFunc)
		obj := handlerFuncVal.Call(args)
		for _, value := range obj {
			err, ok := value.Interface().(ErrorResult)
			if ok {
				c.JSON(http.StatusOK, gin.H{"errno": err.Code, "errmsg": err.Msg})
			} else {
				c.JSON(http.StatusOK, gin.H{"errno": "0", "errmsg": "OK", "data": value.Interface()})
			}
		}
	}
}
