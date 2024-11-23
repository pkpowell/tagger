package tagger

import (
	"strings"
	"sync"
)

type TaggerMap struct {
	tags map[string]struct{}
	mtx  *sync.RWMutex
	min  int
}

// Creates a new tagger.
func New() *TaggerMap {
	return &TaggerMap{
		tags: make(map[string]struct{}),
		mtx:  new(sync.RWMutex),
		min:  3,
	}
}

// Add adds a tag to the tagger without any transforms.
// Best for UUIDs, ip addresses etc where you want to preserve punctuation
func (t *TaggerMap) AddExact(str string) {
	t.add(str)
}

// Add parses and adds a tag or its sub tags to the tagger.
// Best for regular text
func (t *TaggerMap) Add(str string) {
	var newTag string

	for _, newTag = range strings.Fields(Replacer.Replace(str)) {
		t.add(newTag)
	}
}

func (t *TaggerMap) add(newTag string) {
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

	if !known {
		t.tags[newTag] = struct{}{}
	}
}

// returns a slice of all tags.
func (t *TaggerMap) Get() []string {
	var tag string
	var tags = make([]string, 0)

	for tag = range t.tags {
		tags = append(tags, tag)
	}

	return tags
}

// returns a space delimited string of all tags.
func (t *TaggerMap) String() string {
	var tag string
	var tags = make([]string, 0)

	for tag = range t.tags {
		tags = append(tags, tag)
	}

	return strings.Join(tags, " ")
}
