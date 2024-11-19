package tagger

import (
	"strings"
	"unique"
)

type tagType interface {
	string | unique.Handle[string]

	Val() string
}

type Tagger[T tagType] struct {
	tags map[T]struct{}
}

func NewTags[T tagType]() *Tagger[T] {
	return &Tagger[T]{
		tags: make(map[T]struct{}),
	}
}

func (t *Tagger[T]) Add(tag T) {
	// tag := unique.Make(strings.ToLower(s))
	t.tags[tag] = struct{}{}
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
