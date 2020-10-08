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

// LRUCache is a simple LRU interface
type LRUCache interface {
	// Add a value to the cache
	Add(key, value interface{}) bool
	// Return key's value from the cache
	Get(key interface{}) (value interface{}, ok bool)
	// Check if a key exists in cache
	Contains(key interface{}) (ok bool)
	// Return key's value without update
	Peek(key interface{}) (value interface{}, ok bool)
	// Remove a key from the cache
	Remove(key interface{}) bool
	// Removes the oldest element from the cache
	RemoveOldest(interface{}, interface{}, bool)
	// Get the oldest element from the cache
	GetOldest(interface{}, interface{}, bool)
	// Return a slice of the keys in the cache
	Keys() []interface{}
	// Return the number of items in the cache
	Len() int
	// Clear all cache element
	Purge()
	// Resize the cache
	Resize(int) int
}
