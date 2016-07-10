package main

import "github.com/kataras/iris"

type UserAPI struct {
	*iris.Context
}

// Get /users
func (u UserAPI) Get() {
	u.Write("Get from /users")
	// u.JSON(iris.StatusOK, myDb.AllUsers())
}

// Get /:param1 which its value passed to the id argument
func (u UserAPI) GetBy(id string) { // id equals to u.Param("param1")
	u.Write("Get from /users/%s", id)
	// u.JSON(iris.StatusOK, myDb.GetUserById(id))
}

// PUT /users
func (u UserAPI) Put() {
	name := u.FormValue("name") // you can still use the whole context's feature!
	// myDb.InsertUser(...)
	println(string(name))
	println("Put from /users")
}

// POST /users/:param1
func (u UserAPI) PostBy(id string) {
	name := u.FormValue("name") // you can still use the whole context's feature!
	// myDb.UpdateUser(...)
	println(string(name))
	println("Post from /users/" + id)
}

// DELETE /users/:param1
func (u UserAPI) DeleteBy(id string) {
	// myDb.DeleteUser(id)
	println("Delete from /" + id)
}

func main() {
	api := iris.New()
	api.API("/users", UserAPI{})
	api.Listen(":8080")
}

func hi(ctx *iris.Context) {
	ctx.Render("hi.html", map[string]interface{}{"Name": "natali"})
}
