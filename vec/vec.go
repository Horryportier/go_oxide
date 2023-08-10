package vec

type Vec[T any] []T 

func (v Vec[T]) New() Vec[T] {
        v = []T{}
        return  v
}

type Option[T any]  struct {Some T; None interface{}}

func (v *Vec[T]) Len() int {
    return len(*v)
}

func (v *Vec[T]) From(items ...T) {
    for _,i := range items {
        *v = append(*v, i)
    }
}

func (v *Vec[T]) Append(rhs []T) {
    *v = append(*v, rhs...)
}

func (v *Vec[T]) Push(item T) {
    *v = append(*v, item)
    return 
}

func (v *Vec[T]) Pop()  T {
    var vec Vec[T] = *v
    if vec.Len() == 0 {
        return 
    }
    i := vec[0]
    if i != nil {

    }

}
