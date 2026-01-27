assume cs: code, ds: data

data segment
a dw 2
b dw 3
c dw 66
d dw 2
res1 dw ?
res2 dw ?
res  dw ?
minus_flag db 0
minus db "-$"
res_dec db "0000$"
res_hex db "0000$"
data ends

code segment
start:	
        mov ax, data
        mov ds, ax

        mov ax, a
        imul b
        mov res1, ax

        mov ax, c
        imul d
        mov res2, ax

        shr res2, 1
        shr res2, 1
        
        mov ax, res1
        mov cx, res2
        sub ax, cx
        mov res, ax

        test ax, ax
        jge continue

        neg ax
        mov res, ax
        mov minus_flag, 1

continue:
        mov si, offset res_dec + 4

convert_to_demical:
        xor dx, dx
        mov bx, 10
        div bx

        add dl, '0'
        dec si
        mov [si], dl
        
        test ax, ax
        jne convert_to_demical

        cmp minus_flag, 0
        je print_dec
        mov ah, 09h
        mov dx, offset minus
        int 21h

print_dec:
        mov ah, 09h
        mov dx, si
        int 21h

        mov ah, 2
        mov dl, 10
        int 21h

        mov ax, res
        mov si, offset res_hex + 4

convert_to_hex:
        xor dx, dx
        mov bx, 16
        div bx

        cmp dl, 10
        jge letter

        add dl, '0'
        jmp store

letter:
        add dl, 'A' - 10

store:
        dec si
        mov [si], dl
        test ax, ax
        jne convert_to_hex

print_minus:
        cmp minus_flag, 0
        je print_hex
        mov ah, 09h
        mov dx, offset minus
        int 21h

print_hex:
        mov ah, 09h
        mov dx, si
        int 21h

end_program: 
        mov ax, 4c00h
        int 21h
code ends
end start