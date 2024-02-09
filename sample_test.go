package dbwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SampleFunc(t *testing.T) {
	t.Parallel()

	t.Run("sample test", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, 5, SampleFunc())
	})
}
