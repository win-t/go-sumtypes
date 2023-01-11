package optional

type Type[T any] struct{ v *T }

func (o *Type[T]) SetSome(v T) { o.v = &v }
func (o *Type[T]) SetNone()    { o.v = nil }

func (o Type[T]) IsSome() bool { return o.v != nil }
func (o Type[T]) IsNone() bool { return o.v == nil }

func (o Type[T]) UnwrapOrElse(f func() T) T {
	if o.IsSome() {
		return *o.v
	}
	return f()
}

func (o Type[T]) UnwrapOrDefault() T {
	return o.UnwrapOrElse(func() T { var zero T; return zero })
}

func (o Type[T]) UnwrapOr(v T) T {
	return o.UnwrapOrElse(func() T { return v })
}

func (o Type[T]) Unwrap() T {
	return o.UnwrapOrElse(func() T { panic("called Unwrap on a None value") })
}

func (o Type[T]) Expect(msg string) T {
	return o.UnwrapOrElse(func() T { panic(msg) })
}
