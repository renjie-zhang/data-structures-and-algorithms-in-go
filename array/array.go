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
package array

type Arr struct {
	data []int
	length int
}

func (a Arr) NewArray(count int) *Arr {
	t := Arr{
		data:   make([]int,count,count),
		length: 0,
	}
	return &t
}

func (a Arr) isOutOfIndex(index int)bool{
	if index > cap(a.data){
		return false
	}
	return true
}

func (a Arr) Count() int{
	return len(a.data)
}

// Add 自动进行扩容，并且按照顺序进行存放
func (a Arr) Add(v int) (bool,error){
	// 如果当前len == cap，那么进行扩容为当前的两倍
	if len(a.data) == cap(a.data){
		temp := make([]int,cap(a.data)*2)
		copy(a.data,temp)
		a.data = temp
	}
	// 判断当前的值应该存放在那个index下
	tempIndex := -1
	for i := 0; i < len(a.data); i++ {
		if a.data[i] <= v{
			tempIndex = i +1
		}
	}
	// index等于-1.那么v应该放置到最前面
	if tempIndex == -1{
		a.addFirst(v)
	}
	// 将v放置到tempIndex，并且将后面的元素后移
	for i := len(a.data); i >tempIndex ; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[tempIndex] = v
	return true,nil
}
// addFirst 添加v到数组的最前面
func (a Arr) addFirst(v int) {
	temp := make([]int,len(a.data),cap(a.data))
	temp = append(temp,v)
	temp = append(temp,a.data...)
}
