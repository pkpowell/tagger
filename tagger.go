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

// Add adds a tag to the tagger.
func (t *Tagger) Add(newTags string, exact bool) {
	if len(newTags) < 2 {
		return
	}
	newTags = strings.ToLower(newTags)

	var found bool
	var tags []string

	t.mtx.Lock()
	defer t.mtx.Unlock()

	if exact {
		tags = []string{strings.TrimSpace(newTags)}
	} else {
		tags = strings.Split(strings.TrimSpace(Replacer.Replace(newTags)), " ")
	}

	for _, newTag := range tags {
		found = false
		for knownTag := range t.tags {
			if len(knownTag) > len(newTag) {
				if strings.Contains(knownTag, newTag) {
					found = true
					break
				}
			} else {
				if strings.Contains(newTag, knownTag) {
					found = true
					break
				}
			}
		}

		if !found {
			t.tags[newTag] = struct{}{}
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
