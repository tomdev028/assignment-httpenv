package main
import (
  "encoding/json"
  "net/http"
  "net/http/httptest"
  "testing"
)


func TestServeReturnsJSON(t *testing.T) {
  req := httptest.NewRequest("GET", "/", nil)
  w := httptest.NewRecorder()

  serve(w, req)

  resp := w.Result()
  if resp.StatusCode != http.StatusOK {
    t.Errorf("expected 200 got %d", resp.StatusCode)
  }

  // check valid JSON
  var result map[string]string
  if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
    t.Errorf("response is not valid JSON: %v", err)
  }
}
