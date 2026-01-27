assume cs:code, ds:data

data segment
string1 db 100, 99 dup(0)
dummy   db 0Dh, 0Ah, '$'
string2 db 100, 99 dup(0)
data ends

code segment
read_string macro str
        mov ax, data
        mov ds, ax
        mov dx, offset str
        xor ax, ax
        mov ah, 0Ah
        int 21h

        mov dx, offset dummy
        mov ah, 09h
        int 21h

        mov bl, [str+1]
        mov bh, 0
        lea si, [str+2]
        add si, bx
        mov byte ptr [si], '$'
endm

strcat proc
        push bp
        mov  bp, sp
        push si
        push di

        mov  di, [bp+4]
        mov  si, [bp+6]
find_end_of_string:
        mov  al, [di]
        cmp  al, '$'
        je   copy_string
        inc  di
        jmp  find_end_of_string
copy_string:
        mov  al, [si]
        cmp  al, '$'
        je   end_strcat
        mov  [di], al
        inc  di
        inc  si
        ;;movsb
        jmp  copy_string
end_strcat:
        mov  byte ptr [di], '$'
        mov  ax, [bp+4]
        pop  di
        pop  si
        pop  bp
        ret
strcat endp

start:
        read_string string1
        read_string string2

        lea di, [string1+2]
        lea si, [string2+2]

        push si
        push di
        call strcat

        mov dx, ax
        mov ah, 09h
        int 21h
program_end:
        mov ax, 4C00h
        int 21h
code ends
end start