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
package insertsort

type tree struct{
	value int
	left,right *tree
}

func InsertSortByTree(values []int){
	var root *tree
	for _,v := range values{
		root = add(root,v)
	}
	appendValues(values[:0],root)
}

func appendValues(values []int,t *tree ) []int{
	if t!= nil{
		values = appendValues(values,t.left)
		values =  append(values,t.value)
		values = appendValues(values,t.right)
	}
	return values
}

func add(t *tree,value int) *tree{
	if t == nil{
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value{
		t.left = add(t.left,value)
	}else {
		t.right = add(t.right,value)
	}
	return t
}