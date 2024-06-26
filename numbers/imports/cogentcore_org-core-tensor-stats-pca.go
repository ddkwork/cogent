// Code generated by 'yaegi extract cogentcore.org/core/tensor/stats/pca'. DO NOT EDIT.

package imports

import (
	"cogentcore.org/core/tensor/stats/pca"
	"reflect"
)

func init() {
	Symbols["cogentcore.org/core/tensor/stats/pca/pca"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CovarTableCol":    reflect.ValueOf(pca.CovarTableCol),
		"CovarTableColStd": reflect.ValueOf(pca.CovarTableColStd),
		"CovarTensor":      reflect.ValueOf(pca.CovarTensor),
		"CovarTensorStd":   reflect.ValueOf(pca.CovarTensorStd),
		"TableColRowsVec":  reflect.ValueOf(pca.TableColRowsVec),
		"TensorRowsVec":    reflect.ValueOf(pca.TensorRowsVec),

		// type definitions
		"PCA": reflect.ValueOf((*pca.PCA)(nil)),
		"SVD": reflect.ValueOf((*pca.SVD)(nil)),
	}
}
