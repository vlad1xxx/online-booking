package service

const (
	StepNone        = ""
	StepChooseDate  = "choose_date"
	StepChooseTime  = "choose_time"
	StepInputNumber = "input_number"
	StepConfirm     = "confirm"
)

type StateStore interface {
	Get(chatID int64) *UserState
	Set(chatID int64, step string)
	Update(chatID int64, value string)
	Clear(chatID int64)
}

type UserState struct {
	Step   string
	Date   string
	Time   string
	Number string
}

type MemoryState struct {
	data map[int64]*UserState
}

func NewMemoryState() *MemoryState {
	return &MemoryState{data: make(map[int64]*UserState)}
}

func (m *MemoryState) Set(chatID int64, step string) {
	m.data[chatID] = &UserState{Step: step}
}

func (m *MemoryState) Get(chatID int64) *UserState {
	state, ok := m.data[chatID]
	if !ok {
		return nil
	}
	return state
}

func (m *MemoryState) Update(chatID int64, value string) {
	if s, ok := m.data[chatID]; ok {
		switch s.Step {
		case StepChooseDate:
			s.Date = value
			s.Step = StepChooseTime
		case StepChooseTime:
			s.Time = value
			s.Step = StepInputNumber
		case StepInputNumber:
			s.Number = value
			s.Step = StepConfirm
		}
	}
}

func (m *MemoryState) Clear(chatID int64) {
	delete(m.data, chatID)
}
