package scriptx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScriptRuntime_Execute(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		runtime := NewScriptRuntime()
		ret, err := runtime.Execute(`
function entry(webc) {
	return {
		"pattern": webc.pattern,
		"method": webc.method,
		"path": webc.path,
		"host": webc.host,
		"remoteAddr": webc.remoteAddr,
		"querys": webc.querys,
		"headers": webc.headers,
		"UA": webc.headerVar("User-Agent")
	}
}
`, NewScriptWebContext(r, "/debug"))
		if nil != err {
			t.Fatal(err)
		}
		t.Log(ret)
		bd, err := json.Marshal(ret)
		if nil != err {
			t.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bd)
	}))

	defer srv.Close()
	api := srv.URL
	fmt.Println("url:", api)
	_, err := http.Get(api)
	if nil != err {
		t.Fatal(err)
	}
}
