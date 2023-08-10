package types 


type Usize uint

func (u Usize) S(uint) interface{} {
            
        return u
}

func IntoUsize(i uint) {
    var u Usize
    u = i
    return 
}

type Int int

func (i Int) S() interface{} {
        return i
}

type Int8 int8 

func (i Int8) S() interface{} {
        return i
}

type Int16 int16 

func (i Int16) S() interface{} {
        return i
}


type Int32 int32

func (i Int32) S() interface{} {
        return i
}


type Int64 int64

func (i Int64) S() interface{} {
        return i
}
