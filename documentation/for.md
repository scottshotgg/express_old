# The For Loop

`for` loops in Express are a powerful tool that encompass a multitude of features in other languages. Below you will be shown all of the ways in which a `for` loop can be employed, and an explanation given of the structure, along with a comparison to other languages.

1. The common `for` loop:
  ```
  for (int i = 0; i < 10; i++) {

  }
  ```
Above we see a common `C`-type for loop being used with a standard incrementer variable, groupings, and all of the grammatical markup that you would expect to see. However, this is Express, and we know we don't _need_ to be so explicit:

2. The standard `for` loop
  ```
  for i..10 {

  }
  ```
  A _standard_ `for` loop in Express; this allows a much more consise and idiomatic method of expression without being overtly verbose. Inside the for loop, the current iteration can be accessed using the i variable, which can be changed to a 0 if not needed or if the intention is explicitly from 0 to 10. A variable declaration may also be substituted in the beginning as in the following example.

3. Iterable declararion `for` loop
  ```
  for i:0..10 {
    
  }
  ```
  Similar to the previous for loop, it is functionally the same as declaring or `set`ting `i` to 0 before the loop within the previous example, however, you will notice that within the syntax is actually a different meaning. Here `i` is being declared as an `iterable` variable. Building on this bit of knowledge, there are other implications with iterables and channels.

4. The `for`-ever loop
``` for {

}
```
Inspired by `Go`'s `for`-ever loop, this is functionally the same a a `while` loop discussed in the previous chapter.

5. `while` loop composition
```
for i != true {
  
}
```

Now that we have explored the options in Express of loop expression, below will be other examples of `for` loops within Express:

```
for i:=0, i<10, i+=2 {

}
```
```
for i:0, 10 {

}
```
```
for (){
    int i = 0
    i < 10
    i++ 
    {
      print(i)
    }
}
```