package tagger

import (
	"fmt"
	"strings"
	"sync"
)

type Tagger struct {
	tags map[string]struct{}
	mtx  *sync.RWMutex
	min  int
}

func New() *Tagger {
	return &Tagger{
		tags: make(map[string]struct{}),
		mtx:  &sync.RWMutex{},
		min:  3,
	}
}

// Add adds a tag to the tagger.
func (t *Tagger) AddExact(newTag string) {
	if len(newTag) < t.min {
		return
	}
	t.add(newTag)
}

// Add adds a tag to the tagger.
func (t *Tagger) Add(newTags string) {
	if len(newTags) < t.min {
		return
	}

	tags := strings.Fields(Replacer.Replace(newTags))
	for _, newTag := range tags {

		t.add(newTag)
	}
}

func (t *Tagger) add(newTag string) {
	newTag = strings.ToLower(newTag)
	newLen := len(newTag)
	var found bool

	fmt.Println("considering tag:", newTag)

	t.mtx.Lock()
	defer t.mtx.Unlock()

	for knownTag := range t.tags {
		found = false
		if len(knownTag) <= newLen {
			if strings.Contains(newTag, knownTag) {
				found = true
				break
			}
		} else {
			if strings.Contains(knownTag, newTag) {
				found = true
				break
			}
		}
	}

	if !found {
		fmt.Println("adding tag:", newTag)
		t.tags[newTag] = struct{}{}
	} else {
		fmt.Println("not adding tag:", newTag)
	}
}

func (t *Tagger) Get() []string {
	tags := make([]string, len(t.tags))
	for tag := range t.tags {
		tags = append(tags, tag)
	}
	return tags
}

func (t *Tagger) String() string {
	tags := make([]string, len(t.tags))
	for tag := range t.tags {
		tags = append(tags, tag)
	}
	return strings.Join(tags, " ")
}
