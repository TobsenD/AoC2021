package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	start := time.Now()
	task()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)

}

type Missle struct {
	movX int
	movY int
	posX int
	posY int
}

func (missle *Missle) Move() {
	missle.posX += missle.movX
	missle.posY += missle.movY
	if missle.movX > 0 {
		missle.movX--
	} else if missle.movX < 0 {
		missle.movX++
	}
	missle.movY--
}

type Area struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

func task() {

	area := Area{48, 70, -189, -148}

	//calculating maxy for Task1
	maxY := ((area.y1 + 1) * area.y1) / 2

	fmt.Println("Task01: ", maxY)

	//This will take a while
	hits := calculateRoute(area, maxY)
	fmt.Println("Task02: ", hits)

}

func calculateRoute(area Area, maxY int) int {
	hits := 0
	for x := 1; x <= area.x2; x++ {
		for y := area.y1; y < maxY; y++ {
			missle := Missle{
				movX: x,
				movY: y,
			}
			if fireTarget(missle, area) {
				hits++
			}
		}
	}
	return hits
}

func fireTarget(missle Missle, area Area) bool {
	hit := false

	for !hit {
		missle.Move()

		if missle.posX >= area.x1 && missle.posX <= area.x2 {
			if missle.posY >= area.y1 && missle.posY <= area.y2 {
				hit = true
			}
		}
		if missle.posX > area.x2 || missle.posY < area.y1 {
			break
		}
	}
	return hit
}
