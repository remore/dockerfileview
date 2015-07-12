package main

import (
  "bufio"
  "io"
  "net/http"
  "strings"
  "errors"
)

func getDockerfile(imageId string) (io.Reader, string, error){
  url := ""
  repository := strings.Split(imageId, ":")
  if len(repository)>1 || strings.Index(imageId, "/")<0 {
    if strings.Index(imageId, "/")<0 {
      repository = append(repository, "latest")
    }
    resp, err := http.Get("https://raw.githubusercontent.com/docker-library/official-images/master/library/" + repository[0])
    if err==nil {
      scanner := bufio.NewScanner(resp.Body)
      scanner.Split(bufio.ScanLines)
      for scanner.Scan() {
        image := strings.Split(scanner.Text(), " ")
        if image[0]==repository[1]+":" {
          r := strings.NewReplacer("git://github.com", "https://raw.githubusercontent.com", "@", "/")
          url = r.Replace(image[1])
          if len(image)>2 {
            url += "/" + image[2] + "/Dockerfile"
          } else {
            url += "/Dockerfile"
          }
        }
      }
      if url == "" {
        return nil, url, errors.New("Failed to fetch " + imageId)
      }
    } else {
      return nil, url, errors.New("Failed to fetch " + imageId)
    }
  } else {
    url = "https://registry.hub.docker.com/u/" + imageId + "/dockerfile/raw"
  }
  resp, err := http.Get(url)
  if err==nil && strings.Index(resp.Header["Content-Type"][0], "text/plain")>=0 {
    return resp.Body, url, err
  } else {
    return nil, url, errors.New("Failed to fetch " + imageId)
  }
}
