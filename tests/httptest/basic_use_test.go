//REF: https://bignerdranch.com/blog/using-the-httptest-package-in-golang/
package httptest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//////////////////////////////////Testing an HTTP handler//////////////////////////////
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("API-VERSION", "1.3")
	w.Write([]byte("ok"))
}

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	if want, got := "1.3", w.Result().Header.Get("API-VERSION"); want != got {
		t.Fatalf("expected a %q, instead got: %q", want, got)
	}
}

//////////////////////////////////Using the Test Server//////////////////////////////
type client struct{}

func (c *client) Sunday(url string) bool {
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, _ := client.Do(r)

	bodyBytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	return string(bodyBytes) == `{"isItSunday":true}`
}
func (c *client) GetDay(url string) (string, error) {
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, url+"/token", nil)
	resp1, _ := client.Do(r)

	bodyBytes, _ := io.ReadAll(resp1.Body)
	defer resp1.Body.Close()

	token := string(bodyBytes)
	client2 := &http.Client{}

	r, _ = http.NewRequest(http.MethodGet, url+"/resource", nil)
	r.Header.Set("Auth", token)
	resp2, _ := client2.Do(r)
	var day struct {
		Day string
	}
	dec := json.NewDecoder(resp2.Body)
	defer resp2.Body.Close()
	dec.Decode(&day)
	return day.Day, nil
}

func TestTrueSundayResponseReturnsTrue(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"isItSunday":true}`))
	}))
	defer s.Close()

	c := &client{}

	if !c.Sunday(s.URL) {
		t.Fatalf("expected client to return true!")
	}
}

func TestPseudoOAuth(t *testing.T) {
	secret := fmt.Sprintf("secret_%d", time.Now().Unix())

	mux := http.NewServeMux()

	mux.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Auth")
		if auth != secret {
			http.Error(w, "Auth header was incorrect", http.StatusUnauthorized)
			return
		}
		w.Write([]byte(`{"day":"Sunday"}`))
	})
	mux.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(secret))
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	c := &client{}
	got, err := c.GetDay(s.URL)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if want := "Sunday"; want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}
