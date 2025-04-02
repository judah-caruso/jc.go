package gm

type RectCutDirection int

const (
	FromLeft RectCutDirection = iota
	FromRight
	FromTop
	FromBottom
)

type RectCut struct {
	rect *Rect
	dir  RectCutDirection
}

func PrepareRectCut(r *Rect, dir RectCutDirection) RectCut {
	return RectCut{rect: r, dir: dir}
}

func (r RectCut) Cut(amt Float) Rect {
	switch r.dir {
	case FromLeft:
		return r.rect.CutFromLeft(amt)
	case FromRight:
		return r.rect.CutFromRight(amt)
	case FromTop:
		return r.rect.CutFromTop(amt)
	case FromBottom:
		return r.rect.CutFromBottom(amt)
	default:
		panic("unreachable")
	}
}
