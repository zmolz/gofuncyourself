package list_test

import (
	"math/rand"
	"testing"
)

const (
	maxLevel = 32
	pPercent = 25
	pScaled  = uint32(pPercent * 65536 / 100)
)

func randomLevelFloat(r *rand.Rand) int {
	level := 1
	for r.Float64() < float64(pPercent)/100 && level < maxLevel {
		level++
	}
	return level
}

func randomLevelBitmask(r *rand.Rand) int {
	level := 1
	for (r.Uint32() & 0xFFFF) < pScaled && level < maxLevel {
		level++
	}
	return level
}

func BenchmarkRandomLevel(b *testing.B) {
	r := rand.New(rand.NewSource(42))

	b.Run("Float", func(b *testing.B) {
		for b.Loop() {
			_ = randomLevelFloat(r)
		}
	})

	b.Run("Bitmask", func(b *testing.B) {
		for b.Loop() {
			_ = randomLevelBitmask(r)
		}
	})
}