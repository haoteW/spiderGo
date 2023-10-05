package schedule

import "good-spider/engine"

// SimpleScheduler 持有request channel 用于完成对request请求下发时的监听动作
// 但由于其单channel的限制 容易造成饥饿状态循环等待
// 所以在submit提交请求任务时启动协程完成 解决串行等待问题
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() { s.workerChan <- request }()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
