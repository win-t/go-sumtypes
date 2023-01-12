// Code generated by "go-sumtypes/generate.go". DO NOT EDIT.

package sum5

import "reflect"

type Type[T0, T1, T2, T3, T4 any] struct {
	v any
}

func (s *Type[T0, T1, T2, T3, T4]) Set0(v T0) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4]) As0() (T0, bool) {
	v, ok := s.v.(T0)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4]) Set1(v T1) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4]) As1() (T1, bool) {
	v, ok := s.v.(T1)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4]) Set2(v T2) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4]) As2() (T2, bool) {
	v, ok := s.v.(T2)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4]) Set3(v T3) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4]) As3() (T3, bool) {
	v, ok := s.v.(T3)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4]) Set4(v T4) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4]) As4() (T4, bool) {
	v, ok := s.v.(T4)
	return v, ok
}

func (s Type[T0, T1, T2, T3, T4]) Case(f0 func(T0), f1 func(T1), f2 func(T2), f3 func(T3), f4 func(T4)) {
	switch v := s.v.(type) {
	case T0:
		if f0 == nil {
			panic("no handler for case " + reflect.TypeOf(v).String())
		}
		f0(v)
	case T1:
		if f1 == nil {
			panic("no handler for case " + reflect.TypeOf(v).String())
		}
		f1(v)
	case T2:
		if f2 == nil {
			panic("no handler for case " + reflect.TypeOf(v).String())
		}
		f2(v)
	case T3:
		if f3 == nil {
			panic("no handler for case " + reflect.TypeOf(v).String())
		}
		f3(v)
	case T4:
		if f4 == nil {
			panic("no handler for case " + reflect.TypeOf(v).String())
		}
		f4(v)
	default:
		panic("called Case on an invalid value")
	}
}
