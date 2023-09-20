## How to run Benchmark

* Run this command for Atomic Example.

`$go test -benchmem -run=^$ -bench ^BenchmarkAtomic$`

* Run this command for Mutex Lock Example.

`$go test -benchmem -run=^$ -bench ^BenchmarkMutexLock$`

### Example Results

* **Atomic**
  * goos: darwin
  * goarch: amd64
  * pkg: raceConditionPointersFix
  * cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
  * BenchmarkAtomic-4 --- 4977807 --- 236.2 ns/op --- 24 B/op --- 1 allocs/op
  * PASS
  * ok raceConditionPointersFix 1.895s


* **Mutex Lock**
  * goos: darwin
  * goarch: amd64
  * pkg: raceConditionPointersFix
  * cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
  * BenchmarkMutexLock-4 --- 4626063 --- 260.3 ns/op --- 33 B/op          1 allocs/op
  * PASS
  * ok raceConditionPointersFix 1.592s

