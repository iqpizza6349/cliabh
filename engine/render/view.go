package render

// Observer 인터페이스 정의
type Observer interface {
	Update() //Model 이 변경될 때 호출되는 함수
}

// Model 정의
type Model struct {
	LabelText string
	observers []Observer // 등록된 옵저버 목록
}

// NewModel 은 Model 를 생성합니다.
func NewModel() *Model {
	return &Model{
		LabelText: "Initial Text",
		observers: []Observer{},
	}
}

// RegisterObserver 는 옵저버를 등록합니다.
func (m *Model) RegisterObserver(observer Observer) {
	m.observers = append(m.observers, observer)
}

// NotifyObservers 는 모든 등록된 옵저버들에게 변경 사항을 알립니다.
func (m *Model) NotifyObservers() {
	for _, observer := range m.observers {
		observer.Update()
	}
}

// SetLabelText 는 라벨의 텍스트를 변경하고 변경 사항을 옵저버들에게 알립니다.
func (m *Model) SetLabelText(newText string) {
	m.LabelText = newText
	m.NotifyObservers() // 변경 사항을 알림
}
