package concat_test

import (
	"fmt"

	"github.com/Quasilyte/concat"
)

func ExampleStrings() {
	v := "world!"
	fmt.Println(concat.Strings("hello, ", v))

	// Output: hello, world!
}
