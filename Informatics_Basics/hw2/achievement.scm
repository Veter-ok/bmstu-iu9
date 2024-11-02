(define (list-trim-right lst)
  (define (find-index xs pre-result result)
    (cond
      ((null? xs) (if (null? result) 0 result))
      ((char-whitespace? (car xs)) (find-index (cdr xs) (+ pre-result 1) result))
      (else (find-index (cdr xs) (+ pre-result 1) (+ pre-result 1)))))
  (define (new-list xs index)
    (if (= index 0) '() (cons (car xs) (new-list (cdr xs) (- index 1)))))
  (new-list lst (find-index lst 0 0)))