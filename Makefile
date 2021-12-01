SHELL := /bin/bash

time.db: libfunctors.so
	-rm -f $@
	souffle -D . time.dl

libfunctors.so: functors.go
	go build -buildmode=c-shared -o $@ functors.go

run: libfunctors.so
	souffle -D . time.dl

clean:
	-rm -f *.db *.so *.h
