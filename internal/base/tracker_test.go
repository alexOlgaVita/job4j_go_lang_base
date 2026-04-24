package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("check link leak - there'is 1 element, changing 1 element", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t,
			[]base.Item{item},
			tracker.GetItems(),
		)
	})

	t.Run("check link leak - 2 elements, changing the second element", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)
		item2 := base.Item{
			ID:   "2",
			Name: "Second Item",
		}
		tracker.AddItem(item2)
		res := tracker.GetItems()
		res[1].ID = "2_2"
		res[1].Name = "Second_2 Item"

		assert.Equal(t,
			[]base.Item{item, item2},
			tracker.GetItems(),
		)
	})

	t.Run("check link leak - add empty element. no changing", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{}

		tracker.AddItem(item)

		assert.Equal(t,
			[]base.Item{item},
			tracker.GetItems(),
		)
	})

	t.Run("check link leak - changing empty element", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{}

		tracker.AddItem(item)

		res := tracker.GetItems()

		res[0].ID = "1 Item"
		res[0].Name = "First Item"

		assert.Equal(t,
			[]base.Item{item},
			tracker.GetItems(),
		)
	})
}
