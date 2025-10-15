/*Step 1 — “Errors are values” (basic pattern)
Core idea (chinna summary)
Go lo chāla functions (val, error) return chestāyi.
err != nil ante error occurrindhi; nil ante ok.
Pattern: guard early (error vasthe immediate ga return/handle).
v, err := doThing()
if err != nil {
    // handle or return
}*/