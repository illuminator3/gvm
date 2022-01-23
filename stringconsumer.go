package main

type StringConsumer struct {
	s string
}

func (c *StringConsumer) Next() (s string) {
	s = c.Peek()
	c.s = c.s[1:]

	return
}

func (c *StringConsumer) Peek() (s string) {
	s = string(c.s[0])

	return
}

func (c *StringConsumer) Everything() (s string) {
	s = c.s

	return
}

func (c *StringConsumer) HasNext() bool {
	return len(c.s) > 0
}

func (c *StringConsumer) TakeWhile(f func(string) bool) (s string) {
	for c.HasNext() {
		if !f(c.Peek()) {
			break
		}

		s += c.Next()
	}

	return
}

func (c *StringConsumer) Consume(s string) {
	c.s += s
}

func CreateStringConsumer(s string) *StringConsumer {
	return &StringConsumer{s}
}

func CreateEmptyStringConsumer() *StringConsumer {
	return &StringConsumer{""}
}
