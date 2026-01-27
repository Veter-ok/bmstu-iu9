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
        local next_digit, convert, bad_symbol, done_convert
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
        jg bad_symbol
        sub bl, '0'  
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

sum_arrays macro
        local both_positive, who_bigger, num2_greater, execute_sum, different_signs, num1_greater_sub, num2_greater_sub, num1_minus_num2, num2_minus_num1, done
        mov al, num1_sign
        cmp al, num2_sign
        jne different_signs

        cmp al, 1
        jne both_positive
        mov res_sign, 1
        jmp who_bigger
both_positive:
        mov res_sign, 0
who_bigger:
        mov al, num1_len
        cmp al, num2_len
        jle num2_greater
        
        lea si, [num1_array]
        lea di, [num2_array]
        mov cl, num1_len
        mov ch, num2_len
        jmp execute_sum
num2_greater:
        lea si, [num2_array]
        lea di, [num1_array]
        mov cl, num2_len
        mov ch, num1_len
execute_sum:
        lea bx, [res_array]
        call absolute_sum
        jmp done

different_signs:
        mov al, num1_len
        cmp al, num2_len
        jl num2_greater_sub
num1_greater_sub:
        mov al, num1_sign
        cmp al, 0
        je num1_minus_num2
        mov res_sign, 1
        jmp num1_minus_num2
num2_greater_sub:
        mov al, num2_sign
        cmp al, 0
        je num2_minus_num1
        mov res_sign, 1
        jmp num2_minus_num1
num1_minus_num2:
        lea si, [num1_array]
        lea di, [num2_array]
        lea bx, [res_array]
        mov cl, num1_len
        mov ch, num2_len
        call subtract

        cmp dl, 0
        je done
        xor res_sign, 1
num2_minus_num1:
        lea si, [num2_array]
        lea di, [num1_array]
        lea bx, [res_array]
        mov cl, num2_len
        mov ch, num1_len
        call subtract
done:
endm

absolute_sum proc
        mov ax, 0
next_digit:
        mov al, [si]
        add al, [di]
        add al, ah
        xor ah, ah

        cmp al, 10
        jl store_result
        sub al, 10
        inc ah
store_result:
        mov [bx], al
        add res_len, 1
        inc si
        inc di
        inc bx
        dec cl
        dec ch
        cmp ch, 0
        jne next_digit

        cmp cl, 0
        jne latest_digits
        jmp check_shift
latest_digits:
        mov al, [si]
        add al, ah
        xor ah, ah

        cmp al, 10
        jl store_latest_digits
        sub al, 10
        inc ah
store_latest_digits:
        mov [bx], al
        add res_len, 1
        inc si
        inc bx
        dec cl
        cmp cl, 0
        jne latest_digits
check_shift:
        cmp ah, 0
        je done
        mov [bx], ah
        add res_len, 1
done:
        ret
absolute_sum endp

subtract proc
        mov ax, 0
        mov dx, 0
next_digit_sub:
        mov al, [si]
        mov ah, [di]
        add ah, dl
        xor dl, dl
        cmp al, ah
        jge store_result_sub
        add al, 10
        inc dl
store_result_sub:
        sub al, ah
        mov [bx], al
        add res_len, 1
        inc si
        inc di
        inc bx
        dec cl
        dec ch
        cmp ch, 0
        jne next_digit_sub

        cmp cl, 0
        jne latest_digits_sub
        jmp done_sub
latest_digits_sub:
        mov al, [si]
        mov ah, dl
        xor dl, dl
        cmp al, ah
        jge store_latest_digits_sub
        add al, 10
        inc dl
store_latest_digits_sub:
        sub al, ah
        mov [bx], al
        add res_len, 1
        inc si
        inc bx
        dec cl
        cmp cl, 0
        jne latest_digits_sub
done_sub:
        ret
subtract endp

convert_array_to_string macro arr, length, sign, res_str
        local write_loop, store_digit, updates_vars, num_is_zero, done
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
        add al, '0'
        cmp al, '0'
        je num_is_zero
        mov dx, 1
store_digit:
        mov [di], al
        inc di
updates_vars:
        dec si
        dec bl
        cmp bl, 0
        jne write_loop
        mov byte ptr [di], '$'
        jmp done
num_is_zero:
        cmp dx, 0
        je is_last_digit
        jmp store_digit
is_last_digit:
        cmp bl, 1
        je store_digit
        jmp updates_vars
done:
endm

start:	
        mov ax, data
        mov ds, ax

        read_string string1
        read_string string2

        convert_string_to_array string1, num1_array, num1_sign, num1_len
        convert_string_to_array string2, num2_array, num2_sign, num2_len

        sum_arrays

        convert_array_to_string res_array, res_len, res_sign, res

        mov dx, offset res
        mov ah, 09h
        int 21h
            
        mov ax, 4c00h
        int 21h
code ends
end start