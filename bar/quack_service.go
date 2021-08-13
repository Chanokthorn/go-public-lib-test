package bar

import (
	"fmt"
	"github.com/Chanokthorn/go-public-lib-test/quack"
)

type quackService struct {
	name string
}

func NewQuackService(name string) quack.Service {
	return &quackService{name: name}
}

func (q *quackService) Execute(sound string) error {
	fmt.Printf("%s quacks with 'bar': %s\n", q.name, sound)
	return nil
}
