# Ideas

## .= for vector assignment: <br>
  `array_a = array_b`: In this example `array_a` is a _shallow copy_ of `array_b`. In other words, `array_a` has the same value as `array_b` because it points to the same area of memory.

  `array_a .= array_b`: In this example, we see the usage of the _vector_ assignment operator. This type of assignment will force/allow Express to create/assign `array_a` as a _deep copy_ of `array_b` by copying the memory space.