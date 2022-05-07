package dict

type Dictionary interface {
	Name() string
	Transform(word, homeLanguage, targetLanguage string) ([]string, error)
}

type Source struct {
	list []Dictionary
}

func NewSource(list ...Dictionary) *Source {
	return &Source{
		list: list,
	}
}

func (s *Source) List() []Dictionary {
	return s.list
}

func (s *Source) Add(dict Dictionary) {
	s.list = append(s.list, dict)
}
