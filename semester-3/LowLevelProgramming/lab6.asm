.386p
       
RM_CODE segment     para public 'CODE' use16
        assume      CS:RM_CODE,SS:RM_STACK
@@start:
        mov     AX,03h
        int     10h
               
        in      AL,92h
        or      AL,2
        out     92h,AL
 
        xor     EAX,EAX
        mov     AX,PM_CODE
        shl     EAX,4 
        add     EAX,offset ENTRY_POINT
        mov     dword ptr ENTRY_OFF,EAX
 
        xor     EAX,EAX
        mov     AX,RM_CODE
        shl     EAX,4
        add     AX,offset GDT
 
        mov     dword ptr GDTR+2,EAX
 
        lgdt    fword ptr GDTR
 
        cli
 
        in      AL,70h
        or      AL,80h
        out     70h,AL
 
        mov     EAX,CR0
        or      AL,1
        mov     CR0,EAX
 
            db      66h
            db      0EAh
ENTRY_OFF   dd      ?
            dw      00001000b
 
; ТАБЛИЦА ГЛОБАЛЬНЫХ ДЕСКРИПТОРОВ:
GDT:  
    NULL_descr  db      8 dup(0)
    CODE_descr  db      0FFh,0FFh,00h,00h,00h,10011010b,11001111b,00h
    DATA_descr  db      0FFh,0FFh,00h,00h,00h,10010010b,11001111b,00h
    DESC4_descr db      0Ah,00h,0Bh,1Dh,0Fh,11110111b,11000000b,00h
    DESC5_descr  db     0A1h,0B4h,66h,010h,57h,10001010b,11101111b,12h
    DESC6_descr db      34h,12h,78h,56h,9Ah,00000100b,00001111b,0ABh
    GDT_size    equ     $-GDT
 
GDTR        dw      GDT_size-1
            dd      ?
RM_CODE ends
 
 
 
; СЕГМЕНТ СТЕКА (для Real Mode)
RM_STACK       segment          para stack 'STACK' use16
               db     100h dup(?)
RM_STACK       ends
 
 
 
; СЕГМЕНТ КОДА (для Protected Mode)
PM_CODE segment     para public 'CODE' use32
        assume      CS:PM_CODE,DS:PM_DATA
        
print_descr_data macro descriptor, color
    local save_desc, output_desc4, data_type, system_type, type_done, print_limit, type
    local data_segment_chars, done, c_zero, print_c, r_zero, print_r
    local a_zero_code, print_a_code, ed_zero, print_ed, w_zero, print_w
    local a_zero_data, print_a_data, done_print, bit32, add_inf, add_r

    mov  ESI, descriptor
    
    mov  ECX, 8
    lea  EBX, [desc_buffer]
    
save_desc:
    mov  AL, [ESI]
    mov  [EBX], AL
    inc  ESI
    inc  EBX
    loop save_desc

    lea  ESI, [desc_buffer]

    mov  AL, [ESI+6]
    test AL, 20h
    je  add_check 
    jmp print_error

add_check:
    mov  AL, [ESI+5]
    test AL, 10h
    jnz   output_desc4
    and  AL, 0Fh
    
    cmp  AL, 00h
    je   print_error
    
    cmp  AL, 08h
    je   print_error
    
    cmp  AL, 0Ah
    je   print_error
    
    cmp  AL, 0Dh
    je   print_error

    jmp output_desc4

print_error:
    push ESI
    sub  EDI, 20
    mov  ESI, PM_DATA
    shl  ESI, 4
    add  ESI, offset error_msg
    mov  ECX, error_msg_len
    rep  movsw
    add  EDI, 6
    pop  ESI

output_desc4:

    mov  BL, color
    xor  EAX, EAX
    mov  AL, [ESI+7]
    shl  EAX, 8
    mov  AL, [ESI+4]
    shl  EAX, 16
    mov  AX, [ESI+2]
    call PrintHexDword
    add  EDI, 4
    
    xor  EAX, EAX
    mov  AL, [ESI+6]
    and  AL, 0Fh
    shl  EAX, 16
    mov  AX, [ESI]
    
    test byte ptr [ESI+6], 80h
    jz  print_limit
    
    shl  EAX, 12
    add  EAX, 0FFFh

    cmp EAX, 0FFFFFFFFh 
    jne print_limit

    inc EAX
    add  EDI, 4
    mov  AH, color
    mov  AL, '4'
    call PutChar
    mov  AL, 'G'
    call PutChar
    mov  AL, 'B'
    call PutChar
    add  EDI, 6
    jmp type

print_limit:
    inc EAX
    call PrintHexDword

type:  
    add  EDI, 6  
    mov  AH, color
    mov  AL, [ESI+5]
    test AL, 10h
    jz   system_type
    
    test AL, 08h
    jz   data_type
    
    mov  AL, 'c'
    call PutChar
    mov  AL, 'o'
    call PutChar
    mov  AL, 'd'
    call PutChar
    mov  AL, 'e'
    call PutChar
    jmp  type_done
    
data_type:
    mov  AL, 'd'
    call PutChar
    mov  AL, 'a'
    call PutChar
    mov  AL, 't'
    call PutChar
    mov  AL, 'a'
    call PutChar
    jmp  type_done
    
system_type:
    mov  AL, 's'
    call PutChar
    mov  AL, 'y'
    call PutChar
    mov  AL, 's'
    call PutChar
    mov  AL, 't'
    call PutChar

    add  EDI, 10

    mov  AL, '_'
    call PutChar
    mov  AL, '_'
    call PutChar
    mov  AL, '_'
    call PutChar

    jmp pre_done

type_done:
    add  EDI, 10
    mov  AL, [ESI+5]
    mov  BL, AL
    
    test BL, 08h
    jz   data_segment_chars
    
    test BL, 04h
    jz   c_zero
    mov  AL, 'C'
    jmp  print_c
c_zero:
    mov  AL, '_'
print_c:
    call PutChar
    
    test BL, 02h
    jz   r_zero
    mov  AL, 'R'
    jmp  print_r
r_zero:
    mov  AL, '_'
print_r:
    call PutChar
    
    test BL, 01h
    jz   a_zero_code
    mov  AL, 'A'
    jmp  print_a_code
a_zero_code:
    mov  AL, '_'
print_a_code:
    call PutChar
    jmp  add_inf
    
data_segment_chars:
    test BL, 04h
    jz   ed_zero
    mov  AL, 'E'
    jmp  print_ed
ed_zero:
    mov  AL, '_'
print_ed:
    call PutChar
    
    test BL, 02h 
    jz   w_zero
    mov  AL, 'W'
    jmp  print_w
w_zero:
    mov  AL, '_'
print_w:
    call PutChar
    
    test BL, 01h
    jz   a_zero_data
    mov  AL, 'A'
    jmp  print_a_data
a_zero_data:
    mov  AL, '_'
print_a_data:
    call PutChar
add_inf: 
    mov  AL, [ESI+5]
    test AL, 10h
    jz   pre_done

    mov  AL, '+'
    call PutChar

    mov  AL, [ESI+5]
    test AL, 08h
    jz   add_r
    mov  AL, 'E'
    call PutChar
    mov  AL, 'x'
    call PutChar
    sub EDI, 2
    jmp done
add_r:
    mov  AL, 'R'
    call PutChar
    jmp done
pre_done:
    add EDI, 4
done:
    add  EDI, 10
    mov  DL, [ESI+5]
    shr  DL, 5
    and  DL, 3
    add  DL, '0'
    mov  [EDI], DL
    mov  AH, color  
    mov  byte ptr [EDI+1], AH

    add  EDI, 14
    
    mov  DL, [ESI+5]
    shr  DL, 7
    and  DL, 1
    call PrintBit
    add  EDI, 14

    mov  DL, [ESI+6]
    shr  DL, 4
    and  DL, 1
    call PrintBit
    add  EDI, 14
    
    mov  AL, [ESI+6]
    shr  AL, 6
    and  AL, 1
    cmp  AL, 1
    je   bit32
    
    mov  AL, '1'
    call PutChar
    mov  AL, '6'
    call PutChar
    jmp  done_print
    
bit32:
    mov  AL, '3'
    call PutChar
    mov  AL, '2'
    call PutChar
done_print:
endm

ENTRY_POINT:
    mov    AX,00010000b
    mov    DS,AX
    mov    ES,AX
 
    mov     EDI,00100000h
    mov     EAX,00101007h
    stosd
    mov     ECX,1023
    xor     EAX,EAX
    rep     stosd
    mov     EAX,00000007h
    mov     ECX,1024
fill_page_table:
    stosd
    add EAX,00001000h
    loop  fill_page_table 

    mov  EAX,00100000h
    mov  CR3,EAX

    mov  EAX,CR0
    or   EAX,80000000h
    mov  CR0,EAX

    mov  EAX,000B8007h
    mov  ES:00101000h+0Bh*4,EAX

    mov  EDI, 0B8000h
    mov  ESI, PM_DATA
    shl  ESI, 4
    add  ESI, offset info_msg_1
    mov  ECX, info_msg_1_len
    rep  movsw 

    mov  EDI, 0B8000h + 160
    mov  ESI, PM_DATA
    shl  ESI, 4
    add  ESI, offset info_msg_2
    mov  ECX, info_msg_2_len
    rep  movsw 

    mov  EDI, 0B8000h + 2*160
    mov  ESI, PM_DATA
    shl  ESI, 4
    add  ESI, offset info_msg_3
    mov  ECX, info_msg_3_len
    rep  movsw 

    mov  EDI, 0B8000h + 4*160
    mov  ESI, PM_DATA
    shl  ESI, 4
    add  ESI, offset titleTb
    mov  ECX, titleTb_len
    rep  movsw 

    mov  EBX, offset GDT
    mov  ECX, (GDT_size / 8)
    mov  colorRow, 03h             
    mov  EDI, 0B8000h + 160*5
gdt_loop:
    push ecx
    push EDI
    push EBX
    
    mov  eax, (GDT_size / 8)
    sub  eax, ecx
    add  al, '0'
    
    mov  ah, [colorRow]
    mov  [EDI], al
    mov  [EDI+1], ah
    add  EDI, 26 

    mov  EAX, RM_CODE
    shl  EAX, 4
    add  EAX, EBX
    
    print_descr_data EAX, [colorRow]
    
    pop  EBX
    pop  EDI
    pop  ecx
    add  EDI, 160
    add  EBX, 8
    
    inc  byte ptr [colorRow]
    dec  ecx
    jnz  gdt_loop   

    jmp  $
    
PutChar proc
    mov  [EDI], AL
    mov  byte ptr [EDI+1], AH
    add  EDI, 2
    ret
PutChar endp

PrintHexByte proc
    push ebx
    push eax
    mov  DL, AL
    shr  DL, 4
    call NibbleToHexColored
    pop  eax
    mov  DL, AL
    and  DL, 0Fh
    call NibbleToHexColored
    pop  ebx
    ret
PrintHexByte endp

PrintHexWord proc
    push ebx
    push eax
    mov  AL, AH
    call PrintHexByte
    pop  eax
    call PrintHexByte
    pop  ebx
    ret
PrintHexWord endp

PrintHexDword proc
    push ebx
    push eax
    shr  eax, 16
    call PrintHexWord
    pop  eax
    call PrintHexWord
    pop  ebx
    ret
PrintHexDword endp

NibbleToHexColored proc
    push eax
    mov  AL, DL
    cmp  AL, 9
    jbe  .digit
    add  AL, 7
.digit:
    add  AL, '0'
    mov  [EDI], AL
    mov  [EDI+1], BL
    add  EDI, 2
    pop  eax
    ret
NibbleToHexColored endp

PrintBit proc
    push eax
    mov  AL, DL
    add  AL, '0'
    mov  [EDI], AL
    mov  [EDI+1], AH
    add  EDI, 2
    pop  eax
    ret
PrintBit endp

PM_CODE  ends
 
 
; СЕГМЕНТ ДАННЫХ (для Protected Mode)
PM_DATA   segment        para public 'DATA' use32
          assume         CS:PM_DATA
info_msg_1:
irpc            info_msg_1,     <Attr: E - expansion-direction down, C - conforming, W - write enable>
                db             '&info_msg_1&',02h
endm
info_msg_2:
irpc            info_msg_2,     <      R - read enable, A - accessed, Ex - exucation enable>
                db             '&info_msg_2&',02h
endm
info_msg_3:
irpc            info_msg_3,     <      _ - no such bit (=0) or it is system segment>
                db             '&info_msg_3&',02h
endm
titleTb:
irpc            titleTb,        <Descriptor     Base     Limit     Type     Attr    Priv    Mem     User    BDep >
                db             '&titleTb&',02h
endm
error_msg:
irpc            error_msg,     <Invalid>
                db             '&error_msg&',1Ah
endm
titleTb_len     equ            80
error_msg_len   equ            7
info_msg_1_len  equ            68
info_msg_2_len  equ            58
info_msg_3_len  equ            50
colorRow        db              0
desc_buffer db 8 dup(0)
PM_DATA ends
end  @@start