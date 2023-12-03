package datautil

import (
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/mat"
	"gorgonia.org/tensor"
)

// Convert the given columns of the given dataframe to a tensor.
// The tensor will be of type float64, and the source columns must also have type float64.
func ColumnsToTensor(df dataframe.DataFrame, cols []string) *tensor.Dense {
	return tensor.FromMat64(mat.DenseCopyOf(&gotaMatrix{df.Select(cols)}))
}

// Convert the entire dataframe to a tensor.
// The tensor will be of type float64, and the source columns must also have type float64.
func AllColumnsToTensor(df dataframe.DataFrame) *tensor.Dense {
	return tensor.FromMat64(mat.DenseCopyOf(&gotaMatrix{df}))
}
