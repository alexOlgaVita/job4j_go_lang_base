package tracker_test

import (
	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/tracker"
	"testing"
)

func Test_Domain(t *testing.T) {
	t.Parallel()

	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		tracker1 := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tracker1.UpdateItem(item)
		assert.ErrorIs(t, err, tracker.ErrNotFound)
	})

	t.Run("update - ok", func(t *testing.T) {
		t.Parallel()

		tracker1 := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		err := tracker1.AddItem(item)
		item = tracker.Item{
			ID:   "1",
			Name: "First Item really",
		}
		err = tracker1.UpdateItem(item)
		assert.ErrorIs(t, err, nil)
		assert.Equal(t, item.ID, "1")
		assert.Equal(t, item.Name, "First Item really")
	})

	t.Run("error add - already exists", func(t *testing.T) {
		t.Parallel()

		tracker1 := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tracker1.AddItem(item)
		item = tracker.Item{
			ID:   "1",
			Name: "Second Item",
		}
		err = tracker1.AddItem(item)
		assert.ErrorIs(t, err, tracker.ErrAlreadyExists)
	})

	t.Run("add - ok", func(t *testing.T) {
		t.Parallel()

		tracker1 := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tracker1.AddItem(item)
		assert.ErrorIs(t, err, nil)
		assert.Equal(t, item.ID, "1")
		assert.Equal(t, item.Name, "First Item")
	})
}
