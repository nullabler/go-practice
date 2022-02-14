package range_extr

import (
	"fmt"
	"testing"
)

type DataProvider struct {
	list   []int
	expect string
}

var dataProvides = []DataProvider{
	DataProvider{
		[]int{-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20},
		"-10--8,-6,-3-1,3-5,7-11,14,15,17-20",
	},
	DataProvider{
		[]int{},
		"",
	},
}

func TestSolution(t *testing.T) {
	for _, item := range dataProvides {
		output := Solution(item.list)
		if output != item.expect {
			t.Errorf("Output %q not equal to expected %q", output, item.expect)
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution(dataProvides[0].list)
	}
}

func ExampleSolution() {
	fmt.Println(Solution(dataProvides[0].list))
	// Output: -10--8,-6,-3-1,3-5,7-11,14,15,17-20
}
