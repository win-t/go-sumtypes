// Code generated by "go-sumtypes/generate.go". DO NOT EDIT.

package sum11

import "reflect"

type Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] struct {
	v any
}

func (s *Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set0(v T0) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get0() (T0, bool) {
	v, ok := s.v.(T0)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set1(v T1) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get1() (T1, bool) {
	v, ok := s.v.(T1)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set2(v T2) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get2() (T2, bool) {
	v, ok := s.v.(T2)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set3(v T3) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get3() (T3, bool) {
	v, ok := s.v.(T3)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set4(v T4) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get4() (T4, bool) {
	v, ok := s.v.(T4)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set5(v T5) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get5() (T5, bool) {
	v, ok := s.v.(T5)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set6(v T6) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get6() (T6, bool) {
	v, ok := s.v.(T6)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set7(v T7) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get7() (T7, bool) {
	v, ok := s.v.(T7)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set8(v T8) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get8() (T8, bool) {
	v, ok := s.v.(T8)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set9(v T9) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get9() (T9, bool) {
	v, ok := s.v.(T9)
	return v, ok
}

func (s *Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set10(v T10) {
	s.v = v
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get10() (T10, bool) {
	v, ok := s.v.(T10)
	return v, ok
}

func (s Type[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Case(f0 func(T0), f1 func(T1), f2 func(T2), f3 func(T3), f4 func(T4), f5 func(T5), f6 func(T6), f7 func(T7), f8 func(T8), f9 func(T9), f10 func(T10)) {
	switch v := s.v.(type) {
	case T0:
		if f0 == nil {
			noHandler(v)
		}
		f0(v)
	case T1:
		if f1 == nil {
			noHandler(v)
		}
		f1(v)
	case T2:
		if f2 == nil {
			noHandler(v)
		}
		f2(v)
	case T3:
		if f3 == nil {
			noHandler(v)
		}
		f3(v)
	case T4:
		if f4 == nil {
			noHandler(v)
		}
		f4(v)
	case T5:
		if f5 == nil {
			noHandler(v)
		}
		f5(v)
	case T6:
		if f6 == nil {
			noHandler(v)
		}
		f6(v)
	case T7:
		if f7 == nil {
			noHandler(v)
		}
		f7(v)
	case T8:
		if f8 == nil {
			noHandler(v)
		}
		f8(v)
	case T9:
		if f9 == nil {
			noHandler(v)
		}
		f9(v)
	case T10:
		if f10 == nil {
			noHandler(v)
		}
		f10(v)
	default:
		panic("called Case on an invalid value")
	}
}

func noHandler(v any) {
	panic("no handler for case " + reflect.TypeOf(v).String())
}
