import "math/rand/v2"

type RandomizedSet struct {
    vals []int
    idxByVal map[int]int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        vals: make([]int, 0),
        idxByVal: make(map[int]int),
    }
}


func (s *RandomizedSet) Insert(val int) bool {
    if _, found := s.idxByVal[val]; found {
        return false
    }

    s.vals = append(s.vals, val)
    s.idxByVal[val] = len(s.vals) - 1

    return true
}


func (s *RandomizedSet) Remove(val int) bool {
    idx, found := s.idxByVal[val]
    if !found {
        return false
    }
    
    // replace value with last
    // cut array
    // update index for previus last value
    // delete val idx in map (important to be last in case of last element delete)

    lastVal := s.vals[len(s.vals) - 1]
    s.vals[idx] = lastVal
    s.vals = s.vals[:len(s.vals) - 1]
    s.idxByVal[lastVal] = idx
    delete(s.idxByVal, val)

    return true
}


func (s *RandomizedSet) GetRandom() int {
    return s.vals[rand.IntN(len(s.vals))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
