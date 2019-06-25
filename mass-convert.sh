#!/bin/sh

# This will attempt to convert evert .dca in the current directory to a .wav
# Does not work for files with spaces. Why do you have spaces in .dca's?

for FILE in *.dca
{
   ./dca-dec ${FILE} "$(basename "${FILE}" .dca).wav"
}
