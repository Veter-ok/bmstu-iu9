(define (my-range a b d) ;; O((b-a)/d)
  (if (>= a b)
      '()
      (cons a (my-range (+ a d) b d))))


(define (my-flatten xs) ;; O(n) n - длина списка и всех вложенных списков 
  (define (f lst result)
    (cond
      ((null? lst) result)
      ((pair? (car lst)) (f (cdr lst) (f (car lst) result)))
      (else (f (cdr lst) (cons (car lst) result)))))
  (reverse (f xs '())))


(define (my-element? x xs) ;; O(|xs|)
  (and (not (null? xs))
       (or (equal? x (car xs))
           (my-element? x (cdr xs)))))


(define (my-filter pred? xs) ;; O(|xs|) 
  (cond ((null? xs) '())
        ((pred? (car xs)) (cons (car xs) (my-filter pred? (cdr xs))))
        (else (my-filter pred? (cdr xs)))))


(define (my-fold-left op xs) ;; O(|xs|) 
  (define (f a xxs)
    (if (null? xxs)
        a
        (f (op a (car xxs)) (cdr xxs))))
  (if (null? xs) '() (f (car xs) (cdr xs))))


(define (my-fold-right op xs) ;; O(|xs|) 
  (if (null? (cdr xs))
      (car xs)
      (op (car xs) (my-fold-right op (cdr xs)))))


(define (reverse! xs) ;; O(|xs|^2) 
  (define (f count)
    (cond ((< count 0) '())
          (else (cons (car (list-tail xs count)) (f (- count 1))))))
  (let ((ans (f (- (length xs) 1))))
    (if (not (null? xs))
        (begin
          (set-cdr! xs '())
          (set-car! xs '())))
    ans))


;;(define (reverse! xs)
;;  (define (f count len)
;;    (cond ((< count 0)
;;           (begin
;;             (set-car! xs (car (list-tail xs len)))
;;             (set-cdr! xs (cdr (list-tail xs len)))
;;             xs))
;;          (else
;;          (begin
;;             (set-cdr! xs (append (cdr xs) (list (car (list-tail xs count)))))
;;             (f (- count 1) len)))))
;;  (f (- (length xs) 1) (length xs)))


(define (append! . lists) ;; O(n) n = |xs|+|ys|+... 
  (let loop ((xs lists))
    (cond
      ((null? xs) '())
      ((null? (cdr xs)) (car xs))
      ((null? (car xs)) (loop (cdr xs)))
      (else
       (let ((last (car xs)))
         (define (find-last lst)
           (if (null? (cdr lst))
               lst
               (find-last (cdr lst))))
         (set-cdr! (find-last last) (loop (cdr xs)))
         (car xs))))))