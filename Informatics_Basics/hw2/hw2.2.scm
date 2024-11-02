(define (list->set xs) ;; O(|xs|^2)
  (define (f old-list new-list)
    (cond
      ((null? old-list) new-list)
      ((member (car old-list) new-list) (f (cdr old-list) new-list))
      (else (f (cdr old-list) (append new-list (list (car old-list)))))))
  (f xs '()))


(define (set? xs) ;; O(|xs|^2)
  (= (length xs) (length (list->set xs))))


(define (union xs ys) ;; O(|xs|^2)
  (define (f old-list new-list)
    (cond
      ((null? old-list) new-list)
      ((member (car old-list) new-list) (f (cdr old-list) new-list))
      (else (f (cdr old-list) (append new-list (list (car old-list)))))))
  (f xs ys))


(define (intersection xs ys) ;; O(max(|xs|, |ys|)^2)
  (define (f old-list new-list)
    (cond
      ((null? old-list) new-list)
      ((not (member (car old-list) ys)) (f (cdr old-list) new-list))
      (else (f (cdr old-list) (append new-list (list (car old-list)))))))
  (f xs '()))


(define (difference xs ys) ;; O(max(|xs|, |ys|)^2)
  (define (f old-list new-list)
    (cond
      ((null? old-list) new-list)
      ((member (car old-list) ys) (f (cdr old-list) new-list))
      (else (f (cdr old-list) (append new-list (list (car old-list)))))))
  (f xs '()))


(define (symmetric-difference xs ys) ;; O(max(|xs|, |ys|)^2)
  (union (difference xs ys) (difference ys xs)))


(define (set-eq? xs ys)  ;; O(max(|xs|, |ys|)^2)
  (and (= (length xs) (length ys)) (null? (difference xs ys))))