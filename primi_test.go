package primi_test

import (
	"github.com/lvajxi03/primi"
	"testing"
)

func TestMin1(t *testing.T) {
	a := primi.Min(2, 3)
	if a != 2 {
		t.Fatal("Expected 2, got:", a)
	}
}

func TestMin2(t *testing.T) {
	a := primi.Min(3, 2)
	if a != 2 {
		t.Fatal("Expected 2, got:", a)
	}
}

func TestMax1(t *testing.T) {
	a := primi.Max(2, 3)
	if a != 3 {
		t.Fatal("Expected 3, got:", a)
	}
}

func TestMax2(t *testing.T) {
	a := primi.Max(4, 3)
	if a != 4 {
		t.Fatal("Expected 4, got:", a)
	}
}

func TestNew1(t *testing.T) {
	r := primi.NewRect(10, 10, 20, 20)
	if r == nil {
		t.Fatal("Expected non-nil rect")
	}
}

func TestNew2(t *testing.T) {
	r := primi.NewRect(5, 5, -3, 2)
	if r != nil {
		t.Fatal("Expected nil rect")
	}
}

func TestNew3(t *testing.T) {
	r := primi.NewRect(5, 5, 2, -3)
	if r != nil {
		t.Fatal("Expected nil rect")
	}
}

func TestContains1(t *testing.T) {
	r := primi.NewRect(5, 5, 5, 5)
	if ! r.Contains(6, 6) {
		t.Fatal("Rect shall contain <6, 6>")
	}
}

func TestContains2(t *testing.T) {
	r := primi.NewRect(5, 5, 5, 5)
	if ! r.Contains(5, 5) {
		t.Fatal("Rect shall contain <5, 5>")
	}
}

func TestContains3(t *testing.T) {
	r := primi.NewRect(5, 5, 5, 5)
	if r.Contains(4, 5) {
		t.Fatal("Rect shall not contain <4, 5>")
	}
}

func TestContains4(t *testing.T) {
	r := primi.NewRect(5, 5, 5, -1)
	if r.Contains(2, 2) {
		t.Fatal("Rect shall not exist")
	}
}

func TestSetSize1(t *testing.T) {
	r := primi.NewRect(5, 5, 5, 5)
	r.SetSize(6, 6)
	if r.W != 6 {
		t.Fatal("Rect width shall be 6, got:", r.W)
	}
	if r.H != 6 {
		t.Fatal("Rect height shall be 6, got:", r.H)
	}
}

func TestSetSize2(t *testing.T) {
	r := primi.NewRect(5, 5, 5, 5)
	r.SetSize(-1, 3)
	if r.W != 5 {
		t.Fatal("Rect width shall not change")
	}
	if r.H != 5 {
		t.Fatal("Rect height shall not change")
	}
}

func TestMoveTo1(t *testing.T) {
	r := primi.NewRect(5, 5, 5, 5)
	r.MoveTo(6, 6)
	if r.X != 6 {
		t.Fatal("Rect shall be moved to <6,6>")
	}
	if r.Y != 6 {
		t.Fatal("Rect shall be moved to <6, 6>")
	}
}

func TestCollides1(t *testing.T) {
	r := primi.NewRect(5, 5, 5, 5)
	s := primi.NewRect(11, 11, 11, 11)
	if r.Collides(s) {
		t.Fatal("Rects shall not collide")
	}  
}

func TestCollides2(t *testing.T) {
	r := primi.NewRect(5, 5, 5, 5)
	s := primi.NewRect(6, 6, 6, 6)
	if ! r.Collides(s) {
		t.Fatal("Rects shall collide!")
	}
}

func TestCombine1(t *testing.T) {
	r := primi.NewRect(1, 1, 1, 1)
	s := primi.NewRect(3, 3, 1, 1)
	u := r.Combine(s)
	if u.W != 3 {
		t.Fatal("New width shall be 3")
	}
	if u.H != 3 {
		t.Fatal("New height shall be 3")
	}
}

func TestCombine2(t *testing.T) {
	r := primi.NewRect(100, 100, 100, 200)
	s := primi.NewRect(50, 150, 30, 30)
	u := r.Combine(s)
	if u.W != 150 {
		t.Fatal("New width shall be 150")
	}
	if u.H != 200 {
		t.Fatal("Will shall still be 200")
	}
}

func TestCombine3(t *testing.T) {
	r := primi.NewRect(0, 0, 1, 1)
	s := primi.NewRect(0, 0, 0.1, 1.1)
	u := r.Combine(s)
	if u.W != 1 {
		t.Fatal("Width shall still be 1")
	}
	if u.H != 1.1 {
		t.Fatal("New width shall be 1.1 now")
	}
}

func TestCombineIfCollide1(t *testing.T) {
	r := primi.NewRect(1, 1, 1, 1)
	s := primi.NewRect(3, 3, 3, 3)
	u := r.CombineIfCollide(s)
	if u != nil {
		t.Fatal("Expected nil, found new Rect")
	}
}

func TestCombineIfCollide(t *testing.T) {
	r := primi.NewRect(1, 1, 1, 1)
	s := primi.NewRect(2, 2, 2, 2)
	u := r.Combine(s)
	if u.W != 3 {
		t.Fatal("New combined width shall be 3")
	}
	if u.H != 3 {
		t.Fatal("New combined height shall be 3")
	}
}

func TestGrowBy1(t *testing.T) {
	r := primi.NewRect(10, 10, 10, 10)
	r.GrowBy(5, 5)
	if r.W != 15 {
		t.Fatal("Expected new width = 15, found", r.W)
	}
	if r.H != 15 {
		t.Fatal("Expexted new heidht = 15, found", r.H)
	}
}
