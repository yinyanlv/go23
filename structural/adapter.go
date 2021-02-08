package structural

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecialRequest() string
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

type adapteeImpl struct{}

func (a *adapteeImpl) SpecialRequest() string {
	return "adaptee special request"
}

type Adapter struct {
	Adaptee
}

func (a *Adapter) Request() string {
	return a.Adaptee.SpecialRequest()
}

func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{
		Adaptee: adaptee,
	}
}
