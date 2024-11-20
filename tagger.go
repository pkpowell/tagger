package tagger

import (
	"strings"
	"unique"
)

type stringTag struct{ string }
type uniqueTag struct{ unique.Handle[string] }

type tagType interface {
	*stringTag | *uniqueTag

	Val() string
}

func (s *stringTag) Val() string {
	return s.string
}

func (s *uniqueTag) Val() string {
	return s.Value()
}

type Tagger[T tagType] struct {
	tags map[T]struct{}
}

func NewTags[T tagType]() *Tagger[T] {
	return &Tagger[T]{
		tags: make(map[T]struct{}),
	}
}

func (t *Tagger[T]) Add(s T) {
	// var tag T
	// tag := &T{strings.ToLower(s)}
	// tag := unique.Make(strings.ToLower(s))
	t.tags[s] = struct{}{}
}

func (t *Tagger[T]) Get() (tags []string) {
	for tag := range t.tags {
		tags = append(tags, tag.Val())
	}
	return
}

func (t *Tagger[T]) String() string {
	var tags []string
	for tag := range t.tags {
		tags = append(tags, tag.Val())
	}
	return strings.Join(tags, " ")
}
