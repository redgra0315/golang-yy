package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// 定义通道与主题
type (
	subscriber chan interface{}
	topicFunc  func(v interface{}) bool
)

// Publisher 发布者对象
type Publisher struct {
	m           sync.RWMutex // 读写锁
	buffer      int          // 订阅队列的缓存大小
	timeout     time.Duration
	subscribers map[subscriber]topicFunc // 订阅者信息
}

// NewPublisher 构建一个发布者对象，可以设置发布超时时间和缓存队列的长度
func NewPublisher(publisherTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publisherTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// Subscribe 添加一个新的订阅者，订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// SubscribeTopic 添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	defer p.m.Unlock()
	return ch
}

// Evict 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

// Publish 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// Close 关闭发布者对象，同时关闭所有的订阅者通道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// sendTopic 发送主题，可以容热一定时间的超时
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func main() {
	p := NewPublisher(100*time.Microsecond, 10)
	defer p.Close()
	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	p.Publish("hello, world!")
	p.Publish("hello,golang")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}

	}()
	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()
	time.Sleep(3 * time.Second)

}
