
package main

import (
	"bytes"
	"fmt"
	"os"
)

var examples = []string{
	"players",
	"players.list",
	"players.name",
	"players.exists(name)",
	"players.list[name].get(hp, mana, life)",
	"obj.func(p1, p2, p3)",
	`
players.list.{
 name1.get(hp, mana);
 name2.rename('derperina');
 name3.{
  move(x,y,z);
  hp.set(1);
 }
}
`,
}

var (
	errLen = fmt.Errorf("Unexpected EOF.")
)

func main() {
	fmt.Println("Results:")
	for i, ex := range examples {
		out, err := run(ex)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Printf("[%d]: %s \n", i, out)
	}
}

// run the automata.
// <name>
// <name>.<name>
//
func run(w string) (string, error) {
	var buf bytes.Buffer
	// var err error
	// i := 0
	// l := len(w)
	for _, r := range w {
		x := state1(r)
		if x != ' ' {
			buf.WriteRune(x)
		}
	}
	return buf.String(), nil
	//	return "", err
}

func state1(r rune) rune {
	switch r {
	case '.':
		return '>'
	case ',':
		return '&'
	case '{', '(', ')', '}':
		return r
	default:
		return rune(183)
		// 183 middle dot
		// 9251 open box
	}
}

func parseDot() {}
