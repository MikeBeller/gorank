package main

import (
	"bufio"
	_ "fmt"
	//_ "github.com/pkg/profile"
	"math"
	"os"
	"runtime/debug"
	"strconv"
)

type Node struct {
	Left  *Node
	Val   int
	Num   int
	Right *Node
}

func leaf(v int) *Node {
	return &Node{nil, v, 1, nil}
}

func size(t *Node) int {
	if t == nil {
		return 0
	}
	return t.Num
}

func node(l *Node, v int, r *Node) *Node {
	return &Node{l, v, size(l) + size(r) + 1, r}
}

func treeFromVals(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	} else if len(vals) == 1 {
		return leaf(vals[0])
	} else {
		m := len(vals) / 2
		return node(treeFromVals(vals[:m]), vals[m], treeFromVals(vals[m+1:]))
	}
}

func inorder(n *Node) []int {
	if n == nil {
		return make([]int, 0)
	} else {
		return append(append(inorder(n.Left), n.Val), inorder(n.Right)...)
	}
}

func splay(t *Node, i int) *Node {
	if t == nil {
		panic("empty tree")
	} else {
		rr := t.Num - size(t.Right)
		if i == rr {
			return t
		} else if i < rr {
			x, y, c := splay(t.Left, i), t, t.Right
			a, b := x.Left, x.Right
			return node(a, x.Val, node(b, y.Val, c))
		} else {
			a, x, y := t.Left, t, splay(t.Right, i-rr)
			b, c := y.Left, y.Right
			return node(node(a, x.Val, b), y.Val, c)
		}
	}
}

func splitLeft(t *Node, i int) (*Node, *Node) {
	t2 := splay(t, i)
	return t2.Left, node(nil, t2.Val, t2.Right)
}

func splitRight(t *Node, i int) (*Node, *Node) {
	t2 := splay(t, i)
	return node(t2.Left, t2.Val, nil), t2.Right
}

func join(tl *Node, tr *Node) *Node {
	if tl == nil && tr == nil {
		return nil
	} else if tr == nil {
		return tl
	} else if tl == nil {
		return tr
	} else {
		ll := splay(tl, size(tl))
		rr := splay(tr, 1)
		return node(ll.Left, ll.Val, rr)
	}
}

func treeMove(t *Node, side, i, j int) *Node {
	tLeft, tTmp := splitLeft(t, i)
	tMid, tRight := splitRight(tTmp, j-i+1)
	if side == 1 {
		return join(tMid, join(tLeft, tRight))
	} else {
		return join(tLeft, join(tRight, tMid))
	}
}

func scanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	v, _ := strconv.Atoi(scanner.Text())
	return v
}

func writeInt(io *bufio.Writer, i int) {
	d := strconv.Itoa(i)
	io.WriteString(d)
}

func main() {
	debug.SetGCPercent(-1)
	//defer profile.Start(profile.MemProfile).Stop()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	n := scanInt(scanner)
	m := scanInt(scanner)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt(scanner)
	}
	t := treeFromVals(a)
	for k := 0; k < m; k++ {
		d := scanInt(scanner)
		i := scanInt(scanner)
		j := scanInt(scanner)
		t = treeMove(t, d, i, j)
	}
	r := inorder(t)
	io := bufio.NewWriter(os.Stdout)
	writeInt(io, int(math.Abs(float64(r[0]-r[len(r)-1]))))
	io.WriteString("\n")
	for _, v := range r {
		writeInt(io, v)
		io.WriteString(" ")
	}
	io.WriteString("\n")
	io.Flush()
}
