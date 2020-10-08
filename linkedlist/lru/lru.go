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

import (
	"container/list"
	"errors"
)

// EvictCallBack is used to get a callback when a cache element is evicted
type EvictCallBack func(key interface{}, value interface{})

type LRU struct {
	size      int
	evictList *list.List
	items     map[interface{}]*list.Element
	onEvict   EvictCallBack
}

type entry struct {
	key   interface{}
	value interface{}
}

func NewLRU(size int, onEvict EvictCallBack) (*LRU, error) {
	if size <= 0 {
		return nil, errors.New("size must bigger than 0")
	}
	c := &LRU{
		size:      size,
		evictList: list.New(),
		items:     make(map[interface{}]*list.Element),
		onEvict:   onEvict,
	}
	return c, nil
}

func (c *LRU) Purge() {
	for k, v := range c.items {
		if c.onEvict != nil {
			c.onEvict(k, v.Value.(*entry).value)
		}
		delete(c.items, k)
	}
	c.evictList.Init()
}

func (c *LRU) Add(key, value interface{}) (evicted bool) {
	// check existing item
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.(*entry).value = value
		return false
	}
	// Add new item
	ent := &entry{key, value}
	entry := c.evictList.PushFront(ent)
	c.items[key] = entry

	evict := c.evictList.Len() > c.size
	if evict {
		c.removeOldest()
	}
	return evict
}

func (c *LRU) Get(key interface{}) (value interface{},ok bool)  {
	if ent,ok := c.items[key];ok{
		c.evictList.MoveToFront(ent)
		if ent.Value.(*entry) == nil{
			return nil,false
		}
		return ent.Value.(*entry).value,true
	}
	return
}

func (c *LRU) Contains(key interface{}) (ok bool){
	_,ok = c.items[key]
	return
}

func (c *LRU) Peek(key interface{}) (value interface{},ok bool){
	var ent *list.Element
	if ent,ok = c.items[key];ok{
		return ent.Value.(*entry).value,true
	}
	return nil,ok
}

func (c *LRU) Remove(key interface{}) (present bool){
	if ent,ok := c.items[key];ok{
		c.removeElement(ent)
		return true
	}
	return false
}

func (c *LRU) RemoveOldest() (key,value interface{} ,ok bool){
	ent := c.evictList.Back()
	if ent != nil{
		c.removeElement(ent)
		kv := ent.Value.(*entry)
		return kv.key,kv.value,true
	}
	return nil,nil,false
}


func (c *LRU) GetOldest() (key,value interface{},ok bool){
	ent := c.evictList.Back()
	if ent != nil{
		kv := ent.Value.(*entry)
		return kv.key,kv.value,true
	}
	return nil,nil,false
}



func (c *LRU) Keys() []interface{}{
	keys := make([]interface{},len(c.items))
	i := 0
	for ent := c.evictList.Back();ent != nil;ent = ent.Prev(){
		keys[i] = ent.Value.(*entry).key
		i++
	}
	return keys
}

func (c *LRU) Resize(size int) (evicted int) {
	diff := c.Len() - size
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff; i++ {
		c.removeOldest()
	}
	c.size = size
	return diff
}

func (c *LRU) Len() int {
	return c.evictList.Len()
}

func (c *LRU) removeElement(e *list.Element) {
	c.evictList.Remove(e)
	kv := e.Value.(*entry)
	delete(c.items, kv.key)
	if c.onEvict != nil {
		c.onEvict(kv.key, kv.value)
	}
}

func (c *LRU) removeOldest() {
	back := c.evictList.Back()
	if back != nil {
		c.removeElement(back)
	}
}
