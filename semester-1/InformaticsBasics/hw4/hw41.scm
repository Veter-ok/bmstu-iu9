(define (read-words)
  (let loop ((pre-char #\space) (char (read-char)) (ans '()) (word '()))
    (cond ((eof-object? char)
           (if (null? word)
               ans
               (append ans (list (list->string word)))))
          ((char-whitespace? char)
           (if (char-whitespace? pre-char)
               (loop char (read-char) ans word)
               (loop char (read-char) (append ans (list (list->string word))) '())))
          (else (loop char (read-char) ans (append word (list char)))))))

(read-words)