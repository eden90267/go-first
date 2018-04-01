package main

import (
  "fmt"
  "strings"
)

func getUserListSQL(username, email string, sexy int) string {
  sql := "select * from user"
  where := []string{}

  if username != "" {
    where = append(where, fmt.Sprintf("username = '%s'", username))
  }
  if email != "" {
    where = append(where, fmt.Sprintf("email = '%s'", email))
  }
  if sexy != 0 {
    where = append(where, fmt.Sprintf("sexy = '%d'", sexy))
  }

  return sql + " where " + strings.Join(where, " or ")
}

type searchOpts struct {
  username string
  email string
  sexy int
}

func getUserListOptsSQL(opts searchOpts) string {
  sql := "select * from user"
  where := []string{}

  if opts.username != "" {
    where = append(where, fmt.Sprintf("username = '%s'", opts.username))
  }
  if opts.email != "" {
    where = append(where, fmt.Sprintf("email = '%s'", opts.email))
  }
  if opts.sexy != 0 {
    where = append(where, fmt.Sprintf("sexy = '%d'", opts.sexy))
  }

  return sql + " where " + strings.Join(where, " or ")
}

func main() {
  fmt.Println(getUserListSQL("eden90267", "", 0))
  fmt.Println(getUserListSQL("eden90267", "eden90267@gmail.com", 1))

  fmt.Println(getUserListOptsSQL(searchOpts{
    username: "eden90267",
    email: "eden90267@gmail.com",
  }))

  fmt.Println(getUserListOptsSQL(searchOpts{
    sexy: 2,
  }))
}
