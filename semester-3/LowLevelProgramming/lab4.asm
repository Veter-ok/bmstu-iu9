assume cs: code, ds: data

data segment
loop_msg db "Hello world!$"
error_msg db "Something went wrong($"
succes_msg db "Everything is fine$"
A dd 4294967000
B dd 0
data ends

code segment
PUSHM macro X
        IFNB <X>
            IFDEF X
                push [X]
                push [X+2]
            ENDIF
        ENDIF
endm

POPM macro X
        IFNB <X>
            IFDEF X
                pop [X+2]
                pop [X]
            ENDIF
        ENDIF
endm

CALLM macro P
        IFNB <P>
            IFDEF P
                push offset $+12
                jmp P
            ENDIF
        ENDIF
endm

RETM macro N
        pop ax
        IFNB <N>
            IFDEF N
                add sp, N
            ENDIF
        ENDIF
        jmp ax
endm

LOOPM macro L
        IFNB <L>
            IFDEF L
                dec cx
                jnz L
            ENDIF
        ENDIF
endm

test_print proc
        push bp
        mov bp, sp

        PUSHM A 
        POPM B
        ; mov ax, word ptr A+2
        ; inc ax
        ; mov word ptr A+2, ax

        mov cx, 5
repeat_print:
        mov ah, 09h
        mov dx, offset loop_msg
        int 21h
        mov ah, 2
        mov dl, 10
        int 21h
        LOOPM repeat_print

        pop bp
        RETM
test_print endp

start:	mov ax, data
        mov ds, ax
        
        CALLM test_print
        CALLM test_print

        mov ax, word ptr A
        mov bx, word ptr B
        cmp ax, bx
        jne fail

        mov ax, word ptr A + 2
        mov bx, word ptr B + 2
        cmp ax, bx
        jne fail

        mov dx, offset succes_msg
        mov ah, 09h
        int 21h

        jmp program_end
fail:
        mov dx, offset error_msg
        mov ah, 09h
        int 21h
program_end:
        mov ax, 4C00h
        int 21h
code ends
end start