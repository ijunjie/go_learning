package pipefilter

type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

//func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
//	return &StraightPipeline{
//		Name:    name,
//		Filters: &filters,
//	}
//}

func (f *StraightPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data) // 务必注意这里的 =, 不是 :=
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
