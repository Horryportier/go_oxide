package option_test

import (
	. "Horryportier/go_oxide/option"
	. "Horryportier/go_oxide/types"
	"testing"
)

func TestOption(t *testing.T) {
    var option Option[Usize]
    var u Usize  = 2
    option.Some = u;
    switch option.IsSome() {
    case true:  
        return
    case false:
        t.Error("option of Usize is not Some")
    }
}


func TestUnwrap(t *testing.T) {
    var option Option[Usize] = 
    some :=  option.Unwrap()
    if some != 3 {
        t.Error("Unwrap didin't return anything")
    }
}
