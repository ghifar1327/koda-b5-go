package payments

type History struct {
	list []string
}

func NewHistory() *History {
	return &History{
		list: []string{},
	}
}

func (h *History) Add(record string) {
	h.list = append(h.list, record)
}

func (h *History) GetHistory() []string {
	return h.list
}
