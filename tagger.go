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

// type TaggerSlice struct {
// 	tags []string
// 	mtx  *sync.RWMutex
// 	min  int
// }

func New() *TaggerMap {
	return &TaggerMap{
		tags: make(map[string]struct{}),
		mtx:  &sync.RWMutex{},
		min:  3,
	}
}

// func NewSlice() *TaggerSlice {
// 	return &TaggerSlice{
// 		tags: make([]string, 0),
// 		mtx:  &sync.RWMutex{},
// 		min:  3,
// 	}
// }

// Add adds a tag to the tagger (without any transforms).
func (t *TaggerMap) AddExact(str string) {
	t.add(str)
}

// // Add adds a tag to the tagger (without any transforms).
// func (t *TaggerSlice) AddExact(str string) {
// 	t.add(str)
// }

// Add parses and adds a tag (or multiple sub tags) to the tagger.
func (t *TaggerMap) Add(str string) {
	var newTag string

	for _, newTag = range strings.Fields(Replacer.Replace(str)) {
		t.add(newTag)
	}
}

// // Add parses and adds a tag (or multiple sub tags) to the tagger.
// func (t *TaggerSlice) Add(str string) {
// 	var newTag string

// 	for _, newTag = range strings.Fields(Replacer.Replace(str)) {
// 		t.add(newTag)
// 	}
// }

// func remove(slice []string, s int) []string {
// 	return append(slice[:s], slice[s+1:]...)
// }

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

// func (t *TaggerSlice) add(newTag string) {
// 	if len(newTag) < t.min {
// 		return
// 	}

// 	var known bool
// 	var knownTag string
// 	var i int
// 	newTag = strings.ToLower(newTag)
// 	var newLen = len(newTag)

// 	t.mtx.Lock()
// 	defer t.mtx.Unlock()

// 	for i, knownTag = range t.tags {
// 		known = false

// 		if len(knownTag) >= newLen {
// 			// if knownTag >= newTag check if knownTag contains newTag
// 			// and if so, ignore newTag
// 			if strings.Contains(knownTag, newTag) {
// 				known = true
// 				break
// 			}

// 		} else {
// 			// else check if newTag contains knownTag
// 			// if so delete knownTag and add newTag
// 			if strings.Contains(newTag, knownTag) {
// 				remove(t.tags, i)
// 				break
// 			}
// 		}
// 	}

// 	if !known {
// 		t.tags = append(t.tags, newTag)
// 	}
// }

func (t *TaggerMap) Get() []string {
	var tag string
	var tags = make([]string, 0)

	for tag = range t.tags {
		tags = append(tags, tag)
	}

	return tags
}

// func (t *TaggerSlice) Get() []string {
// 	return t.tags
// }

func (t *TaggerMap) String() string {
	var tag string
	var tags = make([]string, 0)

	for tag = range t.tags {
		tags = append(tags, tag)
	}

	return strings.Join(tags, " ")
}

// func (t *TaggerSlice) String() string {
// 	return strings.Join(t.tags, " ")
// }
