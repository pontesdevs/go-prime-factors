package go_prime_factors_test

import (
	. "github.com/pontesdevs/go-prime-factors"
	"reflect"
	"testing"
)

func TestPrimeFactorsOf1IsEmptySlice(t *testing.T) {
	want := []int{}
	got := PrimeFactors(1)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PrimeFactorsOf1 should be %v but we got %v", want, got)
	}
}

func TestPrimeFactorsOf2Is2(t *testing.T) {
	want := []int{2}
	got := PrimeFactors(2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PrimeFactorsOf1 should be %v but we got %v", want, got)
	}
}
