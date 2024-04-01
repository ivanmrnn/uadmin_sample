package main

import (
    "github.com/ivanmrnn/uadmin_sample/test/models"
    "github.com/uadmin/uadmin"
)

func main() {
    // Register Models
    uadmin.Register(
        models.Todo{},
        models.Category{},
        models.Friend{},
        models.Item{},
    )

    // Register Inlines
    uadmin.RegisterInlines(models.Category{}, map[string]string{
        "Todo": "CategoryID",
    })
    uadmin.RegisterInlines(models.Friend{}, map[string]string{
        "Todo": "FriendID",
    })
    uadmin.RegisterInlines(models.Item{}, map[string]string{
        "Todo": "ItemID",
    })

    // Call InitializeRootURL function to change the RootURL value in the Settings model.
    InitializeRootURL()

    // API Handler
    http.HandleFunc("/api/", uadmin.Handler(api.Handler))

    uadmin.StartServer()
}