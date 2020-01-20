package repository

import "github.com/koron/todo-noif-mock/do"

type Repository struct {
}

func (r *Repository) Insert(todo *do.Todo) (*do.Todo, error) {
    // 典型的なDB操作(insert)がココに入る
    return todo, nil
}
