//
//  middleware_test.go
//  aych
//
//  Created by karim-w on 10/07/2025.
//

package aych_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/karim-w/aych"
)

func TestSimple(t *testing.T) {
	response := aych.TTP(t.Context(), "https://example.com").AddBasicAuth("mo", "lester").Use(func(tx aych.TTPContext) {
		start := time.Now()
		tx.AddHeader("X-Test", "test-value")
		tx.AddQuery("query", "value")
		tx.AddQueryArray("array", []string{"value1", "value2"})
		tx.AddBodyRaw([]byte("This is a test body"))
		tx.AddCookie(&http.Cookie{Name: "test_cookie", Value: "cookie_value"})
		tx.AddBearerAuth("your_token_here")
		tx.AddHeader("Content-Type", "application/json")
		tx.AddHeader("X-Another-Header", "another-value")
		tx.AddHeaders(map[string]string{
			"X-Header1": "value1",
			"X-Header2": "value2",
		})

		tx.JSONBody(map[string]any{
			"key1": "value1",
			"key2": "value2",
		})

		tx.Next()

		response := tx.Body()
		statusCode := tx.StatusCode()

		elapsed := time.Since(start)

		t.Logf("Request processed in %s got %s with status code %d and err: %v", elapsed, response, statusCode, tx.Error())
	}).Get()

	defer response.Close()

	body := response.Body()
	success := response.Success()
	headers := response.Header()
	curlCommand := response.CURL()

	t.Logf("Response Body: %s status code: %d, success: %v curl: %s headers: %v", body, response.StatusCode(), success, curlCommand, headers)
}

func TestPost(t *testing.T) {
	response := aych.TTP(t.Context(), "https://example.com").Post()
	if response.Error() != nil {
		t.Fatalf("Post request failed: %v", response.Error())
	}
}

func TestPut(t *testing.T) {
	response := aych.TTP(t.Context(), "https://example.com").Put()
	if response.Error() != nil {
		t.Fatalf("Put request failed: %v", response.Error())
	}
}

func TestDelete(t *testing.T) {
	response := aych.TTP(t.Context(), "https://example.com").Del()
	if response.Error() != nil {
		t.Fatalf("Delete request failed: %v", response.Error())
	}
}

func TestPatch(t *testing.T) {
	response := aych.TTP(t.Context(), "https://example.com").Patch()
	if response.Error() != nil {
		t.Fatalf("Patch request failed: %v", response.Error())
	}
}
