package cat

import "os"

func Cat(arg string, f *os.File, opt ...int) (string, os.Error) {
	// Default
	bufSize := 8192

	// Optional
	if len(opt) > 0 && opt[0] > 0 {
		bufSize = opt[0]
	}

	// Create slice and array
	buf := make([]byte, bufSize)

	// Copy loop
	for {
		switch nr, er := f.Read(buf); {
		case nr > 0:
			if _, ew := os.Stdout.Write(buf[0:nr]); ew != nil {
				return arg, ew
			}
			continue
		case er != nil && er != os.EOF:
			return arg, er
		}
		break
	}
	return arg, nil

}
