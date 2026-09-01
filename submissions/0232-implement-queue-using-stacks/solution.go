type MyQueue struct {
    w Stack
    r Stack
}

type Stack struct {
    data []int
}

func NewStack() Stack {
    return Stack{data:[]int{}}
}

// push to top
func (s *Stack) Push(x int) {
    s.data = append(s.data, x)
    return
}

func (s *Stack) Peek() int {
    return s.data[len(s.data) - 1]
}

func (s *Stack) Pop() int {
    res := s.data[len(s.data) - 1]
    s.data = s.data[:len(s.data) - 1]
    return res
}

func (s *Stack) Size() int {
    return len(s.data)
}

func (s *Stack) Empty() bool {
    return len(s.data) == 0
}


func Constructor() MyQueue {
    return MyQueue{
        w: NewStack(),
        r: NewStack(),
    }
}


func (q *MyQueue) Push(x int) {
    q.w.Push(x)
    return
}


func (q *MyQueue) Pop() int {
    if q.r.Empty() {
        for !q.w.Empty() {
            q.r.Push(q.w.Pop())
        }
    }

    return q.r.Pop()
}


func (q *MyQueue) Peek() int {
    if q.r.Empty() {
        for !q.w.Empty() {
            q.r.Push(q.w.Pop())
        }
    }

    return q.r.Peek()
}


func (q *MyQueue) Empty() bool {
    return q.w.Empty() && q.r.Empty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
