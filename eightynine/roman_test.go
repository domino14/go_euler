package eightynine

import "testing"

type romantest struct {
	num int
	rom string
}

var fromRomanTests = []romantest{
	{4672, "MMMMDCLXXII"},
	{1121, "MCXXI"},
	{1837, "MDCCCXXXVII"},
	{2367, "MMCCCLXVII"},
	{1435, "MCDXXXV"},
	{233, "CCXXXIII"},
	{920, "CMXX"},
	{3164, "MMMCLXIV"},
	{1386, "MCCCLXXXVI"},
	{898, "DCCCXCVIII"},
	{3934, "MMMDCCCCXXXIV"},
	{419, "CDXVIIII"},
	{2235, "MMCCXXXV"},
	{1832, "MDCCCXXXII"},
	{4500, "MMMMD"},
	{2769, "MMDCCLXIX"},
}

func TestFromRoman(t *testing.T) {
	for _, pair := range fromRomanTests {
		f := FromRoman(pair.rom)
		if f != pair.num {
			t.Error("For", pair.rom, "expected", pair.num, "got", f)
		}
	}
}

var toRomanTests = []romantest{
	{4672, "MMMMDCLXXII"},
	{1121, "MCXXI"},
	{1837, "MDCCCXXXVII"},
	{2367, "MMCCCLXVII"},
	{1435, "MCDXXXV"},
	{233, "CCXXXIII"},
	{920, "CMXX"},
	{3164, "MMMCLXIV"},
	{1386, "MCCCLXXXVI"},
	{898, "DCCCXCVIII"},
	{3934, "MMMCMXXXIV"},
	{419, "CDXIX"},
	{2235, "MMCCXXXV"},
	{1832, "MDCCCXXXII"},
	{4500, "MMMMD"},
	{2769, "MMDCCLXIX"},
}

func TestToRoman(t *testing.T) {
	for _, pair := range toRomanTests {
		s := ToRoman(pair.num)
		if s != pair.rom {
			t.Error("For", pair.num, "expected", pair.rom, "got", s)
		}
	}
}
