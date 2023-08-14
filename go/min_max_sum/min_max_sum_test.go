package main

import "testing"

func TestMinMaxSumCase1(t *testing.T) {
	input := [5]int32{1, 3, 5, 7, 9}
	expected := "16 24"
	got := minMaxSum(input)
	if expected != got {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}

func TestMinMaxSumCase2(t *testing.T) {
	input := [5]int32{1, 2, 3, 4, 5}
	expected := "10 14"
	got := minMaxSum(input)
	if expected != got {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}

func TestMinMinBig(t *testing.T) {
	input := [5]int32{256741038, 623958417, 467905213, 714532089, 938071625}
	expected := "2063136757 2744467344"
	got := minMaxSum(input)
	if expected != got {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}

func TestMinMaxSame(t *testing.T) {
	input := [5]int32{5, 5, 5, 5, 5}
	expected := "20 20"
	got := minMaxSum(input)
	if expected != got {
		t.Errorf("expected: %s, got %s", expected, got)
	}
}
