# test

### encoding/binary - slowest
```
BenchmarkComposeNameBinary-8   	 5889136	       180.9 ns/op	     208 B/op	      11 allocs/op
BenchmarkParseNameBinary-8   	 4058863	       272.6 ns/op	     160 B/op	      15 allocs/op
```

### encoding/json - faster
```
BenchmarkComposeNameJSON-8   	 6511203	       165.6 ns/op	     160 B/op	       3 allocs/op
BenchmarkParseNameJSON-8   	 4152084	       263.6 ns/op	    1032 B/op	       9 allocs/op
```

### google.golang.org/protobuf/proto - fastest
```
BenchmarkComposeNamePB-8   	10318105	       103.0 ns/op	     104 B/op	       2 allocs/op
BenchmarkParseNamePB-8   	16157970	        62.19 ns/op	      80 B/op	       1 allocs/op
```
