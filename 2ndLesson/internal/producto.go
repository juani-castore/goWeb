package producto


type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	Expiration  time.Time `json:"expiration"`
	IsPublished bool      `json:"is_published"`
	Price       float64   `json:"price"`
}



type Storage struct {
	Products []Product `json:"products"`
}

func (s *Storage) PrintInfo() {
	fmt.Println("Products: ", s.Products)
}

func (s *Storage) GetProductByName(name string) ([]Product, error) {
	var productos []Product
	for _, p := range s.Products {
		if p.Name == name {
			productos = append(productos, p)
		}
	}
	if len(productos) == 0 {
		return nil, errors.New("no se encontraron productos")
	}
	return productos, nil
}

func (s *Storage) GetAll(ctx context.Context) ([]Product, error) {
	user, ok := ctx.Value("user").(string)
	if !ok && user != "" {
		fmt.Println("Context value in package producto: ", user)
		return nil, errors.New("invalid user")
	}
	return s.Products, nil
}

func (s *Storage) GetProductosMayorPrecio(precio float64) []Producto {
	var resultado []Producto

	for _, producto := range s.Productos {
		if producto.Price >= precio {
			resultado = append(resultado, producto)
		}
	}

	return resultado
}