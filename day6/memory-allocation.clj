(defn getMax
  [vect]

  (let [max (reduce #(Math/max %1 %2) 0 vect)
        index (first (keep-indexed #(when (= %2 max) %1) vect))]
    index
    ))

(defn updateVector
  [vect index]
  (let [val (get vect index)
        newVect (assoc vect index 0)]
    
    (loop [n val v newVect i (mod (inc index) (count vect))]
      (if (zero? n)
          v
          (recur
           (dec n)
           (assoc v i (inc (get v i)))
           (mod (inc i) (count vect)))))
    ))

(defn shift-vector-left
  [vect]

  (let [val       (first vect)
        lastIndex (dec (count vect))]
    (def ^:dynamic updatingVect vect)
    (doall
    (for [i (range lastIndex)]
      (let [foo (assoc updatingVect i (get updatingVect (inc i)))]
        (def updatingVect foo))))

    (assoc updatingVect lastIndex val)
    )
  )

(defn run-iteration
  [existingVectors iteration currentVector]

  (let [updatedVector    (updateVector currentVector (getMax currentVector))
        leftSortedVector (reduce (fn [vect index] shift-vector-left vect) updatedVector (range (getMax updatedVector)))]
    (if (contains? existingVectors leftSortedVector)
      (inc iteration)
      #(run-iteration
        (conj existingVectors leftSortedVector)
        (inc iteration)
        updatedVector))))

(trampoline run-iteration (hash-set) 0 [11 11 13 7 0 15 5 5 4 4 1 1 7 1 15 11])
