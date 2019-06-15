  1 package main
  2 
  3 import (
  4 	"flag"
  5 	"log"
  6 
  7 	//"fmt"
  8 	"io"
  9 	"os"
 10 
 11 	"engo.io/audio"
 12 	"github.com/jonas747/dca"
 13 	"layeh.com/gopus"
 14 )
 15 
 16 var (
 17 	InFile string
 18 	//OutFile string = "pipe:1"
 19 	err error
 20 
 21 	format = &audio.Format{
 22 		NumChannels: 2,
 23 		SampleRate:  48000,
 24 	}
 25 )
 26 
 27 func init() {
 28 	flag.StringVar(&InFile, "i", "pipe:0", "infile")
 29 	flag.Parse()
 30 }
 31 
 32 func main() {
 33 	if len(flag.Args()) < 1 {
 34 		log.Println("usage:", os.Args[0], "[input]", "[output]")
 35 	}
 36 
 37 	InFile = flag.Args()[1]
 38 	OutFile = flag.Args()[2]
 39 
 40 	// Open the file
 41 	inputReader, err := os.Open(InFile)
 42 	if err != nil { // error check
 43 		panic(err) // crash and burn
 44 	}
 45 
 46 	// Close the file on finish
 47 	defer inputReader.Close()
 48 
 49 	outputWriter, err := os.OpenFile(OutFile, os.O_WRONLY, os.ModePerm)
 50 	if err != nil {
 51 		panic(err)
 52 	}
 53 
 54 	defer outputWriter.Close()
 55 
 56 	// Make a new decoder
 57 	decoder := dca.NewDecoder(inputReader)
 58 
 59 	var PCM []int16
 60 
 61 	opusDecoder, err := gopus.NewDecoder(
 62 		format.SampleRate,  // sampling rate
 63 		format.NumChannels, // channels
 64 	)
 65 
 66 	if err != nil {
 67 		panic(err)
 68 	}
 69 
 70 	encoder := wav.NewEncoder(
 71 		outputWriter,
 72 		format.SampleRate, 16, format.NumChannels, 1,
 73 	)
 74 
 75 	defer encoder.Close()
 76 
 77 	var intBuf = &audio.IntBuffer{
 78 		Format:         format,
 79 		SourceBitDepth: 16,
 80 	}
 81 
 82 	for {
 83 		frame, err := decoder.OpusFrame()
 84 		if err != nil {
 85 			// Error happened before finishing
 86 			if err != io.EOF {
 87 				panic(err)
 88 			}
 89 
 90 			break
 91 		}
 92 
 93 		pcm, err := opusDecoder.Decode(frame, 960, false)
 94 		if err != nil {
 95 			panic(err)
 96 		}
 97 
 98 		ints := make([]int, len(pcm))
 99 		for j, i := range pcm {
100 			ints[j] = int(i)
101 		}
102 
103 		intBuf.Data = ints
104 
105 		if err := encoder.Write(intBuf); err != nil {
106 			panic(err)
107 		}
108 	}
109 }
