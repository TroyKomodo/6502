VALUE = $0000 ; 8 bytes 
MOD10 = VALUE + 8 ; 1 byte
STRING_BUILDER = MOD10 + 1 ; 32 bytes
STRING_BUILDER_IDX = STRING_BUILDER + 32 ; 1 byte
VALUE_PTR = STRING_BUILDER_IDX + 1 ; 2 bytes
STDOUT = $8000


	.org $8100

number: .byte 0x00, 0x00, 0x00, 0x12, 0x00, 0x55, 0x00, 0x11 ; 8 bytes 

main:
	lda #$00
	sta VALUE_PTR

	ldy #$00
	set_defaults:
		lda number,y
		sta (VALUE_PTR),y
		iny
		cpy #$08
		bne set_defaults

	lda #$00 ; load a with 0
	ldx #32
	mini_loop:
		sta STRING_BUILDER,x
		dex
		bne mini_loop
	
	lda #01 ; load a with 4
	sta STRING_BUILDER_IDX

	jsr to_decimal ; convert a to decimal
	ldx STRING_BUILDER_IDX ; load x with STRING_BUILDER_IDX
	jsr print_msg ; print the string

	lda #$0A ; load a with new line
	sta STDOUT ; print a new line

	rti

print_msg:
	dex ; increment x
	beq print_msg_return ; if a is 0, jump to return

	lda STRING_BUILDER,x ; load a with STRING_BUILDER[x]
	sta STDOUT ; store a in STDOUT
	jmp print_msg ; loop back to print the next character

	print_msg_return:
		rts


to_decimal:
	; convert a binary number to decimal
	; input: binary number in VALUE
	; clobbers a, x, y
	lda #$00 ; load a with 0
	sta MOD10 ; save a to MOD10

	clc ; clear carry

	ldx #64 ; load x with 32

	to_decimal_loop:
		; txa ; transfer x to a
		; tay ; transfer a to y
		ldy #$00 ; load y with 0
		php
		y_ptr_loop_1:
			plp
			lda (VALUE_PTR),y ; load a with VALUE[y]
			rol A
			php
			sta (VALUE_PTR),y ; store a in VALUE[y]
			iny ; increment y
			cpy #$08 ; compare y with 8
			bne y_ptr_loop_1 ; if y is not 8, loop back
		plp

		rol MOD10 ; rotate left MOD10

		; a,y = divident - divsor
		sec
		lda MOD10
		sbc #10
		bcc to_decimal_ignore_result ; branch if divident < divsor
		sta MOD10 ; save a to MOD10

	to_decimal_ignore_result:
		dex ; decrement x
		bne to_decimal_loop ; branch if x != 0

		ldy #$00 ; load y with 0
		php
		y_ptr_loop_2:
			plp
			lda (VALUE_PTR),y ; load a with VALUE[y]
			rol A
			php
			sta (VALUE_PTR),y ; store a in VALUE[y]
			iny ; increment y
			cpy #$08 ; compare y with 8
			bne y_ptr_loop_2 ; if y is not 8, loop back
		plp

		lda MOD10 ; load a with MOD10
		clc
		adc #"0"
		ldx STRING_BUILDER_IDX ; load x with STRING_BUILDER_IDX
		sta STRING_BUILDER,x ; save a to STRING_BUILDER[x]
		inc STRING_BUILDER_IDX ; increment STRING_BUILDER_IDX

		lda VALUE
		ldy #$00 ; load y with 0
		php
		y_ptr_loop_3:
			plp
			ora (VALUE_PTR),y ; load a with VALUE[y]
			php
			iny ; increment y
			cpy #$08 ; compare y with 8
			bne y_ptr_loop_3 ; if y is not 8, loop back
		plp
		bne to_decimal ; branch if VALUE != 0

		rts


loop:
	jmp loop ; loop forever

end:
	.org $fffc
	.word loop
	.word main
