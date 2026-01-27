assume cs: code, ds: data

data segment
one_seq dw 0
ans dw 0
len dw 3
arr dw 0, 0, 0
res_dec db "0000$"
data ends

code segment
start:	mov ax, data
        mov ds, ax

        mov si, 0
        mov cx, len
next_num:
        mov ax, arr[si]
        add si, 2

        cmp ax, 1
        je num_is_one
check_seq:
        cmp one_seq, 0
        jg store_and_exit
        mov one_seq, 0
        jmp continue
num_is_one:
        inc one_seq
continue:
        loop next_num
        mov ax, one_seq
        mov ans, ax
        mov si, offset res_dec + 4
        jmp convert_to_demical
store_and_exit:
        mov ax, one_seq
        mov ans, ax
        mov si, offset res_dec + 4
        jmp convert_to_demical

convert_to_demical:
        xor dx, dx
        mov bx, 10
        div bx

        add dl, '0'
        dec si
        mov [si], dl
        
        test ax, ax
        jne convert_to_demical
print_dec:
        mov ah, 09h
        mov dx, si
        int 21h

        mov ax, 4c00h
        int 21h
code ends
end start