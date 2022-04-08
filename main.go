package main

import (
	"errors"
	"flag"
	"os"

	"github.com/b4sen/pngme/chunk"
	"github.com/b4sen/pngme/chunk_type"
)

func main() {
	infile := flag.String("i", "Dice.png", "Path to the input image")
	outfile := flag.String("o", "test_out.png", "Path to the output image")
	typeName := flag.String("n", "teSt", `Name of the chunk. The first letter has to be lowercase and
	the third letter uppercase.`)
	msg := flag.String("m", "default message", "Message to be written into the file")
	flag.Parse()

	ct := chunk_type.FromString(*typeName) // 3rd needs to be uppercase
	if !ct.IsValid() || ct.IsCritical() {
		panic(errors.New("Reserved bit is invalid!"))
	}

	data, err := os.ReadFile(*infile)
	if err != nil {
		panic(err)
	}
	c := chunk.New(&ct, []byte(*msg))
	bytes := c.ToBytes()
	data = append(data, bytes...)

	os.WriteFile(*outfile, data, 0644)

}
