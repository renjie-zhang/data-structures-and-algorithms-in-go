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
package SingleLinkedList

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	list.Print()
}

func TestSinglyLinkedList_InsertFirst(t *testing.T) {
	list := New(9, 7, 5, 3, 1)
	list.InsertFirst(0)
	list.Print()
}

func TestSinglyLinkedList_InsertLast(t *testing.T) {
	list := New(9, 7, 5, 3, 1)
	list.InsertLast(0)
	list.Print()
}

func TestSinglyLinkedList_Add(t *testing.T) {
	list := New(1, 2, 3, 4)
	list.Add(5)
	list.Print()
}

func TestSinglyLinkedList_GetCount(t *testing.T) {
	list := New(1, 2, 3, 4)
	fmt.Println(list.GetCount())
}

func TestSinglyLinkedList_GetItems(t *testing.T) {
	list := New(1, 2, 3, 4)
	fmt.Println(list.GetItems())
}
