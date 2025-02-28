(load "unit-tests.scm")

(define (interpret program stack)
  (let ((xs '((+ +) (- -) (* *) (/ /) (mod remainder)))
        (defines '()))
    (define (add-define program name cur-define)
      (if (equal? (car program) 'end)
          (begin
            (set! defines (cons (list name (append (cons #t cur-define) (list (car program)))) defines))
            program)
          (add-define (cdr program) name (append cur-define (list (car program))))))
    (define (exit-define program) (if (equal? (car program) 'end) program (exit-define (cdr program))))
    (define (end-if program) (if (equal? (car program) 'endif) program (end-if (cdr program))))
    (do ((program (vector->list program) (cdr program))) ((null? program) stack)
      (cond ((null? program) stack)
            ((assoc (car program) defines) (set! program (append (cadr (assoc (car program) defines)) (cdr program))))
            ((equal? (car program) 'define) (set! program (add-define (cddr program) (cadr program) '())))
            ((equal? (car program) 'exit) (set! program (exit-define program)))
            ((number? (car program)) (set! stack (cons (car program) stack)))
            ((member (car program) '(+ - * / mod))
             (let ((oper (eval (cadr (assoc (car program) xs)) (interaction-environment))))
               (set! stack (cons (oper (cadr stack) (car stack)) (cddr stack)))))
            ((member (car program) '(= > <))
             (let ((oper (eval (car program) (interaction-environment))))
               (set! stack (cons (if (oper (cadr stack) (car stack)) -1 0) (cddr stack)))))
            ((equal? (car program) 'neg) (set! stack (cons (- (car stack)) (cdr stack))))
            ((equal? (car program) 'and)
             (set! stack (cons (if (and (not (= (car stack) 0)) (not (= (cadr stack) 0))) -1 0)
                               (cddr stack))))
            ((equal? (car program) 'or)
             (set! stack (cons (if (or  (not (= (car stack) 0)) (not (= (cadr stack) 0))) -1 0)
                               (cddr stack))))
            ((equal? (car program) 'not) (set! stack (cons ((if (= (car stack) 0) -1 0)) (cdr stack))))
            ((equal? (car program) 'drop) (set! stack (cdr stack)))
            ((equal? (car program) 'swap) (set! stack (cons (cadr stack) (cons (car stack) (cddr stack)))))
            ((equal? (car program) 'dup)  (set! stack (cons (car stack) stack)))
            ((equal? (car program) 'over) (set! stack (cons (cadr stack) stack)))
            ((equal? (car program) 'rot)  (set! stack (cons (car stack) (cons (cadr stack) (cddr stack)))))
            ((equal? (car program) 'depth)(set! stack (cons (length stack) stack)))
            ((equal? (car program) 'if) (if (not (= (car stack) 0))
                                            (set! stack (cdr stack))
                                            (begin (set! stack (cdr stack))
                                                   (set! program (end-if program)))))))))


(define the-tests
  (list (test (interpret #(2 3 * 4 5 * +) '()) (26))
        (test (interpret #(define -- 1 - end  5 -- --) '()) (3))
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
                            234 8100 gcd) '()) (18 9))))

(run-tests the-tests)