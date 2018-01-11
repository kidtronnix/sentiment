package response

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	A int
	B string
}

func TestInternalError(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()

	InternalError(w)

	assert.Equal(500, w.Code)
	assert.Equal("application/json", w.Header().Get("Content-Type"))
	assert.Equal(`{"statusCode":500,"error":"Internal Error!"}`, w.Body.String())
}

func TestBadRequest(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()

	BadRequest(w, "POOP")

	assert.Equal(400, w.Code)
	assert.Equal("application/json", w.Header().Get("Content-Type"))
	assert.Equal(`{"statusCode":400,"error":"Bad Request!","message":"POOP"}`, w.Body.String())
}

func TestJSON(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()

	test := Test{
		1,
		"BOOP",
	}

	JSON(w, test)

	assert.Equal(200, w.Code)
	assert.Equal("application/json", w.Header().Get("Content-Type"))
	assert.Equal(`{"A":1,"B":"BOOP"}`, w.Body.String())
}

type BadJSONMarshaller struct{}

func (s BadJSONMarshaller) MarshalJSON() ([]byte, error) {
	return nil, errors.New("woops")
}

func TestJSONMarshalError(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()

	JSON(w, BadJSONMarshaller{})

	assert.Equal(500, w.Code)
	assert.Equal("application/json", w.Header().Get("Content-Type"))
	assert.Equal(`{"statusCode":500,"error":"Internal Error!"}`, w.Body.String())
}
