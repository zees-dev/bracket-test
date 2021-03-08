# Simple bracket test

## Problem

- Given a string, determine if its brackets are properly nested.

Examples:

- Input: `"{[]()}"`
- Output: `true`

- Input: `"{[(])}"`
- Output: `false`

- Input: `"{[}"`
- Output: `false`

Approach:

- Use a stack to keep track of matching parenthesis as we iterate
through the string.

## Implementation

Implementation details in [main.go](./main.go)

Known issues:
`BracketStack` data structure is no concurrency-safe (multiple go-routines writing to this stack will cause isses).
To resolve this, a mutex should be used upon reading/writing to/from the slice data structure (stack).

## Tests

Tests in [main_test.go](./main_test.go)

```sh
GO111MODULE=auto go test -v .
```
