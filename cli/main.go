package main

import (
	"github.com/Chanokthorn/go-public-lib-test/bar"
	"github.com/Chanokthorn/go-public-lib-test/foo"
	"github.com/Chanokthorn/go-public-lib-test/quack"
)

func main() {
	b := bar.NewQuackService("john")
	f := foo.NewQuackService("jane", 24)

	quack.Set(b)
	err := quack.Execute("quack!!!")
	if err != nil {
		panic(err)
	}

	quack.Set(f)
	err = quack.Execute("QUACK")
	if err != nil {
		panic(err)
	}

}
