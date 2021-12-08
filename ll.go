package ll

import "sync"

// Parallel For loop with panic catching/forwarding and built-in waiting. 
// Ex: ll.For(len(fooList), func(i int){ fmt.Println(fooList[i]) })
// Note, it's most often coupled with an mx-protected output.
func For(i int, f func(i int)) {
  wg := sync.WaitGroup{}
  wg.Add(i)
  var mx sync.Mutex
  var toPanic interface{}
  for v := 0; v < i; v++ {
    go func(v int){
      defer wg.Done()
      defer func(){
        if r := recover(); r != nil {
          mx.Lock()
          toPanic = r
          mx.Unlock()
        }
      }()
      f(v)
    }(v)
  }
  wg.Wait()
  if toPanic != nil {
    panic(toPanic)
  }
}
