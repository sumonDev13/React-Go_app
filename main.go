package main
import ("fmt"
		"log"
		"github.com/gofiber/fiber/v2"
		
)

func main() {
	fmt.Println("Hello Go Application");

	var firstName string = "Sumon";
	const secondName string = "Kumar";
	thirdName := "Mondal";

	fmt.Println(firstName);
	fmt.Println(secondName);
	fmt.Println(thirdName);

	app := fiber.New()
	log.Fatal(app.Listen(":4000"));
}