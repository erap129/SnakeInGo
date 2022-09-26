package main

type SnakeBody struct {
	X      int
	Y      int
	Xspeed int
	Yspeed int
}

func (sb *SnakeBody) ChangeDir(vertical int, horizontal int) {
	sb.Yspeed = vertical
	sb.Xspeed = horizontal
}

func (sb *SnakeBody) Update() {
	sb.X += sb.Xspeed
	sb.Y += sb.Yspeed
}
