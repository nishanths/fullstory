package fullstory

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	router *http.ServeMux
	srv    *httptest.Server
	client *Client

	testdata = map[string][]byte{
		"sessions": []byte(`
[{
    "UserId": 1234567890,
    "SessionId": 1234567890,
    "CreatedTime": 1411492739,
    "FsUrl": "https://www.fullstory.com/ui/ORG_ID/discover/session/1234567890:1234567890"
}]`),
		"exportList": []byte(`
{"exports": [{
    "Start": 1447984800,
    "Stop": 1448071200,
    "ID": 123456789
    },{
    "Start": 1448071200,
    "Stop": 1448157600,
    "ID": 987654321
    },{
    "Start": 1448157600,
    "Stop": 1448244000,
    "ID": 456789123
    }]
}`),
	}
)

func TestMain(m *testing.M) {
	flag.Parse()
	setupTest()
	defer teardownTest()
	os.Exit(m.Run())
}

func setupTest() {
	// Test router.
	router = http.NewServeMux()
	srv = httptest.NewServer(router)

	// FullStory test API client.
	client = NewClient("xyz")
	client.BaseURL = srv.URL

	mw := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			if r.Header.Get("Authorization") != "Basic xyz" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			next(w, r)
		}
	}

	router.HandleFunc("/sessions", mw(func(w http.ResponseWriter, r *http.Request) {
		if r.Form.Get("email") == "john@example.com" {
			w.Write(testdata["sessions"])
			return
		}
		w.Write([]byte("{}"))
	}))

	router.HandleFunc("/export/list", mw(func(w http.ResponseWriter, r *http.Request) {
		w.Write(testdata["exportList"])
	}))

	router.HandleFunc("/export/get", mw(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{}"))
	}))
}

func teardownTest() {
	srv.Close()
}
