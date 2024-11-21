package tagger

import (
	"fmt"
	"testing"
)

func BenchmarkTags(b *testing.B) {
	tagger := New()

	for i := range b.N {
		s := fmt.Sprintf("tag1-%d", i)
		tagger.Add(s)
	}
	b.Log(tagger.Get())
}

func TestADD(t *testing.T) {
	tagger := New()
	tagger.Add("phil")
	tagger.Add("philip")
	tagger.Add("philip.powell")
	tagger.Add("phil.local")
	tagger.Add("ph")
	tagger.Add("0acd511a-4baa-5094-b68e-a330009c09e9")
	tagger.Add("0acd511a-4baa")

	t.Log(tagger.Get())
}

func BenchmarkTagsMisc(b *testing.B) {
	t := New()
	t.Add("philip.powell")
	t.Add("phil.local")
	// t.Add("phil", false)
	// t.Add("phil", false)
	// t.Add("    ", false)
	t.Add("phil")
	// t.Add("phil", true)
	t.Add("philip")
	t.Add("ph")
	t.Add("0acd511a-4baa-5094-b68e-a330009c09e9")
	t.Add("0acd511a-4baa-")
	// t.Add("4baa-5094-b68e-", false)
	b.Log(t.Get())
}
func BenchmarkTag(b *testing.B) {
	t := New()
	for i := range b.N {
		t.Add(fmt.Sprintf("tag1-%d", i))
	}
}
