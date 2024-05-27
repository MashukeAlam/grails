package internals

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/handlers"
)

func SetupRoutes(app *fiber.App, dbGorm *gorm.DB) {
	// Route to render the Slim template
	app.Get("/", func(c *fiber.Ctx) error {
		// Pass the title to the template
		return c.Render("index", fiber.Map{
			"Title": "Hello, Fiber with Slim!",
		}, "layouts/main")
	})

	// Create a /game endpoint
	game := app.Group("/game")
	game.Get("/", handlers.GetGames(dbGorm))
	game.Get("/insert", handlers.InsertGame())
	game.Post("/", handlers.CreateGame(dbGorm))
	game.Get("/:id/edit", handlers.EditGame(dbGorm))
	game.Get("/:id/delete", handlers.DeleteGame(dbGorm))
	game.Post("/:id", handlers.UpdateGame(dbGorm))
	game.Delete("/:id", handlers.DestroyGame(dbGorm))

	// Car routes
	Car := app.Group("/Car")
	Car.Get("/", handlers.GetCars(dbGorm))
	Car.Get("/insert", handlers.InsertCar())
	Car.Post("/", handlers.CreateCar(dbGorm))
	Car.Get("/:id/edit", handlers.EditCar(dbGorm))
	Car.Post("/:id", handlers.UpdateCar(dbGorm))
	Car.Get("/:id/delete", handlers.DeleteCar(dbGorm))
	Car.Delete("/:id", handlers.DeleteCar(dbGorm))

	// Bike routes
	Bike := app.Group("/Bike")
	Bike.Get("/", handlers.GetBikes(dbGorm))
	Bike.Get("/insert", handlers.InsertBike())
	Bike.Post("/", handlers.CreateBike(dbGorm))
	Bike.Get("/:id/edit", handlers.EditBike(dbGorm))
	Bike.Put("/:id", handlers.UpdateBike(dbGorm))
	Bike.Get("/:id/delete", handlers.DeleteBike(dbGorm))
	Bike.Delete("/:id", handlers.DestroyBike(dbGorm))

	// Animal routes
	Animal := app.Group("/Animal")
	Animal.Get("/", handlers.GetAnimals(dbGorm))
	Animal.Get("/insert", handlers.InsertAnimal())
	Animal.Post("/", handlers.CreateAnimal(dbGorm))
	Animal.Get("/:id/edit", handlers.EditAnimal(dbGorm))
	Animal.Put("/:id", handlers.UpdateAnimal(dbGorm))
	Animal.Get("/:id/delete", handlers.DeleteAnimal(dbGorm))
	Animal.Delete("/:id", handlers.DestroyAnimal(dbGorm))

	// Human routes
	Human := app.Group("/Human")
	Human.Get("/", handlers.GetHumans(dbGorm))
	Human.Get("/insert", handlers.InsertHuman())
	Human.Post("/", handlers.CreateHuman(dbGorm))
	Human.Get("/:id/edit", handlers.EditHuman(dbGorm))
	Human.Put("/:id", handlers.UpdateHuman(dbGorm))
	Human.Get("/:id/delete", handlers.DeleteHuman(dbGorm))
	Human.Delete("/:id", handlers.DestroyHuman(dbGorm))

	// Language routes
	Language := app.Group("/Language")
	Language.Get("/", handlers.GetLanguages(dbGorm))
	Language.Get("/insert", handlers.InsertLanguage())
	Language.Post("/", handlers.CreateLanguage(dbGorm))
	Language.Get("/:id/edit", handlers.EditLanguage(dbGorm))
	Language.Put("/:id", handlers.UpdateLanguage(dbGorm))
	Language.Get("/:id/delete", handlers.DeleteLanguage(dbGorm))
	Language.Delete("/:id", handlers.DestroyLanguage(dbGorm))

	// Bird routes
	Bird := app.Group("/Bird")
	Bird.Get("/", handlers.GetBirds(dbGorm))
	Bird.Get("/insert", handlers.InsertBird())
	Bird.Post("/", handlers.CreateBird(dbGorm))
	Bird.Get("/:id/edit", handlers.EditBird(dbGorm))
	Bird.Put("/:id", handlers.UpdateBird(dbGorm))
	Bird.Get("/:id/delete", handlers.DeleteBird(dbGorm))
	Bird.Delete("/:id", handlers.DestroyBird(dbGorm))

	// Girl routes
	Girl := app.Group("/Girls")
	Girl.Get("/", handlers.GetGirls(dbGorm))
	Girl.Get("/insert", handlers.InsertGirl())
	Girl.Post("/", handlers.CreateGirl(dbGorm))
	Girl.Get("/:id/edit", handlers.EditGirl(dbGorm))
	Girl.Put("/:id", handlers.UpdateGirl(dbGorm))
	Girl.Get("/:id/delete", handlers.DeleteGirl(dbGorm))
	Girl.Delete("/:id", handlers.DestroyGirl(dbGorm))

	// Laptop routes
	Laptop := app.Group("/Laptops")
	Laptop.Get("/", handlers.GetLaptops(dbGorm))
	Laptop.Get("/insert", handlers.InsertLaptop())
	Laptop.Post("/", handlers.CreateLaptop(dbGorm))
	Laptop.Get("/:id/edit", handlers.EditLaptop(dbGorm))
	Laptop.Put("/:id", handlers.UpdateLaptop(dbGorm))
	Laptop.Get("/:id/delete", handlers.DeleteLaptop(dbGorm))
	Laptop.Delete("/:id", handlers.DestroyLaptop(dbGorm))

	// Electronic routes
	Electronic := app.Group("/Electronic")
	Electronic.Get("/", handlers.GetElectronics(dbGorm))
	Electronic.Get("/insert", handlers.InsertElectronic())
	Electronic.Post("/", handlers.CreateElectronic(dbGorm))
	Electronic.Get("/:id/edit", handlers.EditElectronic(dbGorm))
	Electronic.Put("/:id", handlers.UpdateElectronic(dbGorm))
	Electronic.Get("/:id/delete", handlers.DeleteElectronic(dbGorm))
	Electronic.Delete("/:id", handlers.DestroyElectronic(dbGorm))

	// Mobile routes
	Mobile := app.Group("/Mobiles")
	Mobile.Get("/", handlers.GetMobiles(dbGorm))
	Mobile.Get("/insert", handlers.InsertMobile())
	Mobile.Post("/", handlers.CreateMobile(dbGorm))
	Mobile.Get("/:id/edit", handlers.EditMobile(dbGorm))
	Mobile.Put("/:id", handlers.UpdateMobile(dbGorm))
	Mobile.Get("/:id/delete", handlers.DeleteMobile(dbGorm))
	Mobile.Delete("/:id", handlers.DestroyMobile(dbGorm))

	// Smartphone routes
	Smartphone := app.Group("/Smartphones")
	Smartphone.Get("/", handlers.GetSmartphones(dbGorm))
	Smartphone.Get("/insert", handlers.InsertSmartphone())
	Smartphone.Post("/", handlers.CreateSmartphone(dbGorm))
	Smartphone.Get("/:id/edit", handlers.EditSmartphone(dbGorm))
	Smartphone.Put("/:id", handlers.UpdateSmartphone(dbGorm))
	Smartphone.Get("/:id/delete", handlers.DeleteSmartphone(dbGorm))
	Smartphone.Delete("/:id", handlers.DestroySmartphone(dbGorm))

	// Monitor routes
	Monitor := app.Group("/Monitors")
	Monitor.Get("/", handlers.GetMonitors(dbGorm))
	Monitor.Get("/insert", handlers.InsertMonitor())
	Monitor.Post("/", handlers.CreateMonitor(dbGorm))
	Monitor.Get("/:id/edit", handlers.EditMonitor(dbGorm))
	Monitor.Put("/:id", handlers.UpdateMonitor(dbGorm))
	Monitor.Get("/:id/delete", handlers.DeleteMonitor(dbGorm))
	Monitor.Delete("/:id", handlers.DestroyMonitor(dbGorm))

	// Pen routes
	Pen := app.Group("/Pens")
	Pen.Get("/", handlers.GetPens(dbGorm))
	Pen.Get("/insert", handlers.InsertPen())
	Pen.Post("/", handlers.CreatePen(dbGorm))
	Pen.Get("/:id/edit", handlers.EditPen(dbGorm))
	Pen.Put("/:id", handlers.UpdatePen(dbGorm))
	Pen.Get("/:id/delete", handlers.DeletePen(dbGorm))
	Pen.Delete("/:id", handlers.DestroyPen(dbGorm))

	// Tissue routes
	Tissue := app.Group("/Tissues")
	Tissue.Get("/", handlers.GetTissues(dbGorm))
	Tissue.Get("/insert", handlers.InsertTissue())
	Tissue.Post("/", handlers.CreateTissue(dbGorm))
	Tissue.Get("/:id/edit", handlers.EditTissue(dbGorm))
	Tissue.Put("/:id", handlers.UpdateTissue(dbGorm))
	Tissue.Get("/:id/delete", handlers.DeleteTissue(dbGorm))
	Tissue.Delete("/:id", handlers.DestroyTissue(dbGorm))

	// Perfume routes
	Perfume := app.Group("/Perfumes")
	Perfume.Get("/", handlers.GetPerfumes(dbGorm))
	Perfume.Get("/insert", handlers.InsertPerfume())
	Perfume.Post("/", handlers.CreatePerfume(dbGorm))
	Perfume.Get("/:id/edit", handlers.EditPerfume(dbGorm))
	Perfume.Put("/:id", handlers.UpdatePerfume(dbGorm))
	Perfume.Get("/:id/delete", handlers.DeletePerfume(dbGorm))
	Perfume.Delete("/:id", handlers.DestroyPerfume(dbGorm))

	// Card routes
	Card := app.Group("/Cards")
	Card.Get("/", handlers.GetCards(dbGorm))
	Card.Get("/insert", handlers.InsertCard())
	Card.Post("/", handlers.CreateCard(dbGorm))
	Card.Get("/:id/edit", handlers.EditCard(dbGorm))
	Card.Put("/:id", handlers.UpdateCard(dbGorm))
	Card.Get("/:id/delete", handlers.DeleteCard(dbGorm))
	Card.Delete("/:id", handlers.DestroyCard(dbGorm))

	// Dev routes
	Dev := app.Group("/dev")
	Dev.Get("/", handlers.GetDevView())
	Dev.Post("/", handlers.ProcessIncomingScaffoldData(dbGorm))

}
