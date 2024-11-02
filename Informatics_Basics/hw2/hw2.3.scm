(define (string-trim-left str) ;; O(|str|^2)
  (define (f sttr)
    (if (equal? sttr "")
        sttr
        (if (or (equal? (string-ref sttr 0) #\tab)
                (equal? (string-ref sttr 0) #\space)
                (equal? (string-ref sttr 0) #\newline))
            (f (substring sttr 1))
            sttr)))
  (f str))


(define (string-trim-right str) ;; O(|str|^2)
  (define (f sttr len)
    (if (equal? sttr "")
        sttr
        (if (or (equal? (string-ref sttr len) #\tab)
                (equal? (string-ref sttr len) #\space)
                (equal? (string-ref sttr len) #\newline))
            (f sttr (- len 1))
            (substring sttr 0 (+ len 1)))))
  (f str (- (string-length str) 1)))


(define (string-trim str) (string-trim-right (string-trim-left str))) ;; O(|str|^2)


(define (string-prefix? a b) ;; O(|a|)
  (let ((lenA (string-length a))
        (lenB (string-length b)))
    (and (>= lenB lenA)
         (equal? a (substring b 0 lenA)))))


(define (string-suffix? a b) ;; O(|a|)
  (let ((lenA (string-length a))
        (lenB (string-length b)))
    (and (>= lenB lenA)
         (equal? a (substring b (- lenB lenA))))))


(define (string-infix? a b) ;; O(|a|*(|b| - |a|))
  (define (f index lenA lenB)
    (and (not (or (< lenB lenA) (> (+ index lenA) lenB)))
         (or (equal? a (substring b index (+ index lenA)))
             (f (+ index 1) lenA lenB))))
  (f 0 (string-length a) (string-length b)))


(define (string-split str sep) ;; O(|str| * |sep|)
  (define (f index last-index)
    (cond ((< (- (string-length str) index) 2) (list (substring str last-index)))
          ((not (equal? sep (substring str index (+ index (string-length sep))))) (f (+ index 1) last-index))
          (else (append (list(substring str last-index index))
                        (f (+ index (string-length sep)) (+ index (string-length sep)))))))
  (f 0 0))