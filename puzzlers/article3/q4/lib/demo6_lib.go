package lib

import (
	"os"
	in "puzzlers/article3/q4/lib/internalx"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
