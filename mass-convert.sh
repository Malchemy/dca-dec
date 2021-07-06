#!/bin/bash

# This will attempt to convert evert .dca in the current directory to a .wav
# Does not work for files with spaces. Why do you have spaces in .dca's?

# dca-dec must be in $PATH
for FILE in *.dca
{
   dca-dec ${FILE} "${FILE%%.*}.wav"
}

# Comment out the top and use these lines if you'd rather have dca-dec
# in the current folder instead of $PATH or just add ./ in front of dca-dec above..
#for FILE in *.dca
#{
#   ./dca-dec ${FILE} "${FILE%%.*}.wav"
#}
