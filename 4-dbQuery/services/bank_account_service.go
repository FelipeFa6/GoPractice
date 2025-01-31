package services

import (
    "dbQuery/db"
    "dbQuery/models"
)

// retrieves all accounts
func GetBankAccounts() ([]models.BankAccount, error) {
    var accounts []models.BankAccount

    db, err := db.InitConnection()
    if err != nil {
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query("SELECT id, owner, money FROM bank_account")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var account models.BankAccount
        if err := rows.Scan(&account.ID, &account.Owner, &account.Money); err != nil {
            return nil, err
        }
        accounts = append(accounts, account)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return accounts, nil
}

