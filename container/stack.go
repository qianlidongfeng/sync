package container

import (
	"sync"
)


type snode struct{
	data interface{}
	pre *snode
}

type SyncStack struct{
	top *snode
	lock *sync.Mutex
	cond *sync.Cond
	len int64
}

func NewStack() SyncStack{
	l:= SyncStack{lock:&sync.Mutex{}}
	l.cond=sync.NewCond(l.lock)
	return l
}

func (this *SyncStack) Push(data interface{}){
	n:=snode{data:data}
	this.lock.Lock()
	if this.top==nil{
		this.top=&n
		this.cond.Signal()
	}else{
		n.pre=this.top
		this.top=&n
	}
	this.len++
	this.lock.Unlock()
}


func (this *SyncStack) Pop() (elem interface{}){
	this.lock.Lock()
	for{
		if this.top == nil{
			this.cond.Wait()
		}
		if this.top == nil{
			continue
		}else{
			elem=this.top.data
			break
		}
	}
	this.top=this.top.pre
	this.len--
	this.lock.Unlock()
	return
}

func (this *SyncStack) TryPop() (elem interface{},ok bool){
	this.lock.Lock()
	if this.top == nil{
		ok=false
	}else{
		ok = true
		elem=this.top.data
		this.top=this.top.pre
	}
	this.lock.Unlock()
	return
}

func (this *SyncStack) Len() int64{
	return this.len
}


