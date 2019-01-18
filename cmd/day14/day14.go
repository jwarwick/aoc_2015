package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	max, points := distance(string(input), 2503)
	fmt.Println("Part 1: ", max)
	fmt.Println("Part 2: ", points)
}

func distance(input string, steps int) (int, int) {
	deer := buildSpeeds(input)

	for _, d := range deer {
		d.initialize()
	}

	for t := 0; t < steps; t++ {
		for _, d := range deer {
			d.step()
		}
		step_max := maxDistance(deer)
		for _, d := range deer {
			d.maybeTakePoints(step_max)
		}
	}

	max := maxDistance(deer)
	points := maxPoints(deer)

	return max, points
}

func maxDistance(deer []*reindeer) int {
	max := 0
	for _, d := range deer {
		if d.distance > max {
			max = d.distance
		}
	}
	return max
}

func maxPoints(deer []*reindeer) int {
	max := 0
	for _, d := range deer {
		if d.points > max {
			max = d.points
		}
	}
	return max
}

type state int

const (
	flying state = iota
	resting
)

type reindeer struct {
	name      string
	speed     int
	on        int
	off       int
	distance  int
	doing     state
	remaining int
	points    int
}

func (r *reindeer) initialize() {
	r.distance = 0
	r.doing = flying
	r.remaining = r.on
}

func (r *reindeer) maybeTakePoints(dist int) {
	if r.distance >= dist {
		r.points++
	}
}

func (r *reindeer) step() {
	if r.doing == flying {
		r.distance += r.speed
	}

	r.remaining -= 1
	if r.remaining <= 0 {
		r.doing = toggleState(r.doing)
		switch r.doing {
		case flying:
			r.remaining = r.on
		case resting:
			r.remaining = r.off
		}
	}
}

func toggleState(s state) state {
	if s == flying {
		return resting
	} else {
		return flying
	}
}

func buildSpeeds(input string) []*reindeer {
	lines := splitInput(input)
	r := make([]*reindeer, len(lines))
	for idx, l := range lines {
		name, speed, on, off := parseLine(l)
		r[idx] = &reindeer{name: name, speed: speed, on: on, off: off}
	}
	return r
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return lines
}

func parseLine(input string) (name string, speed int, on int, off int) {
	words := strings.Split(input, " ")
	name = words[0]
	speed = parseInt(words[3])
	on = parseInt(words[6])
	off = parseInt(words[13])
	return
}

func parseInt(input string) int {
	i, err := strconv.Atoi(input)
	if nil != err {
		panic("Not a number when parsing")
	}
	return i
}
