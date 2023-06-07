package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Rafel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rafel")
		}
	})

	b.Run("Lino", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Lino")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Rafel)",
			request: "Rafel",
		},
		{
			name:    "HelloWorld(Jono)",
			request: "Jono",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("rafel")
	}
}

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

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot Run On windows")
	}
	result := HelloWorld("rafel")
	require.Equal(t, "hello rafel", result, "they should be equal")
}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")

	m.Run()

	fmt.Println("Setelah Unit Test")
}

func TestSubTest(t *testing.T) {
	t.Run("rafel", func(t *testing.T) {
		result := HelloWorld("rafel")
		require.Equal(t, "hello rafel", result, "harusnya sama")
	})

	t.Run("lino", func(t *testing.T) {
		result := HelloWorld("rafel")
		require.Equal(t, "hello rafel", result, "harusnya sama")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(rafel)",
			request:  "rafel",
			expected: "hello rafel",
		},
		{
			name:     "HelloWorld(idang)",
			request:  "idang",
			expected: "hello idang",
		},
		{
			name:     "HelloWorld(tono)",
			request:  "tono",
			expected: "hello rafel",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
