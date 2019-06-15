package main

import (
	"flag"
	"log"

	"fmt"
	"io"
	"os"

	"github.com/go-audio/wav"
	"github.com/go-audio/audio"
	//"github.com/Malchemy/audio"
	//"engo.io/audio"
	"github.com/jonas747/dca"
	"layeh.com/gopus"
)

var (
	/*InFile string
	OutFile string*/
	//OutFile string = "pipe:1"
	err error
)
	var format = &audio.Format {
		NumChannels: 2,
		SampleRate:  48000,
	}

/*func init() {
	flag.StringVar(InFile, "i", "pipe:0", "infile")
	flag.Parse()
}*/

func main() {
	var InFile string = flag.Args()[1]
	var OutFile string = flag.Args()[2]

	if len(flag.Args()) < 1 {
		log.Println("usage:", os.Args[0], "[input]", "[output]")
	}
	
	fmt.Printf("%#v\n", flag.Args())
	fmt.Println("Length of `x` is", len(flag.Args()))
	fmt.Println("Third element of `x` is", flag.Args()[2])

	// Open the file
	inputReader, err := os.Open(InFile)
	if err != nil { // error check
		panic(err) // crash and burn
	}

	// Close the file on finish
	//defer inputReader.Close()

	outputWriter, err := os.OpenFile(OutFile, os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	//defer outputWriter.Close()

	// Make a new decoder
	decoder := dca.NewDecoder(inputReader)

	var pcm []int16

	opusDecoder, err := gopus.NewDecoder(
		format.SampleRate,  // sampling rate
		format.NumChannels, // channels
	)

	if err != nil {
		panic(err)
	}

	encoder := wav.NewEncoder(
		outputWriter,
		format.SampleRate, 16, format.NumChannels, 1,
	)

	defer encoder.Close()

	var intBuf = &audio.IntBuffer{
		Format:         format,
		SourceBitDepth: 16,
	}

	for {
		frame, err := decoder.OpusFrame()
		if err != nil {
			// Error happened before finishing
			if err != io.EOF {
				panic(err)
			}

			break
		}

		pcm, err := opusDecoder.Decode(frame, 960, false)
		if err != nil {
			panic(err)
		}

		ints := make([]int, len(pcm))
		for j, i := range pcm {
			ints[j] = int(i)
		}

		intBuf.Data = ints

		if err := encoder.Write(intBuf); err != nil {
			panic(err)
		}
	}
	_ = pcm
}
