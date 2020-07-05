package collection

// Collection キュー・スタックのインターフェイス
type Collection interface {
	Push(interface{}) error //値を追加
	Pop() (interface{}, error) //値を取り出し&削除
	Peek() (interface{},error) //値を取り出し
	Size() (int, error) //サイズ
	Empty() (bool,error) //空か
}