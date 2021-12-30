package solutions

import "testing"

func TestSonarDepth(t *testing.T) {
	got := NumberOfIncreasesInIntSlice([]int{10, 20, 30, 5, 10})
	if got != 3 {
		t.Errorf("NumberOfIncreasesInIntSlice() got %d, want %d", got, 3)
	}
}

func TestSlidingSonarDepth(t *testing.T) {
	got := SlidingNumberOfIncreasesInIntSlice([]int{607, 618, 618, 617, 647, 716, 769, 792})
	want := 5

	if got != want {
		t.Errorf("SlidingNumberOfIncreasesInIntSlice() got %d, want %d", got, want)
	}
}
