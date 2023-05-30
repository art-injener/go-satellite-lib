package device

import "io"

type InputDevice interface {
	Find()
	Open() error
	io.ReadCloser
}
