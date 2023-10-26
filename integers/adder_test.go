package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	got := Add(1, 2)
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got , want)
	}
}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	//Output: 4
}