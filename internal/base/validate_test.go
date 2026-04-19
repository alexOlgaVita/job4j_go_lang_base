package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("{UserID: 'e0eea7a7', Title: 'first', Description: 'Первый элемент'} - all fields are inizialized", func(t *testing.T) {
		t.Parallel()
		exp := make([]string, 0)

		in := &base.ValidateRequest{
			UserID:      "e0eea7a7",
			Title:       "first",
			Description: "Первый элемент",
		}
		rsl := base.Validate(in)

		assert.Equal(t, exp, rsl)
	})

	t.Run("{Title: 'first', Description: 'Первый элемент'} - field 'UserID' isn't inizialized", func(t *testing.T) {
		t.Parallel()
		exp := make([]string, 0)
		exp = append(exp, "Поле 'UserID' не инициализировано")

		in := &base.ValidateRequest{
			Title:       "first",
			Description: "Первый элемент",
		}
		rsl := base.Validate(in)

		assert.Equal(t, exp, rsl)
	})

	t.Run("{UserID: 'e0eea7a7', Description: 'Первый элемент'} - field 'Title' isn't inizialized", func(t *testing.T) {
		t.Parallel()
		exp := make([]string, 0)
		exp = append(exp, "Поле 'Title' не инициализировано")

		in := &base.ValidateRequest{
			UserID:      "e0eea7a7",
			Description: "Первый элемент",
		}
		rsl := base.Validate(in)

		assert.Equal(t, exp, rsl)
	})

	t.Run("{UserID: 'e0eea7a7', Description: 'Первый элемент'} - field 'Description' isn't inizialized", func(t *testing.T) {
		t.Parallel()
		exp := make([]string, 0)
		exp = append(exp, "Поле 'Description' не инициализировано")

		in := &base.ValidateRequest{
			UserID: "e0eea7a7",
			Title:  "first",
		}
		rsl := base.Validate(in)

		assert.Equal(t, exp, rsl)
	})

	t.Run("{} - all fields aren't inizialized", func(t *testing.T) {
		t.Parallel()
		exp := make([]string, 0)
		exp = append(exp, "Поле 'UserID' не инициализировано")
		exp = append(exp, "Поле 'Title' не инициализировано")
		exp = append(exp, "Поле 'Description' не инициализировано")

		in := &base.ValidateRequest{}
		rsl := base.Validate(in)

		assert.Equal(t, exp, rsl)
	})

	t.Run("{UserID: 'e0eea7a7'} - fields 'Title' and 'Description' aren't inizialized", func(t *testing.T) {
		t.Parallel()
		exp := make([]string, 0)
		exp = append(exp, "Поле 'Title' не инициализировано")
		exp = append(exp, "Поле 'Description' не инициализировано")

		in := &base.ValidateRequest{
			UserID: "e0eea7a7",
		}
		rsl := base.Validate(in)

		assert.Equal(t, exp, rsl)
	})

	t.Run("{Title: 'first'} - fields 'UserID' and 'Description' aren't inizialized", func(t *testing.T) {
		t.Parallel()
		exp := make([]string, 0)
		exp = append(exp, "Поле 'UserID' не инициализировано")
		exp = append(exp, "Поле 'Description' не инициализировано")

		in := &base.ValidateRequest{
			Title: "first",
		}
		rsl := base.Validate(in)

		assert.Equal(t, exp, rsl)
	})

	t.Run("{Description: 'Первый элемент'} - fields 'UserID' and 'Title' aren't inizialized", func(t *testing.T) {
		t.Parallel()
		exp := make([]string, 0)
		exp = append(exp, "Поле 'UserID' не инициализировано")
		exp = append(exp, "Поле 'Title' не инициализировано")

		in := &base.ValidateRequest{
			Description: "Первый элемент",
		}
		rsl := base.Validate(in)

		assert.Equal(t, exp, rsl)
	})
}
