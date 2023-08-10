package vec_test

import (
	. "Horryportier/go_oxide/vec"
	"testing"
)

func EqualSlice[T comparable](lhs []T, rhs []T) bool {
        if len(lhs) != len(rhs) {
            return false
        }
        for i,l := range lhs {
            if l != rhs[i] {
                return false
            }
        }
        return true
}


func TestNew(t *testing.T)  {
    var s []int
    var v Vec[int]
    v.New()
    if !EqualSlice(v, s) {
        t.Errorf("v is not []int %v", v)
    }
}

func TestLen(t *testing.T)  {
    var v Vec[string] = []string{"0", "1", "2"}
    if v.Len() != 3 {
        t.Errorf("len (%v) of v is not 3", v.Len())
    }
}

func TestFrom(t *testing.T) {
    var s []int  = []int{1, 2, 3, 4}
    var v Vec[int] 
    v.From(s...)
    if !EqualSlice(v, s) {
        t.Errorf("v: %v is not equal s: %v", v, s)
    }
}

func TestAppend(t *testing.T) {
    var f []int = []int{1, 2, 3, 6}
    var l []int = []int{1, 2}
    var r []int = []int{3, 6}

    var v Vec[int]
    v.New()
    v.Append(l)
    v.Append(r)

    if !EqualSlice(v, f) {
        t.Errorf("v: %v not equal to f: %v ", v, f)
    }

}
