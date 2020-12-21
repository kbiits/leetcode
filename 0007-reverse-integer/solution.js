/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
  let rev =0
  while(x<0||x>0){   // for  -ve and -ve both integer.  
   rev =  (rev*10)+(x%10);
   x = ~~(x/10);  // alternative of floor method .
  }
  return (rev>(2**31-1) || rev< -(2**31))?  0  :  rev;
};
