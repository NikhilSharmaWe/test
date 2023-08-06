package main

import (
	"testing"
)

func BenchmarkComposeNameBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name.ComposeNameBinary()
	}
}

func BenchmarkParseNameBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseNameBinary(nameBytes)
	}
}

func BenchmarkComposeNameJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name.ComposeNameJSON()
	}
}

func BenchmarkParseNameJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseNameJSON(nameBytes)
	}
}

func BenchmarkComposeNamePB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ComposeNamePB(naam)
	}
}

func BenchmarkParseNamePB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseNamePB(nameBytes)
	}
}
