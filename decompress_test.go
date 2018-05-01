package decompress

import (
  "testing"
)

func TestDecompressSimpleString(t *testing.T) {
  decompressed, _ := Decompress([]byte("abc"), false)

  if string(decompressed) != "abc" {
    t.Fail()
  }

}

func TestDecompressSingleCompression(t *testing.T) {
  decompressed, _ := Decompress([]byte("3[abc]"), false)

  if string(decompressed) != "abcabcabc" {
	t.Fail()
  }

}

func TestDecompressSingleCompressionPlusSimpleString(t *testing.T) {
  decompressed, _ := Decompress([]byte("3[abc]xyz"), false)

  if string(decompressed) != "abcabcabcxyz" {
	t.Fail()
  }

}

func TestDecompresshMultipleLinearCompression(t *testing.T) {
  decompressed, _ := Decompress([]byte("3[abc]2[efg]"), false)

  if string(decompressed) != "abcabcabcefgefg" {
	t.Fail()
  }

}

func TestDecompressMultipleLinearCompressionWithSimpleStringBetween(t *testing.T) {
  decompressed, _ := Decompress([]byte("3[abc]gosh2[efg]"), false)

  if string(decompressed) != "abcabcabcgoshefgefg" {
	t.Fail()
  }

}

func TestDecompressSimpleNested(t *testing.T) {
  decompressed, _ := Decompress([]byte("3[2[a]]"), false)

  if string(decompressed) != "aaaaaa" {
	t.Fail()
  }

}

func TestDecompressMultipleNested(t *testing.T) {
  decompressed, _ := Decompress([]byte("3[2[3[a]]]"), false)

  if string(decompressed) != "aaaaaaaaaaaaaaaaaa" {
	t.Fail()
  }

}

func TestDecompressSimpleNestedAndLinearString(t *testing.T) {
  decompressed, _ := Decompress([]byte("3[2[a]c]"), false)

  if string(decompressed) != "aacaacaac" {
	t.Fail()
  }

}

func TestDecompressSimpleNestedAndMultipleLinear(t *testing.T) {
  decompressed, _ := Decompress([]byte("3[2[a]2[c]]"), false)

  if string(decompressed) != "aaccaaccaacc" {
	t.Fail()
  }

}