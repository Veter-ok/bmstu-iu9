assume cs: code, ds: data

data segment
max_digits db 50
string1    db 100, 50 dup(0)
dummy      db 0Dh, 0Ah, '$'
string2    db 100, 50 dup(0)

num1_array db 50 dup(0)
num2_array db 50 dup(0)
res_array  db 50 dup(0)

num1_sign  db 0
num2_sign  db 0
res_sign   db 0

num1_len   db 0
num2_len   db 0
res_len    db 0

carry    db 0
res        db 100, 50 dup(0)
error_msg db 'Error: Invalid input.$'
data ends

code segment
read_string macro str
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

convert_string_to_array macro str, arr, sign, length
        local convert, next_digit, digit_0_9, store_digit, bad_symbol, done_convert
        mov al, [str + 1]
        mov ah, 0

        lea si, [str + 2]
        mov bl, [si]
        cmp bl, '-'
        jne convert
        mov sign, 1
        dec ax
        inc si
convert:
        mov length, al
        add si, ax
        dec si
        lea di, [si]
        lea si, [arr]
next_digit:
        mov bl, [di]
        cmp bl, 30h
        jl bad_symbol
        cmp bl, 39h
        jle digit_0_9

        cmp bl, 'A'
        jl bad_symbol
        cmp bl, 'F'
        jg bad_symbol

        sub bl, 'A' - 10
        jmp store_digit
digit_0_9:
        sub bl, '0'  
store_digit:
        mov [si], bl
        dec di
        inc si
        dec ax
        cmp ax, 0
        jne next_digit
        jmp done_convert
bad_symbol:
        mov dx, offset error_msg
        mov ah, 09h
        int 21h
            
        mov ax, 4c00h
        int 21h
done_convert:
endm

mul_arrays macro
        lea si, [num1_array]
        lea di, [num2_array]
        lea bx, [res_array]
        mov cl, num1_len
        mov ch, num2_len
        
        multiply
        mov al, num1_sign
        cmp al, num2_sign
        jne set_negative
        jmp done_mul
set_negative:
        mov res_sign, 1
done_mul:
endm

multiply macro
        local loop_1, loop_2, loop_minus_16, updates_vars, check_borrow, end_loop2, end_loop1
        mov ax, 0
        mov dx, 0
        mov res_len, cl
loop_1:
        mov dh, [di]
loop_2:
        mov al, [si]
        mul dh
        add al, carry
        mov carry, 0
        add [bx], al
        cmp byte ptr [bx], 10h
        jb updates_vars
loop_minus_16:
        sub byte ptr [bx], 10h
        inc carry
        cmp byte ptr [bx], 10h
        jnb loop_minus_16
updates_vars:
        inc si
        inc bx
        dec cl
        cmp cl, 0
        jne loop_2
check_borrow:
        cmp carry, 0
        je end_loop2
        mov al, carry
        mov [bx], al
        add res_len, 1
        mov carry, 0
end_loop2:
        inc di
        add res_len, 1
        lea si, [num1_array]
        mov cl, num1_len
        sub bl, num1_len
        add bx, 1
        dec ch
        cmp ch, 0
        jne loop_1
end_loop1:
endm

convert_array_to_string macro arr, length, sign, res_str
        local write_loop, digit_0_9, store_digit, num_is_zero, check_if_last, store_last_zero, updates_vars, not_all_zeros, done
        lea si, [arr]
        lea di, [res_str]
        mov dx, 0

        mov bl, length
        mov bh, 0
        add si, bx
        dec si

        cmp sign, 1
        jne write_loop
        mov byte ptr [di], '-'
        inc di
write_loop:
        mov al, [si]
        cmp al, 10
        jl digit_0_9
        add al, 'A' - 10
        jmp store_digit
digit_0_9:
        add al, '0'
store_digit:
        cmp al, '0'
        je num_is_zero
        mov dx, 1
        mov [di], al
        inc di
        jmp updates_vars
num_is_zero:
        cmp dx, 0
        je check_if_last
        mov [di], al
        inc di
        jmp updates_vars
check_if_last:
        cmp bl, 1
        je store_last_zero
        jmp updates_vars
store_last_zero:
        mov [di], al
        inc di
updates_vars:
        dec si
        dec bl
        cmp bl, 0
        jne write_loop
        cmp dx, 0
        jne not_all_zeros
        mov byte ptr [di], '0'
        inc di
not_all_zeros:
        mov byte ptr [di], '$'
done:
endm

start:	
        mov ax, data
        mov ds, ax

        read_string string1
        read_string string2

        convert_string_to_array string1, num1_array, num1_sign, num1_len
        convert_string_to_array string2, num2_array, num2_sign, num2_len

        mul_arrays

        convert_array_to_string res_array, res_len, res_sign, res

        mov dx, offset res
        mov ah, 09h
        int 21h
            
        mov ax, 4c00h
        int 21h
code ends
end start