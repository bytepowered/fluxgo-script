package scriptx

import (
	"net/http"
	"net/url"
)

type GetVarFunc func(key string) string

type ScriptContext struct {
	RequestPattern    string      `json:"pattern"`
	RequestMethod     string      `json:"method"`
	RequestPath       string      `json:"path"`
	RequestHost       string      `json:"host"`
	RequestRemoteAddr string      `json:"remoteAddr"`
	HeaderValues      http.Header `json:"headers"`
	FormValues        url.Values  `json:"forms"`
	QueryValues       url.Values  `json:"query"`
	// Function
	GetHeaderVarFunc GetVarFunc `json:"headerVar"`
	GetFormVarFunc   GetVarFunc `json:"formVar"`
	GetQueryVarFunc  GetVarFunc `json:"queryVar"`
}

func NewScriptContext(r *http.Request, pattern string) ScriptContext {
	query := r.URL.Query()
	return ScriptContext{
		RequestPattern:    pattern,
		RequestMethod:     r.Method,
		RequestPath:       r.RequestURI,
		RequestHost:       r.Host,
		RequestRemoteAddr: r.RemoteAddr,
		HeaderValues:      r.Header,
		FormValues:        r.PostForm,
		QueryValues:       query,
		GetHeaderVarFunc: func(key string) string {
			return r.Header.Get(key)
		},
		GetFormVarFunc: func(key string) string {
			return r.PostForm.Get(key)
		},
		GetQueryVarFunc: func(key string) string {
			return query.Get(key)
		},
	}
}
