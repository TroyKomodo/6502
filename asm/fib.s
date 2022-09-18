VALUE = $0000 ; 8 bytes 

X_VALUE = VALUE + 8 ; 8 bytes
Y_VALUE = X_VALUE + 8 ; 8 bytes
MOD10 = Y_VALUE + 8 ; 1 byte
STRING_BUILDER = MOD10 + 1 ; 32 bytes
STRING_BUILDER_IDX = STRING_BUILDER + 32 ; 1 byte
X_VALUE_PTR = STRING_BUILDER_IDX + 1 ; 2 bytes
Y_VALUE_PTR = X_VALUE_PTR + 2 ; 2 bytes

STDOUT = $8000

	.org $8100

main:
	lda #$00
	ldy #$00
	vector:
		sta VALUE,y
		sta X_VALUE,y
		sta Y_VALUE,y
		iny
		cpy #$08
		bne vector

	lda #$00
	sta X_VALUE
	lda #$01
	sta Y_VALUE

	lda #<X_VALUE
	sta X_VALUE_PTR
	lda #<Y_VALUE
	sta Y_VALUE_PTR
	
	ldx #$00
	add_loop:
		brk
		txa 
		pha

		lda X_VALUE_PTR
		tax
		lda Y_VALUE_PTR
		sta X_VALUE_PTR
		stx Y_VALUE_PTR

		jsr add_numbers

		ldy #$00
		copy_loop:
			lda (X_VALUE_PTR),y
			sta VALUE,y
			iny
			cpy #$08
			bne copy_loop

		jsr print_to_decimal ; convert a to decimal
		ldx STRING_BUILDER_IDX ; load x with STRING_BUILDER_IDX
		jsr print_msg ; print the string

		lda #$0A ; load a with new line
		sta STDOUT ; print a new line

		pla
		tax
		inx
		cpx #90
		bne add_loop

	jmp loop

add_numbers:
	ldy #$00
	clc
	php
	add_numbers_inner:
		; add the numbers from the pointers
		plp
		lda (X_VALUE_PTR),y
		adc (Y_VALUE_PTR),y
		sta (X_VALUE_PTR),y
		php
		iny
		cpy #$08 ; check if y is 8
		bne add_numbers_inner ; if not, loop
	plp
	rts

print_msg:
	dex ; increment x
	beq print_msg_return ; if a is 0, jump to return

	lda STRING_BUILDER,x ; load a with STRING_BUILDER[x]
	sta STDOUT ; store a in STDOUT
	jmp print_msg ; loop back to print the next character

	print_msg_return:
		rts


print_to_decimal:
	; convert a binary number to decimal
	; input: binary number in VALUE
	; clobbers a, x, y
	lda #$00 ; load a with 0
	ldx #32
	mini_loop:
		sta STRING_BUILDER,x
		dex
		bne mini_loop
	
	lda #01 ; load a with 4
	sta STRING_BUILDER_IDX
	to_decimal:
		; convert a binary number to decimal
		; input: binary number in VALUE
		; clobbers a, x, y
		lda #$00 ; load a with 0
		sta MOD10 ; save a to MOD10

		clc ; clear carry

		ldx #64 ; load x with 64

		to_decimal_loop:
			rol VALUE ; rotate left
			rol VALUE+1
			rol VALUE+2
			rol VALUE+3
			rol VALUE+4
			rol VALUE+5
			rol VALUE+6
			rol VALUE+7
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

			rol VALUE ; rotate left
			rol VALUE+1
			rol VALUE+2
			rol VALUE+3
			rol VALUE+4
			rol VALUE+5
			rol VALUE+6
			rol VALUE+7

			lda MOD10 ; load a with MOD10
			clc
			adc #"0"
			ldx STRING_BUILDER_IDX ; load x with STRING_BUILDER_IDX
			sta STRING_BUILDER,x ; save a to STRING_BUILDER[x]
			inc STRING_BUILDER_IDX ; increment STRING_BUILDER_IDX

			lda VALUE ; load a with VALUE[0]
			ora VALUE+1 ; or a with VALUE[1]
			ora VALUE+2 ; or a with VALUE[2]
			ora VALUE+3 ; or a with VALUE[3]
			ora VALUE+4 ; or a with VALUE[4]
			ora VALUE+5 ; or a with VALUE[5]
			ora VALUE+6 ; or a with VALUE[6]
			ora VALUE+7 ; or a with VALUE[7]
			bne to_decimal ; branch if VALUE != 0

			rts


loop:
	jmp loop ; loop forever

interupt:
	rti

end:
	.org $fffc
	.word main
	.word interupt
