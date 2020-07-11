package collection

// Collection キュー・スタックのインターフェイス
type Collection interface {
	Push(*Node) error //値を追加
	Pop() (*Node, error) //値を取り出し&削除
	Peek() (*Node,error) //値を取り出し
	Size() (int, error) //サイズ
	Empty() (bool,error) //空か
	String() string
}