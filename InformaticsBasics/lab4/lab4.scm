;; Задание 1
(define a #f)

(define-syntax use-assertions
  (syntax-rules ()
    ((_)
     (call-with-current-continuation (lambda (cc)
                                       (set! a cc))))))

(define-syntax assert
  (syntax-rules ()
    ((_ expr)
     (if expr
         #t
         (begin
           (display "FAILED: ")
           (a (write 'expr)))))))

(use-assertions)

(define (1/x x)
  (assert (not (zero? x)))
  (/ 1 x))

(map 1/x '(1 2 3 4 5))
(map 1/x '(-2 -1 0 1 2))
(newline)

;; Задание 2
(define (select-rec source dest)
  (call-with-input-file source
    (lambda (port1)
      (call-with-output-file dest
        (lambda (port2) (select-rec-loop port1 port2))))))

(define (cool-member element lst)
  (cond ((null? lst) #f)
        ((equal? element (car lst)) #t)
        ((list? (car lst))
         (or (cool-member element (car lst))
             (cool-member element (cdr lst))))
        (else (cool-member element (cdr lst)))))

(define (select-rec-loop input output)
  (let loop ((expr (read input)))
    (if (not (eof-object? expr))
        (begin
          (if (and (list? expr)
                   (equal? (car expr) 'define)
                   (list? (cadr expr)))
              (if (cool-member (caadr expr) (cddr expr))
                  (begin
                    (write expr output)
                    (newline output))))
          (loop (read input))))))

;;(select-rec "lab4.scm" "result.scm")

(define (text-stat source)
  (call-with-input-file source check-file))

(define (check-file port)
  (let ((words 0) (nums 0) (formulas 0))
    (let loop ((cur-char (read-char port)) (letters 0) (digits 0) (math-symbols 0))
      (cond
        ((or (eof-object? cur-char) (char-whitespace? cur-char))
         (begin
           (cond
             ((and (> letters 0) (= digits 0) (= math-symbols 0)) (set! words (+ words 1)))
             ((and (= letters 0) (> digits 0) (= math-symbols 0)) (set! nums (+ nums 1)))
             (else (set! formulas (+ formulas 1))))
           (if (not (eof-object? cur-char)) (loop (read-char port) 0 0 0))))
        ((char-numeric? cur-char)
         (loop (read-char port) letters (+ digits 1) math-symbols))
        ((member cur-char '(#\= #\+ #\- #\* #\/))
         (loop (read-char port) letters digits (+ math-symbols 1)))
        (else (loop (read-char port) (+ letters 1) digits math-symbols))))
    (begin
      (display "Слова:   ")
      (display words)
      (newline)
      (display "Числа:   ")
      (display nums)
      (newline)
      (display "Формулы: ")
      (display formulas)
      (newline))))

;;(text-stat "text.txt")

;; Задание 3
(define (trib n)
  (cond ((< n 2) 0)
        ((= n 3) 1)
        (else (+ (trib (- n 3)) (trib (- n 2)) (trib (- n 1))))))

(define (trib-memo n)
  (let ((known-results '()))
    (let loop ((num n))
      (cond ((< num 2) 0)
            ((= num 3) 1)
            (else
             (let* ((res (assoc num known-results)))
               (if res
                   (cadr res)
                   (let ((res (+ (loop (- num 1)) (loop (- num 2)) (loop (- num 3)))))
                     (set! known-results
                           (cons (list num res) known-results))
                     res))))))))

;; Задание 4
(define-syntax lazy-cons
  (syntax-rules ()
    ((lazy-cons a b)
     (cons a (delay b)))))

(define (lazy-car p) (car p))
(define (lazy-cdr p) (force (cdr p)))

(define (lazy-head xs k)
  (if (<= k 0)
      '()
      (cons (car xs) (lazy-head (lazy-cdr xs) (- k 1)))))

(define (lazy-ref xs k)
  (if (= k 0)
      (lazy-car xs)
      (lazy-ref (lazy-cdr xs) (- k 1))))

(define (lazy-map proc xs)
  (lazy-cons (proc (lazy-car xs))
             (lazy-map proc (lazy-cdr xs))))

(define (lazy-zip xs ys)
  (if (or (null? xs) (null? ys))
      '()
      (lazy-cons (list (lazy-car xs) (lazy-car ys))
                 (lazy-zip (lazy-cdr xs) (lazy-cdr ys)))))


(define naturals (lazy-cons 1 (lazy-map (lambda (x) (+ x 1)) naturals)))
(define factorials
  (lazy-cons 1
             (lazy-map 
              (lambda (pair) (* (car pair) (cadr pair)))
              (lazy-zip naturals factorials))))

;; Задание 5
(define-syntax unless
  (syntax-rules ()
    ((unless cond? . exprs)
     (if (not cond?) (begin . exprs)))))

(define x 1)
(unless (= x 0) (display "x != 0") (newline))

(define-syntax while
  (syntax-rules ()
    ((while cond? . exprs)
     (let loop () 
       (if cond? (begin (begin . exprs) (loop)))))))

(let ((p 0)
      (q 0))
  (while (< p 3)
         (set! q 0)
         (while (< q 3)
                (display (list p q))
                (newline)
                (set! q (+ q 1)))
         (set! p (+ p 1))))

(define-syntax cout
  (syntax-rules (<< endl)
    ((cout << endl) (newline))
    ((cout << endl . exprs) (begin (newline) (cout . exprs)))
    ((cout << expr) (display expr))
    ((cout << expr . exprs) (begin (display expr) (cout . exprs)))))

(cout << "a = " << 1 << endl << "b = " << 2 << endl)