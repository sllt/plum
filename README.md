Plum [![Build Status](https://travis-ci.org/sllt/plum.svg?branch=master)](https://travis-ci.org/sllt/plum)
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

(define max 
        (fn* [l] 
            (if (= (count l) 1) 
                (first l) 
                (if (> (first l) (max (rest l)))
                    (car l)
                    (max (cdr l))))))
;; macros

(defmacro unless (fn* [prd a b] `(if (not ~prd) ~a ~b)))

(defmacro defn (fn* [name args body] `(define ~name (fn* ~args ~body))))
`))

;; apply

(funcall + 1 2 3)
(apply + (list 1 2 3))

```

目标
==================================
* 支持宏
* 尾递归优化
* 并发
* 可扩展
