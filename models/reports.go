package models

import (
    "time"
    )

type QA struct {
    Question string `json:"question" required`
    Answers []string `json:"answers" required`
}

type Report struct {
    Id string `json:"id"`
    Date time.Time  `json:"date" required`
    Responses []QA `json:"responses" required`
}

type Reports struct {
    Reports []Report `json:"reports" required`
    Total float64 `json:"total"`
}
