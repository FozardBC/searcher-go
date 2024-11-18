package database

import "io"

type IDatabase interface {
	io.Writer
	//io.Reader
}
