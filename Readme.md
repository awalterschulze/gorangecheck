# Range checking

Benchmarking and Assembler comparing for several implementations of encode fixed 64 in golang.

Learnings:

  - The range check hint seems to require a constant
  - The unsafe is still faster than a range check hint, but by very little.

