// +build mock

package service

import "github.com/koron/todo-noif-mock/do"

type Repository struct {
    // モックに必要なフィールドをここに置く
}

func (m *Repository) Insert(todo *do.Todo) (*do.Todo, error) {
    // モックのコードをココに書く
    return todo, nil
}
