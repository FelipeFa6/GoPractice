package main

import (
    "fmt"
    "log"
    "dbQuery/services"
)

func main() {
    accounts, err := services.GetBankAccounts()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Bank accounts:")
    for _, account := range accounts {
        fmt.Printf("ID: %d, Owner: %s, Money: %d\n", account.ID, account.Owner, account.Money)
    }
}

