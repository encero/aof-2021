package helpers

type Vector struct {
	X, Y int
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.X + v2.X, v.Y + v2.Y}
}

func (v Vector) Equal(end Vector) bool {
	return v.X == end.X && v.Y == end.Y
}

func (v Vector) ManhattanDistance(end Vector) int {
	return AbsInt(v.X-end.X) + AbsInt(v.Y-end.Y)
}

var (
	UpVector    = Vector{0, -1}
	DownVector  = Vector{0, 1}
	LeftVector  = Vector{-1, 0}
	RightVector = Vector{1, 0}

	DownRightVector = Vector{1, 1}
	DownLeftVector  = Vector{-1, 1}

	UpRightVector = Vector{1, -1}
	UpLeftVector  = Vector{-1, -1}

	ZeroVector = Vector{0, 0}

	EightDirections = []Vector{UpVector, DownVector, LeftVector, RightVector, DownRightVector, DownLeftVector, UpRightVector, UpLeftVector}
	FourDirections  = []Vector{UpVector, DownVector, LeftVector, RightVector}
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
