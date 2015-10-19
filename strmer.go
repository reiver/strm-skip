package verboten


import (
	"github.com/reiver/go-strm/driver"
)


const (
	// SKIP is a (string) constant that this Beginner driver
	// is registered under.
	SKIP = "SKIP"

	defaultSkip = 0
)


func init() {
	strmDriver := newStrmer()

	strmdriver.RegisterStrmer(SKIP, strmDriver)
}


type internalStrmer struct{}


func newStrmer() strmdriver.Strmer {
	strmDriver := internalStrmer{

	}

	return &strmDriver
}



func (strmDriver *internalStrmer) Strm(src <-chan []interface{}, dst chan<- []interface{}, args ...interface{}) {

	// Parse args.
	if 1 < len(args) {
//@TODO: Throw something better.
		panic("Too many parameters.")
	}

	var skip int = defaultSkip
	if 1 == len(args) {
		arg0 := args[0]
		switch n := arg0.(type) {
		case int:
			skip = n
		case int8:
			skip = int(n)
		case int16:
			skip = int(n)
		case int32:
			skip = int(n)
		case int64:
			skip = int(n)
		case uint:
			skip = int(n)
		case uint8:
			skip = int(n)
		case uint16:
			skip = int(n)
		case uint32:
			skip = int(n)
		case uint64:
			skip = int(n)
		default:
//@TODO: Throw something better.
			panic("Wrong type for skip.")
		}
	}

	// Execute.
	i :=0
	for datum := range src {
		if i >= skip {
			dst <- datum
		}

		i++
	}
	close(dst)
}
