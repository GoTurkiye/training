# Benchmark

* Benchmark
    * Benchmarkları çalıştırır ve çıktı olarak b.ReportAllocs() yoksa memory ve allocation çıktısı vermez
        * go test -bench .
    * Benchmark çıktısını memory ile beraber oluşturur aynı zamanda testin içinde b.ReportAllocs() ilede bunu elde edebiliriz
        * go test -bench . -benchmem
    * BenchmarkGzipWithoutPoll-{x} x değeri işlemcide ki çekirdek sayısını gösterir -cpu ile bunu ayarlayabiliri
        * go test -bench . -cpu 2
    * -benchtime ile benchmarkların çalışma süresini ya da sayısını ayarlayabiliriz. 3s ile süreyi 100x ile sayıyı belirtebiliriz
        * go test -bench . -benchtime 3s
        * go test -bench . -benchtime 100x
    * Testlerde default timeout kapalıdır fakat uzun sürmesinin önüne geçmek için bunu açabiliriz
        * go test -bench . -timeout 1s
    * Compiler optimizasyonlarını kapatarak daha doğru test sonucu elde edebiliriz
        * go test -bench . -gcflags='-N -l'
    * Testleri -count ile tekrarlı çalıştırabiliriz
        * go test -bench . -count 10  -gcflags='-N -l'
    * Testlerimizi compile edip tekrarlı çalıştırabiliriz
        * go test -c -o gzip_bench.test
    * Benchstat ile ortalamayı bulabiliriz
        * ./gzip_bench.test -test.bench . -test.benchmem -test.count 10 > result.txt
        * bencstat result.txt (go install golang.org/x/perf/cmd/benchstat@latest komutu ile benchstati bilgisayariniza indirebilirsiniz)
    * Aynı zamanda bunları karşılaştırabiliriz
        * benchstat old.txt new.txt
* PProf 
    * Benchmarklarımızdan profile alarak pprofu çalıştırabiliriz
        * ./gzip_bench.test -test.bench . -test.benchmem -test.count 10 -test.memprofile=mem.profile -test.cpuprofile=cpu.profile
    * Profilelarımızla pprofu çalıştırabilir ve analiz edebiliriz
        * go tool proof mem.profile
        * go tool pprof cpu.profile
    * İki farklı profile ı karşılaştırabiliriz
        *  go tool pprof -base old-mem.profile new-mem.profile  
    * list ile karşılaştırma sonucunu metinsel görebiliriz list arguman olarak regex alır
        * (pprof) list Zip
    * web ile karşılaştırma sonucunu görebiliriz, iyileşmeler yeşil, kötüleşmeler kırmızı
        * (pprof) web
    * Binary vererek disassembly elde edebiliriz
        * go tool pprof gzip_bench.test mem.profile
    * top 10 ile en çok kaynak harcanan 10 bloğu görebiliriz
        * (pprof) top 10
    * peek ile regex match olan yerdeki en çok kaynak kullanan yeri görebiliriz
        * (pprof) peek Zip
    * weblist ile list komutunun grafiksel çıktısını elde edebiliriz
        * (pprof) weblist Zip
* Trace
    * Benchmark testlerinden trace ciktisi alabiliriz
        * go test -bench . -trace=trace.out