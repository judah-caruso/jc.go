package array

import (
	"iter"
)

const DefaultElementsPerBucket = 32

// Stable is a resizable array whose values will never move in memory.
// This means it is safe to take a pointer to a value within the array
// while continuing to append to it.
type Stable[T any] struct {
	buckets           []bucket[T]
	last              int
	elementsPerBucket int
}

func (s *Stable[T]) Init() {
	s.InitWithSize(DefaultElementsPerBucket)
}

func (s *Stable[T]) InitWithSize(elementsPerBucket int) {
	if elementsPerBucket <= 0 {
		elementsPerBucket = DefaultElementsPerBucket
	}

	s.elementsPerBucket = elementsPerBucket

	s.buckets = s.buckets[:0]
	s.buckets = append(s.buckets, make(bucket[T], 0, s.elementsPerBucket))
	s.last = 0
}

func (s *Stable[T]) Append(value T) *T {
	if len(s.buckets) == 0 {
		s.Init()
	}

	if len(s.buckets[s.last]) == cap(s.buckets[s.last]) {
		s.buckets = append(s.buckets, make(bucket[T], 0, s.elementsPerBucket))
		s.last += 1
	}

	s.buckets[s.last] = append(s.buckets[s.last], value)
	return &s.buckets[s.last][len(s.buckets[s.last])-1]
}

func (s *Stable[T]) AppendMany(values ...T) (first *T) {
	if len(values) == 0 {
		return nil
	}

	first = s.Append(values[0])

	if len(values) > 1 {
		for _, v := range values[1:] {
			s.Append(v)
		}
	}

	return
}

func (s *Stable[T]) Get(index int) *T {
	bucket := s.buckets[index/s.elementsPerBucket]
	return &bucket[index%s.elementsPerBucket]
}

func (s *Stable[T]) Len() int {
	return s.Cap() - (cap(s.buckets[s.last]) - len(s.buckets[s.last]))
}

func (s *Stable[T]) Cap() int {
	return len(s.buckets) * s.elementsPerBucket
}

func (s *Stable[T]) Pointers() iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		for bi := range s.buckets {
			startIdx := bi * s.elementsPerBucket
			for i := range s.buckets[bi] {
				if !yield(startIdx+i, &s.buckets[bi][i]) {
					return
				}
			}
		}
	}
}

func (s *Stable[T]) Values() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for bi, bucket := range s.buckets {
			startIdx := bi * s.elementsPerBucket
			for i := range bucket {
				if !yield(startIdx+i, bucket[i]) {
					return
				}
			}
		}
	}
}

type bucket[T any] = []T
