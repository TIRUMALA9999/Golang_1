package calculator

import (
	//"errors"
	"testing"
)

// ----------- Test Add (subtests example) -----------
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed sign", -2, 3, 1},
		{"zero", 0, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// ----------- Test Sub -----------
func TestSub(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		got := Sub(10, 3)
		if got != 7 {
			t.Errorf("Sub(10,3) = %d; want 7", got)
		}
	})
	t.Run("negative", func(t *testing.T) {
		got := Sub(-2, -3)
		if got != 1 {
			t.Errorf("Sub(-2,-3) = %d; want 1", got)
		}
	})
}

// ----------- Test Mul -----------
func TestMul(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-2, 3, -6},
	}

	for _, tt := range tests {
		if got := Mul(tt.a, tt.b); got != tt.want {
			t.Errorf("Mul(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

// ----------- Test Div with Error Handling -----------
func TestDiv(t *testing.T) {
	t.Run("normal division", func(t *testing.T) {
		got, err := Div(10, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 5 {
			t.Errorf("Div(10,2) = %d; want 5", got)
		}
	})

	/*t.Run("divide by zero", func(t *testing.T) {
		_, err := Div(10, 1)
		if !errors.Is(err, errors.New("divide by zero")) {
			if err == nil {
				t.Error("expected error, got nil")
			} else {
				t.Errorf("unexpected error: %v", err)
			}
		}
	})*/
}

// ----------- Benchmark -----------
func BenchmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul(100, 200)
	}
}
