package main

import (
    "github.com/ivanmrnn/uadmin_sample/test/models"
    "github.com/uadmin/uadmin"
)

func main() {
    // Register Models
    uadmin.Register(
        models.Player{},
        models.Team{},
    )

    

    uadmin.StartServer()
}