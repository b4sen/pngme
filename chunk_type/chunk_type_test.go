package chunk_type

import (
	"testing"
)

func TestIsCritical(t *testing.T) {
	ct := FromString("ASDF")
	if !ct.IsCritical() {
		t.Errorf("Chunk should be critical!")
	}
}

func TestIsNotCritical(t *testing.T) {
	ct := FromString("aSDF")
	if ct.IsCritical() {
		t.Errorf("Chunk should not be critical!")
	}
}

func TestIsPublic(t *testing.T) {
	ct := FromString("ASDF")
	if !ct.IsPublic() {
		t.Errorf("Chunk should be public!")
	}
}

func TestIsNotPublic(t *testing.T) {
	ct := FromString("AsDF")
	if ct.IsPublic() {
		t.Errorf("Chunk should not be public!")
	}
}

func TestIsReservedBitValid(t *testing.T) {
	ct := FromString("ASDF")
	if !ct.IsReservedBitValid() {
		t.Errorf("Reserved bit should be valid!")
	}
}

func TestIsReservedBitInvalid(t *testing.T) {
	ct := FromString("ASdF")
	if ct.IsReservedBitValid() {
		t.Errorf("Reserved bit should be invalid!")
	}
}

func TestValidChunkIsValid(t *testing.T) {
	ct := FromString("ASDF")
	if !ct.IsValid() {
		t.Errorf("Chunk should be valid!")
	}
}

func TestChunkToString(t *testing.T) {
	ct := FromString("ASDF")
	if ct.ToString() != "ASDF" {
		t.Errorf("Invalid ToString()")
	}
}

func TestChunkTypeEq(t *testing.T) {
	ct1 := FromString("ASDF")
	ct2 := FromString("ASDF")
	if !ct1.Eq(&ct2) {
		t.Errorf("Chunk's not equal!")
	}
}
