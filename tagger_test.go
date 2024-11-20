package tagger

import (
	"fmt"
	"testing"
)

func BenchmarkTags(b *testing.B) {
	tagger := NewTags()

	for i := range b.N {
		s := fmt.Sprintf("tag1-%d", i)
		tagger.Add(s)
	}
	b.Log(tagger.Get())
}

func BenchmarkTagsNormal(b *testing.B) {
	t := NewTags()
	t.Add("phil.local")
	t.Add("phil")
	t.Add("phil")
	t.Add("phil")
	t.Add("phil")
	t.Add("philip")
	t.Add("philip.powell")
	t.Add("0acd511a-4baa-5094-b68e-a330009c09e9")
	t.Add("0acd511a-4baa-")
	b.Log(t.Get())
}
