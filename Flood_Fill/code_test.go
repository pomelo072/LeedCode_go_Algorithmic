package Flood_Fill

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloodFillBFS(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr, sc, newColor := 1, 1, 2
	v := FloodFillBFS(image, sr, sc, newColor)
	if assert.Equal(t, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}, v) {
		t.Log(v)
	} else {
		t.Log(v)
		t.Error("Error.")
	}
}

func TestFloodFillDFS(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr, sc, newColor := 1, 1, 2
	v := FloodFillDFS(image, sr, sc, newColor)
	if assert.Equal(t, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}, v) {
		t.Log(v)
	} else {
		t.Log(v)
		t.Error("Error.")
	}
}

func BenchmarkFloodFillBFS(b *testing.B) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr, sc, newColor := 1, 1, 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FloodFillBFS(image, sr, sc, newColor)
	}
	b.StopTimer()
}

func BenchmarkFloodFillDFS(b *testing.B) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr, sc, newColor := 1, 1, 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FloodFillDFS(image, sr, sc, newColor)
	}
	b.StopTimer()
}
