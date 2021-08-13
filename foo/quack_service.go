package foo

import (
	"fmt"
	"github.com/Chanokthorn/go-public-lib-test/quack"
)

type quackService struct {
	name string
	age  int
}

func NewQuackService(name string, age int) quack.Service {
	return &quackService{name: name, age: age}
}

func (q *quackService) Execute(sound string) error {
	fmt.Printf("duck name: %s, age: %d quacks with 'foo': %s\n", q.name, q.age, sound)
	return nil
}
