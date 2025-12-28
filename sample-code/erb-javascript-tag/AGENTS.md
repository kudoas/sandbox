# Repository Guidelines

## Project Structure & Module Organization
- `package.json`: VS Code extension manifest (language id, grammar path, configuration).
- `language-configuration.json`: comment, bracket, and auto-closing rules for the ERB language.
- `syntaxes/erb.tmLanguage.json`: TextMate grammar that powers syntax highlighting.
- `README.md` and `CHANGELOG.md`: user docs and release notes.
- `vsc-extension-quickstart.md`: setup and local testing tips.

## Build, Test, and Development Commands
- No npm scripts are configured in `package.json`.
- Run the extension locally using VS Code:
  - Open the folder in VS Code and press `F5` to launch an Extension Development Host.
  - Reload the host window (`Cmd+R` / `Ctrl+R`) after grammar or configuration changes.
- Manual install for ad‑hoc testing: copy the repo to `~/.vscode/extensions/` and restart VS Code.
- Install dependencies after changes to `package.json`:
  - `npm install`

## Coding Style & Naming Conventions
- Keep JSON formatting consistent with each file:
  - `syntaxes/erb.tmLanguage.json` currently uses tabs for indentation.
  - `package.json` and `language-configuration.json` use 2‑space indentation.
- Use clear, descriptive keys and regex names in the grammar (e.g., `keyword.control.erb`).
- When adding new scopes, follow TextMate naming patterns (e.g., `string.quoted.double.erb`).

## Testing Guidelines
- No automated tests are present.
- Minimum manual check before PRs:
  - Launch the Extension Development Host.
  - Open a `.erb` file and verify highlighting, comments, and bracket behavior.
  - Spot-check any new grammar rules with representative ERB snippets.

## Commit & Pull Request Guidelines
- Recent commits are short and lowercase (e.g., `add server, client`, `generate codes`).
- Keep commit messages concise and action‑oriented.
- PRs should include:
  - A brief summary and rationale.
  - Manual test notes (steps and observed behavior).
  - Updates to `README.md`/`CHANGELOG.md` for user-facing changes.

## Notes for Contributors
- This repo is a lightweight language extension; avoid adding heavy dependencies without clear value.
- Prefer small, focused changes to the grammar and configuration files.
- JavaScript completions inside `javascript_tag` are powered by a TypeScript LanguageService in `extension.js`.
  This avoids creating real files in the workspace; do not generate or write temp `.js` files.
