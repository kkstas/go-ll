package main

import (
	"errors"
	"fmt"
)

type Node struct {
	next  *Node
	value string
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(value string) {
	newNode := Node{value: value}

	if l.head == nil {
		l.head = &newNode
		return
	}

	currentHead := l.head
	for currentHead.next != nil {
		currentHead = currentHead.next
	}
	currentHead.next = &newNode
}

func (l *LinkedList) Print() {
	if l.head == nil {
		return
	}

	currentHead := l.head
	var counter uint64 = 0

	for currentHead != nil {
		fmt.Printf("%v: %s\n", counter, currentHead.value)
		currentHead = currentHead.next
		counter++
	}
}

func (l *LinkedList) Find(s string) (*Node, error) {
	if l.head == nil {
		return nil, errors.New("linked list is empty")
	}
	currentHead := l.head
	var counter uint64 = 0

	for currentHead != nil {
		if currentHead.value == s {
			return currentHead, nil
		}
		currentHead = currentHead.next
		counter++
	}
	return nil, errors.New("string " + s + " not found in list")
}

func (l *LinkedList) Delete(s string) {
	if l.head == nil {
		return
	}
	if l.head.value == s {
		l.head = l.head.next
		return
	}

	prevHead := l.head
	currentHead := l.head.next
	found := false

	for currentHead != nil && !found {
		if currentHead.value == s {
			found = true
			prevHead.next = currentHead.next
			return
		}
		currentHead = currentHead.next
	}
}

func (l *LinkedList) InsertAfter(s, value string) {
	if l.head == nil {
		return
	}

	currentHead := l.head

	for currentHead.next != nil {
		if currentHead.value == s {
			currentHead.next = &Node{value: value, next: currentHead.next}
			return
		}
		currentHead = currentHead.next
	}

	currentHead.next = &Node{value: value}
}

func (l *LinkedList) FindByIndex(i uint) (*Node, error) {
	currentHead := l.head
	var counter uint = 0

	for currentHead != nil {
		if counter != i {
			currentHead = currentHead.next
			counter++
		} else {
			return currentHead, nil
		}

	}

	return nil, errors.New("Index too large.")
}
