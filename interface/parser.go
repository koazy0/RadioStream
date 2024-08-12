package _interface

import (
	"RadioStream/interface/Link"
	"RadioStream/interface/Stream"
)

type Parser interface {
	Link.Link
	Stream.Stream
	Parse(url string)
}
