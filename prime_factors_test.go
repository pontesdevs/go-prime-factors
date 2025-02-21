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
		t.Errorf("PrimeFactors should be %v but we got %v", want, got)
	}
}

func TestPrimeFactorsOf2Is2(t *testing.T) {
	want := []int{2}
	got := PrimeFactors(2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PrimeFactors should be %v but we got %v", want, got)
	}
}

func TestPrimeFactorsOf3Is3(t *testing.T) {
	want := []int{3}
	got := PrimeFactors(3)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PrimeFactors should be %v but we got %v", want, got)
	}
}

func TestPrimeFactorsOf4Is2_2(t *testing.T) {
	want := []int{2, 2}
	got := PrimeFactors(4)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PrimeFactors should be %v but we got %v", want, got)
	}
}

func TestPrimeFactorsOf6Is2_3(t *testing.T) {
	want := []int{2, 3}
	got := PrimeFactors(6)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PrimeFactors should be %v but we got %v", want, got)
	}
}

func TestPrimeFactorsOf8Is2_2_2(t *testing.T) {
	want := []int{2, 2, 2}
	got := PrimeFactors(8)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PrimeFactors should be %v but we got %v", want, got)
	}
}

func TestPrimeFactorsOf9Is3_3(t *testing.T) {
	want := []int{3, 3}
	got := PrimeFactors(9)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PrimeFactors should be %v but we got %v", want, got)
	}
}
