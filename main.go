package main



import "github.com/jorkle/trapalert/fsnotify"



func main() {

  watcher, err := fsnotify.NewWatcher()
  if err != nil {
    log.Fatal(err)
  }
  defer watcher.Close()

  go func() {
    for {
      select {
      case event, ok := <-watcher.Events:event
        if !ok {
          return
        }
        fmt.Println("event:", event)
        if event.Has(fsnotify.Read)
      }
    }
  }
}
