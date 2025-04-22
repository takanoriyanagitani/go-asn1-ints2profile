package ints2prof

import (
	"bufio"
	"encoding/asn1"
	"io"
	"iter"
	"os"
	"strconv"
)

type IntegerProfile struct {
	Max int64
	Min int64
	Sum int64
	Cnt int64
}

func (i IntegerProfile) ToDerBytes() ([]byte, error) {
	return asn1.Marshal(i)
}

func (i IntegerProfile) DerToWriter(w io.Writer) error {
	der, e := i.ToDerBytes()
	if nil != e {
		return e
	}
	_, e = w.Write(der)
	return e
}

func (i IntegerProfile) DerToStdout() error {
	return i.DerToWriter(os.Stdout)
}

type Strings iter.Seq[string]

func (s Strings) ToIntegers() iter.Seq2[int, error] {
	return func(yield func(int, error) bool) {
		for line := range s {
			i, e := strconv.Atoi(line)
			if !yield(i, e) {
				return
			}
		}
	}
}

type Integers iter.Seq2[int, error]

func (i Integers) ToStat() (IntegerProfile, error) {
	var stat IntegerProfile
	for item, e := range i {
		if nil != e {
			return stat, e
		}

		if 0 == stat.Cnt {
			stat.Max = int64(item)
			stat.Min = int64(item)
		}

		stat.Cnt += 1
		stat.Sum += int64(item)

		stat.Max = max(int64(item), stat.Max)
		stat.Min = min(int64(item), stat.Min)
	}
	return stat, nil
}

func StdinToStrings() Strings {
	return func(yield func(string) bool) {
		var s *bufio.Scanner = bufio.NewScanner(os.Stdin)
		for s.Scan() {
			if !yield(s.Text()) {
				return
			}
		}
	}
}

func StdinToIntegersToStatsToDerToStdout() error {
	var lines Strings = StdinToStrings()
	var ints Integers = Integers(lines.ToIntegers())
	prof, e := ints.ToStat()
	if nil != e {
		return e
	}
	return prof.DerToStdout()
}
