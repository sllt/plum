(define version "1.0.0")

(defmacro unless (fn* [prd a b] `(if (not ~prd) ~a ~b)))