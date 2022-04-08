package chunk

import (
	"encoding/binary"
	"hash/crc32"

	"github.com/b4sen/pngme/chunk_type"
)

type Chunk struct {
	Length uint
	Type   *chunk_type.ChunkType
	Data   []byte
	CRC    uint32
}

func (c *Chunk) calculateCRC() uint32 {
	data := append(c.Type.Name, c.Data...) // append is a variadic function, need to spread the slice!
	return crc32.ChecksumIEEE(data)
}

func (c *Chunk) calculateLength() uint {
	return uint(len(c.Data))
}

func New(ct *chunk_type.ChunkType, data []byte) Chunk {
	c := Chunk{}
	c.Type = ct
	c.Data = data
	c.Length = c.calculateLength()
	c.CRC = c.calculateCRC()
	return c
}

func (c *Chunk) ToBytes() []byte {
	var bytes []byte
	length := []byte{0, 0, 0, byte(c.Length)}
	crc := make([]byte, 4)
	binary.BigEndian.PutUint32(crc, c.CRC)
	bytes = append(bytes, length...)
	bytes = append(bytes, c.Type.Name...)
	bytes = append(bytes, c.Data...)
	bytes = append(bytes, crc...)

	return bytes
}
