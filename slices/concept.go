/*A slice is a dynamic view of an underlying array.

Unlike arrays, slices can grow/shrink.

Syntax: []T (without length).

Has length (len(slice)) and capacity (cap(slice)). */

/* Common functions:

append(slice, value) → add elements

copy(dst, src) → copy elements

slice[a:b] → sub-slice from index a to b-1*/