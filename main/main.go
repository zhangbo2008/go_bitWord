package main
import (


  bolt "bbolt"
)
func main() {
  db, err := bolt.Open("my.db", 0600, nil)
  print(db)
  if err != nil {
print(11111)
  }
  defer db.Close()
  println("1212121")
  const wordSize = 32 << (^uint(0) >> 63)
  println(wordSize)
  println((^uint(0) >> 63))// unit(0)是全0的,
  println(^(uint(0)))
}
