package tagger

import (
	"fmt"
	"testing"
)

func BenchmarkTags(b *testing.B) {
	tagger := New()

	for i := range b.N {
		s := fmt.Sprintf("tag1-%d", i)
		tagger.Add(s, true)
	}
	b.Log(tagger.Get())
}

func BenchmarkTagsMisc(b *testing.B) {
	t := New()
	t.Add("phil.local", false)
	t.Add("phil", false)
	t.Add("phil", false)
	t.Add("phil", false)
	t.Add("phil", false)
	t.Add("ph", false)
	t.Add("philip", false)
	t.Add("philip.powell", true)
	t.Add("0acd511a-4baa-5094-b68e-a330009c09e9", true)
	t.Add("0acd511a-4baa-", false)
	t.Add("4baa-5094-b68e-", false)
	b.Log(t.Get())
}
func BenchmarkTag(b *testing.B) {
	t := New()
	for i := range b.N {
		t.Add(fmt.Sprintf("tag1-%d", i), true)
	}
}
