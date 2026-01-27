assume cs: code, ds: data

data segment
a db 10
b db 3
c db 2
d db 2
res1 db ?
res2 db ?
res  db ?
data ends

code segment
start:	mov ax, data
        mov ds, ax

        mov al, a
        mul b
        mov res1, al

        mov al, c
        mul d
        mov res2, al

        shr res2, 1
        shr res2, 1
        
        mov al, res1
        mov bl, res2
        sub al, bl

        mov res, al
code ends
end start