package sample

import "testing"

// NOTE: 前提として、DBがないので、DBに繋げるというのはできない
func TestUser(t *testing.T) {
	t.Fail()
	// テストユーザのサブテストを書く
	// 検査対象は、email を送ってるということと
	// DBに保存されている内容
	// 保存することにフォーカスするなら、元の Userオブジェクトの値を更新するんじゃなくて、新しくUserという構造体を作って更新する方法もありそう？
	// 構造体のフィールド値が変更することも含まれそう
	t.Run(
		"ユーザが Email を変更できる", func(t *testing.T) {
			// 検証するもの
			// 準備フェーズ
			beforeEmail := "test@test.com"
			afterEmail := "test1@test.com"
			u := User{
				UserID: 123,
				Email:  beforeEmail,
				Type:   Customer,
			}
			// 実行フェーズ(act)
			u.ChangeEmail(u.UserID, afterEmail)

			// 検証フェーズ
			if u.Email != afterEmail {
				t.Fail()
			}
		},
	)
}

// TODO: どういうテストケースを書いていくか？
// テストケース出していくと、

//
