package datautil

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func CategoricalSeries(src series.Series, classes []string) map[string]series.Series {
	catSeries := make(map[string]series.Series)
	for _, class := range classes {
		sName := src.Name + ":" + class
		catSeries[sName] = series.New(src.Compare(series.Eq, class), series.Float, sName)
	}
	return catSeries
}

func SplitCategoricalColumn(df dataframe.DataFrame, srcColName string, classes []string) dataframe.DataFrame {
	catSeries := CategoricalSeries(df.Col(srcColName), classes)
	for _, series := range catSeries {
		df = df.Mutate(series)
	}
	return df
}
