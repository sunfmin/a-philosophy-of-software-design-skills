# A Philosophy of Software Design — skills

Claude Code skills distilled from John Ousterhout's *A Philosophy of Software Design* (2nd ed.).

| Skill | Use it to |
|---|---|
| `complexity-review` | Review code or a diff against the book's red flags |
| `design-it-twice` | Compare 2–3 radically different interfaces before implementing |
| `define-errors-out` | Shrink error handling: define away, mask, aggregate, or crash |
| `comments-first` | Write interface comments first as a design probe; audit comments & names |
| `strategic-change` | Modify existing code so it looks designed-in, not patched on; check the diff for design debt |
| `split-or-join` | Decide whether code belongs together or apart; fix pass-through layers |

```bash
npx skills add sunfmin/a-philosophy-of-software-design-skills -g
```
