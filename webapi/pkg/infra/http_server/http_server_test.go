package httpserver

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Should_Execute_Setup_Correctly(t *testing.T) {
	sut := NewHttpServerToTest()

	sut.server.Setup()
	calledTimes := reflect.ValueOf(sut.loggerSpy).Elem().Field(0).Interface()

	assert.Equal(t, calledTimes, 1, "Should called GetHandleFunc once")
}

func Test_Should_Execute_RegistreRoute_POST_Correctly(t *testing.T) {
	sut := NewHttpServerToTest()
	sut.server.Setup()

	sut.server.RegistreRoute("POST", "/api/v1/books", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, "created")
	})

	response, err := sut.doRequest("POST", "/api/v1/books")

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, response.StatusCode, http.StatusCreated, fmt.Sprintf("Expected to get status %d but instead got %d\n", http.StatusOK, response.StatusCode))
}

func Test_Should_Execute_RegistreRoute_GET_Correctly(t *testing.T) {
	sut := NewHttpServerToTest()
	sut.server.Setup()

	sut.server.RegistreRoute("GET", "/api/v1/books", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "created")
	})

	response, err := sut.doRequest("GET", "/api/v1/books")

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, response.StatusCode, http.StatusOK, fmt.Sprintf("Expected to get status %d but instead got %d\n", http.StatusOK, response.StatusCode))
}

func Test_Should_Execute_RegistreRoute_PUT_Correctly(t *testing.T) {
	sut := NewHttpServerToTest()
	sut.server.Setup()

	sut.server.RegistreRoute("PUT", "/api/v1/books", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "created")
	})

	response, err := sut.doRequest("PUT", "/api/v1/books")

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, response.StatusCode, http.StatusOK, fmt.Sprintf("Expected to get status %d but instead got %d\n", http.StatusOK, response.StatusCode))
}

func Test_Should_Execute_RegistreRoute_DELETE_Correctly(t *testing.T) {
	sut := NewHttpServerToTest()
	sut.server.Setup()

	sut.server.RegistreRoute("DELETE", "/api/v1/books", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "created")
	})

	response, err := sut.doRequest("DELETE", "/api/v1/books")

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, response.StatusCode, http.StatusOK, fmt.Sprintf("Expected to get status %d but instead got %d\n", http.StatusOK, response.StatusCode))
}

func Test_Should_Execute_RegistreRoute_Error(t *testing.T) {
	sut := NewHttpServerToTest()
	sut.server.Setup()

	err := sut.server.RegistreRoute("Something", "/api/v1/books", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "created")
	})

	assert.EqualError(t, err, "http method not allowed", "Expected to get error when try to register route with wrong method")
}

func Test_Should_Execute_Run_Correctly(t *testing.T) {
	sut := NewHttpServerToTest()
	sut.server.Setup()

	go func(t *testing.T) {
		err := sut.server.Run()

		assert.NoError(t, err, "Server should run without panic")
	}(t)

	time.Sleep(time.Microsecond * 5)
}

func Test_Should_Execute_RegisterMiddleware_Correctly(t *testing.T) {
	sut := NewHttpServerToTest()
	sut.server.Setup()

	sut.server.RegisterMiddleware(func(ctx *gin.Context) {})
}
