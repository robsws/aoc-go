package twelve

import (
	"aoc-go/utils"
	"aoc-go/vector"
	"log"
)

type ferry struct {
	Pos          vector.Vec2
	Dir          vector.Vec2
	UsesWaypoint bool
}

// MakeFerry - initialise a new ferry at 0,0 facing east
func makeFerry(usesWaypoint bool) ferry {
	pos := vector.Vec2{X: 0, Y: 0}
	var dir vector.Vec2
	if usesWaypoint {
		dir = vector.Vec2{X: 10, Y: -1}
	} else {
		dir = vector.Vec2{X: 1, Y: 0}
	}
	return ferry{Pos: pos, Dir: dir, UsesWaypoint: usesWaypoint}
}

func (f *ferry) TakeCommand(command rune, value int) {
	switch command {
	case 'N':
		if f.UsesWaypoint {
			f.Dir.Y -= value
		} else {
			f.Pos.Y -= value
		}
	case 'S':
		if f.UsesWaypoint {
			f.Dir.Y += value
		} else {
			f.Pos.Y += value
		}
	case 'E':
		if f.UsesWaypoint {
			f.Dir.X += value
		} else {
			f.Pos.X += value
		}
	case 'W':
		if f.UsesWaypoint {
			f.Dir.X -= value
		} else {
			f.Pos.X -= value
		}
	case 'L':
		f.rotate(true, value)
	case 'R':
		f.rotate(false, value)
	case 'F':
		f.Pos = f.Dir.Mul(value).Add(f.Pos)
	default:
		log.Fatal("Unsupported command", command)
	}
}

func (f *ferry) rotate(toLeft bool, degrees int) {
	// (1,0) <-> (0,-1) <-> (-1,0) <-> (0,1)
	if !toLeft {
		degrees = 360 - degrees
	}
	for i := 0; i < degrees; i += 90 {
		dir := f.Dir
		f.Dir.X = dir.Y
		f.Dir.Y = dir.X * -1
	}
}

func (f *ferry) ManhattanDistance() int {
	return utils.AbsInt(f.Pos.X) + utils.AbsInt(f.Pos.Y)
}
