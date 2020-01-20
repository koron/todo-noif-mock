package service

import (
	"reflect"
	"testing"

	"github.com/koron/todo-noif-mock/do"
)

func TestServiceCreate(t *testing.T) {
	r := &Repository{
		// モック用のRepositoryの初期化をここに書く
		Insert_Rs: []*RepositoryInsert_R{
			{&do.Todo{ID: 123, Task: "mytask1"}, nil},
		},
	}
	srv := &Service{repos: r}
	d, err := srv.PostTodo("in_task1")
	if err != nil {
		t.Fatal(err)
	}
	// dをチェックする
	if !reflect.DeepEqual(d, &do.Todo{ID: 123, Task: "mytask1"}) {
		t.Fatal("unexpected response")
	}
	// Insertが正しく呼ばれたかチェック
	if !reflect.DeepEqual(r.Insert_Ps, []*RepositoryInsert_P{
		{&do.Todo{Task: "in_task1"}},
	}) {
		t.Fatal("unexpected Insert calls")
	}
}
