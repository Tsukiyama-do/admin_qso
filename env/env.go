package env

import (
  "log"
  "os"
)

var smode bool = true

func S_mode() bool {
   return smode
}

func S_mset(b bool) {
    smode = b
    log.Printf("Mode is changed to %t. \n", b)
}

func S_host() string {
  if s_hostname, _ := os.Hostname(); s_hostname == "yuichi-linux" {
    return "localhost:50005"
  } else {
    return "www.jj1pow.com:50005"
  }
}
