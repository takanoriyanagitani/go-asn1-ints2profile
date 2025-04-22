#!/bin/sh

input(){
	seq 1 100
}

input |
	./ints2profile |
	xxd -ps |
	tr -d '\n' |
	python3 \
		-m asn1tools \
		convert \
		-i der \
		-o jer \
		./iprof.asn \
		IntegerProfile \
		-
