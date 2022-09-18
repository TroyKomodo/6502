STDOUT = $8000

	.org $8100

main:
	ldx #00 ; load x with 0

print_msg:
	lda message,x ; load a with message[x]
	beq done ; if a is 0, jump to done
	sta STDOUT ; store a in STDOUT
	inx ; increment x
	jmp print_msg ; loop back to print the next character

done:
	rti ; return from interrupt

loop:
	jmp loop ; loop forever

; store the message with a new line at the end and a null byte
message: .byte "Hello, World!", $0A, $00

	.org $fffc
	.word loop
	.word main
