// Code generated by entc, DO NOT EDIT.

package ogent

import "ariga.io/ogent/example/todo/ent"

func NewTodoCreate(e *ent.Todo) *TodoCreate {
	if e == nil {
		return nil
	}
	return &TodoCreate{
		ID:    e.ID,
		Title: e.Title,
		Done:  NewOptBool(e.Done),
	}
}

func NewTodoCreates(es []*ent.Todo) []TodoCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]TodoCreate, len(es))
	for i, e := range es {
		r[i] = NewTodoCreate(e).Elem()
	}
	return r
}

func (t *TodoCreate) Elem() TodoCreate {
	if t == nil {
		return TodoCreate{}
	}
	return *t
}

func NewTodoList(e *ent.Todo) *TodoList {
	if e == nil {
		return nil
	}
	return &TodoList{
		ID:    e.ID,
		Title: e.Title,
		Done:  NewOptBool(e.Done),
	}
}

func NewTodoLists(es []*ent.Todo) []TodoList {
	if len(es) == 0 {
		return nil
	}
	r := make([]TodoList, len(es))
	for i, e := range es {
		r[i] = NewTodoList(e).Elem()
	}
	return r
}

func (t *TodoList) Elem() TodoList {
	if t == nil {
		return TodoList{}
	}
	return *t
}

func NewTodoRead(e *ent.Todo) *TodoRead {
	if e == nil {
		return nil
	}
	return &TodoRead{
		ID:    e.ID,
		Title: e.Title,
		Done:  NewOptBool(e.Done),
	}
}

func NewTodoReads(es []*ent.Todo) []TodoRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]TodoRead, len(es))
	for i, e := range es {
		r[i] = NewTodoRead(e).Elem()
	}
	return r
}

func (t *TodoRead) Elem() TodoRead {
	if t == nil {
		return TodoRead{}
	}
	return *t
}

func NewTodoUpdate(e *ent.Todo) *TodoUpdate {
	if e == nil {
		return nil
	}
	return &TodoUpdate{
		ID:    e.ID,
		Title: e.Title,
		Done:  NewOptBool(e.Done),
	}
}

func NewTodoUpdates(es []*ent.Todo) []TodoUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]TodoUpdate, len(es))
	for i, e := range es {
		r[i] = NewTodoUpdate(e).Elem()
	}
	return r
}

func (t *TodoUpdate) Elem() TodoUpdate {
	if t == nil {
		return TodoUpdate{}
	}
	return *t
}
