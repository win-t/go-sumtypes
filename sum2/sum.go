// Code generated by "go-sumtypes/generate.go". DO NOT EDIT.

package sum2

type Type[T0, T1 any] struct {
	v any
}

func (s *Type[T0, T1]) Set0(v T0) {
	s.v = v
}

func (Type[T0, T1]) New0(v T0) Type[T0, T1] {
	return Type[T0, T1]{v}
}

func (s Type[T0, T1]) As0() (T0, bool) {
	v, ok := s.v.(T0)
	return v, ok
}

func (s *Type[T0, T1]) Set1(v T1) {
	s.v = v
}

func (Type[T0, T1]) New1(v T1) Type[T0, T1] {
	return Type[T0, T1]{v}
}

func (s Type[T0, T1]) As1() (T1, bool) {
	v, ok := s.v.(T1)
	return v, ok
}

func (s Type[T0, T1]) Underlying() any {
	return s.v
}

func (s Type[T0, T1]) Case(f0 func(T0), f1 func(T1)) {
	switch v := s.v.(type) {
	case T0:
		f0(v)
	case T1:
		f1(v)
	default:
		panic("called Case on an invalid value")
	}
}
