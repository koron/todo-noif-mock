package service

import "testing"

func TestServiceCreate(t *testing.T) {
	r := &Repository{
		// モック用のRepositoryの初期化をここに書く
	}
	srv := &Service{repos: r}
	d, err := srv.PostTodo("テストのタスク1")
	if err != nil {
		t.Fatal(err)
	}

	// TODO: dをチェックする
	_ = d

	// TODO: Createが正しく呼ばれたか`r`の中身をチェックするなど
}
