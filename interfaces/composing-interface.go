package main

type Reader interface {
	Read(p string)
}

type Writer interface {
	Write() string
}

/*
Instead of Writing like below
type ReaderWriter interface {
	Read(p string)
	Write() string
}
*/

type ReaderWriter interface {
	Reader
	Writer
}

// Go allows to combine interfaces by embedding them inside another interface.

// interface embedding is not inheritance