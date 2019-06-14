package main

import (
	"flag"
	"fmt"
	"github.com/jonas747/dca"
	"io"
	"os"
	"time"
)

func main() {

	InFile = os.Args[1]
  OutFile = os.Args[2]
  
  var output = os.Stdout

  // inputReader is an io.Reader, like a file for example
  decoder := dca.NewDecoder(InFile)

  for {
      frame, err := decoder.OpusFrame()
      if err != nil {
          if err != io.EOF {
              // Handle the error
          }
        
          break
      }
    
      // Do something with the frame, in this example were sending it to discord
      select{
          case io.Copy(output) <- frame:
          case <-time.After(time.Second):
              // We haven't been able to send a frame in a second, assume the connection is borked
              return
      }
  }
}
