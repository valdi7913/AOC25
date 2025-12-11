package main

import (
	"fmt"
)

type UF struct {
	count int
	id []int
}

func NewUnion(n int) *UF {
	id := make([]int, n)
	for i:= range n {
		id[i] = i
	}
	return &UF{count: n, id: id}
}

func (u *UF) Count() int {
	return u.count;
}

func (u *UF) Find(p int) int {
	err := u.validate(p);
	if err != nil {
		return -1
	}
	return u.id[p];
}

func (u *UF) validate(p int) error {
	n := len(u.id)
	if (p < 0 || p >= n) {
		return fmt.Errorf("Illegal argument index %d is not between 0 and %d", p, n-1)
	}
	return nil
}

func (u *UF) Connected(p, q int) bool {
	pErr := u.validate(p)
	qErr := u.validate(q)
	if (pErr != nil || qErr != nil) {
		return false
	}
	return u.id[p] == u.id[q];
}

func (u *UF) Union(p, q int) error {
	pErr := u.validate(p);
	qErr := u.validate(q);
	if pErr != nil {
		return pErr;
	}
 	if qErr != nil {
  	return qErr;
	}

	pID := u.id[p]
	qID := u.id[q]

	if (pID == qID) {return nil}

	for i:=0; i < len(u.id); i++ {
		if u.id[i] == pID {
			u.id[i] = qID
		}
	}
	u.count--
	return nil
}

func (u *UF) String() {
	fmt.Printf("UF: count=%d, id=%v\n", u.count, u.id)
}
