# go test
## benchmark
### param
b *testing.B
### code
for i:=0;i<b.N;i++ {}
### instruction
test all 
#### func1
go test -bench .

#### output a cpuprofile
go test -bench . -cpuprofile cpu.out

#### pprof the file
go tool pprof cpu.out

#### choose one way to show
web 