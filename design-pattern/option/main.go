package main

import "fmt"

//选项模式
func main() {
	m1 := New(1001, "张三", "ddddd", "122222222222")
	m1.String()

	fmt.Println()
	fmt.Println("=========")
	m2 := NewOption(WithId(1002), WithName("lisi"))
	m2.String()
}

type Message struct {
	Id      int
	Name    string
	Address string
	Phone   string
}

func (m *Message) String() {
	fmt.Printf("Id:%s;Name:%s;Address:%s;Phone:%s;", m.Id, m.Name, m.Address, m.Phone)
}
func New(Id int, Name, Address, Phone string) Message {
	return Message{
		Id:      Id,
		Name:    Name,
		Address: Address,
		Phone:   Phone,
	}
}

type Option func(message *Message)

var defaultMessage = Message{
	Id:      -1,
	Name:    "-1",
	Address: "-1",
	Phone:   "-1",
}

func WithId(id int) Option {
	return func(message *Message) {
		message.Id = id
	}
}

func WithName(name string) Option {
	return func(message *Message) {
		message.Name = name
	}
}

func WithAddress(address string) Option {
	return func(message *Message) {
		message.Address = address
	}
}

func WithPhone(phone string) Option {
	return func(message *Message) {
		message.Phone = phone
	}
}

func NewOption(options ...Option) Message {
	message := defaultMessage
	for _, o := range options {
		o(&message)
	}
	return message
}
