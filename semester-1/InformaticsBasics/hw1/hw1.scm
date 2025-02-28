(define (day-of-week day month year)
  (define (cool-dev a b)
    (if (integer? (/ a b))
        (/ a b)
        (cool-dev (- a 1) b)))
  (if (> month 2)
      (remainder (- (+ day
                       (cool-dev ( - (* 13 (- month 2)) 1) 5)
                       year
                       (cool-dev year 4)
                       (cool-dev year 400))
                    (cool-dev year 100)) 7)
      (remainder (- (+ day
                       (cool-dev ( - (* 13 (+ month 10)) 1) 5)
                       (- year 1)
                       (cool-dev (- year 1) 4)
                       (cool-dev (- year 1) 400))
                    (cool-dev (- year 1) 100)) 7)))

(define (linear-equation a11 a12 a21 a22 b1 b2)
  (if (or (= (- (* a11 a22) (* a21 a12)) 0) (= (- (* a12 a21) (* a11 a22)) 0))
      '()
      (list (/ (- (* b1 a22) (* b2 a12)) (- (* a11 a22) (* a21 a12)))
            (/ (- (* b1 a21) (* b2 a11)) (- (* a12 a21) (* a11 a22))))))

(define (my-gcd a b)
  (define (f a b d)
    (if (and (= (remainder a d) 0) ( = (remainder b d) 0))
        d
        (f a b (- d 1))))
  (f a b (min (abs a) (abs b))))

(define (my-lcm a b)
  (define (f a b d)
    (if (and (= (remainder d a) 0) ( = (remainder d b) 0))
        d
        (f a b (+ d 1))))
  (f a b (max (abs a) (abs b))))

(define (prime? n)
  (define (f n d)
    (or (= d n) (and (not ( = (remainder n d) 0)) (f n (+ d 1)))))
  (f n 2))