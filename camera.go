package main

import "github.com/go-gl/mathgl/mgl32"

// CameraMovement stored as int
type CameraMovement int

// forward is 0, backward is 1, left is 2, right is 3
const (
	MoveForward CameraMovement = iota
	MoveBackward
	MoveLeft
	MoveRight
)

// Camera object stores camera state
type Camera struct {
	pos    mgl32.Vec3
	up     mgl32.Vec3
	right  mgl32.Vec3
	front  mgl32.Vec3
	wfront mgl32.Vec3

	rotatex, rotatey float32

	Sens float32 // speed of rotations?

	flying bool
}

// NewCamera returns a new camera object
func NewCamera(pos mgl32.Vec3) *Camera {
	c := &Camera{
		pos:     pos,
		front:   mgl32.Vec3{0, 0, -1},
		rotatey: 0,
		rotatex: -90,
		Sens:    0.14,
		flying:  false,
	}
	c.updateAngles()
	return c
}

// Restore sets camera state to relevant values from PlayerState
func (c *Camera) Restore(state PlayerState) {
	c.pos = mgl32.Vec3{state.X, state.Y, state.Z}
	c.rotatex = state.Rx
	c.rotatey = state.Ry
	c.updateAngles()
}

// State serializes camera state into a PlayerState object
func (c *Camera) State() PlayerState {
	return PlayerState{
		X:  c.pos.X(),
		Y:  c.pos.Y(),
		Z:  c.pos.Z(),
		Rx: c.rotatex,
		Ry: c.rotatey,
	}
}

// Matrix returns a 4x4 transform matrix from world space->eye space
func (c *Camera) Matrix() mgl32.Mat4 {
	return mgl32.LookAtV(c.pos, c.pos.Add(c.front), c.up)
}

// SetPos sets the camera location.
func (c *Camera) SetPos(pos mgl32.Vec3) {
	c.pos = pos
}

// Pos is a getter for Camera.pos
func (c *Camera) Pos() mgl32.Vec3 {
	return c.pos
}

// Front is a getter for Camera.front
func (c *Camera) Front() mgl32.Vec3 {
	return c.front
}

// FlipFlying toggles Camera.flying (on TAB press)
func (c *Camera) FlipFlying() {
	c.flying = !c.flying
}

// Flying is a getter for Camera.flying
func (c *Camera) Flying() bool {
	return c.flying
}

// OnAngleChange rotates the camera
func (c *Camera) OnAngleChange(dx, dy float32) {
	if mgl32.Abs(dx) > 200 || mgl32.Abs(dy) > 200 {
		return
	}
	c.rotatex += dx * c.Sens
	c.rotatey += dy * c.Sens
	if c.rotatey > 89 {
		c.rotatey = 89
	}
	if c.rotatey < -89 {
		c.rotatey = -89
	}
	c.updateAngles()
}

// OnMoveChange moves the camera in XYZ space
func (c *Camera) OnMoveChange(dir CameraMovement, delta float32) {
	if c.flying {
		delta = 5 * delta
	}
	switch dir {
	case MoveForward:
		if c.flying {
			c.pos = c.pos.Add(c.front.Mul(delta))
		} else {
			c.pos = c.pos.Add(c.wfront.Mul(delta))
		}
	case MoveBackward:
		if c.flying {
			c.pos = c.pos.Sub(c.front.Mul(delta))
		} else {
			c.pos = c.pos.Sub(c.wfront.Mul(delta))
		}
	case MoveLeft:
		c.pos = c.pos.Sub(c.right.Mul(delta))
	case MoveRight:
		c.pos = c.pos.Add(c.right.Mul(delta))
	}
}

// updateAngles gets the new angles after rotations
func (c *Camera) updateAngles() {
	front := mgl32.Vec3{
		cos(radian(c.rotatey)) * cos(radian(c.rotatex)),
		sin(radian(c.rotatey)),
		cos(radian(c.rotatey)) * sin(radian(c.rotatex)),
	}
	c.front = front.Normalize()
	c.right = c.front.Cross(mgl32.Vec3{0, 1, 0}).Normalize()
	c.up = c.right.Cross(c.front).Normalize()
	c.wfront = mgl32.Vec3{0, 1, 0}.Cross(c.right).Normalize()
}
