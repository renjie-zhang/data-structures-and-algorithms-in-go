/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package lru

type Element struct {
	prev,next *Element
	Key interface{}
	Value interface{}
}

func (e *Element) Next() *Element{
	return e.next
}

func (e *Element) Prev() *Element{
	return e.prev
}


type LRUCache struct {
	cache map[interface{}]*Element
	head *Element
	tail *Element
	capacity int
}

func New(capacity int) *LRUCache {
	return &LRUCache{
		cache:    make(map[interface{}]*Element),
		head:     nil,
		tail:     nil,
		capacity: capacity,
	}
}

func(l *LRUCache) Put(key interface{},value interface{}){
	// 如果在cache中存在，那么将此value放到头部
	if e,ok := l.cache[key];ok{
		e.Value = value
		l.refresh(e)
		return
	}
	// 如果容量为0，那么直接返回，如果容量已经满了，此时将删除最后一个节点，然后将value添加到头部
	if l.capacity == 0{
		return
	}else if len(l.cache) >= l.capacity {
		delete(l.cache,l.tail.Key)
		l.remove(l.tail)
	}
	e := &Element{nil,l.head,key,value}
	l.cache[key] = e
	if len(l.cache) != 1{
		l.head.prev = e
	}else {
		l.tail = e
	}
	l.head = e
	l.capacity++
}

func (l *LRUCache) Get(key interface{}) (interface{},bool){
	if e,ok := l.cache[key];ok{
		l.refresh(e)
		return e.Value,true
	}
	return nil,false
}

func (l *LRUCache) Delete(key interface{}){
	if e,ok := l.cache[key];ok{
		delete(l.cache,key)
		l.remove(e)
	}
}

func (l *LRUCache) Range(f func(key,value interface{}) bool)  {
	for i:= l.head;i != nil; i= i.Next(){
		if !f(i.Key,i.Value){
			break
		}
	}
}

func (l *LRUCache) Front() *Element{
	return l.head
}

func(l *LRUCache) Back() *Element{
	return l.tail
}

func(l *LRUCache) Len() int{
	return len(l.cache)
}

func (l *LRUCache) Capacity() int{
	return l.capacity
}

func (l *LRUCache) refresh(e *Element){
	if e.prev != nil{
		e.prev.next = e.next
		if e.next == nil{
			l.tail = e.prev
		}else {
			e.next.prev = e.Prev()
		}
		e.prev = nil
		e.next = l.head
		l.head.prev = e
		l.head = e
	}
}

func (l *LRUCache) remove(e *Element){
	if e.prev == nil{
		l.head = e.next
	}else {
		e.prev.next = e.next
	}
	if e.next == nil{
		l.tail = e.prev
	}else {
		e.next.prev = e.prev
	}
	l.capacity--
}