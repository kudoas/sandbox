# README

## Tailwind class sorting for ERB

This project does not use `herb-format` for ERB rewriting.
It runs only the Tailwind class sorter against HTML-ish ERB templates via `bun`.

```sh
bun run sort:tailwind
bun run sort:tailwind:check
```

Defaults:

- Scans `app/views`
- Targets `*.html.erb`, `*.html+*.erb`, `*.turbo_stream.erb`, and `*.rhtml`
- Leaves non-HTML ERB files such as `*.json.erb` untouched
- Rewrites static quoted `class=` and `class:` values only

You can also pass files or directories explicitly:

```sh
bun scripts/sort-tailwind-classes.mjs app/views/posts
bun scripts/sort-tailwind-classes.mjs app/views/layouts/application.html.erb --check
```
