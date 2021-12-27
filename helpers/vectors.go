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

    ZeroVector = Vector{0,0}

	EightDirections = []Vector{UpVector, DownVector, LeftVector, RightVector, DownRightVector, DownLeftVector, UpRightVector, UpLeftVector}
    FourDirections = []Vector{UpVector, DownVector, LeftVector, RightVector}
)

