package helpers

import (
	"testing"

	is_ "github.com/matryer/is"
)

func TestVec3RotateTop(t *testing.T) {
	is := is_.New(t)

	v := Vec3{X: 3, Y: 1}

	v = v.RotateZ(1)
	is.Equal(v, Vec3{X: 1, Y: -3}) // rotate 1

	v = v.RotateZ(1)
	is.Equal(v, Vec3{X: -3, Y: -1}) // rotate 2

	v = v.RotateZ(1)
	is.Equal(v, Vec3{X: -1, Y: 3}) // rotate 3

	v = v.RotateZ(1)
	is.Equal(v, Vec3{X: 3, Y: 1})
}

func TestVec3RotateSide(t *testing.T) {
	is := is_.New(t)

	v := Vec3{Z: 3, Y: 1}

	v = v.RotateX(1)
	is.Equal(v, Vec3{Z: 1, Y: -3}) // rotate 1

	v = v.RotateX(1)
	is.Equal(v, Vec3{Z: -3, Y: -1}) // rotate 2

	v = v.RotateX(1)
	is.Equal(v, Vec3{Z: -1, Y: 3}) // rotate 3

	v = v.RotateX(1)
	is.Equal(v, Vec3{Z: 3, Y: 1}) // rotate 4
}

func TestVec3RotateY(t *testing.T) {
	is := is_.New(t)

	v := Vec3{X: 3, Z: 1}

	v = v.RotateY(1)
	is.Equal(v, Vec3{X: 1, Z: -3}) // rotate 1

	v = v.RotateY(1)
	is.Equal(v, Vec3{X: -3, Z: -1}) // rotate 2

	v = v.RotateY(1)
	is.Equal(v, Vec3{X: -1, Z: 3}) // rotate 3

	v = v.RotateY(1)
	is.Equal(v, Vec3{X: 3, Z: 1}) // rotate 4
}

func TestVec3Add(t *testing.T) {
	is := is_.New(t)

	v := Vec3{X: 1, Y: 1, Z: 1}

	is.Equal(v.Add(Vec3{X: 1, Y: 2, Z: 3}), Vec3{X: 2, Y: 3, Z: 4})
}

func TestVec3_UniqueRotations(t *testing.T) {
	is := is_.New(t)

	init := Vec3{X: 1, Y: 2, Z: 3}

	set := make(map[Vec3]struct{})

	yRotations := []int{1, 3}

	for z := 0; z < 4; z++ {
		for x := 0; x < 4; x++ {
			v := init.RotateZ(z).RotateX(x)

			if _, ok := set[v]; ok {
				is.Fail() // already in set
			}

			set[v] = struct{}{}
		}
	}

	for _, y := range yRotations {
		for x := 0; x < 4; x++ {
			v := init.RotateY(y).RotateX(x)

			if _, ok := set[v]; ok {
				is.Fail() // already in set
			}

			set[v] = struct{}{}
		}
	}

	is.Equal(len(set), 24) // should produce 24 unique coordinates
}
