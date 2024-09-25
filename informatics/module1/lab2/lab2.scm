(define (uniq xs)
  (if (null? xs)
      '()
      (if (null? (cdr xs))
          xs
          (if (equal? (car xs) (cadr xs))
              (uniq (cdr xs))
              (cons (car xs) (uniq (cdr xs)))))))


(define (delete pred? xs)
  (if (null? xs)
      '()
      (if (not (pred? (car xs)))
          (cons (car xs) (delete pred? (cdr xs)))
          (delete pred? (cdr xs)))))


(define (polynom coef x)
  (if (null? coef)
      '()
      (if (null? (cdr coef))
          (car coef)
          (+ (* (car coef) (expt x (- (length coef) 1))) (polynom (cdr coef) x)))))


(define (intersperse e xs)
  (if (null? xs)
      '()
      (if (null? (cdr xs)) xs
          (append (cons (car xs) (cons e '())) (intersperse e (cdr xs))))))


(define (all? pred? xs)
  (or (null? xs)
      (and (pred? (car xs))
           (all? pred? (cdr xs)))))


(define (f x) (+ x 2))
(define (g x) (* x 3))
(define (h x) (- x))


(define (o . funcs)
  (if (null? funcs)
      (lambda (x) x)
      (lambda (x) ((car funcs) (apply (o (cdr funcs)) x)))))