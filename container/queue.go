package container

import (
	"sync"
)

type node struct{
	data interface{}
	next *node
}



type SyncQueue struct{
	head *node
	tail *node
	lock *sync.Mutex
	cond *sync.Cond
}

func NewQueue() SyncQueue{
	q:= SyncQueue{lock:&sync.Mutex{}}
	q.cond=sync.NewCond(q.lock)
	return q
}

func (this *SyncQueue) Put(data interface{}){
	n:=node{data:data}
	this.lock.Lock()
	if this.head==nil{
		this.head=&n
	}
	if this.tail!=nil{
		this.tail.next=&n
	}
	this.tail=&n
	this.cond.Signal()
	this.lock.Unlock()
}


func (this *SyncQueue) Get() (elem interface{}){
	this.lock.Lock()
	for{
		if this.head == nil{
			this.cond.Wait()
		}
		if this.head == nil{
			continue
		}else{
			elem=this.head.data
			break
		}
	}
	this.head=this.head.next
	if this.head == nil{
		this.tail=nil
	}
	this.lock.Unlock()
	return
}

func (this *SyncQueue) TryGet() (elem interface{},ok bool){
	this.lock.Lock()
	if this.head == nil{
		ok=false
	}else{
		ok = true
		elem=this.head.data
		this.head=this.head.next
		if this.head == nil{
			this.tail=nil
		}
	}
	this.lock.Unlock()
	return
}


