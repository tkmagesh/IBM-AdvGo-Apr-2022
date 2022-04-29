package utils

import "testing"

func Benchmark_GeneratePrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePrimes(3, 100)
	}
}

func Benchmark_IsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(97)
	}
}

func Benchmark_IsPrime_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime_2(97)
	}
}

func Benchmark_IsPrime_2_A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime_2_A(97)
	}
}

func Benchmark_IsPrime_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime_3(97)
	}
}
