(load "unit-tests.scm")

(define (make-stream items . eos)
  (if (null? eos)
      (make-stream items #f)
      (list items (car eos))))

(define (peek stream)
  (if (null? (car stream))
      (cadr stream)
      (caar stream)))

(define (next stream)
  (let ((n (peek stream)))
    (if (not (null? (car stream)))
        (set-car! stream (cdr (car stream))))
    n))

(define (chars->integer xs)
  (define (create-num xs num)
    (cond ((null? xs) num)
          ((equal? #\/ (car xs))
           (/ num (chars->integer (cdr xs))))
          (else
           (create-num (cdr xs)
                       (+ (* num 10)
                          (- (char->integer (car xs))
                             (char->integer #\0)))))))
  (if (equal? (car xs) #\-)
      (- (create-num (cdr xs) 0))
      (create-num (if (member (car xs) '(#\- #\+))
                      (cdr xs)
                      xs)
                  0)))

;; <frac> ::= <signed-int> "/" <unsigned-int>
;; <signed-int> ::= +<unsigned-int> | -<unsigned-int> | <unsigned-int>
;; <unsigned-int> ::= DIGIT <unsigned-int> | DIGIT
(define (valid-frac? string)
  ;; <frac> ::= <signed-int> "/" <unsigned-int>
  (define (frac stream error)
    (and  (signed-num stream error)
          (if (equal? (peek stream) #\/)
              (next stream)
              (error #f)) 
          (unsigned-num stream error)))
  ;; <signed-int> ::= +<unsigned-int> | -<unsigned-int> | <unsigned-int>
  (define (signed-num stream error)
    (cond ((equal? #\+ (peek stream))
           (next stream)
           (unsigned-num stream error))
          ((equal? #\- (peek stream))
           (next stream)
           (unsigned-num stream error))
          (else (unsigned-num stream error))))
  ;; <unsigned-int> ::= DIGIT <unsigned-int> | DIGIT
  (define (unsigned-num stream error)
    (cond ((and (char? (peek stream))
                (char-numeric? (peek stream)))
           (next stream)
           (unsigned-num stream error))
          (else #t)))
  (define stream (make-stream (string->list string) 'EOF))
  (call-with-current-continuation
   (lambda (error)
     (frac stream error)
     (eqv? (peek stream) 'EOF))))


;; <fracs> ::= <space><frac><fracs> | <empty>
;; <frac> ::= <signed-int> "/" <unsigned-int>
;; <signed-int> ::= +<unsigned-int> | -<unsigned-int> | <unsigned-int>
;; <unsigned-int> ::= DIGIT <unsigned-int> | DIGIT
;; <space> ::= SPACE <spaces> | <empty>
(define (valid-many-fracs? string)
  ;; <fracs> ::= <space><frac><fracs> | <empty>
  (define (frac-list stream error)
    (or (equal? (peek stream) 'EOF)
        (and (spaces stream error)
             (frac stream error)
             (frac-list stream error))))
  ;; <space> ::= SPACE <spaces> | <empty>
  (define (spaces stream error)
    (cond ((and (char? (peek stream))
                (char-whitespace? (peek stream)))
           (next stream)
           (spaces stream error))
          (else #t)))
  ;; <frac> ::= <signed-int> "/" <unsigned-int>
  (define (frac stream error)
    (and  (signed-num stream error)
          (if (equal? (peek stream) #\/)
              (next stream)
              (error #f)) 
          (unsigned-num stream error)))
  ;; <signed-int> ::= +<unsigned-int> | -<unsigned-int> | <unsigned-int>
  (define (signed-num stream error)
    (cond ((equal? #\+ (peek stream))
           (next stream)
           (unsigned-num stream error))
          ((equal? #\- (peek stream))
           (next stream)
           (unsigned-num stream error))
          (else (unsigned-num stream error))))
  ;; <unsigned-int> ::= DIGIT <unsigned-int> | DIGIT
  (define (unsigned-num stream error)
    (cond ((and (char? (peek stream))
                (char-numeric? (peek stream)))
           (next stream)
           (unsigned-num stream error))
          (else #t)))
  (define stream (make-stream (string->list string) 'EOF))
  (call-with-current-continuation
   (lambda (error)
     (define res (frac-list stream error))
     (and (frac-list stream error) (eqv? (peek stream) 'EOF)))))


;; <frac> ::= <signed-int> "/" <unsigned-int>
;; <signed-int> ::= +<unsigned-int> | -<unsigned-int> | <unsigned-int>
;; <unsigned-int> ::= DIGIT <unsigned-int> | DIGIT
(define (scan-frac string)
  ;; <frac> ::= <signed-int> "/" <unsigned-int>
  (define (frac stream error)
    (define part1 (signed-num stream error))
    (if (equal? (peek stream) #\/)
        (next stream)
        (error #f)) 
    (append part1 '(#\/) (unsigned-num stream error)))
  ;; <signed-int> ::= +<unsigned-int> | -<unsigned-int> | <unsigned-int>
  (define (signed-num stream error)
    (cond ((equal? #\+ (peek stream))
           (next stream)
           (unsigned-num stream error))
          ((equal? #\- (peek stream))
           (next stream)
           (cons #\- (unsigned-num stream error)))
          (else (unsigned-num stream error))))
  ;; <unsigned-int> ::= DIGIT <unsigned-int> | DIGIT
  (define (unsigned-num stream error)
    (cond ((and (char? (peek stream))
                (char-numeric? (peek stream)))
           (cons (next stream)
                 (unsigned-num stream error)))
          (else '())))
  (define stream (make-stream (string->list string) 'EOF))
  (call-with-current-continuation
   (lambda (error)
     (define res (frac stream error))
     (if (eqv? (peek stream) 'EOF) (chars->integer res) #f))))


;; <fracs> ::= <space><frac><fracs> | <empty>
;; <frac> ::= <signed-int> "/" <unsigned-int>
;; <signed-int> ::= +<unsigned-int> | -<unsigned-int> | <unsigned-int>
;; <unsigned-int> ::= DIGIT <unsigned-int> | DIGIT
;; <space> ::= SPACE <spaces> | <empty>
(define (scan-many-fracs str)
  ;; <fracs> ::= <space><frac><fracs> | <empty>
  (define (frac-list stream error)
    (cond ((and (char? (peek stream))
                (or (char-whitespace? (peek stream))
                    (equal? (peek stream) #\+)
                    (equal? (peek stream) #\-)
                    (char-numeric? (peek stream))))
           (spaces stream error)
           (let ((new-frac (frac stream error)))
             (spaces stream error)
             (cons new-frac (frac-list stream error))))
          (else '())))
  ;; <space> ::= SPACE <spaces> | <empty>
  (define (spaces stream error) 
    (cond ((and (char? (peek stream))
                (char-whitespace? (peek stream)))
           (next stream)
           (spaces stream error))
          (else #t)))
  ;; <frac> ::= <signed-int> "/" <unsigned-int>
  (define (frac stream error)
    (define part1 (signed-num stream error))
    (if (equal? (peek stream) #\/)
        (next stream)
        (error #f)) 
    (append part1 '(#\/) (unsigned-num stream error)))
  ;; <signed-int> ::= +<unsigned-int> | -<unsigned-int> | <unsigned-int>
  (define (signed-num stream error)
    (cond ((equal? #\+ (peek stream))
           (next stream)
           (unsigned-num stream error))
          ((equal? #\- (peek stream))
           (next stream)
           (cons #\- (unsigned-num stream error)))
          (else (unsigned-num stream error))))
  ;; <unsigned-int> ::= DIGIT <unsigned-int> | DIGIT
  (define (unsigned-num stream error)
    (cond ((and (char? (peek stream))
                (char-numeric? (peek stream)))
           (cons (next stream)
                 (unsigned-num stream error)))
          (else '())))
  (define stream (make-stream (string->list str) 'EOF))
  (call-with-current-continuation
   (lambda (error)
     (define res (frac-list stream error))
     (define (list->list-integer xs)
       (if (null? xs)
           '()
           (cons (chars->integer (car xs))
                 (list->list-integer (cdr xs)))))
     (and (eqv? (peek stream) 'EOF)
          (list->list-integer res)))))


;; <Program>  ::= <Articles> <Body> .
;; <Articles> ::= <Article> <Articles> | .
;; <Article>  ::= define word <Body> end .
;; <Body>     ::= if <Body> endif <Body>
;;             | while <Body> do <Body> wend <Body>
;;             | integer <Body> | word <Body> | .

(define (valid? xs)
  (define (in-stream? stream element error)
    (if (equal? (peek stream) element)
        (next stream)
        (error #f)))
  (define (word? stream error)
    (if (member (peek stream) '(define end if endif while do wend))
        (error #f)
        #t))
  ;; <Program>  ::= <Articles> <Body> .
  (define (progam stream error)
    (and (articles stream error)
         (body stream error)))
  ;; <Articles> ::= <Article> <Articles> | .
  (define (articles stream error)
    (cond ((equal? (peek stream) 'define)
           (article stream error)
           (articles stream error))
          (else #t)))
  ;; <Article>  ::= define word <Body> end .
  (define (article stream error) 
    (cond ((equal? (peek stream) 'define)
           (in-stream? stream 'define error)
           (word? stream error)
           (body stream error)
           (in-stream? stream 'end error))
          (else #t)))
  ;; <Body> ::= <if> | <while> | integer <Body> | word <Body> | .
  (define (body stream error)
    (cond ((null? (car stream)) #t)
          ((equal? (peek stream) 'if)
           (scan-if stream error))
          ((equal? (peek stream) 'while)
           (scan-while stream error))
          ((number? (peek stream))
           (next stream)
           (body stream error))
          ((word? stream error)
           (next stream)
           (body stream error))
          (else #t)))
  ;; <if> ::= if <Body> endif <Body>
  (define (scan-if stream error)
    (and (in-stream? stream 'if error)
         (body stream error)
         (in-stream? stream 'endif error)
         (body stream error)))
  ;; <while> ::= while <Body> do <Body> wend <Body>
  (define (scan-while stream error)
    (and (in-stream? stream 'while error)
         (body stream error)
         (in-stream? stream 'do error)
         (body stream error)
         (in-stream? stream 'wend error)
         (body stream error)))
  (define stream (make-stream (vector->list xs) 'EOF))
  (call-with-current-continuation
   (lambda (error)
     (progam stream error)
     (eqv? (peek stream) 'EOF))))


;; <Program>  ::= <Articles> <Body> .
;; <Articles> ::= <Article> <Articles> | .
;; <Article>  ::= define word <Body> end .
;; <Body>     ::= if <Body> endif <Body>
;;             | while <Body> do <Body> wend <Body>
;;             | integer <Body> | word <Body> | .
(define (parse xs)
  (define (in-stream? stream element error)
    (if (equal? (peek stream) element)
        (next stream)
        (error #f)))
  (define (word? stream error)
    (if (member (peek stream) '(define end if endif while do wend))
        #f
        '()))
  ;; <Program>  ::= <Articles> <Body> .
  (define (progam stream error)
    (define arts (articles stream error))
    (if (null? arts)
        (cons '() (list (body stream error)))
        (list (append arts
                      (list (body stream error))))))
  ;; <Articles> ::= <Article> <Articles> | .
  (define (articles stream error)
    (cond ((equal? (peek stream) 'define)
           (append (list (article stream error))
                   (articles stream error)))
          (else '())))
  ;; <Article>  ::= define word <Body> end .
  (define (article stream error)
    (if (equal? (peek stream) 'define)
        (append
         (if (in-stream? stream 'define error) '())
         (if (word? stream error)
             (list (next stream) (body stream error)))
         (if (in-stream? stream 'end error) '())
         '())))
  ;; <Body> ::= <if> | <while> | integer <Body> | word <Body> | .
  (define (body stream error)
    (cond ((null? (car stream)) '())
          ((equal? (peek stream) 'if)
           (append (list (scan-if stream error)) (body stream error)))
          ((equal? (peek stream) 'while)
           (append (list (scan-while stream error)) (body stream error)))
          ((number? (peek stream))
           (cons (next stream) (body stream error)))
          ((word? stream error)
           (cons (next stream) (body stream error)))
          (else '())))
  ;; <if> ::= if <Body> endif <Body>
  (define (scan-if stream error)
    (append (list (in-stream? stream 'if error))
            (list (body stream error))
            (if (in-stream? stream 'endif error) '())))
  ;; <while> ::= while <Body> do <Body> wend <Body>
  (define (scan-while stream error)
    (append (list (in-stream? stream 'while error))
            (list (body stream error))
            (if (in-stream? stream 'do error)
                (list (body stream error)))
            (if (equal? (peek stream) 'wend)
                (begin (next stream) '())
                '())))
  (define stream (make-stream (vector->list xs) 'EOF))
  (call-with-current-continuation
   (lambda (error)
     (define res  (progam stream error))
     (and (eqv? (peek stream) 'EOF)
          res))))


(define the-tests
  (list (test (valid-frac? "110/111") #t)
        (test (valid-frac? "-4/3")  #t)
        (test (valid-frac? "+5/10") #t)
        (test (valid-frac? "5.0/10") #f)
        (test (valid-frac? "FF/10") #f)
        (test (valid-frac? "12/+10") #f)
        (test (valid-many-fracs? "\t1/2 1/3\n\n0/8") #t)
        (test (valid-many-fracs? "\t1/2 1/3\n\n10/8") #t)
        (test (valid-many-fracs? "\t1/2 1/3\n\n2/-5") #f)
        (test (valid-many-fracs? "+1/2-3/4") #t)
        (test (scan-frac "110/111") 110/111)
        (test (scan-frac "-4/3")  -4/3)
        (test (scan-frac "+5/10")  1/2)
        (test (scan-frac "5.0/10") #f)
        (test (scan-frac "FF/10") #f)
        (test (scan-frac "2/-5") #f)
        (test (scan-many-fracs "\t1/2 1/3\n\n10/8") (1/2 1/3 5/4))
        (test (scan-many-fracs "\t1/2 1/3\n\n2/-5") #f)
        (test (scan-many-fracs "+1/2-3/4") (1/2 -3/4))
        (test (valid? #(1 2 +)) #t)
        (test (valid? #(define 1 2 end)) #f)
        (test (valid? #(define x if end endif)) #f)
        (test (parse #(1 2 +)) (() (1 2 +)))
        (test (parse #(x dup 0 swap if drop -1 endif)) (() (x dup 0 swap (if (drop -1)))))
        (test (parse #(x dup while dup 0 > do 1 - swap over * swap)) (() (x dup (while (dup 0 >) (1 - swap over * swap)))))
        (test (parse #(define abs dup 0 < if neg endif end 9 abs -9 abs)) (((abs (dup 0 < (if (neg)))) (9 abs -9 abs))))
        (test (parse #( define -- 1 - end
                         define =0? dup 0 = end
                         define =1? dup 1 = end
                         define factorial
                         =0? if drop 1 exit endif
                         =1? if drop 1 exit endif
                         1 swap
                         while dup 0 > do
                         1 - swap over * swap
                         wend
                         drop
                         end
                         0 factorial
                         1 factorial
                         2 factorial
                         3 factorial
                         4 factorial )) (((-- (1 -))
                                          (=0? (dup 0 =))
                                          (=1? (dup 1 =))
                                          (factorial
                                           (=0? (if (drop 1 exit)) =1? (if (drop 1 exit))
                                                1 swap (while (dup 0 >) (1 - swap over * swap)) drop))
                                          (0 factorial 1 factorial 2 factorial 3 factorial 4 factorial))))
        (test (parse #(define word w1 w2 w3)) #f)))

(run-tests the-tests)