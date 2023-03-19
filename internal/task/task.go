package task

import (
	"context"
	"errors"
	"strconv"
)

//Reglas de Dominio o Negocio

var ErrInvalidTaskID = errors.New("Invalid Task ID")

type ID struct {
	value int
}

func NewTaskID(value string) (ID, error) {

	v, err := strconv.Atoi(value)

	if err != nil {
		return ID{}, ErrInvalidTaskID
	}

	return ID{value: v}, nil
}

var ErrInvalidTaskName = errors.New("Invalid Task Name")

type Name struct {
	value string
}

func NewtaskName(value string) (Name, error) {
	if value == "" {
		return Name{}, ErrInvalidTaskName
	}

	return Name{value: value}, nil
}

var ErrInvalidTaskContent = errors.New("Invalid Task Content")

type Content struct {
	value string
}

func NewTaskContent(value string) (Content, error) {
	if value == "" {
		return Content{}, ErrInvalidTaskContent
	}

	return Content{value: value}, nil
}

type Task struct {
	Id      ID
	Name    Name
	Content Content
}

func NewTask(id, name, content string) (Task, error) {
	Id, errId := NewTaskID(id)
	Name, errName := NewtaskName(name)
	Content, errContent := NewTaskContent(content)

	if errId != nil {
		return Task{}, errId
	} else if errName != nil {
		return Task{}, errName
	} else if errContent != nil {
		return Task{}, errContent
	}

	return Task{
		Id:      Id,
		Name:    Name,
		Content: Content,
	}, nil
}

func (c Task) IDs() ID {
	return c.Id
}

func (c Task) Names() Name {
	return c.Name
}

func (c Task) Contents() Content {
	return c.Content
}

func (id ID) String() string {
	return strconv.Itoa(id.value)
}

func (name Name) String() string {
	return name.value
}

func (content Content) String() string {
	return content.value
}

//Interface para guardar course
type TaskSaveRepository interface {
	Save(ctx context.Context, task Task) error
}

type TaskUpdateRepository interface {
	Update(ctx context.Context, task Task) error
}

type TaskDeleteRepository interface {
	Delete(ctx context.Context, Id string) error
}

type TaskGetAllRepository interface {
	ReadAll(ctx context.Context) ([]TaskDB, error)
}

type TaskGetIdRepository interface {
	ReadID(ctx context.Context, Id string) ([]TaskDB, error)
}

type TaskDB struct {
	ID      int
	Name    string
	Content string
}
