package main

import "fmt"

type QueryResultsIterator interface {
	first() string
	next() string
	currentItem() string
	isDone() bool
}

type QueryResults interface {
	createIterator() QueryResultsIterator
	fetchData()
}

type QueryResultsIteratorImpl struct {
	results [] string
	cursor int
	max int
}

func (this *QueryResultsIteratorImpl) SetQueryResultsIteratorImpl(r []string)  {
	this.results = r
	this.cursor = 0
	this.max = len(r)
}

func (this *QueryResultsIteratorImpl)First() string{
	return this.results[0]
}
func (this *QueryResultsIteratorImpl) Next()string{
	this.cursor++
	if !this.IsDone(){
		return this.results[this.cursor]
	}else {
		return "<end>"
	}
}

func (this *QueryResultsIteratorImpl) CurrentItem() string  {
	return this.results[this.cursor]
}

func (this *QueryResultsIteratorImpl) IsDone()bool{
	return (this.cursor == this.max)
}

type QueryResultsImpl struct {
	results [] string
}

func (this *QueryResultsImpl) FetchData()  {
	this.results = append(this.results,"Item 1","Item 2","Item 3","Item 4","Item 5","Item 6","Item 7","Item 8","Item 9","Item 10")
}

func (this *QueryResultsImpl) CreateIterator() QueryResultsIteratorImpl  {
	q := new(QueryResultsIteratorImpl)
	q.SetQueryResultsIteratorImpl(this.results)
	return *q
}

func main(){
	query := "select * from test"
	fmt.Println("Executing query: ", query)
	q := new(QueryResultsImpl)
	q.FetchData()
	iter := q.CreateIterator()
	for !iter.IsDone() {
		fmt.Println(iter.CurrentItem())
		iter.Next()
	}
}