Plum
=======================
Yet another lisp-like programming language.

Overview
=======================

```racket
;; functions
(+ 1 2 3 4 5 6 7 8 9 10)
(define inc (fn* [n] (+ n 1)))

;; apply
(apply + (list 1 2 3 4 5))

;; lambda
((fn* [a b] (+ a b)) 1 2)

;; macros
(define max 
        (fn* [l] 
            (if (= (count l) 1) 
                (first l) 
                (if (> (first l) (max (rest l)))
                    (car l)
                    (max (cdr l))))))

```