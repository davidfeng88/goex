# Time comparison: strings.Join vs. string concatenation

## strings.Join

Modified echo3:

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
    REPEAT := 10000
	for i := 0; i < REPEAT; i++ {
		fmt.Println(strings.Join(os.Args[1:], " "))
	}
}
```

    $ go build
    $ time ./echo3 a b c > dev\null

    real    0m0.059s
    user    0m0.016s
    sys     0m0.040s

## string concatenation

Modified echo2:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	REPEAT := 10000
	for i := 0; i < REPEAT; i++ {
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
	}
}
```

    $ go build
    $ time ./echo2 a b c > dev\null

    real    0m0.613s
    user    0m0.198s
    sys     0m0.328s

## Conclusion

`strings.Join` is faster. For 10000 loops, it took 0.059s, vs. 0.613s by string concatenation.