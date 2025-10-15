package even_odd

import (
	"testing"
	"fmt"
)

func TestIsEven(t *testing.T) {
	tests := []struct {
		a    int
		want string
	}{
		{3, "odd"},
		{4, "even"},
		{4542, "even"},
	}

	for _, val := range tests {
		got := IsEven(val.a)
		if got != val.want { // ✅ correct boolean comparison
			t.Errorf("IsEven(%d) = %v; want %v", val.a, got, val.want)
		} else {
			fmt.Println("✅ Test passed for:", val.a, "Expected:", val.want)
		}
	}
}
