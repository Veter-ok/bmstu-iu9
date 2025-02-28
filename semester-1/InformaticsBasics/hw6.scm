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

(define (chars->integer xs) (string->number (list->string xs)))
(define (list->symbol xs)
  (define (char-lower-case c)
    (if (and (char>=? c #\A) (char<=? c #\Z))
        (integer->char (+ (char->integer c) 32))
        c))
  (string->symbol (list->string (map char-lower-case xs))))

;;<expression> ::= <space><term> | <space><term><operator><expression>
;;<term> ::= <number> | <variable> | "(" <expression> ")"
;;<number> ::=  DIGIT <number> | DIGIT
;;<variable> ::=  LETTER <variable> | LETTER
;;<operator> ::= "+" | "-" | "*" | "/" | "^"
;;<space> ::= SPACE <spaces> | <empty>
(define (tokenize string)
  ;;<expression> ::= <term> | <term><operator><expression>
  (define (expression stream error)
    (cond ((equal? (peek stream) 'EOF) '())
          ((char-numeric? (peek stream))
           (cons (chars->integer (number stream error)) (expression stream error)))
          ((char-whitespace? (peek stream))
           (spaces stream error)
           (expression stream error))
          ((member (peek stream) '(#\+ #\- #\* #\/ #\^))
           (cons (list->symbol (list (next stream))) (expression stream error)))
          ((equal? (peek stream) #\()
           (next stream)
           (cons "(" (expression stream error)))
          ((equal? (peek stream) #\))
           (next stream)
           (cons ")" (expression stream error)))
          ((char-alphabetic? (peek stream))
           (cons (list->symbol (variable stream error)) (expression stream error)))
          (else (error #f))))
  ;;<number> ::=  DIGIT <number> | DIGIT
  (define (number stream error)
    (define (chaeckNumber wasE? wasPoint?)
      (cond ((and (char? (peek stream))
                  (char-numeric? (peek stream)))
             (cons (next stream)
                   (chaeckNumber wasE? wasPoint?)))
            ((equal? (peek stream) #\.)
             (if (or wasPoint? wasE?)
                 (error #f)
                 (cons (next stream)
                       (chaeckNumber wasE? #t))))
            ((equal? (peek stream) #\e)
             (if wasE?
                 (error #f)
                 (begin
                   (next stream)
                   (append (list #\e (if (equal? (peek stream) #\-)
                                         (begin (next stream) #\-)
                                         #\+))
                           (chaeckNumber #t wasPoint?)))))
            (else '())))
    (chaeckNumber #f #f))
  ;;<variable> ::=  LETTER <variable> | LETTER
  (define (variable stream error)
    (cond ((null? (car stream)) '())
          ((and (char? (peek stream))
                (char-alphabetic? (peek stream)))
           (cons (next stream)
                 (variable stream error)))
          ((char-numeric? (peek stream)) (error #f))
          (else '())))
  ;;<space> ::= SPACE <spaces> | <empty>
  (define (spaces stream error)
    (cond ((and (char? (peek stream))
                (char-whitespace? (peek stream)))
           (next stream)
           (spaces stream error))))
  (define stream (make-stream (string->list string) 'EOF))
  (call-with-current-continuation
   (lambda (error)
     (define res  (expression stream error))
     (and (eqv? (peek stream) 'EOF)
          res))))


(define (parse xs)
  ;; Expr ::= Term Expr'
  (define (expr stream error)
    (if (equal? (peek stream) 'EOF)
        (error #f)
        (expr+ stream error (term stream error))))
  ;; Expr' ::= AddOp Term Expr' | .
  (define (expr+ stream error left)
    (if (member (peek stream) '(+ -))
        (let ((oper (next stream))
              (right (term stream error)))
          (if (and right (not (equal? right 'EOF)))
              (expr+ stream error (list left oper right))
              (error #f)))
          left))
  ;; Term ::= Factor Term'
  (define (term stream error) (term+ stream error (factor stream error)))
  ;; Term' ::= MulOp Factor Term' | .
  (define (term+ stream error left)
    (if (member (peek stream) '(* /))
        (term+ stream error (list left (next stream) (factor stream error)))
        left))
  ;; Factor ::= Power Factor'
  (define (factor stream error) (factor+ stream error (power stream error)))
  ;; Factor' ::= PowOp Power Factor' | .
  (define (factor+ stream error left)
    (if (equal? (peek stream) '^)
        (list left (next stream) (factor stream error))
        left))
  ;; Power ::= value | "(" Expr ")" | unaryMinus Power
  (define (power stream error)
    (cond ((equal? (peek stream) "(")
           (next stream)
           (let ((expr-result (expr stream error)))
             (if (and expr-result (equal? (peek stream) ")"))
                 (begin (next stream) expr-result)
                 (error #f))))
          ((equal? (peek stream) ")") (error #f))
          ((equal? (peek stream) '-) (next stream) (list '- (power stream error)))
          (else (next stream))))
  (define stream (make-stream xs 'EOF))
  (call-with-current-continuation
   (lambda (error)
     (define res  (expr stream error))
     (and (eqv? (peek stream) 'EOF)
          res))))


(define (tree->scheme xs)
  (let ((op (if (equal? (cadr xs) '^) 'expt (cadr xs)))
        (first (if (list? (car xs)) (tree->scheme (car xs)) (car xs)))
        (second (if (list? (caddr xs)) (tree->scheme (caddr xs)) (caddr xs))))
    (list op first second)))