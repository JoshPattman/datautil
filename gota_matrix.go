package datautil

import (
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/mat"
)

// This is a type (copied from gorgonia docs) that implements the mat.Matrix interface for a gota dataframe.
// This makes it easier to convert dataframe to tensor.
type gotaMatrix struct {
	dataframe.DataFrame
}

func (m gotaMatrix) At(i, j int) float64 {
	return m.Elem(i, j).Float()
}

func (m gotaMatrix) T() mat.Matrix {
	return mat.Transpose{Matrix: m}
}
