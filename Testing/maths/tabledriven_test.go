//Go developers often use this style — one test function that runs multiple cases.
package maths
import "testing"
func TestAdd1(t *testing.T){
	tests:=[]struct{
		a,b,want int
	}{
		{2,3,5},
		{-1,1,0},
		{0,0,0},
	}

	for _,tt:=range tests{
		got:=Add(tt.a,tt.b)
		if got != tt.want{
			t.Errorf("Add(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}