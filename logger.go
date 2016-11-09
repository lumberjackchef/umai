package main

import (
  "fmt"
  "log"
  "net/http"
  "time"

  "github.com/fatih/color"
)

func Logger(inner http.Handler, name string) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    l := log.New(
      new(logWriter),
      "",
      log.Lshortfile,
    )

    method := color.New(color.FgCyan).SprintFunc()

    inner.ServeHTTP(w, r)

    l.Printf(
      "%s\t%s\t",
      method(r.Method),
      r.RequestURI,
    )
  })
}

type logWriter struct {}

func (writer logWriter) Write(bytes []byte) (int, error) {
  timey := color.New(color.FgHiMagenta).SprintFunc()

  return fmt.Print(timey(time.Now().UTC().Format("2006-01-02 15:04:05 MST ")) + "| " + string(bytes))
}
