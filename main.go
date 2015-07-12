package main

import (
  "os"
  "strings"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "dockerfileview"
  app.Version = "0.1.0"
  app.HideHelp = true
  app.Flags = []cli.Flag{
      cli.HelpFlag,
      cli.BoolFlag{
        Name: "text, t",
        Usage: "text output without color and syntax highlight",
      },
  }
  app.Usage = "a public Dockerfile viewer"
  app.Action = func(c *cli.Context) {
    if len(os.Args)>1 {
      colorFlag := true
      if c.Bool("text") {
        colorFlag = false
      }
      if strings.Index(os.Args[len(os.Args)-1], "Dockerfile")>=0 {
        inFile, _ := os.Open(os.Args[len(os.Args)-1])
        defer inFile.Close()
        parseDockerfile(os.Args[len(os.Args)-1], "", inFile, colorFlag)
      } else {
        parseDockerfile("(Direct Input from CLI)", "", strings.NewReader("FROM " +os.Args[len(os.Args)-1]), colorFlag)
      }
    }
  }
  app.Run(os.Args)
}
