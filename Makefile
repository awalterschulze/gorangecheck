report:
	make clean
	make bench > benchmarks.txt
	make asm
	cat benchmarks.txt

clean:
	(cd onebyone && rm main.s || true)
	(cd withbinary && rm main.s || true)
	(cd withunsafe && rm main.s || true)
	(cd withhint && rm main.s || true)
	(cd betterhint && rm main.s || true)
	(cd onebyone && rm main || true)
	(cd withbinary && rm main || true)
	(cd withunsafe && rm main || true)
	(cd withhint && rm main || true)
	(cd betterhint && rm main || true)

asm:
	(cd onebyone && go build main.go && go tool objdump -S -s encodeFixed64 ./main > main.s)
	(cd withbinary && go build main.go && go tool objdump -S -s encodeFixed64 ./main > main.s)
	(cd withunsafe && go build main.go && go tool objdump -S -s encodeFixed64 ./main > main.s)
	(cd withhint && go build main.go && go tool objdump -S -s encodeFixed64 ./main > main.s)
	(cd betterhint && go build main.go && go tool objdump -S -s encodeFixed64 ./main > main.s)

bench:
	(cd onebyone && go test -bench=.)
	(cd withbinary && go test -bench=.)
	(cd withunsafe && go test -bench=.)
	(cd withhint && go test -bench=.)
	(cd betterhint && go test -bench=.)

