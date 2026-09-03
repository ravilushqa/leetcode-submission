func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    stack := NewStack()

    for i := range temperatures {
        for stack.Top() >= 0 && temperatures[stack.Top()] < temperatures[i] {
            unanswered := stack.Pop()
            res[unanswered] = i - unanswered
        }
        stack.Push(i)
    }

    return res
}

type Stack struct {
    data []int
}

func NewStack() Stack {
    return Stack{data: []int{}}
}

func (s *Stack) Push(x int) {
    s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
    if len(s.data) == 0 {
        return -1
    }
    res := s.data[len(s.data) - 1]
    s.data = s.data[:len(s.data) - 1]

    return res
}


func (s *Stack) Top() int {
    if len(s.data) == 0 {
        return -1
    }
    return s.data[len(s.data) - 1]
}
