# 6502 Golang Emulator

This is a emulator for the 6502 microprossor
I built it because I didnt have a 6502 on hand and I wanted to play around with one.

Its not perfect, there are some features/opeations I got wrong, however in the test programs I created the 6502 emulator runs fine.

I watched a lot of [Ben Eater's](https://www.youtube.com/c/BenEater) videos which inspired me to build an emulator to play around with it.

StdIn is at READ on memory address `$6000`

StdErr is at WRITE on memory address `$6000`

Currently the clock is disabled (pulsing infinitely fast) but you can test it on what ever clock cycle you want and or make it so stdin causes a pulse, just modify the top level code in main.go

You can compile the assembly code using
<http://www.compilers.de/vasm.html>

Have fun :)
