package tagger

import (
	"strings"
)

type Tagger struct {
	tags map[string]struct{}
}

func NewTags() *Tagger {
	return &Tagger{
		tags: make(map[string]struct{}),
	}
}

func (t *Tagger) Add(tag string) {
	var found bool
	tags := Replacer.Replace(tag)
	for _, tag := range strings.Split(tags, " ") {
		for t := range t.tags {
			if len(t) > len(tag) {
				if strings.Contains(t, tag) {
					found = true
					break
				}
			} else {
				if strings.Contains(tag, t) {
					found = true
					break
				}
			}
		}

		if !found {
			t.tags[tag] = struct{}{}
		}
	}
}

func (t *Tagger) Get() (tags []string) {
	for tag := range t.tags {
		tags = append(tags, tag)
	}
	return
}

func (t *Tagger) String() string {
	var tags []string
	for tag := range t.tags {
		tags = append(tags, tag)
	}
	return strings.Join(tags, " ")
}
