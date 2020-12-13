package scriptx

import (
	"fmt"
	"github.com/dop251/goja"
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

func (c *ScriptRuntime) Execute(scripts string, context interface{}) (v interface{}, err error) {
	_, rerr := c.vm.RunScript("scriptx:entry", scripts)
	if nil != rerr {
		return nil, fmt.Errorf("compile script, error: %w", rerr)
	}
	var entry func(goja.Value) interface{}
	verr := c.vm.ExportTo(c.vm.Get(ScriptEntryFunName), &entry)
	if verr != nil {
		return nil, fmt.Errorf("bind runtime entry function, error: %w", verr)
	}
	defer func() {
		if r := recover(); nil != r {
			err = fmt.Errorf("executing script, error: %s", r)
		}
	}()
	c.vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	return entry(c.vm.ToValue(context)), nil
}
