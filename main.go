package main

import (
	"flag"
	"fmt"
	"github.com/jonas747/dca"
	"io"
	"os"
	"time"
)

var (
	InFile      string
	OutFile string = "pipe:1"
)

func init() {	
	flag.StringVar(&InFile, "i", "pipe:0", "infile")
	flag.Parse()
}

func main() {

  InFile = os.Args[1]
  OutFile = os.Args[2]
  
  var output = os.Stdout

  inputReader, err := os.Open(InFile)
	
  // inputReader is an io.Reader, like a file for example
  decoder := dca.NewDecoder(inputReader)

  for {
      frame, err := decoder.OpusFrame()
      if err != nil {
          if err != io.EOF {
              // Handle the error
          }
        
          break
      }
      select{
          case os.Stdout <- frame:
          case <-time.After(time.Second):
              // We haven't been able to send a frame in a second, assume the connection is borked
              return
      }
  }
}
