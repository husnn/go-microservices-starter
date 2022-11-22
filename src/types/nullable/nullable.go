package nullable

type Nullable[T any] struct {
	Ptr *T
}

func (n *Nullable[T]) Null() bool {
	return n.Ptr == nil
}

func (n *Nullable[T]) NotNull() bool {
	return n.Ptr != nil
}

func (n *Nullable[T]) Value() *T {
	return n.Ptr
}

func NewNull[T any]() Nullable[T] {
	return Nullable[T]{Ptr: nil}
}

func (n *Nullable[any]) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	i := src.(any)
	n.Ptr = &i
	return nil
}

type String = Nullable[string]

func NewString(val string) String {
	if val == "" {
		return NewNull[string]()
	}
	return String{Ptr: &val}
}

func (n *Nullable[string]) ValueOrEmpty() string {
	var val string
	if n.NotNull() {
		val = *n.Ptr
	}
	return val
}

type Int64 = Nullable[int64]

func NewInt64(val int64) Int64 {
	if val == 0 {
		return NewNull[int64]()
	}
	return Int64{Ptr: &val}
}

func (n *Nullable[int64]) ValueOrZero() int64 {
	var val int64
	if n.NotNull() {
		val = *n.Ptr
	}
	return val
}
