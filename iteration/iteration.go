package iteration


func Repeat(char string, repeatedCount int) string {
  var repeated string
  for i:=0; i<repeatedCount; i++ {
    repeated+=char
  }
  return repeated
}
