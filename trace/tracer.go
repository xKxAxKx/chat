package trace

/*
  コード内でのイベントを記録できるオブジェクトを表すインタフェース
*/
type Tracer interface {
	Trace(...interface{})
}
