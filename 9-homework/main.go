package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Age        int
	Field      string
	Gender     string
	IsEmployee bool
}

func main() {
	// Database connection string
	dsn := "host=localhost user=postgres password=1111 dbname=demo port=5432 sslmode=disable"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto migrate the schema
	db.AutoMigrate(&Product{}, &User{})

	// 1. Create a product
	db.Create(&Product{Code: "P001", Price: 1000})

	// 2. Create a user
	db.Create(&User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Password: "password", Age: 30, Field: "Engineering", Gender: "Male", IsEmployee: true})

	// 3. Retrieve first product
	var product Product
	db.First(&product)
	fmt.Println("First product:", product)

	// 4. Retrieve first user
	var user User
	db.First(&user)
	fmt.Println("First user:", user)

	// 5. Find product by primary key
	db.First(&product, 1)
	fmt.Println("Product by ID 1:", product)

	// 6. Find user by primary key
	db.First(&user, 1)
	fmt.Println("User by ID 1:", user)

	// 7. Find products by code
	db.Where("code = ?", "P001").Find(&product)
	fmt.Println("Product by code P001:", product)

	// 8. Find users by email
	db.Where("email = ?", "john.doe@example.com").Find(&user)
	fmt.Println("User by email john.doe@example.com:", user)

	// 9. Update product price
	db.Model(&product).Update("Price", 1500)
	fmt.Println("Updated product price:", product)

	// 10. Update user field
	db.Model(&user).Update("Field", "Marketing")
	fmt.Println("Updated user field:", user)

	// 11. Update multiple fields
	db.Model(&user).Updates(User{Age: 31, IsEmployee: false})
	fmt.Println("Updated multiple user fields:", user)

	// 12. Delete product
	db.Delete(&product, 1)
	fmt.Println("Deleted product by ID 1")

	// 13. Delete user
	db.Delete(&user, 1)
	fmt.Println("Deleted user by ID 1")

	// 14. Bulk insert products
	products := []Product{
		{Code: "P002", Price: 2000},
		{Code: "P003", Price: 3000},
	}
	db.Create(&products)
	fmt.Println("Bulk inserted products")

	// 15. Bulk insert users
	users := []User{
		{FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com", Password: "password", Age: 25, Field: "Design", Gender: "Female", IsEmployee: true},
		{FirstName: "Bob", LastName: "Smith", Email: "bob.smith@example.com", Password: "password", Age: 28, Field: "Engineering", Gender: "Male", IsEmployee: false},
	}
	db.Create(&users)
	fmt.Println("Bulk inserted users")

	// 16. Find all products
	var allProducts []Product
	db.Find(&allProducts)
	fmt.Println("All products:", allProducts)

	// 17. Find all users
	var allUsers []User
	db.Find(&allUsers)
	fmt.Println("All users:", allUsers)

	// 18. Count products
	var productCount int64
	db.Model(&Product{}).Count(&productCount)
	fmt.Println("Number of products:", productCount)

	// 19. Count users
	var userCount int64
	db.Model(&User{}).Count(&userCount)
	fmt.Println("Number of users:", userCount)

	// 20. Find products with limit and offset
	var limitedProducts []Product
	db.Limit(1).Offset(1).Find(&limitedProducts)
	fmt.Println("Limited products:", limitedProducts)

	// 21. Find users with limit and offset
	var limitedUsers []User
	db.Limit(1).Offset(1).Find(&limitedUsers)
	fmt.Println("Limited users:", limitedUsers)

	// 22. Order products by price
	var orderedProducts []Product
	db.Order("price desc").Find(&orderedProducts)
	fmt.Println("Ordered products by price:", orderedProducts)

	// 23. Order users by age
	var orderedUsers []User
	db.Order("age desc").Find(&orderedUsers)
	fmt.Println("Ordered users by age:", orderedUsers)

	// 24. Find users with condition
	var specificUsers []User
	db.Where("age > ? AND gender = ?", 25, "Female").Find(&specificUsers)
	fmt.Println("Specific users:", specificUsers)

	// 25. Update using SQL expression
	db.Model(&User{}).Where("email = ?", "jane.doe@example.com").Update("age", gorm.Expr("age + ?", 1))
	db.Where("email = ?", "jane.doe@example.com").First(&user)
	fmt.Println("Updated age using expression:", user)
}
