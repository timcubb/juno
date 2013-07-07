package system

import(
    "sync"
    )

type Node struct {
    Host string
    Port int
    App string
    Status string
	//StartTime string
	//EndTime string
}

type System struct {
    lock *sync.RWMutex
    n map[string]Node
}

func NewSystem() *System {
    return &System{n: make(map[string]Node), lock: new(sync.RWMutex)}
}

func (s *System) AddNode(Name string, node Node) {
    s.lock.Lock()
    defer s.lock.Unlock()
    s.n[Name] = node
}

func (s *System) GetNode(Name string) Node {
    s.lock.Lock()
    defer s.lock.Unlock()
    r := s.n[Name]
    return r
}
