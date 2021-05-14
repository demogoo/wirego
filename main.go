package main

func main() {
	server := CreateServer()
	server.SetupRoutes()
	server.Engine.Run(server.Port)
}
