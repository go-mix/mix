package sample

type Sample struct {
	Values []Value
}

func New(values []Value) Sample {
	return Sample{
		Values: values,
	}
}
