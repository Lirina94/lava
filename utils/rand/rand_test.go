package rand

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func isUniformlyDistributed(t *testing.T, f func(n int64) int64) {
	t.Helper()

	const iterations = 100000
	const maxValue = 50

	numberToCount := make(map[int64]int64, maxValue)

	for i := 0; i < iterations; i++ {
		out := f(maxValue)
		numberToCount[out]++
	}

	target := iterations / maxValue
	const maxPercentDivergence = 0.07

	for i := 0; i < maxValue; i++ {
		count := numberToCount[int64(i)]
		require.InEpsilon(t, target, count, maxPercentDivergence, "slot "+strconv.Itoa(i))
	}
}

func TestDeterminism(t *testing.T) {
	// copied from the output of the loop below
	expected := []int64{
		34, 18, 63, 3, 16, 24, 81, 49,
		49, 65, 81, 66, 10, 89, 6, 85,
		71, 68, 89, 52, 7, 4, 3, 21,
		17, 90, 2, 36, 5, 4, 75, 93,
	}

	rng := New([]byte("pre-determined-data"))

	for i := 0; i < 32; i++ {
		random := rng.Int63n(100)
		require.Equal(t, expected[i], random, "item "+strconv.Itoa(i))
	}

	Seed(rng, []byte("other-determined-data"))

	count := 0
	for i := 0; i < 32; i++ {
		random := rng.Int63n(100)
		if expected[i] == random {
			count++
		}
	}
	require.Less(t, count, 32)
}

func TestDistribution(t *testing.T) {
	rng := New([]byte("pre-determined-data"))
	t.Run("new rand", func(t *testing.T) {
		isUniformlyDistributed(t, rng.Int63n)
	})

	Seed(rng, []byte("other-determined-data"))
	t.Run("new seed", func(t *testing.T) {
		isUniformlyDistributed(t, rng.Int63n)
	})
}
