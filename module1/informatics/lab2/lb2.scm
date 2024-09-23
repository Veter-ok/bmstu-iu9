#lang R5RS
(define (count x xs)
  (if (null? xs) 0
      (if (equal? x (car xs))
          (+ (count x (cdr xs)) 1)
          (count x (cdr xs)))))

(define (delete pred? xs)
  (if (null? xs) '()
      (if (not (pred? (car xs)))
          (cons (car xs) (delete pred? (cdr xs)))
          (delete pred? (cdr xs)))))

(define (iterate f x n)
  (if (= n 0) '()
      (cons x (iterate f (f x) (- n 1)))))

(define (intersperse e xs)
  (if (null? xs) '()
      (if (null? (cdr xs)) xs
          (cons (cons (car xs) (cons e '())) (intersperse e (cdr xs))))))

(define (any? pred? xs)
  (and (not (null? xs))
       (or (pred? (car xs))
           (any? pred? (cdr xs)))))

(define (all? pred? xs)
  (or (null? xs)
      (and (pred? (car xs))
           (all? pred? (cdr xs)))))
           
(define (o . func)
  (if (null? func)
      (lambda (x) x)
      (lambda (x) ((car func) ((apply o (cdr func)) x)))))