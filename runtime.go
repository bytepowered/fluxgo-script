package scriptx

import (
	"fmt"
	"github.com/dop251/goja"
	"reflect"
)

const (
	ScriptEntryFunName = "entry"
)

type ScriptRuntime struct {
	vm *goja.Runtime
}

func NewScriptRuntime() *ScriptRuntime {
	return &ScriptRuntime{
		vm: goja.New(),
	}
}

func (c *ScriptRuntime) Execute(scripts string, structContext interface{}) (v interface{}, err error) {
	return c.Call(scripts, ScriptEntryFunName, structContext)
}

func (c *ScriptRuntime) Call(scripts string, funName string, structContext interface{}) (v interface{}, err error) {
	rv := reflect.ValueOf(structContext)
	if rv.IsNil() || rv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("context MUST be struct, was: %s", rv.Kind().String())
	}
	_, rerr := c.vm.RunScript("scriptx:"+funName, scripts)
	if nil != rerr {
		return nil, fmt.Errorf("compile script, error: %w", rerr)
	}
	var entry func(goja.Value) interface{}
	verr := c.vm.ExportTo(c.vm.Get(funName), &entry)
	if verr != nil {
		return nil, fmt.Errorf("bind runtime entry function, error: %w", verr)
	}
	defer func() {
		if r := recover(); nil != r {
			err = fmt.Errorf("executing script, error: %s", r)
		}
	}()
	c.vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	return entry(c.vm.ToValue(structContext)), nil
}
