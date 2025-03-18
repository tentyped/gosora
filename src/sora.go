package sora
import (
    "fmt"
    "sora/src/utils"
)

func TestSettings() {
  settings := utils.Settings{}
  err := settings.Load()
  if err != nil {
    fmt.Println("Error loading settings: ", err)
    return
  }

  fmt.Println("test_value: ", settings.TestValue)

  settings.TestValue = "blue"
  err = settings.Save()
  if err != nil {
    fmt.Println("Error saving settings: ", err)
    return
  }
}
