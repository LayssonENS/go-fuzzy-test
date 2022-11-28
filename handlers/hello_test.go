package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHello(t *testing.T) {
	type args struct {
		Name     string
		Expected string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				Name:     "Laysson",
				Expected: "Hello Laysson",
			},
		},
		{
			name: "test2",
			args: args{
				Name:     "Ana",
				Expected: "Hello Ana",
			},
		},
	}

	srv := httptest.NewServer(http.HandlerFunc(HandleHello))
	defer srv.Close()

	for _, tt := range tests {
		path := fmt.Sprintf("%s/?name=%s", srv.URL, tt.args.Name)
		req, err := http.NewRequest("GET", fmt.Sprintf("%s", path), nil)
		if err != nil {
			t.Fatal(err)
		}

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if response.StatusCode != http.StatusOK {
			t.Fatalf("expect: http 200, got: http")
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatal(err)
		}

		var respBody map[string]string
		err = json.Unmarshal(body, &respBody)
		if err != nil {
			t.Fatal(err)
		}

		if respBody["message"] != tt.args.Expected {
			t.Fatalf("expected: %s, got: %s", tt.args.Expected, respBody["message"])
		}
	}
}
