package abstractfactory

type ISportsFactory interface {
	MakeShort() IShort
	MakeShoe() IShoe
}

func GetSportsFactory(brand string) ISportsFactory {
	switch brand {
	case "adidas":
		return &Adidas{}
	case "nike":
		return &Nike{}
	}
	return nil
}
