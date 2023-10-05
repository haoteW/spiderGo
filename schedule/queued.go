package schedule

import "good-spider/engine"

// QueuedScheduler 中有两个channel 一个用于请求传递 一个用于worker空闲状态标识
// 当worker空闲时，通过WorkerReady 将自身接收请求channel插入workerChan 标识自身空闲
// QueuedScheduler 运行时维护worker队列和request队列 用于存储请求和空闲worker标识
// 当两个队列同时存在元素时，则表示可唤醒worker 将请求放入worker中channel唤醒
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

/*
Run 完成

	1、channel 构建
	2、worker与request队列构建维护
	3、监听worker与request channel
	4、满足条件时唤醒worker
*/
func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
