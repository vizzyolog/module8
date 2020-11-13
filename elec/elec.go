package elec

type Phone interface {
	Brand() string
	Model() string
	Type() string
}
type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string
}

type baseSmartphone struct {
	phBrand string
	phModel string
}

type applePhone struct {
	baseSmartphone
}

type androidPhone struct {
	baseSmartphone
}

type radioPhone struct {
	phBrand string
	phModel string

	buttonsAmount int
}

func (p *applePhone) Brand() string { return "Apple" }
func (p *applePhone) Model() string { return p.phModel }
func (p *applePhone) Type() string  { return "smartphone" }
func (p *applePhone) OS() string    { return "iOS" }

func (p *androidPhone) Brand() string { return p.phBrand }
func (p *androidPhone) Model() string { return p.phModel }
func (p *androidPhone) Type() string  { return "smartphone" }
func (p *androidPhone) OS() string    { return "Android" }

func (p *radioPhone) Brand() string { return p.phBrand }
func (p *radioPhone) Model() string { return p.phModel }
func (p *radioPhone) Type() string  { return "station" }

func (p *radioPhone) ButtonsCount() int { return p.buttonsAmount }

//NewApplePhone - make iPhone :)
func NewApplePhone(model string) *applePhone {
	iphone := new(applePhone)
	iphone.phBrand = model

	return iphone
}

//NewAndroidPhone - make androidPhone :(
func NewAndroidPhone(brand, model string) *androidPhone {
	android := new(androidPhone)
	android.phBrand = brand
	android.phModel = model

	return android
}

//NewRadioPhone - make nosmartphone
func NewRadioPhone(brand string, model string, buttonsAmount int) *radioPhone {
	radioPhone := new(radioPhone)
	radioPhone.buttonsAmount = buttonsAmount
	radioPhone.phBrand = brand
	radioPhone.phModel = model
	return radioPhone
}
