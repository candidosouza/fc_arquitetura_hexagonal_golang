package application

type ProductInterface interface {
	IsValid() (boll, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLE = "enable"
)

type Product struct {
	ID string
	Name string
	Price float64
	Status string
}

fuc (p *Product) IsValid() (bool, error) {

}

fuc (p *Product) Enable() error {
	
}

fuc (p *Product) Disable() error {
	
}

fuc (p *Product) GetID() string {
	
}

fuc (p *Product) GetName() string {
	
}

fuc (p *Product) GetStatus() string {
	
}

fuc (p *Product) GetPrice() float64 {
	
}