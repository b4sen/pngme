package chunk

import (
	"reflect"
	"testing"

	"pngme/chunk_type"
)

func TestNewChunk(t *testing.T){
	ct := chunk_type.FromString("ASDF")
	c := Chunk{}
	c.Type = &ct
	c.Data = []byte("Hello World!")
	c.Length = 12
	c.CRC = 3519565349

	c2 := New(&ct, []byte("Hello World!"))
	
	if !reflect.DeepEqual(c, c2){
		t.Errorf("Chunks are not equal!")
	}
}
