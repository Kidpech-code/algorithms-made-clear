# Algorithms, Made Clear

**Sorting algorithms, explained through everyday stories, pictures, diagrams, and runnable Go.**

Books of different heights, a pile of receipts, and numbered pickup tickets can all be put in order. There is more than one way to do it: compare neighbors, sort smaller groups and combine them, or split items around a chosen reference number. Those are the ideas behind Bubble Sort, Merge Sort, and Quick Sort.

![Three sorting methods illustrated as neighbor swaps, split and merge, and partitioning around a pivot](assets/cover.svg)

This first chapter covers **only three sorting algorithms**. Start with the story and picture, then work through the classroom-style calculations and run the Go code. You can understand the idea before reading the equations.

## Sorting in one minute

Sorting puts items in a chosen order. Here, the rule is **smaller number first**:

```text
Before: 5  1  4  2  8  3
After:  1  2  3  4  5  8
```

For real items, decide what the number means: a book's height, a receipt's total, or a pickup ticket's number. The sorting algorithm follows that rule; it does not decide the rule for you.

| Algorithm | Picture to keep in mind | Read the visual guide |
| --- | --- | --- |
| **Bubble Sort** | Compare neighboring books by height; swap those in the wrong order. | [Bubble Sort](guides/bubble-sort.md) |
| **Merge Sort** | Split a pile of receipt totals, then combine sorted piles in order. | [Merge Sort](guides/merge-sort.md) |
| **Quick Sort** | Put pickup tickets on either side of one reference ticket; repeat within each side. | [Quick Sort](guides/quick-sort.md) |

```mermaid
flowchart LR
  U["5 · 1 · 4 · 2 · 8 · 3"] --> B["Bubble: compare neighbors"]
  U --> M["Merge: split and combine"]
  U --> Q["Quick: partition around a pivot"]
  B --> S["1 · 2 · 3 · 4 · 5 · 8"]
  M --> S
  Q --> S
```

## Work through it like a classroom problem

All three guides use the same input. Number the positions from **0**, as Go does:

| Position `i` | 0 | 1 | 2 | 3 | 4 | 5 |
| --- | --- | --- | --- | --- | --- | --- |
| Value `A[i]` | 5 | 1 | 4 | 2 | 8 | 3 |

Read each guide in four ways: **picture** (what moves), **trace** (the row after each step), **pseudocode and Go** (how to repeat it), and **math** (why it works and how much work it does). Pause before each row and predict the next result. For example, Bubble Sort first compares `A[0] = 5` with `A[1] = 1`; because `5 > 1`, it swaps them and the row becomes `[1, 5, 4, 2, 8, 3]`.

In the calculations, `n` is the number of values, `A[i]` is the value at position `i`, and `T(n)` means the work needed to sort `n` values. The final answer is always `[1, 2, 3, 4, 5, 8]`; the steps and costs differ.

## Which one should I choose?

| Algorithm | Best | Typical | Worst | Extra memory | Keeps equal items in their original order? |
| --- | --- | --- | --- | --- | --- |
| Bubble Sort | O(n) | O(n²) | O(n²) | O(1) | Yes |
| Merge Sort | O(n log n) | O(n log n) | O(n log n) | O(n) | Yes |
| Quick Sort here | O(n log n) | O(n log n) | O(n²) | O(log n) typical; O(n) worst, for recursion | No |

- **Learning the idea?** Start with Bubble Sort: every comparison is easy to see. It becomes slow as the list grows.
- **Need reliable performance and the original order of ties?** Merge Sort is a good model. It needs another array while sorting.
- **Want to see how partitioning can sort without a full second array?** Study Quick Sort. Its pivot choice matters: this teaching version chooses the last value, so already sorted or all-equal input can take O(n²) time.

“Stable” means ties keep their earlier order. If two receipts both total $20 and receipt A came before receipt B, a stable sort leaves A before B. “Extra memory” counts space beyond the original list; the Quick Sort row includes its recursive calls.

**For a real Go application, use the standard library** rather than these teaching implementations: [`slices.Sort`](https://pkg.go.dev/slices#Sort) for ordered values, [`slices.SortFunc`](https://pkg.go.dev/slices#SortFunc) for a custom comparison, or [`slices.SortStableFunc`](https://pkg.go.dev/slices#SortStableFunc) when ties must retain their original order. The `slices` package is part of Go's standard library starting with [Go 1.21](https://go.dev/doc/go1.21).

## Run the same example

Install Go 1.22 or newer, then run this from the repo root on **macOS, Linux, or Windows (PowerShell)**:

```sh
go run ./examples/sorting
```

```text
Before: [5 1 4 2 8 3]
Bubble: [1 2 3 4 5 8]
Merge : [1 2 3 4 5 8]
Quick : [1 2 3 4 5 8]
```

Each function changes the slice you pass to it. The example copies the starting values so that all three methods get the same input. Read the [Go implementations](sorting/sort.go) or run the checks with `go test ./...`.

## A few words you will meet

- **Pass:** one trip through the values being examined.
- **Pivot:** the reference value used to split a Quick Sort section.
- **Comparison:** checking which of two values should come first. **Swap:** exchanging their positions.
- **`log₂ n`:** roughly how many times `n` can be halved before reaching 1. For example, `8 → 4 → 2 → 1` takes three halvings, so `log₂ 8 = 3`.
- **`O(f(n))`:** an upper bound on how work grows as `n` grows, ignoring constant factors. `Θ(f(n))` means a matching upper and lower growth bound. The guides derive `O(n)`, `O(n²)`, and `O(n log n)`; these are growth descriptions, not exact seconds.
- **Stable:** items with equal sort values keep their earlier relative order.

The diagrams show the same numbers as the runnable code, so you can compare each step with the final output.

## License

[MIT](LICENSE)
