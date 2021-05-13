package remote__test

import (
	mm "github.com/easierway/concurrent_map"
	cmap "github.com/orcaman/concurrent-map"
	"testing"
	//"github.com/orcaman/concurrent-map"
)
func TestConcurrentMap(t *testing.T) {
	m := mm.CreateConcurrentMap(99)
	m.Set(mm.StrKey("key"), 109)
	t.Log(m.Get(mm.StrKey("key")))

	concurrentMap := cmap.New()
	concurrentMap.Set("foo", "bar")
	if value, status := concurrentMap.Get("foo"); status {
		bar := value.(string)
		t.Log(bar)
	}
	concurrentMap.Remove("foo")
}