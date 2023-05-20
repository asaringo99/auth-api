package domain

type Email struct {
	value string
}

func (d *Email) ToValue() string {
	return d.value
}

func NewEmail(value string) Email {
	return Email{
		value: value,
	}
}
