package tagger

import (
	"fmt"
	"testing"
	"unique"
)

func BenchmarkTags(b *testing.B) {
	tagger := NewTags[*stringTag]()

	for i := range b.N {
		s := fmt.Sprintf("tag1-%d", i)
		tagger.Add(&stringTag{s})
	}
	b.Log(tagger.Get())
}
func BenchmarkTagsU(b *testing.B) {
	tagger := NewTags[*uniqueTag]()

	for i := range b.N {
		s := unique.Make(fmt.Sprintf("tag1-%d", i))
		tagger.Add(&uniqueTag{s})
	}
	// tagger.Get()
	b.Log(tagger.Get())
}
