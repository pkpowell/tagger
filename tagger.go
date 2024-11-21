package tagger

import (
	"strings"
	"sync"
)

type Tagger struct {
	tags map[string]struct{}
	mtx  *sync.RWMutex
}

func New() *Tagger {
	return &Tagger{
		tags: make(map[string]struct{}),
		mtx:  &sync.RWMutex{},
	}
}

func (t *Tagger) Add(newTag string, exact bool) {
	if len(newTag) < 2 {
		return
	}
	newTag = strings.ToLower(strings.TrimSpace(newTag))

	var found bool
	var tags []string

	t.mtx.Lock()
	defer t.mtx.Unlock()

	if exact {
		tags = []string{newTag}
	} else {
		tags = strings.Split(Replacer.Replace(newTag), " ")
	}

	for _, tag := range tags {
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
