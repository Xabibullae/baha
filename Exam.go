package main

import (
	"errors"
	"fmt"
)

// Structlar e'lon qilish
type Car struct {
	ID     int
	Brand  string
	Rented bool
}

type Bike struct {
	ID     int
	Brand  string
	Rented bool
}

type Skate struct {
	ID     int
	Brand  string
	Rented bool
}

// Vehicle interfeysi
type Vehicle interface {
	Rent() error
	Return() error
	Details() (string, error)
}

// Car metodlari
func (c *Car) Rent() error {
	if c.Rented {
		return errors.New("Ushbu mashina allaqachon ijaraga olingan")
	}
	c.Rented = true
	return nil
}

func (c *Car) Return() error {
	if !c.Rented {
		return errors.New("Bu mashina ijaraga olinmagan")
	}
	c.Rented = false
	return nil
}

func (c *Car) Details() (string, error) {
	if c.ID == 0 {
		return "", errors.New("Mashina haqida ma'lumot mavjud emas")
	}
	return fmt.Sprintf("Mashina - ID: %d, Brend: %s, Ijarada: %v", c.ID, c.Brand, c.Rented), nil
}

// Bike metodlari
func (b *Bike) Rent() error {
	if b.Rented {
		return errors.New("Ushbu velosiped allaqachon ijaraga olingan")
	}
	b.Rented = true
	return nil
}

func (b *Bike) Return() error {
	if !b.Rented {
		return errors.New("Bu velosiped ijaraga olinmagan")
	}
	b.Rented = false
	return nil
}

func (b *Bike) Details() (string, error) {
	if b.ID == 0 {
		return "", errors.New("Velosiped haqida ma'lumot mavjud emas")
	}
	return fmt.Sprintf("Velosiped - ID: %d, Brend: %s, Ijarada: %v", b.ID, b.Brand, b.Rented), nil
}

// Skate metodlari
func (s *Skate) Rent() error {
	if s.Rented {
		return errors.New("Ushbu skeyt allaqachon ijaraga olingan")
	}
	s.Rented = true
	return nil
}

func (s *Skate) Return() error {
	if !s.Rented {
		return errors.New("Bu skeyt ijaraga olinmagan")
	}
	s.Rented = false
	return nil
}

func (s *Skate) Details() (string, error) {
	if s.ID == 0 {
		return "", errors.New("Skeyt haqida ma'lumot mavjud emas")
	}
	return fmt.Sprintf("Skeyt - ID: %d, Brend: %s, Ijarada: %v", s.ID, s.Brand, s.Rented), nil
}

// Asosiy funksiya
func main() {
	var vehicles []Vehicle

	for {
		fmt.Println("Harakatni tanlang: \n1)Vehicle qo'shish\n2)ijara\n3)qaytar\n4)ko'rish\n5)chiq ")
		var x int
		fmt.Scanln(&x)

		switch x {
		case 1:
			fmt.Println("Transport turini kiriting \n1)mashina\n2)velosiped\n3)skeyt ")
			var vehicleType int
			fmt.Scanln(&vehicleType)

			fmt.Println("ID raqamini kiriting: ")
			var id int
			fmt.Scanln(&id)

			fmt.Println("Brend nomini kiriting: ")
			var brand string
			fmt.Scanln(&brand)

			switch vehicleType {
			case 1:
				vehicles = append(vehicles, &Car{ID: id, Brand: brand})
			case 2:
				vehicles = append(vehicles, &Bike{ID: id, Brand: brand})
			case 3:
				vehicles = append(vehicles, &Skate{ID: id, Brand: brand})
			default:
				fmt.Println("Noto'g'ri transport turi!")
			}

		case 2:
			fmt.Println("Ijaraga olish uchun ID raqamini kiriting: ")
			var id int
			fmt.Scanln(&id)

			found := false
			for _, v := range vehicles {
				if details, err := v.Details(); err == nil && v.Rent() == nil {
					fmt.Println(details, "- Muvaffaqiyatli ijaraga olindi")
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Transport topilmadi yoki allaqachon ijaraga olingan!")
			}

		case 3:
			fmt.Println("Qaytarish uchun ID raqamini kiriting: ")
			var id int
			fmt.Scanln(&id)

			found := false
			for _, v := range vehicles {
				if details, err := v.Details(); err == nil && v.Return() == nil {
					fmt.Println(details, "- Muvaffaqiyatli qaytarildi")
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Transport topilmadi yoki ijaraga olinmagan!")
			}

		case 4:
			for _, v := range vehicles {
				if details, err := v.Details(); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(details)
				}
			}
		case 5:
			fmt.Println("Dasturdan chiqmoqda...")
			return

		default:
			fmt.Println("Noto'g'ri harakat!")
		}
	}
}
