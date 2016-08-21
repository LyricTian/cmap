package cmap_test

import (
	"testing"

	"github.com/LyricTian/cmap"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMap(t *testing.T) {
	Convey("test all the map function", t, func() {
		m := cmap.NewMap()
		m.Set("foo", "bar")
		m.Set("foo1", "bar1")
		v, ok := m.Get("foo")
		So(v, ShouldEqual, "bar")
		So(ok, ShouldBeTrue)
		So(m.Count(), ShouldEqual, 2)
		So(m.Items(), ShouldContainKey, "foo")

		m.Remove("foo")
		So(m.Count(), ShouldEqual, 1)

		m.Clear()
		So(m.Count(), ShouldEqual, 0)
	})
}
