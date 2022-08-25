;; Practice in sicp learning

;;The first program
486

(+ 137 349)

(+ 2.7 10)

(+ 21 35 12 7)

(+ (* 3 5) (- 10 6))

(+ (* 3
      (+ (* 2 4)
	 (+ 3 5)))
   (+ (- 10 7)
      6))

(+ (* 2 4)
   (- 10 5)
   8)

(define size 2)

size

(* 3 size)

(define pi 3.1415926)

(define radius 10)

(* pi
   (* radius radius))

(define cir (* 2 radius pi))

cir

(define (squ x) (* x x))

(squ 4)

(squ (+ 2 5))

(define (sum-of-square x y)
  (+ (squ x) (squ y)))

(sum-of-square 3 4)

; condition expression

(define (abs x)
  (cond ((> x 0) x)
	((= x 0) 0)
	((< x 0) (- x))))

(abs -3)

(abs 4)

; Unspecified return value
(cond ((> 3 6) 7))

;;else
(cond ((> 3 6) 3)
      (else 6))

;;if
(if (> 3 6)
    (+ 3 (* 1 2))
    (+ 1 (* 2 4)))
;; and
(and (> 2 1) #t #f)
(and (> 4 1) 78)

;; or
(or (> 2 1) (< 5 2))
(or #f #t #f)
(or (> 3 5) 57 #f)

;; not
(not #f)
(not (> 5 3))
(not 4)

;; define >=
(define (>= x y)
  (or (> x y)
      (= x y)))
(>= 6 4)

(define (<= x y)
  (not (> x y)))
(<= 9 6)

;; sum of square of large two of there number
;; solution's method
(define (bigger x y)
  (if (> x y) x y))
(define (smaller x y)
  (if (< x y) x y))
(define (another-bigger x y z)
  (bigger(smaller (x y) z)))
(define (two-large-sum x y z)
  (sum-of-square (bigger x y)
		 (bigger (smaller x y) z)))
;; my stupid method
(define (larger-sum x y z)
  (cond((and (< z y)(< z x))(sum-of-square x y))
       ((and (< y z)(< y x))(sum-of-square x z))
       ((and (< x z)(< x y))(sum-of-square y z))))
;;test
(larger-sum 2 1 3)

;; treat "+", "-" as function
;; treat function as value
(if (> 4 3) + -)
-
+

(define (a-plus-abs-b a b)
  ((if (> b 0) + -) a b))

(a-plus-abs-b 3 -4)
(a-plus-abs-b 3 6)

;; implement sqrt
(define (good-enough? guess x)
  (< (abs(- (square guess) x)) 0.001))
(define (average x y)
  (/ (+ x y) 2))
(define (improve guess x)
  (average guess (/ x guess)))
(define (sqrt-iter guess x)
  (if (good-enough? guess x)
      guess
      (sqrt-iter (improve guess x)
		 x)))
(define (sqrt x)
  (sqrt-iter 1.0 x))
;; test
(sqrt 9)
(sqrt (+ (sqrt 2) (sqrt 3)))

;; new-if something wrong
(define (new-if pre then-clause else-clause)
 (cond (pre then-clause)
	(else else-clause)))
(new-if (= 2 3) 0 5)
(new-if (= 1 1) 0 5)
;;maximum recursion depth exceeded
(define (sqrt-iter guess x)
 (new-if (good-enough? guess x)
     guess
     (sqrt-iter (improve guess x)
		x)))

(new-if (= 2 3) (display "good") (display "bad"))

;; small number
(sqrt 2)

;; big number; endless loop
(sqrt 20000000000000000000000000000000)

;;1.7  new good-enough
(define (good-enough? old-guess new-guess)
  (> 0.01
     (/ (abs (- new-guess old-guess))
	old-guess)))
(define (sqrt-iter guess x)
  (if (good-enough? guess (improve guess x))
      (improve guess x)
      (sqrt-iter (improve guess x)
		 x)))

;; cube-root 
(define (improve y x)
  (/ (+ (/ x (* y y))
	(* 2 y))
     3))
(define (cube-root x)
  (cube-root-iter 1.0 x))
(define (cube-root-iter guess x)
  (if (good-enough? guess x)
      guess
      (cube-root-iter (improve guess x)
		      x)))
(define (good-enough? guess x)
  (< (abs(- (*  guess guess guess) x)) 0.001))

(cube-root 27)
 
;;block structure
(define (sqrt x)
  (define (good-enough? guess x)
    (< (abs (- (square guess) x)) 0.001))
  (define (improve guess x)
    (average guess (/ x guess)))
  (define (sqrt-iter guess x)
    (if (good-enough? guess x)
	guess
	(sqrt-iter (improve guess x) x)))
  (sqrt-iter 1.0 x))

(sqrt 4)

;; lexical scoping.
;; allow x to be a free variable in the internal definitions

(define (sqrt x)
  (define (good-enough? guess)
    (< (abs(- (square guess) x)) 0.001))
  (define (improve guess)
    (average guess (/ x guess)))
  (define (sqrt-iter guess)
    (if (good-enough? guess)
	guess
	(sqrt-iter (improve guess))))
  (define (average x y)
    (/ (+ x y) 2))
  (sqrt-iter 1.0))

;; procedure definition
;; lambda way, just like to define a data: (define a (expression))
(define add (lambda (x y)(+ x y)))
;; normal way. for lisp it's the same thing,
;; this is a syntactic sugar of lambda way 
(define (add x y)(+ x y))

(add 3 5)

;; test override buildin procedure
(define + (lambda (x y) (- x y)))
(+ 8 3)

;; factorial method 1
(define (factorial n)
  (if (= n 1)
      n
      (* n (factorial (- n 1)))))

(factorial 6000)

;; factorial method 2
(define (factorial n)
  (define (fact-iter product count)
    (if (> count n)
	product
	(fact-iter (* count product)
		   (+ 1 count))))
  (fact-iter 1 1))

;; execrise 1.10
(define (ackermann x y)
  (cond ((= y 0) 0)
	((= x 0) (* 2 y))
	((= y 1) 2)
	(else (ackermann (- x 1)
			 (ackermann x (- y 1))))))

(ackermann 1 10)
(ackermann 2 4)
(+ (ackermann 3 3) 10000)
(ackermann 2 4)

;; fibonacci method 1 Tree Recursion
(define (fib n)
  (cond ((= n 0) 0)
	((= n 1) 1)
	(else (+ (fib (- n 1))
		 (fib (- n 2))))))
(fib 6)

;; fibonacci method 2 Linear Iteration, little difference with book
(define (fib n)
  (define (fib-iter a b count)
    (if (= count n)
	b
	(fib-iter (+ a b) a (+ count 1))))
  (fib-iter 1 0 0))

;; count change
;; leetcode: 518
(define (count-change amount)
  (cc amount 5))
(define (cc amount kinds-of-coins)
  (cond ((= amount 0) 1)
	((or (< amount 0) (= kinds-of-coins 0)) 0)
	(else (+ (cc amount (- kinds-of-coins 1))
		 (cc (- amount (first-denomination kinds-of-coins))
		     kinds-of-coins)))))
(define (first-denomination kinds-of-coins)
  (cond ((= kinds-of-coins 1) 1)
	((= kinds-of-coins 2) 5)
	((= kinds-of-coins 3) 10)
	((= kinds-of-coins 4) 25)
	((= kinds-of-coins 5) 50)))

(count-change 100)

;; exercise 1.11
;; recursion version
(define (f n)
  (if (< n 3)
      n
      (+ (f (- n 1))
	 (* 2 (f (- n 2)))
	 (* 3 (f (- n 3))))))
(f 5)
;; iteration version
(define (fi n)
  (define (fi-iter a b c count)
    (if (= count n)
	a
	(fi-iter b
		 c
		 (+ c
		    (* 2 b)
		    (* 3 a))
		 (+ count 1))))
  (fi-iter 0 1 2 0))
(fi 5)

;; exercise 1.12
;; make it 0-based
(define (pascal row col)
  (if (or (= col 0) (= col row))
      1
      (+ (pascal (- row 1) (- col 1))
	 (pascal (- row 1) col))))
(pascal 4 2)

;; exercise 1.15
(define (cube x) (* x x x))
(define (p x) (- (* 3 x) (* 4 (cube x))))
(define (sine angle) 
  (if (not (> (abs angle) 0.1))
      angle
      (p (sine (/ angle 3.0)))))

(sine 90)

;; Exponentiation
;; recursion
(define (expt b n)
  (if (= n 0)
      1
      (* b (expt b (- n 1)))))
(expt 3 3)

;; iteration
(define (expt b n)
  (define (expt-iter b count product)
    (if(= count 0)
       product
       (expt-iter b (- count 1) (* b product))))
  (expt-iter b n 1))

;; faster version
(define (fast-expt b n)
  (define (even? n)
    (= (remainder n 2) 0))
  (cond ((= n 0) 1)
	((even? n) (square (fast-expt b (/ n 2))))
	(else (* b (fast-expt b (- n 1))))))

(fast-expt 2 10)

;; exercise 1.16
;; iteration fast version
(define (fast-expt b n)
  (expt-iter b n 1))
;; b^n * a
;; at first a = 1
;; transfer data from b,n to a
;; at last, a = result
(define (expt-iter b n a)
  (cond ((= n 0) a)
	((even? n) (expt-iter (square b)
			      (/ n 2)
			      a))
	((odd? n) (expt-iter b
			     (- n 1)
			     (* b a)))))

;; exercise 1.17

(define (double n)
  (+ n n))
(define (halve n)
  (/ n 2))
(halve 3)
;; recursion version fast multi
(define (mul a b)
  (cond ((= b 0) 0)
	((even? b) (double (mul a (halve b))))
	(else (+ a (mul a (- b 1))))))

(mul 2 7)

;; exercise 1.18
;; iteration verison fast multi
(define (mul a b)
  (mul-iter a b 0))
;; a * b + t
;; at first, t = 0
;; step by step, transfer data from a,b to t
;; at the end, t = result
(define (mul-iter a b t)
  (cond ((= b 0) t)
	((even? b) (mul-iter (double a) (halve b) t))
	(else (mul-iter a (- b 1) (+ a t)))))

(mul 0 5)

;; Greatest Common Divisors
(define (gcd a b)
  (if (= b 0)
      a
      (gcd b (remainder a b))))
(gcd 16 28)

;; smallest divisor
(define (smallest-divisor n)
  (find-divisor n 2))
(define (find-divisor n test-divisor)
  (cond ((> (square test-divisor) n) n)
	((divides? test-divisor n) test-divisor)
	(else (find-divisor n (+ test-divisor 1)))))
(define (divides? a b)
  (= (remainder b a) 0))
(define (prime? n)
  (= n (smallest-divisor n)))

(prime? 11)

;; fermat test
(define (expmod base exp m)
  (cond ((= exp 0) 1)
	((even? exp)
	 (remainder (square (expmod base (/ exp 2) m))
		    m))
	(else
	 (remainder (* base (expmod base (- exp 1) m))
		    m))))

(define (fermat-test n)
  (define (try-it a)
    (= (expmod a n n) a))
  (try-it (+ 1 (random (- n 1)))))
(define (fast-prime? n times)
  (cond ((= times 0) true)
	((fermat-test n) (fast-prime? n (- times 1)))
	(else false)))
(fast-prime? 1105 30)

;; test multi procedure
(define (test-multi-step)
  (display "hello")
  (newline) ;; output new line
  (display "world")
  (+ 3 5)
  77
  (+ 6 8));; the value of last procedure will be value of this whole procedure

(test-multi-step)

;; test runtime fucntion
;; (runtime) stands for time cost on computing sine last start
(runtime)
(define (test-runtime)
  (display (runtime))
  (newline)
  (fib 30)
  (display (runtime)))
(test-runtime)

;; exercise 1.22
;; in new version of mit-scheme, runtime count by second. so we use real-time-clock instead

;; (define (timed-prime-test n)
;;   (newline)
;;   (display n)
;;   (start-prime-test n (runtime)))

;; (define (start-prime-test n start-time)
;;   (if (prime? n)
;;       (report-prime (- (runtime) start-time))))
(define (timed-prime-test n)
  (start-prime-test n (real-time-clock)))

(define (start-prime-test n start-time)
  (if (prime? n)
      (report-prime n  (- (real-time-clock) start-time))))


(define (report-prime num  elapsed-time)
  (newline)
  (display num)
  (display " *** ")
  (display elapsed-time))

(timed-prime-test 104395301)

(define (search-for-primes a b)
  (cond ((> a b) 0)
	((even? a)(search-for-primes (+ a 1) b))
	(else (timed-prime-test a)(search-for-primes (+ a 2) b))))
(search-for-primes 10000000 10000020)

;; 1.3

(define (sum term a next b)
  (if (> a b)
      0
      (+ (term a)
	 (sum term (next a) next b))))
(define (inc n) (+ n 1))
(define (cube a)
  (* a a a))
(define (sum-cubes a b)
  (sum cube a inc b))
(sum-cubes 1 10)

(define (pi-sum a b)
  (define (pi-term x)
    (/ 1.0 (* x (+ x 2))))
  (define (pi-next x)
    (+ x 4))
  (sum pi-term a pi-next b))

(* 8 (pi-sum 1 100000))

(define (integral f a b dx)
  (define (add-dx x) (+ x dx))
  (* (sum f (+ a (/ dx 2.0)) add-dx b)
     dx))

(integral cube 0 1 0.001)

;; 1.29
(define (simpson f a b n)
  (define h (/ (- b a) n))
  (define (y x) (f (+ a (* x h))))
  (define (term-simpson k)
    (cond((or (= k 0) (= k n)) (y k))
	 ((= 0 (remainder k 2)) (* 2 (y k)))
	 (else (* 4 (y k)))))
  (* (/ h 3)
     (sum term-simpson 0 inc n)))
(simpson cube 0 1 10.0)

;; 1.30
(define (sum term a next b)
  (define (iter a res)
    (if (> a b)
	res
	(iter (next a)
	      (+ res
		 (term a)))))
  (iter a 0))

;; 1.31
;; a)
(define (product a b term next)
  (if (> a b)
      1
      (* (term a)
	 (product (next a) b term next))))

(define (factorial b n)
  (define (term-b x)
    b)
  (define (inc x)
    (+ x 1))
  (product 1 n term-b inc))

(factorial 4 2)

(define (cal-pi a b)
  (define (term x)
    (/ (* (- x 1) (+ x 1))
       (* x x)))
  (define (next x)
    (+ x 2))
  (* 4
     (product a b term next)))

(cal-pi 3.0 11111)

;; b)
(define (product a b term next)
  (define (iter c res)
    (if (> c b)
	res
	(iter (next c) (* res (term c)))))
  (iter a 1))

;; 1.32
;; a)
(define (accumulate combiner null-value term a next b)
  (if (> a b)
      null-value
      (combiner (term a)
		(accumulate combiner null-value term (next a) next b))))
;; test
(define (term x)
  x)
(define (next x)
  (+ x 1))
(accumulate * 1 term 2 next 4)
;; test done

;; new version of sum
(define (sum term a next b)
  (accumulate + 0 term a next b))
;; b)
(define (accumulate combiner null-value term a next b)
  (define (iter t res)
    (if (> t b)
	res
	(iter (next t)
	      (combiner (term t)
			res))))
  (iter a null-value))

;; 1.33
(define (filtered-accumulate combiner null-value term a next b filter)
  (if (> a b)
      null-value
      (if (filter (term a))
	  (combiner (term a)
		    (filtered-accumulate combiner null-value term (next a) next b filter))
	  (filtered-accumulate combiner null-value term (next a) next b filter))))
;; a)
(filtered-accumulate + 0 term 1 next 10 prime?)
;; b)
(define (test n)
  (define (filter x)
    (= (GCD x n) 1))
  (filtered-accumulate * 1 term 1 next (- n 1) filter))

(test 10)

;; 1.3.2
(define (pi-sum a b)
  (sum (lambda (x) (/ 1.0 (* x (+ x 2))))
       a
       (lambda (x) (+ x 4))
       b))

(* 8 (pi-sum 1 4500))

;; define procedure in lambda way
(define plus4 (lambda (x) (+ x 4)))
(plus4 6)
;; define a procedure and call it
((lambda (x y) (+ x y)) 3 4)

;; local variable
;; method 1: inner procedure
(define (func x y)
  (define (func-helper a b)
    (+ (* x (square a))
       (* y b)
       (* a b)))
  (func-helper (+ 1 (* x y))
	       (- 1 y)))

;; test value: 4
(func 1 2)

;; method 2: inner procedure in lambda way
(define (func x y)
  ((lambda (a b)
     (+ (* x (square a))
	(* y b)
	(* a b)))
   (+ 1 (* x y))
   (- 1 y)))

;; method 3: let
(define (func x y)
  (let ((a (+ 1 (* x y)))
	(b (- 1 y)))
    (+ (* x (square a))
       (* y b)
       (* a b))))

;;---

((lambda (x)
  (+ (let ((x 3))
     (+ x (* x 10)))
     x)) 5)

((lambda (x)
   (let ((x 3)
	 (y (+ x 2)))
     (* x y))) 2)

;; 1.34
(define (f g)
  (g 2))

(f square)
(f (lambda (z) (* z (+ z 1))))

(f f)

;; 1.3.3
;; half interval method
(define (search f neg pos)
  (let ((mid (average neg pos)))
    (if (close-enough? neg pos)
	mid
	(let ((test (f mid)))
	  (cond ((positive? test)
		 (search f neg mid))
		((negative? test)
		 (search f mid pos))
		(else mid))))))

(define (close-enough? x y)
  (< (abs (- y x)) 0.001))

(define (average a b)
  (/ (+ a b)
     2))

(define (half-interval-method f a b)
  (let ((a-val (f a))
	(b-val (f b)))
    (cond ((and (negative? a-val) (positive? b-val))
	   (search f a b))
	  ((and (positive? a-val) (negative? b-val))
	   (search f b a))
	  (else (error "wrong picked value" a b)))))

(half-interval-method sin 2.0 4.0)
(half-interval-method (lambda (x) (- (* x x x) (* 2 x) 3))
		      1.0
		      2.0)

;; fixed point
(define tolerance 0.000001)

(define (fixed-point f first-guess)
  (define (close-enough? v1 v2)
    (< (abs (- v1 v2)) tolerance))
  (define (try guess)
    (let ((next (f guess)))
      (if (close-enough? guess next)
	  next
	  (try next))))
  (try first-guess))

(fixed-point cos 1.0)
(fixed-point (lambda (y) (+ (sin y) (cos y)))
	     1.0)
;; loop forever
(define (sqrt x)
  (fixed-point (lambda (y) (/ x y))
	       1.0))
;; improved version
(define (sqrt x)
  (fixed-point (lambda (y) (average y (/ x y)))
	       1.0))
(sqrt 4)

;; 1.35
(fixed-point (lambda (x) (+ 1 (/ 1 x)))
	     1.0)

;; exercise 1.36
;; without average damp
(define (fixed-point f first-guess)
  (define (close-enough? v1 v2)
    (< (abs (- v1 v2)) tolerance))
  (define (try guess)
    (newline)
    (display guess)
    (let ((next (f guess)))
      (if (close-enough? guess next)
	  next
	  (try next))))
  (try first-guess))

(fixed-point (lambda (x) (/ (log 1000)
			    (log x)))
	     2.0)

;; with average dump

;; (define (fixed-point f first-guess)
;;   (define (close-enough? v1 v2)
;;     (< (abs (- v1 v2)) tolerance))
;;   (define (average x y)
;;     (/ (+ x y)
;;        2))
;;   (define (try guess)
;;     (newline)
;;     (display guess)
;;     (let ((next (f guess)))
;;       (if (close-enough? guess next)
;; 	  next
;; 	  (try (average next guess)))))
;;   (try first-guess))

(fixed-point (lambda (x) (average x (/ (log 1000) (log x)))) 2.0)

;; 1.37
;; recursion version
(define (cont-frac n d k)
  (define (do-it i)
    (if (= i k)
	(/ (n i)
	   (d i))
	(/ (n i)
	   (+ (d i)
	      (do-it (+ i 1))))))
  (do-it 1))

;; test
(/ 1
   (cont-frac (lambda (i) 1.0)
	   (lambda (i) 1.0)
	   1000))

;; iteration version
;; TODO 

;; 1.38
(+ 2
   (cont-frac (lambda (i) 1)
	      (lambda (i) (cond ((= 0 (remainder (+ i 1) 3))
				 (* 2.0
				    (/ (+ i 1)
				       3.0)))
				(else 1)))
	      50.0))

;; 1.39
(define (tan-cf x k)
  (cont-frac (lambda (i)
	       (cond ((= i 1) x)
		     (else (- (square x)))))
	     (lambda (i)
	       (- (* i 2.0) 1))
	     k))

(tan-cf 10 100.0)

;; 1.3.4
(define (average-damp f)
  (lambda (x) (average (f x) x)))

(define (average x y)
  (/ (+ x y)
     2))
((average-damp square) 10)

(define (sqrt x)
  (fixed-point (average-damp (lambda (y) (/ x y)))  1.0))
(sqrt 4)

(define (cube-root x)
  (fixed-point (average-damp (lambda (y) (/ x (square y)))) 1.0))
(cube-root 27)

;; newton's method
;; derivative
(define (deriv g)
  (lambda (x)
    (/ (- (g (+ x dx)) (g x))
       dx)))

(define dx 0.0001)
(define (cube x)
  (* x x x))
((deriv cube) 5)

;;newton-transform
(define (newton-transform g)
  (lambda (x)
    (- x (/ (g x) ((deriv g) x)))))

(define (newtons-method g guess)
  (fixed-point (newton-transform g) guess))

(define (sqrt x)
  (newtons-method (lambda (y) (- (square y) x)) 1.0))
(sqrt 9)

;; fixed-point-of-transform
(define (fixed-point-of-transform g transform guess)
  (fixed-point (transform g) guess))

(define (sqrt x)
  (fixed-point-of-transform (lambda (y) (/ x y))
			    average-damp
			    1.0))
(define (sqrt x)
  (fixed-point-of-transform (lambda (y) (- (square y) x))
			    newton-transform
			    1.0))

;; 1.40
(define (cubic a b c)
  (lambda (x)
    (+ (* x x x)
       (* a (square x))
       (* b x)
       c)))
(newtons-method (cubic 2 5 5) 1.0)

;; 1.41
(define (double f)
  (lambda (x)
    (let ((a (f x)))
      (f a))))

;; do not need 'let' expression, just use (f (f x)) 

(((double (double double)) (lambda (x)(+ x 1))) 5 )

;; 1.42
(define (compose f g)
  (lambda (x)
    (f (g x))))

((compose square (lambda (x) (+ x 1))) 6)

;; 1.43
;; wrong version 
;; (define (repeated f n)
;;   (if (= n 1)
;;       (lambda (x) (f x))
;;       (repeated (compose f f) (- n 1))))

(define (repeated f n)
  (if (= n 1)
      f
      (compose f (repeated f (- n 1)))))

((repeated (lambda (x) (+ x 1)) 2) 5)

;; 1.44
(define (smooth f)
  (lambda (x)
    (/ (+ (f (- x dx))
	  (f x)
	  (f (+ x dx)))
       3)))

(define (smooth-n-times f n)
  ((repeated smooth n) f))

(square 5.0)
((smooth square) 5 )
((smooth-n-times square 10) 5)

;; ch2
;;2.1.1
;; rational number operations
(define (add-rat x y)
  (make-rat (+ (* (numer x) (denom y))
	       (* (numer y) (denom x)))
	    (* (denom x) (denom y))))
(define (div x y)
  (make-rat (* (numer x) (denom y))
	    (* (denom x) (numer y))))
(define (equal-rat? x y)
  (= (* (numer x) (denom y))
     (* (numer y) (denom x))))

;; pairs
(define x (cons 1 4))
(car x)
(cdr x)
(define y (cons 2 5))
(define z (cons x y))
(car (cdr z))

;; Representing rational numbers
(define (make-rat n d) (cons n d))
(define (numer x) (car x))
(define (denom x) (cdr x))

;; print rat
(define (print-rat x)
  (newline)
  (display (numer x))
  (display "/")
  (display (denom x)))

;; test
(define one-half (make-rat 1 2))
(print-rat one-half)
(define one-third (make-rat 1 3))
(print-rat (add-rat one-third one-third))

;; to reduce rational numbers to lowest terms
(define (make-rat n d)
  (let ((g (gcd n d)))
    (cons (/ n g) (/ d g))))

;; exercise 2.1
(define (make-rat n d)
  (if (or (and (< n 0) (< d 0))
	  (and (> n 0) (> d 0)))
      (cons (abs n) (abs d))
      (cons (- (abs n)) (abs d))))
(print-rat (make-rat 2 (- 3)))

;; exercise 2.2
(define (make-segment start-point end-point)
  (cons start-point end-point))
(define (start-segment segment)
  (car segment))
(define (end-segment segment)
  (cdr segment))

(define (make-point x y)
  (cons x y))
(define (x-point point)
  (car point))
(define (y-point point)
  (cdr point))

(define (midpoint-segment segment)
  (make-point (average (x-point (start-segment segment))
		       (x-point (end-segment segment)))
	      (average (y-point (start-segment segment))
		       (y-point (end-segment segment)))))

(define (print-point p)
  (newline)
  (display "(")
  (display (x-point p))
  (display ",")
  (display (y-point p))
  (display ")"))

;; test
(define p1 (make-point 1 3))
(define p2 (make-point 4 3))
(define seg (make-segment p1 p2))
(print-point (midpoint-segment seg))

;; exercise 2.3
(define (perimeter rect)
  (* (+ (height rect) (width rect))
     2))
(define (area rect)
  (* (height rect)
     (width rect)))

;; constructor and selector 1
(define (make-rect point-left-top point-right-bottom)
  (cons point-left-top point-right-bottom))
(define (left-top-point rect) (car rect))
(define (right-bottom-point rect) (cdr rect))
(define (height rect)
  (- (y-point (left-top-point rect))
     (y-point (right-bottom-point rect))))
(define (width rect)
  (- (x-point (right-bottom-point rect))
     (x-point (left-top-point rect))))

;; test 1
(define rect0 (make-rect (make-point 0 2) (make-point 3 0)))
(perimeter rect0)
(area rect0)

;; constructor and selector 2
(define (make-rect left-segment top-segment)
  (cons left-segment top-segment))
(define (left-segment rect) (car rect))
(define (top-segment rect) (cdr rect))
(define (height rect)
  (abs (- (y-point (start-segment (left-segment rect)))
	  (y-point (end-segment (left-segment rect))))))
(define (width rect)
  (abs (- (x-point (start-segment (top-segment rect)))
	  (x-point (end-segment (top-segment rect))))))

;; test 2
(define rect1 (make-rect (make-segment (make-point 0 0)
				       (make-point 0 3))
			 (make-segment (make-point 0 0)
				       (make-point 4 0))))
(perimeter rect1)
(area rect1)

;; 2.1.3
;; implement pairs in procedure way
(define (cons x y)
  (define (dispatch m)
    (cond ((= m 0) x)
	  ((= m 1) y)
	  (else (error "arg error"))))
  dispatch)
(define (car z) (z 0))
(define (cdr z) (z 1))

;; test
(define pair (cons 3 5))
(car pair)

;; exercise 2.4
(define (cons x y)
  (lambda (m) (m x y)))
(define (car z)
  (z (lambda (p q) p)))
(define (cdr z)
  (z (lambda (p q) q)))

;; exercise 2.5
(define (cons x y)
  (* (expt 2 x)
     (expt 3 y)))
(define (car z)
  (define (count num cnt)
    (newline)
    (display num)
    (if (= (remainder num 2)
	   0)
	(count (/ num 2) (+ cnt 1))
	cnt))
  (count z 0))

(define (cdr z)
  (define (count num cnt)
    (if (= (remainder num 3)
	   0)
	(count (/ num 3) (+ cnt 1))
	cnt))
  (count z 0))

;; 2.1.4
(define (add-interval x y)
  (make-interval (+ (lower-bound x) (lower-bound y))
		 (+ (upper-bound x) (upper-bound y))))
(define (mul-interval x y)
  (let ((p1 (* (lower-bound x) (lower-bound y)))
	(p2 (* (lower-bound x) (upper-bound y)))
	(p3 (* (upper-bound x) (lower-bound y)))
	(p4 (* (upper-bound x) (upper-bound y))))
    (make-interval (min p1 p2 p3 p4)
		   (max p1 p2 p3 p4))))
(define (div-interval x y)
  (mul-interval x
		(make-interval (/ 1.0 (upper-bound y))
			       (/ 1.0 (lower-bound y)))))

;; exercise 2.7
(define (make-interval a b) (cons a b))
(define (lower-bound c) (car c))
(define (upper-bound c) (cdr c))

;; test
(define test-interval (make-interval 2 3))
(add-interval test-interval test-interval)
(mul-interval test-interval test-interval)
(div-interval test-interval test-interval)

;;exercise 2.8
(define (sub-interval a b)
  (make-interval (- (lower-bound a) (upper-bound b))
		 (- (upper-bound a) (lower-bound b))))

;; test
(define test2 (make-interval -1 5))
(sub-interval test2 test-interval)

;; exercise 2.10
(define (div-interval x y)
  (if (<= (* (lower-bound y) (upper-bound y)) 0)
      (error "division error, interval spans 0 " y)
      (mul-interval x
		(make-interval (/ 1.0 (upper-bound y))
			       (/ 1.0 (lower-bound y))))))

(div-interval test-interval test2)

;; 2.11
;; ben's advice is awful

;; exercise 2.12
(define (make-center-percent cen per)
  (let ((w (* cen per)))
    (make-interval (- cen w) (+ cen w))))
(define (center i)
  (/ (+ (lower-bound i) (upper-bound i)) 2))
(define (percent i)
  (let ((w (/ (- (upper-bound i) (lower-bound i)) 2)))
    (/ w (center i))))

;; test
(define test (make-center-percent 5.4 0.11))
(center test)
(percent test)

;; ch 2.2
;; 2.2.1
;; list
(define one-four (list  1 2 3 4))
(car (cdr one-four))
(cons 8 one-four)
(cons one-four 9)

;; list ref
(define (list-ref items n)
  (if (= n 0)
      (car items)
      (list-ref (cdr items) (- n 1))))

(list-ref one-four 2)

;; length recursion
(define (length items)
  (if (null? items)
      0
      (+ 1 (length (cdr items)))))

(length one-four)

;; length iteration
(define (length items)
  (define (len-iter a count)
    (if (null? a)
    count
    (len-iter (cdr a) (+ count 1))))
  (len-iter items 0))

;; append
(define (append list1 list2)
  (if (null? list1)
      list2
      (cons (car list1) (append (cdr list1) list2))))
;; test
(define list2 (list 9 8 7))
(append list2 one-four)

;; ex 2.17
(define (last-pair items)
  (if (null? (cdr items))
      items
      (last-pair (cdr items))))
;; test
(last-pair (list 23 45 53 22))

;; 2.18
(define (reverse items)
  (define (reverse-iter remained res)
    (if (null? remained)
	res
	(reverse-iter (cdr remained) (cons (car remained) res))))
  (reverse-iter items '()))
(reverse (list 3 5 7))

;; 2.20
(define (same-parity x . nums)
  (define (is-same-parity x y)
    (= (remainder x 2)
       (remainder y 2)))
  (define (iter ret nums)
    (if (null? nums)
	(reverse ret)
	(if (is-same-parity x (car nums))
	    (iter (cons (car nums) ret) (cdr nums))
	    (iter ret (cdr nums)))))
  (newline)
  (display x)
  (display " ")
  (display nums)
  (iter '() (cons x nums)))

(same-parity  2 3 4 3 6 8 5 10)

;; scale list
(define (scale-list items factor)
  (if (null? items)
      '()
      (cons (* (car items) factor)
	    (scale-list (cdr items) factor))))
;; test
(scale-list (list 2 3 4 5) 10)

;; scheme internal map experiment
(map (lambda (x y) (+ x (* 2 y)))
     (list 1 2 3)
     (list 4 5 6))

;; my-map
(define (my-map proc items)
  (if (null? items)
      '()
      (cons (proc (car items))
	    (my-map proc (cdr items)))))
;; test
(my-map (lambda (x) (+ x 3))
	(list 4 5 6))
;; re-define scale-list by my-map
(define (scale-list items factor)
  (my-map (lambda (x) (* x factor))
	  items))
;; test
(scale-list (list 1 2 3) 5)

;; exercise 2.21
;; method 1
(define (square-list items)
  (if (null? items)
      '()
      (cons (square (car items))
	    (square-list (cdr items)))))
;; test
(square-list (list 1 2 3 4))

;; method 2
(define (square-list items)
  (my-map square items))

;; exercise 2.22
(define (square-list items)
  (define (iter things answer)
    (if (null? things)
	answer
	(iter (cdr things)
	      (cons answer
		    (square (car things))))))
  (iter items '()))

;; 2.23
(define (for-each proc items)
  (define (iter proc items-left current)
    (if (null? items-left)
	#t
	(iter proc (cdr items-left) (proc (car items-left)))))
  (iter proc items (car items)))

(for-each (lambda (x) (newline) (display x))
	  (list 34 52 45))

;; ch 2.2.2
(define test_list (cons (list 1 2) (list 3 4)))

(define (count-leaves x)
  (cond ((null? x) 0)
	((not (pair? x)) 1)
	(else (+ (count-leaves (car x))
		 (count-leaves (cdr x))))))

(length test_list)
(count-leaves test_list)
(count-leaves (list test_list test_list))

;; exercise 2.25
(define ex25_1 (list 1 3 (list 5 7) 9))
ex25_1
(car (cdr (car (cdr (cdr ex25_1)))))

(define ex25_2 (list (list 7)))
(car (car ex25_2))

(define ex25_3 (list 1 (list 2 (list 3 (list 4 (list 5 (list 6 7)))))))
(car (cdr (car (cdr (car (cdr (car (cdr (car (cdr (car (cdr ex25_3))))))))))))

;; exercise 2.26
(define x (list 1 2 3))
(define y (list 4 5 6))
(append x y)
(cons x y)
(list x y)

;; exercise 2.27
(define (deep-reverse items)
  (define (iter x ret)
    (if (null? x)
	ret
	(if (pair? (car x))
	    (iter (cdr x) (cons (deep-reverse (car x)) ret))
	    (iter (cdr x) (cons (car x) ret)))))
  (iter items '()))

;; test
(deep-reverse (list x y))

;; 2.28 
(define (fringe items)
  (cond ((null? items) '())
	((pair? (car items))
	 (append (fringe (car items)) (fringe (cdr items))))
	(else (cons (car items) (fringe (cdr items))))))

;; test
(define x (list (list 1 2) 5 6 (list 3 4)))
(fringe x)
(fringe (list x x))

;; scale tree
(define (scale-tree tree factor)
  (cond ((null? tree) '())
	((not (pair? tree)) (* tree factor))
	(else (cons (scale-tree (car tree) factor)
		    (scale-tree (cdr tree) factor)))))
;; test
(scale-tree (list 1 (list 2 (list 3 4) 5) (list 6 7))
	    10)

;; scale tree in another way

(define (scale-tree tree factor)
  (map (lambda (sub-tree)
	 (if (pair? sub-tree)
	     (scale-tree sub-tree factor)
	     (* sub-tree factor)))
       tree))

;; exercise 2.21
(define (square-tree tree)
  (cond ((null? tree) '())
	((not (pair? tree)) (square tree))
	(else (cons (square-tree (car tree))
		    (square-tree (cdr tree))))))
;; test
(square-tree (list 1 (list 2 (list 3 4) 5) (list 6 7)))

;; map style
(define (square-tree tree)
  (map (lambda (sub-tree)
	 (if (pair? sub-tree)
	     (square-tree sub-tree)
	     (square sub-tree)))
       tree))

;; 2.31
;; there are 2 ways in 2.30, I choose the latter one
(define (tree-map proc tree)
  (map (lambda (sub-tree)
	 (if (pair? sub-tree)
	     (tree-map proc sub-tree)
	     (proc sub-tree)))
       tree))
(define (square-tree tree) (tree-map square tree))

;; 2.32

(define (subsets s)
  (if (null? s)
      (list '()) ;; => (()) -> a empty list has one subset not zero, that is itself
      (let ((rest (subsets (cdr s))))
	(append rest (map (lambda (x)
			    (cons (car s) x))
			  rest)))))

(subsets (list 1 2 3))

;; ch 2.2.3

;; sum-odd-square
(define (sum-odd-squares tree)
  (cond ((null? tree) 0)
	((not (pair? tree))
	 (if (odd? tree) (square tree) 0))
	(else (+ (sum-odd-squares (car tree))
		 (sum-odd-squares (cdr tree))))))
;; test
(sum-odd-squares (list 1 (list 2 (list 3 4) 5) (list 6 7)))

;; even fibs
(define (even-fibs n)
  (define (next k)
    (if (> k n)
	'()
	(let ((f (fib k)))
	  (if (even? f)
	      (cons f (next (+ k 1)))
	      (next (+ k 1))))))
  (next 0))

(even-fibs 3)

;; use list to implement stream
;; map
(map square (list 1 2 3))

;; filter
(define (filter predicate sequence)
  (cond ((null? sequence) '())
	((predicate (car sequence))
	 (cons (car sequence)
	       (filter predicate (cdr sequence))))
	(else (filter predicate (cdr sequence)))))

(filter odd? (list 1 2 3 4))

;; accumulate
(define (accumulate op initial sequence)
  (if (null? sequence)
      initial
      (op (car sequence)
	  (accumulate op initial (cdr sequence)))))

(accumulate + 0 (list 1 2 3))

;; enumerate

(define (enumerate-interval low high)
  (if (> low high)
      '()
      (cons low (enumerate-interval (+ low 1) high))))

(enumerate-interval 2 5)

(define (enumerate-tree tree)
  (cond ((null? tree) '())
	((not (pair? tree)) (list tree))
	(else (append (enumerate-tree (car tree))
		      (enumerate-tree (cdr tree))))))

(enumerate-tree (list 1 (list 2 (list 3 4)) 5 ))

;; re-write sum-odd-square and even-fibs in signal-flow way
(define (sum-odd-square tree)
  (accumulate +
	      0
	      (map square
		   (filter odd?
			   (enumerate-tree tree)))))
;; test 
(sum-odd-square (list 1 (list 2 (list 3 4) 5) (list 6 7)))

(define (even-fibs n)
  (accumulate cons
	      '()
	      (filter even?
		      (map fib
			   (enumerate-interval 0 n)))))

(even-fibs 7)

;; other programs using sequence as signal flow
(define (list-fib-square n)
  (accumulate cons
	      '()
	      (map square
		   (map fib
			(enumerate-interval 0 n)))))
(list-fib-square 10)

(define (product-odd-square seq)
  (accumulate *
	      1
	      (map square
		   (filter odd? seq))))
(product-odd-square (list 1 2 3 4 5))

;; ex 2.33
(define (my-map p seq)
  (accumulate (lambda (x y) (cons (p x) y)) '() seq))
;; test
(my-map square (list 1 3 4))

(define (my-append seq1 seq2)
  (accumulate cons seq2 seq1))
;; test
(my-append (list 2 4) (list 1 5))

(define (my-length seq)
  (accumulate (lambda (x y) (+ y 1)) 0 seq))
;; test
(my-length (list 1 2 3 4 5))

;; ex 2.34

(define (hornor-eval x coefficient-seq)
  (accumulate (lambda (this-coeff higher-terms)
		(+ this-coeff (* x higher-terms)))
	      0
	      coefficient-seq))

(hornor-eval 2 (list 1 3 0 5 0 1))

;; ex 3.35
(define (count-leaves tree)
  (accumulate +
	      0
	      (map (lambda (x)
		     (cond ((null? x) 0)
			   ((pair? x) (count-leaves x))
			   (else 1)))
		   tree)))

(count-leaves (list 1 (list 2 (list 3 4))))

;; ex 2.36
(define (accumulate-n op init seqs)
  (if (null? (car seqs))
      '()
      (cons (accumulate op init (map (lambda (seq) (car seq))
				     seqs))
	    (accumulate-n op init (map (lambda (seq) (cdr seq))
				       seqs)))))
(accumulate-n + 0 (list (list 1 2 3) (list 4 5 6) (list 7 8 9) (list 10 11 12)))

;; ex 2.37
;; test map multi list arguments
(map * (list 1 3) (list 4 5) (list 2 2))

(define (dot-product v w)
  (accumulate + 0 (map * v w)))
(dot-product (list 1 3) (list 2 4))

;; ex 2.38
(define fold-right accumulate)

(define (fold-left op init seq)
  (define (iter result rest)
    (if (null? rest)
	result
	(iter (op result (car rest))
	      (cdr rest))))
  (iter init seq))

(fold-right / 1 (list 1 2 3))
(fold-left / 1 (list 1 2 3))
(fold-right list '() (list 1 2 3))
(fold-left list '() (list 1 3 4))

;; ex 2.39
(define (reverse seq)
  (fold-right (lambda (x y) (append y (list x)))
	      '()
	      seq))

(define (reverse seq)
  (fold-left (lambda (x y) (append (list y) x))
	     '()
	     seq))

(reverse (list 1 2 3))
