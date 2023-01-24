/*
Copyright © 2023 Loa212 <anobileloris@gmail.com>
*/
package main

import (
	"github.com/Loa212/todo2/cmd"
	"github.com/Loa212/todo2/initializers"
)

func main() {

	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDatabase()

	println("Hold on, I'm looking for the db...")

	cmd.Execute()
}
