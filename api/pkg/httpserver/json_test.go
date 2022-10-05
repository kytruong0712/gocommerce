package httpserver

import (
	"bytes"
	"net/http/httptest"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseJSON(t *testing.T) {
	// Given:
	given := []byte(`{"message":"OK"}`)
	r := httptest.NewRequest(http.MethodPost, "/some/path", bytes.NewReader(given))
	reader := r.Body
	var obj Success

	// When:
	err := ParseJSON(reader, &obj)

	// Then:
	require.Nil(t, err)
	require.Equal(t, "OK", obj.Message)
}
