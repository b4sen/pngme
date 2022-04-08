package chunk_type

import (
	"unicode"
)

type ChunkType struct {
	Name  []byte
	Props []int
}

func (ct *ChunkType) IsCritical() bool {
	return ct.Props[0] == 0
}

func (ct *ChunkType) IsPublic() bool {
	return ct.Props[1] == 0
}

func (ct *ChunkType) IsReservedBitValid() bool {
	return ct.Props[2] == 0
}

func (ct *ChunkType) IsSafeToCopy() bool {
	return ct.Props[3] != 0
}

func (ct *ChunkType) IsValid() bool {

	for _, chr := range ct.Name {
		str := rune(chr)
		if unicode.IsDigit(str) {
			return false
		}
	}
	return ct.IsReservedBitValid()
}

func (ct *ChunkType) Eq(other *ChunkType) bool {
	return ct.ToString() == other.ToString()
}

func (ct *ChunkType) ToString() string {
	return string(ct.Name)
}

func FromString(str string) ChunkType {

	name := make([]byte, 4)
	props := make([]int, 4)

	for idx, char := range str[:4] { // only take the first 4 chars
		if unicode.IsUpper(char) {
			props[idx] = 0
		} else {
			props[idx] = 1
		}

		name[idx] = byte(char)
	}
	ct := ChunkType{}
	ct.Name = name
	ct.Props = props

	return ct
}

func FromBytes(bytes []byte) ChunkType {
	str := string(bytes)
	ct := FromString(str)
	return ct
}
