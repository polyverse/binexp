package binexp

import (
	"github.com/polyverse/binexp/syntax"
	"testing"
)

func TestRegexBinaryMatchBasic(t *testing.T) {

	// Ensure this is a non-UTF-8 compliant string
	opcode, err := Compile("\x65\xff\x15", None)
	if err != nil {
		t.Error(err)
	}

	data := string([]byte{0x65, 0xff, 0x15, 0xef, 0x65, 0x15, 0xcd, 0x50, 0x65, 0xff, 0x15, 0x25})

	match, err := opcode.FindStringMatchStartingAt(data, 0)
	if err != nil {
		t.Error(err)
	}
	if match.Index != 0 {
		t.Fatalf("First match expected at index 0. Instead found it at: %d", match.Index)
	}
	if match.Length != 3 {
		t.Fatalf("First match length expected to be 3. Instead found it at: %d", match.Length)
	}

	match, err = opcode.FindStringMatchStartingAt(data, 1)
	if err != nil {
		t.Error(err)
	}
	if match.Index != 8 {
		t.Fatalf("Second match (starting from 1) expected at index 8. Instead found it at: %d", match.Index)
	}
	if match.Length != 3 {
		t.Fatalf("Second match (starting from 1) length expected to be 3. Instead found it at: %d", match.Length)
	}

	match, err = opcode.FindStringMatchStartingAt(data, 8)
	if err != nil {
		t.Error(err)
	}
	if match.Index != 8 {
		t.Fatalf("Second match (starting from 8) expected at index 8. Instead found it at: %d", match.Index)
	}
	if match.Length != 3 {
		t.Fatalf("First match (starting from 8) length expected to be 3. Instead found it at: %d", match.Length)
	}

	match, err = opcode.FindStringMatchStartingAt(data, 9)
	if err != nil {
		t.Error(err)
	}
	if match != nil {
		t.Fatalf("No matches were expected after index 9. Found a match.")
	}
}

func TestEnsureRegexStringMatchFail(t *testing.T) {

	// Ensure this is a non-UTF-8 compliant string
	opcode, err := Compile("\xca[\x00-\xff]{2}", None)
	if err != nil {
		t.Error(err)
	}

	rawdata := []byte{0x65, 0xca, 0x05, 0xf4, 0x65, 0xca, 0xaf, 0xca, 0x65, 0xff, 0x15, 0x25}
	data := string(rawdata)
	match, err := opcode.FindStringMatchStartingAt(data, 1)
	if err != nil {
		t.Error(err)
	}

	match, err = opcode.FindBytesMatchStartingAt(rawdata, 1)
	if err != nil {
		t.Error(err)
	}
	if match != nil {
		t.Errorf("Did NOT expect a match")
	}

}

func TestRegexByteMatchSuccess(t *testing.T) {

	// Ensure this is a non-UTF-8 compliant string
	opcode, err := Compile("\xca[\x00-\xff]{2}", syntax.ByteRunes)
	if err != nil {
		t.Error(err)
	}

	rawdata := []byte{0x65, 0xca, 0x05, 0xf4, 0x65, 0xca, 0xaf, 0xca, 0x65, 0xff, 0x15, 0x25}
	data := string(rawdata)
	match, err := opcode.FindStringMatchStartingAt(data, 1)
	if err != nil {
		t.Error(err)
	}
	if match != nil {
		t.Errorf("Did NOT expect a match")
	}

	match, err = opcode.FindBytesMatchStartingAt(rawdata, 1)
	if err != nil {
		t.Error(err)
	}
	if match.Index != 1 {
		t.Errorf("First match expected at index 1. Instead got %d.", match.Index)
	}
	if match.Length != 3 {
		t.Errorf("First match expected length 3. Instead got %d.", match.Length)
	}

	match, err = opcode.FindBytesMatchStartingAt(rawdata, 2)
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}
	if match.Index != 5 {
		t.Errorf("Second match expected at index 5. Instead got %d.", match.Index)
	}
	if match.Length != 3 {
		t.Errorf("second match expected length 3. Instead got %d.", match.Length)
	}

	match, err = opcode.FindBytesMatchStartingAt(rawdata, 5)
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}
	if match.Index != 5 {
		t.Errorf("Second match expected at index 5. Instead got %d.", match.Index)
	}
	if match.Length != 3 {
		t.Errorf("second match expected length 3. Instead got %d.", match.Length)
	}

	match, err = opcode.FindBytesMatchStartingAt(rawdata, 6)
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}
	if match.Index != 7 {
		t.Errorf("Third match expected at index 7. Instead got %d.", match.Index)
	}
	if match.Length != 3 {
		t.Errorf("second match expected length 3. Instead got %d.", match.Length)
	}

}
