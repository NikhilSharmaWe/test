package main

import "testing"

func BenchmarkComposeNameBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name.ComposeNameBinary()
	}
}

func BenchmarkComposeNameJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name.ComposeNameJSON()
	}
}

func BenchmarkParseNameBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseNameBinary(nameBytes)
	}
}

func BenchmarkParseNameJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseNameJSON(nameBytes)
	}
}
