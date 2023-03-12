package main

import (
	"fmt"
	"sync"
)

func main() {

	var a int

	incA := func() {
		a++
	}

	incA()
	incA()
	incA()

	fmt.Println(a)
	//-------------
	var b int
	var once sync.Once

	incB := func() {
		once.Do(func() {
			b++
		})
	}

	incB()
	incB()
	incB()

	fmt.Println(b)
}

type Cache struct {
	data interface{}
	mx   sync.Mutex
	once sync.Once
}

func NewCache() *Cache {
	c := &Cache{}
	c.mx.Lock()
	return c
}

func (c *Cache) WriteData(data interface{}) {
	c.data = data
	c.once.Do(func() {
		c.mx.Unlock()
	})
}

func (c *Cache) GetData() interface{} {
	c.mx.Lock()
	defer c.mx.Unlock()
	return c.data
}
