(load "unit-tests.scm")

(define (interpret program stack)
  (let ((xs '((+ +) (- -) (* *) (/ /) (mod remainder))) (defines '()) (nested-if 0))
    (define (add-define program name cur-define)
      (if (equal? (car program) 'end)
          (begin
            (set! defines (cons (list name (append cur-define (list (car program)))) defines))
            (cdr program))
          (add-define (cdr program) name (append cur-define (list (car program))))))
    (define (exit-define program) (if (equal? (car program) 'end) (cdr program) (exit-define (cdr program))))
    (define (else-if program) (if (member (car program) '(else endif)) (cdr program) (else-if (cdr program))))
    (define (machine program stack)
      (cond ((null? program) stack)
            ((assoc (car program) defines) (machine (cdr program) (machine (cadr (assoc (car program) defines)) stack)))
            ((equal? (car program) 'define) (machine (add-define (cddr program) (cadr program) '()) stack))
            ((equal? (car program) 'exit) (machine (exit-define program) stack))
            ((number? (car program)) (machine (cdr program) (cons (car program) stack)))
            ((member (car program) '(+ - * / mod))
             (let ((oper (eval (cadr (assoc (car program) xs)) (interaction-environment))))
               (machine (cdr program) (cons (oper (cadr stack) (car stack)) (cddr stack)))))
            ((member (car program) '(= > <))
             (let ((oper (eval (car program) (interaction-environment))))
               (machine (cdr program) (cons (if (oper (cadr stack) (car stack)) -1 0) (cddr stack)))))
            ((equal? (car program) 'neg) (machine (cdr program) (cons (- (car stack)) (cdr stack))))
            ((equal? (car program) 'and) (machine (cdr program) (cons (if (and (not (= (car stack) 0)) (not (= (cadr stack) 0))) -1 0) (cddr stack))))
            ((equal? (car program) 'or)  (machine (cdr program) (cons (if (or  (not (= (car stack) 0)) (not (= (cadr stack) 0))) -1 0) (cddr stack))))
            ((equal? (car program) 'not) (machine (cdr program) (cons (if (= (car stack) 0) -1 0)) (cdr stack)))
            ((equal? (car program) 'drop) (machine (cdr program) (cdr stack)))
            ((equal? (car program) 'swap) (machine (cdr program) (cons (cadr stack) (cons (car stack) (cddr stack)))))
            ((equal? (car program) 'dup)  (machine (cdr program) (cons (car stack) stack)))
            ((equal? (car program) 'over) (machine (cdr program) (cons (cadr stack) stack)))
            ((equal? (car program) 'rot)  (machine (cdr program) (cons (car stack) (cons (cadr stack) (cddr stack)))))
            ((equal? (car program) 'depth)(machine (cdr program) (cons (length stack) stack)))
            ((equal? (car program) 'else) (machine (else-if (cdr program)) stack))
            ((equal? (car program) 'if)
             (if (not (= (car stack) 0))
                 (begin (set! nested-if (+ nested-if 1)) (machine (cdr program) (cdr stack)))
                 (begin (set! nested-if (- nested-if 1)) (machine (else-if program) (cdr stack)))))
            (else (machine (cdr program) stack))))
    (machine (vector->list program) stack)))


(define the-tests
  (list (test (interpret #(2 3 * 4 5 * +) '()) (26))
        (test (interpret #(define -- 1 - end  5 -- --) '()) (3))
        (test (interpret #(define abs
                            dup 0 <
                            if neg endif
                            end
                            9 abs
                            -9 abs) (quote ())) (9 9))
        (test (interpret #(define =0? dup 0 = end
                            define <0? dup 0 < end
                            define signum
                            =0? if exit endif
                            <0? if drop -1 exit endif
                            drop
                            1
                            end
                            0 signum -5 signum 10 signum) (quote ())) (1 -1 0))
        (test (interpret #(define -- 1 - end
                            define =0? dup 0 = end
                            define =1? dup 1 = end
                            define factorial
                            =0? if drop 1 exit endif
                            =1? if drop 1 exit endif
                            dup --
                            factorial
                            *
                            end
                            0 factorial
                            1 factorial
                            2 factorial
                            3 factorial
                            4 factorial) (quote ())) (24 6 2 1 1))
        (test (interpret #(define =0? dup 0 = end
                            define =1? dup 1 = end
                            define -- 1 - end
                            define fib
                            =0? if drop 0 exit endif
                            =1? if drop 1 exit endif
                            -- dup
                            -- fib
                            swap fib
                            +
                            end
                            define make-fib
                            dup 0 < if drop exit endif
                            dup fib
                            swap --
                            make-fib
                            end
                            10 make-fib) (quote ())) (0 1 1 2 3 5 8 13 21 34 55))
        (test (interpret #(define =0? dup 0 = end
                            define gcd
                            =0? if drop exit endif
                            swap over mod
                            gcd
                            end
                            90 99 gcd
                            234 8100 gcd) '()) (18 9))
        (test (interpret #(1 if 100 else 200 endif) '()) (100))
        (test (interpret #(0 if 100 else 200 endif) '()) (200))
        (test (interpret #(0 if 1 if 2 endif 3 endif 4) '()) (4))
        (test (interpret #(1 if 2 if 3 endif 4 endif 5) '()) (5 4 3))
        (test (interpret #(1 if 0 if 2 endif 3 endif 4) '()) (4 3))))

(run-tests the-tests)