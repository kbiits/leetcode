func finalValueAfterOperations(operations []string) int {
   x := 0
   for _, v := range operations {
       if v[len(v) - 1] == '+' || v[0] == '+' {
           x++
       } else {
           x--
       }
   }

   return x
}
