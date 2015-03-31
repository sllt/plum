(define version "1.0.0")

(defmacro unless (fn* [prd a b] `(if (not ~prd) ~a ~b)))

(define car first)
(define cdr rest)

(define max (fn* [l] (if (= (count l) 1) (first l) (if (> (first l) (max (rest l))) (car l) (max (cdr l))))))
(define min (fn* [l] (if (= (count l) 1) (first l) (if (< (first l) (min (rest l))) (car l) (min (cdr l))))))

(define inc (fn* [a] (+ a 1)))
(define dec (fn* [a] (- a 1)))

(define length count)

(define reduce
  (fn* (f init xs)
    (if (> (count xs) 0)
      (reduce f (f init (first xs)) (rest xs))
      init)))

(defmacro defn (fn* [name args body] `(define ~name (fn* ~args ~body))))