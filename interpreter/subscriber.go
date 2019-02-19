package interpreter

import (
	"github.com/tzmfreedom/land/ast"
)

type Subscriber func(ctx *Context, n ast.Node)

var subscribers = map[string][]Subscriber{}

func Publish(event string, ctx *Context, n ast.Node) {
	if subs, ok := subscribers[event]; ok {
		for _, s := range subs {
			s(ctx, n)
		}
	}
}

func Subscribe(event string, subscriber Subscriber) {
	if v, ok := subscribers[event]; ok {
		subscribers[event] = append(v, subscriber)
	} else {
		subscribers[event] = []Subscriber{subscriber}
	}
}
