package distributedcalc

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestE2E(t *testing.T) {
	body := strings.NewReader(`{"expression": "X-(2Y+(X/Y))","variables":{"X":8,"Y":4}}`)
	r, err := http.NewRequest(http.MethodPut, os.Getenv("API_ADDR"), body)
	if err != nil {
		t.Fatal(err)
	}
	c := &http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code 200, but got %v", resp.StatusCode)
	}
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if bytes.Equal(byt, []byte(`{"result":-2}`)) {
		t.Error("unexpected body")
	}
}
