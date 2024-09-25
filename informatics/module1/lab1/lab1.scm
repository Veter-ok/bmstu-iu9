#lang R5RS
(define (my-even? x) (if (= (remainder x 2) 0) #t #f))
(define (my-remainder a b) (if (< a b) a (my-remainder (- a b) b)))