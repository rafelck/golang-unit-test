package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("rafel")

	if result != "hello rafel" {
		t.Error("Result must be hello rafel")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldLino(t *testing.T) {
	result := HelloWorld("lino")

	if result != "hello lino" {
		t.Fatal("Result must be hello lino")

	}
	fmt.Println("TestHelloWorldLino Done")

}

func TestHelloWordAssert(t *testing.T) {
	result := HelloWorld("lino")
	assert.Equal(t, "hello rafel", result, "they should be equal")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWordRequire(t *testing.T) {
	result := HelloWorld("lino")
	require.Equal(t, "hello rafel", result, "they should be equal")
	fmt.Println("Tidak dieksekusi")
}
