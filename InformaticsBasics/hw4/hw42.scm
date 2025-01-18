(define-syntax define-struct
  (syntax-rules ()
    ((define-struct name (fields ...))
     (begin
       (eval `(define
              ,(string->symbol (string-append "make-" (symbol->string 'name)))
              (lambda (fields ... ) 
              (list 'name (list (list 'fields fields ) ...)))) (interaction-environment))
       (eval `(define
              ,(string->symbol (string-append (symbol->string 'name) "?"))
              (lambda (p) (and (pair? p) (equal? (car p) 'name)))) (interaction-environment))
       (eval `(define
              ,(string->symbol (string-append (symbol->string 'name) "-" (symbol->string 'fields)))
              (lambda (p) (cadr (assoc 'fields (cadr p))))) (interaction-environment)) ...
       (eval `(define
              ,(string->symbol (string-append "set-" (symbol->string 'name) "-" (symbol->string 'fields) "!"))
              (lambda (p num) 
              (set-car! (cdr (assoc 'fields (cadr p))) num))) (interaction-environment)) ...))))