# SecondBrainManager - Agent Context

## Project Overview

SecondBrainManager is a Go-based CLI tool that integrates and manages multiple "Second Brain" utilities for knowledge management and content processing.

## Architecture

- **Language**: Go (with embedded README.md)
- **Build Tool**: Makefile
- **Dependencies**: godotenv for environment configuration
- **Help System**: `--help` flag opens embedded README.md in browser/viewer

## Current Integrations

### 1. AI News

- **Tech Stack**: Node.js/TypeScript with pnpm
- **Purpose**: Scrapes and summarizes news articles (especially Hacker News)
- **Key Commands**:
  - Local: `pnpm dev <url>`
  - Global: `ai-news <url>` (after `pnpm link --global`)
- **Config**: `AI_GATEWAY_API_KEY` in `.env` (local) or `~/.config/ai-news/.env` (global)
- **Repo**: [ai-news](https://github.com/albibenni/ai-news)

### 2. Kindle Highlights Parser

- **Tech Stack**: Go with Makefile
- **Purpose**: Extracts and organizes Kindle highlights from `My Clippings.txt`
- **Key Commands**:
  - Install: `make install` (→ `~/go/bin/kindle-parser`)
  - Usage: `kindle-parser "Book Title"`
- **Config**: `NOTE_PATH` and `CLIPPING_PATH` in `mac.env` or `~/.config/kindle-highlights/.env`
- **Repo**: [kindle-highlights](https://github.com/albibenni/kindle-highlights)

## Key Files

- `main.go`: Entry point with embedded README and help flag handler
- `Makefile`: Build, install, test, and documentation commands
- `README.md`: User-facing documentation (embedded in binary)
- `AGENTS.md` (this file): AI agent context
- `CLAUDE.md`, `GEMINI.md`: Symlinks to this file (see `make link-agent`)

## Development Workflow

```bash
make build ARGS="--help"  # Build and run with args
make install              # Install to ~/go/bin
make help                 # Show Makefile commands
make help-doc             # Open README.md
```

## Important Context for AI Agents

- **Help Documentation**: The `--help` flag uses Go's `embed` package to include README.md in the compiled binary, creating a temp file and opening it in the default viewer
- **Environment**: Supports both local `.env` and global config files in `~/.config/<tool>/`
- **Cross-Platform**: Currently macOS-focused (uses `open` command)
- **Modular Design**: Each integration maintains its own repo; this is a wrapper/manager CLI
