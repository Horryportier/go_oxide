package option


type None interface{
    N() interface{};
} 

type Some interface{ 
    S() interface{};  
}

type Option[T any] struct {Some;None}

func (o Option[T]) Unwrap() interface{} {
    if o.IsSome() {
        panic("Option is nil")
    }
    return o.Some.S()
}


func (o Option[T]) IsSome() bool {
    if o.Some != nil {
        return true
    }
    return false
}

