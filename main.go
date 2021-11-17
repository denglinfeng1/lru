package main

import "fmt"

type node struct {
	key  interface{}
	val  interface{}
	pre  *node
	next *node
}

type lru struct {
	size   int
	top    *node
	bottom *node
	m      map[interface{}]*node
}

func newLRU(size int) *lru {
	l := &lru{
		size:   size,
		top:    nil,
		bottom: nil,
		m:      map[interface{}]*node{},
	}
	return l
}

func deleteNode(n *node, l *lru)  {
	delete(l.m, n.key)
	p := n.pre
	nex := n.next
	if p != nil {
		p.next = nex
		if nex == nil {
			l.bottom = p
		}
	}
	if nex != nil {
		nex.pre = p
		if p == nil {
			l.top = nex
		}
	}
}

func addNode(n *node, l *lru)  {
	l.m[n.key] = n
	l.top = n
	if l.top != nil {
		top := l.top
		top.pre = n
		n.next = top
	}
	if l.bottom == nil {
		l.bottom = n
	}
}

func (l *lru)get(k interface{}) interface{} {
	n, ok := l.m[k]
	if !ok {
		return nil
	} else {
		deleteNode(n, l)
		addNode(n, l)
		return n.val
	}
}

func (l *lru)set(k, v interface{}) {
	if _,ok := l.m[k]; ok {
		return
	}
	n := &node{
		key:  k,
		val:  v,
	}
	addNode(n, l)
	// more than size
	if len(l.m) > l.size {
		deleteNode(l.bottom, l)
	}
}

func main() {
	l := newLRU(2)
	l.set(1, 2)
	l.set(3, 4)
	l.set(5, 6)
	fmt.Println(l.get(1))
}
