package interpreter

import (
	"fmt"
	"math"
	"os"
	"strings"

	"bufio"

	"strconv"

	"github.com/chzyer/readline"
	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
)

var Debugger = &debugger{}

type debugger struct {
	Enabled bool
	StepOut bool
	Step    int
	Frame   int
}

func (d *debugger) Debug(ctx *Context, n ast.Node) {
	d.Enabled = true
	d.StepOut = false
	d.Step = 0
	d.Frame = 0

	showCurrent(n)
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m>>\033[0m ",
		HistoryFile:     "/tmp/land_debugger.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	for {
		line, err := l.Readline()
		if err != nil {
			panic(err)
		}
		inputs := strings.Split(line, " ")
		cmd := inputs[0]
		switch cmd {
		case "s", "step":
			d.Step = 1
			return
		case "n", "next":
			d.StepOut = true
			d.Step = 1
			return
		case "f", "finish":
			d.StepOut = true
			d.Frame = 1
			d.Step = 1
			return
		case "current":
			showCurrent(n)
		case "exit":
			d.Step = 0
			d.Frame = 0
			d.StepOut = false
			d.Enabled = false
			return
		default:
			if cmd == "_" {
				for k, v := range ctx.Env.Data.All() {
					fmt.Printf("%s => %s\n", k, builtin.String(v))
				}
			} else if cmd != "" {
				varName := cmd
				resolver := NewTypeResolver(ctx)
				obj, err := resolver.ResolveVariable(strings.Split(varName, "."))
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				fmt.Println(builtin.String(obj))
			}
		}
	}
}

func showCurrent(n ast.Node) {
	file := n.GetLocation().FileName
	line := n.GetLocation().Line
	column := n.GetLocation().Column
	fmt.Printf("=== debugger === \n")
	fmt.Printf("filename: %s\n", file)
	fmt.Printf("line: %d, column: %d\n", line, column)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	width := int(math.Floor(math.Log10(float64(line) + 5.0)))
	for i := 1; sc.Scan(); i++ {
		if line == i {
			format := ">> %0" + strconv.Itoa(width) + "d %s\n"
			fmt.Printf(format, i, sc.Text())
		} else if line-5 < i && line+5 > i {
			format := "   %0" + strconv.Itoa(width) + "d %s\n"
			fmt.Printf(format, i, sc.Text())
		}
	}
}

func init() {
	Subscribe("method_start", func(ctx *Context, n ast.Node) {
		if Debugger.Enabled {
			Debugger.Frame++
		}
	})
	Subscribe("method_end", func(ctx *Context, n ast.Node) {
		if Debugger.Enabled {
			if Debugger.Frame > 0 {
				Debugger.Frame--
			}
		}
	})
	Subscribe("line", func(ctx *Context, n ast.Node) {
		if Debugger.Enabled {
			if Debugger.Step > 0 {
				if Debugger.StepOut && Debugger.Frame != 0 {
					return
				}
				Debugger.Step--
			}
			if Debugger.Step == 0 {
				Debugger.Debug(ctx, n)
			}
		}
	})
}
