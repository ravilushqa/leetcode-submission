import "math/rand/v2"

type RandomizedSet struct {
    m map[int]int
    nums []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        m: map[int]int{},
        nums: []int{},
    }
}


func (s *RandomizedSet) Insert(val int) bool {
    if _, found := s.m[val]; found {
        return false
    }

    idx := len(s.nums)
    s.m[val] = idx
    s.nums = append(s.nums, val)

    return true
}


func (s *RandomizedSet) Remove(val int) bool {
    sliceIdx, found := s.m[val]
    if !found {
        return false
    }

    lastVal := s.nums[len(s.nums)-1]
    s.nums[sliceIdx] = lastVal
    s.nums = s.nums[:len(s.nums)-1]

    s.m[lastVal] = sliceIdx
    delete(s.m, val)

    return true
    // remove from slice - done
    // move last on the place of removed value - done
    // shrink slice - done
    // update index for last element
    // remove from map
    // corner case for last element (done by order of 2 previous step)
}


func (s *RandomizedSet) GetRandom() int {
    return s.nums[rand.IntN(len(s.nums))] 
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
