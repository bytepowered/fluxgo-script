package scriptx

import (
	"github.com/bytepowered/flux"
	"net/http"
)

func NewScriptWebContext(webc flux.WebContext, pattern string) ScriptContext {
	header, _ := webc.HeaderValues()
	request := webc.RawWebRequest().(*http.Request)
	return ScriptContext{
		RequestPattern:    pattern,
		RequestMethod:     webc.Method(),
		RequestPath:       webc.RequestURI(),
		RequestHost:       webc.Host(),
		RequestRemoteAddr: request.RemoteAddr,
		HeaderValues:      header,
		FormValues:        webc.FormValues(),
		QueryValues:       webc.QueryValues(),
		GetHeaderVarFunc: func(key string) string {
			return webc.HeaderValue(key)
		},
		GetFormVarFunc: func(key string) string {
			return webc.FormValue(key)
		},
		GetQueryVarFunc: func(key string) string {
			return webc.QueryValue(key)
		},
	}
}
