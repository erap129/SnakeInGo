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

func (sb *SnakeBody) Update(width int, height int) {
	sb.X = (sb.X + sb.Xspeed) % width
	if sb.X < 0 {
		sb.X += width
	}
	sb.Y = (sb.Y + sb.Yspeed) % height
	if sb.Y < 0 {
		sb.Y += height
	}
}
