// Package primi exposes various primitive types
package primi

// Rect type holds top-left corner location of a rectangle along with its width and height
type Rect struct {
	X, Y, W, H float64
}

// Min returns the minimum of two given float64 numbers
func Min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}

// Max raturns the maximum of two given float64 numbers
func Max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}

// NewRect creates new Rect variable
func NewRect(x, y, width, height float64) *Rect {
	if width > 0 && height > 0 {
		r := &Rect{X: x, Y: y, W: width, H: height}
		return r
	}
	return nil
}

// Contains checks if given location is inside a rectangle
func (r *Rect) Contains(x, y float64) bool {
	if r != nil {
		if r.X <= x && x <= r.X+r.W {
			if r.Y <= y && y <= r.Y+r.H {
				return true
			}
		}
	}
	return false
}

// SetSize sets new width and height, if greater than zero
func (r *Rect) SetSize(neww, newh float64) {
	if r != nil {
		if neww > 0 && newh > 0 {
			r.W = neww
			r.H = newh
		}
	}
}

// MoveTo moves the top-left corner coordinates to a new location
func (r *Rect) MoveTo(x, y float64) {
	if r != nil {
		r.X = x
		r.Y = y
	}
}

// MoveBy moves the top-left corner coordinates by given delta values
func (r *Rect) MoveBy(dx, dy float64) {
	if r != nil {
		r.X += dx
		r.Y += dy
	}
}

// GrowBy increases/decreases size of rectangle
// (but only if final size is positive)
func (r *Rect) GrowBy(dw, dh float64) {
	if r != nil {
		neww := r.W + dw
		newh := r.H + dh
		if neww > 0 && newh > 0 {
			r.W = neww
			r.H = newh
		}
	}
}

// Collides checks if two rectangles overlap themselves
func (r *Rect) Collides(s *Rect) bool {
	if r != nil && s != nil {
		if r.X >= s.X+s.W || r.X+r.W <= s.X {
			return false
		}
		if r.Y >= s.Y+s.H || r.Y+r.H <= s.Y {
			return false
		}
		return true
	}
	return false
}


// Combine makes an union of two given rectangles
func (r *Rect) Combine(s *Rect) *Rect {
	if r != nil && s != nil {
		x11 := Min(r.X, s.X + s.W)
		x12 := Min(s.X, r.X + r.W)
		x13 := Min(x11, x12)
		y11 := Min(r.Y, s.Y + s.H)
		y12 := Min(s.Y, r.Y + r.H)
		y13 := Min(y11, y12)

		x21 := Max(r.X, s.X + s.W)
		x22 := Max(r.X + r.W, s.X)
		x23 := Max(x21, x22)

		y21 := Max(r.Y, s.Y + s.H)
		y22 := Max(r.Y + r.H, s.Y)
		y23 := Max(y21, y22)
		w := x23 - x13
		h := y23 - y13
		return &Rect{X: x13, Y: y13, W: w, H: h}
	}
	return nil
}


// CombineIfCollide makes an union of two given rectangles,
// but only if they overlap themselves
func (r *Rect) CombineIfCollide(s *Rect) *Rect {
	if r != nil && s != nil {
		if r.Collides(s) {
			return r.Combine(s)
		}
	}
	return nil
}
