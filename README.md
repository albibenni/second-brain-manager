# SecondBrainManager

## Current Support

### Ai news

AI news summarizer and article generator that uses AI to analyze and summarize content from news websites, with special support for Hacker News.

#### Setup

For a quick start:

- Install pnpm (if needed): `npm install -g pnpm`
- Configure pnpm and restart your shell: `pnpm setup` then reopen your terminal
- Install deps: `pnpm install`
- Configure API key:
  - Global: `mkdir -p ~/.config/ai-news && echo "AI_GATEWAY_API_KEY=your_api_key" > ~/.config/ai-news/.env`
- If you see pnpm store errors when linking globally: `pnpm install -g`

Usage:

- Run locally: `pnpm dev <url>` (e.g. `pnpm dev https://news.ycombinator.com/`)
- Install globally: `pnpm link --global` then run `ai-news <url>`

For the full setup and examples, see the ai-news repository: <https://github.com/albibenni/ai-news>

### Kindle Highlights

#### Setup

Quick start:

- Prereqs: Go (>=1.24) and `make`.
- Clone and enter the repo:

```bash
git clone <repo-url>
cd kindle-highlights
```

- Create global config and copy defaults: `make setup-config`
- Build & install: `make install` (installs `kindle-parser` to `~/go/bin`)
- Ensure Go bin is in PATH:

```bash
export PATH=$HOME/go/bin:$PATH
```

- Configure env (local or global): set `NOTE_PATH` and `CLIPPING_PATH` in `mac.env` or `~/.config/kindle-highlights/.env`

Usage examples:

```bash
kindle-parser help
kindle-parser "Book Title"
kindle-parser test "Book Title"
```

For the full setup refer to [Kindle-Higlights](https://github.com/albibenni/kindle-highlights).

## Testing

- Run the test suite using the Makefile helper:

```bash
make test
```

- Behavior: the `test` target will use `gotestsum` for nicer output when it is installed; otherwise it falls back to `go test ./... -v`.
- Note: Make executes recipe lines with the system shell (usually `/bin/sh`). If you need Bash-specific features in Make recipes, set `SHELL := /bin/bash` at the top of the `Makefile`.

