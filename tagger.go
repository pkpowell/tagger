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

// AddExact adds a tag to the tagger without any transforms.
// Best for UUIDs, ip addresses etc where you want to preserve punctuation
func (t *Tagger) AddExact(str string) {
	t.add(str)
}

// Add parses and adds a tag or its sub tags to the tagger.
// Best for regular text
func (t *Tagger) Add(str string) {
	var newTag string

	for _, newTag = range strings.Fields(Replacer.Replace(str)) {
		t.add(newTag)
	}
}

func (t *Tagger) add(newTag string) {
	if len(newTag) < t.min {
		return
	}

	var known bool
	var knownTag string

	newTag = strings.ToLower(newTag)
	var newLen = len(newTag)

	t.mtx.Lock()
	defer t.mtx.Unlock()

	for knownTag = range t.tags {
		known = false

		if len(knownTag) >= newLen {
			// if knownTag >= newTag check if knownTag contains newTag
			// and if so, ignore newTag
			if strings.Contains(knownTag, newTag) {
				known = true
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
	if !known {
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
