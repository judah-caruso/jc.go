package mg

type Rect struct {
	Min Vec2
	Max Vec2
}

func R(minx, miny, maxx, maxy Float) Rect {
	return Rect{
		Min: V2(minx, miny),
		Max: V2(maxx, maxy),
	}
}

func Ri(minx, miny, maxx, maxy int) Rect {
	return Rect{
		Min: V2i(minx, miny),
		Max: V2i(maxx, maxy),
	}
}

func Rv2(min, max Vec2) Rect {
	return Rect{
		Min: min,
		Max: max,
	}
}

func (r Rect) Width() Float {
	return r.Max.Sub(r.Min).X
}

func (r Rect) Height() Float {
	return r.Max.Sub(r.Min).Y
}

func (r Rect) Center() Vec2 {
	return V2(
		r.Width()/2,
		r.Height()/2,
	)
}

func (r Rect) Expand(amt Float) Rect {
	half := amt / 2
	return Rect{
		Min: r.Min.Subf(half),
		Max: r.Max.Addf(half),
	}
}

func (r Rect) Contract(amt Float) Rect {
	half := amt / 2
	return Rect{
		Min: r.Min.Addf(half),
		Max: r.Max.Subf(half),
	}
}

func (r *Rect) PrepareCut(dir RectCutDirection) RectCut {
	return PrepareRectCut(r, dir)
}

func (r *Rect) CutFromLeft(amt Float) Rect {
	pMinX := r.Min.X
	r.Min.X = min(r.Max.X, r.Min.X+amt)
	return Rect{
		Min: V2(pMinX, r.Min.Y),
		Max: V2(r.Min.X, r.Max.Y),
	}
}

func (r *Rect) CutFromRight(amt Float) Rect {
	pMaxX := r.Max.X
	r.Max.X = max(r.Min.X, r.Max.X-amt)
	return Rect{
		Min: V2(r.Max.X, r.Min.Y),
		Max: V2(pMaxX, r.Max.Y),
	}
}

func (r *Rect) CutFromTop(amt Float) Rect {
	pMinY := r.Min.Y
	r.Min.Y = min(r.Max.Y, r.Min.Y+amt)
	return Rect{
		Min: V2(r.Min.X, pMinY),
		Max: V2(r.Max.X, r.Min.Y),
	}
}

func (r *Rect) CutFromBottom(amt Float) Rect {
	pMaxY := r.Max.Y
	r.Max.Y = max(r.Min.Y, r.Max.Y-amt)
	return Rect{
		Min: V2(r.Min.X, r.Max.Y),
		Max: V2(r.Max.X, pMaxY),
	}
}
