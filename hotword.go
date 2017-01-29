package snowboy

type Hotword struct {
	Name        string
	Path        string
	Sensitivity string
}

func NewHotwords() *Hotwords {
	return &Hotwords{Words: []Hotword{}}
}

type Hotwords struct {
	Words []Hotword
}

func (s *Hotwords) Add(name, path, sensitivity string) {
	s.Words = append(s.Words, Hotword{Name: name, Path: path, Sensitivity: sensitivity})
}

func (s *Hotwords) ModelStr() string {
	var result string
	for _, hw := range s.Words {
		result = result + "," + hw.Path
	}
	return result
}

func (s *Hotwords) SensitivityStr() string {
	var result string
	for _, hw := range s.Words {
		result = result + "," + hw.Sensitivity
	}
	return result
}

func (s *Hotwords) GetNameByIndex(index int) string {
	if index > len(s.Words) || index < 0 {
		return ""
	}
	return s.Words[index].Name
}
