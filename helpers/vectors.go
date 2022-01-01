package helpers

type Vec2 struct {
	X, Y int
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v.X + v2.X, v.Y + v2.Y}
}

func (v Vec2) Equal(end Vec2) bool {
	return v.X == end.X && v.Y == end.Y
}

func (v Vec2) ManhattanDistance(end Vec2) int {
	return AbsInt(v.X-end.X) + AbsInt(v.Y-end.Y)
}

var (
	UpVector    = Vec2{0, -1}
	DownVector  = Vec2{0, 1}
	LeftVector  = Vec2{-1, 0}
	RightVector = Vec2{1, 0}

	DownRightVector = Vec2{1, 1}
	DownLeftVector  = Vec2{-1, 1}

	UpRightVector = Vec2{1, -1}
	UpLeftVector  = Vec2{-1, -1}

	ZeroVector = Vec2{0, 0}

	NineDirections = []Vec2{UpLeftVector, UpVector, UpRightVector, LeftVector, ZeroVector, RightVector, DownLeftVector, DownVector, DownRightVector}
	EightDirections = []Vec2{UpVector, DownVector, LeftVector, RightVector, DownRightVector, DownLeftVector, UpRightVector, UpLeftVector}
	FourDirections  = []Vec2{UpVector, DownVector, LeftVector, RightVector}
)

var (
	Xp = Vec3{X: 1}
	Xm = Vec3{X: -1}
	Yp = Vec3{Y: 1}
	Ym = Vec3{Y: -1}
	Zp = Vec3{Z: 1}
	Zm = Vec3{Z: -1}

	UpsZ = []Vec3{Yp, Ym, Xp, Xm}
	UpsX = []Vec3{Yp, Ym, Zp, Zm}
	UpsY = []Vec3{Xp, Xm, Zp, Zm}
)

type Vec3 struct {
	X, Y, Z int
}

func (v Vec3) RotateZ(count int) Vec3 {
	for i := 0; i < count; i++ {
		v = Vec3{
			Z: v.Z,
			X: v.Y,
			Y: -v.X,
		}
	}

	return v
}

func (v Vec3) RotateX(count int) Vec3 {
	for i := 0; i < count; i++ {
		v = Vec3{
			Z: v.Y,
			X: v.X,
			Y: -v.Z,
		}
	}

	return v
}

func (v Vec3) RotateY(count int) Vec3 {
	for i := 0; i < count; i++ {
		v = Vec3{
			Z: -v.X,
			X: v.Z,
			Y: v.Y,
		}
	}

	return v
}

func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func AllRotations() []Vec3 {
	rotations := make([]Vec3, 0, 24)
	yRotations := []int{1, 3}

	for z := 0; z < 4; z++ {
		for x := 0; x < 4; x++ {
			rotations = append(rotations, Vec3{X: x, Y: 0, Z: z})
		}
	}

	for _, y := range yRotations {
		for x := 0; x < 4; x++ {
			rotations = append(rotations, Vec3{X: x, Y: y, Z: 0})
		}
	}

	return rotations
}

var Rotations = AllRotations()
