;; Задача 1
(define-syntax trace
  (syntax-rules ()
    ((trace code)
     (let ((value code))
       (write 'code)
       (display " => ")
       (write value) 
       (newline)
       value)))) 

(define (zip . xss)
  (if (or (null? xss) (null? (trace (car xss)))) 
      '()
      (cons (map car xss) (apply zip (map cdr (trace xss)))))) 

;; Задача 2
(define-syntax test
  (syntax-rules ()
    ((test code res) '(code res))))

(define (run-test test)
  (let* ((returned-result (eval (car test) (interaction-environment)))
         (result-test
          (if (equal? (cdr test) (list returned-result))
              (begin (write (car test))
                     (display " ok\n")
                     #t)
              (begin (write (car test))
                     (display " FAIL\n")
                     (display "  Expected: ")
                     (write (car (cdr test)))
                     (newline)
                     (display "  Returned: ")
                     (write returned-result)
                     (newline)
                     #f))))
    result-test))

(define (run-tests list-tests)
  (define (f xs)
    (if (null? xs)
        '()
        (cons (run-test (car xs)) (f (cdr xs)))))
  (not (member #f (f list-tests))))

(define (signum x)
  (cond
    ((< x 0) -1)
    ((= x 0)  1)
    (else     1)))

(define the-tests
  (list (test (signum -2) -1)
        (test (signum  0)  0)
        (test (signum  2)  1)))

(run-tests the-tests)

(define counter
  (let ((n 0))
    (lambda ()
      (set! n (+ n 1))
      n)))

(+ (trace (counter))
   (trace (counter)))

(define counter-tests
  (list (test (counter) 3)
        (test (counter) 77)
        (test (counter) 5)))

(run-tests counter-tests)

;; Задача 3
(define (ref xs index . new-element)
  (cond
    ((list? xs)
     (if (null? new-element)
         (and (and (>= index 0) (< index (length xs)))
              (list-ref xs index))
         (and (and (>= index 0) (<= index (length xs)))
              (if (= index 0)
                  (cons (car new-element) xs)
                  (cons (car xs) (ref (cdr xs) (- index 1) (car new-element)))))))
    ((vector? xs)
     (if (null? new-element) 
         (and (and (>= index 0) (< index (vector-length xs)))
              (vector-ref xs index))
         (and (and (>= index 0) (<= index (vector-length xs)))
              (let ((new-vector (make-vector (+ (vector-length xs) 1))))
                (define (copy-vector old new i)
                  (if (< i (vector-length new))
                      (cond
                        ((= i index)
                         (begin
                           (vector-set! new i (car new-element))
                           (copy-vector old new (+ i 1))))
                        ((> i index)
                         (begin
                           (vector-set! new i (vector-ref old (- i 1)))
                           (copy-vector old new (+ i 1))))
                        (else
                         (begin
                           (vector-set! new i (vector-ref old i))
                           (copy-vector old new (+ i 1)))))))
                (copy-vector xs new-vector 0)
                new-vector))))
    ((string? xs)
     (if (null? new-element)
         (and (and (>= index 0) (< index (string-length xs)))
              (string-ref xs index))
         (and (and (>= index 0) (<= index (string-length xs)) (char? (car new-element)))
              (string-append
               (substring xs 0 index)
               (string (car new-element))
               (substring xs index (string-length xs))))))))

;; Задание 4
(define (if->cond expr)
  (define (f code count-if)
    (cond
      ((and (= count-if 0) (= (length (cdr code)) 2))
       `(,(cadr code) ,(caddr code)))
      ((= (length (cdr code)) 2)
       `((,(cadr code) ,(caddr code))))
      ((and (list? (cadddr code)) (= count-if 0))
       `(cond (,(cadr code) ,(caddr code)) ,@(f (car (cdddr code)) 1)))
      ((list? (cadddr code))
       `((,(cadr code) ,(caddr code)) ,@(f (car (cdddr code)) 1)))
      ((and (not (list? (cadddr code))) (= count-if 0))
       `(cond (,(cadr code) ,(caddr code)) (else ,(cadddr code))))
      ((not (list? (cadddr code)))
       `((,(cadr code) ,(caddr code)) (else ,(cadddr code))))))
  (f expr 0))


(define if->cond-tests
  (list (test (if->cond '(if (> x 0) +1 (if (< x 0) -1 0))) (cond ((> x 0) 1) ((< x 0) -1) (else 0)))
        (test (if->cond '(if (equal? (car expr) 'lambda)
                             (compile-lambda expr)
                             (if (equal? (car expr) 'define)
                                 (compile-define expr)
                                 (if (equal? (car expr) 'if)
                                     (compile-if expr)))))
              (cond ((equal? (car expr) 'lambda) (compile-lambda expr))
                    ((equal? (car expr) 'define) (compile-define expr))
                    ((equal? (car expr) 'if) (compile-if expr))))))

(run-tests if->cond-tests)