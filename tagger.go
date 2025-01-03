package tagger

import (
	"strings"
	"sync"
)

type Tagger struct {
	tags map[string]struct{}
	mtx  *sync.RWMutex
	min  int
}

// New returns a new tagger
func New() *Tagger {
	return &Tagger{
		tags: make(map[string]struct{}),
		mtx:  new(sync.RWMutex),
		min:  3,
	}
}

// AddExact adds taga to the tagger without any transforms.
// Best for UUIDs, ip addresses etc
func (t *Tagger) AddExact(str ...string) {
	for _, s := range str {
		t.add(s)
	}
}

// Add parses and adds tags to the tagger while removing spaces and punctuation.
// Best for regular text
func (t *Tagger) Add(str ...string) {
	var newTag string
	for _, s := range str {
		for _, newTag = range strings.Fields(Replacer.Replace(s)) {
			t.add(newTag)
		}
	}
}

func (t *Tagger) add(newTag string) {
	if len(newTag) < t.min {
		return
	}

	t.mtx.Lock()
	defer t.mtx.Unlock()

	var isKnown = false
	var knownTag string

	newTag = strings.ToLower(newTag)
	var newLen = len(newTag)

	for knownTag = range t.tags {
		isKnown = false

		if len(knownTag) >= newLen {
			// if knownTag >= newTag check if knownTag contains newTag
			// and if so, ignore newTag
			if strings.Contains(knownTag, newTag) {
				isKnown = true
				break
			}

		} else {
			// else check if newTag contains knownTag
			// if so delete knownTag and add newTag
			if strings.Contains(newTag, knownTag) {
				delete(t.tags, knownTag)
				break
			}
		}
	}

	// if newTag not known, add it
	if !isKnown {
		t.tags[newTag] = struct{}{}
	}
}

// Get returns tags as a slice
func (t *Tagger) Get() []string {
	var tag string
	var tags = make([]string, 0)

	for tag = range t.tags {
		tags = append(tags, tag)
	}

	return tags
}

// String returns tags as a space delimited string
func (t *Tagger) String() string {
	return strings.Join(t.Get(), " ")
}
