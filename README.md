# språk

Duolingo-like language learning tool for the CLI.

Go, SQLite with Prisma & [Lipgloss](https://github.com/charmbracelet/lipgloss).

- [] Anki deck sync
- [] Google TTS Waveshare voice
- [] Modes mirroring Duo:
  - [] Multliple choice (for words, sentences)
  - [] Type sentence to / from target language
  - [] Listen & write from target to native
  - [] Word matching

Obviously a lot more to be considered like word synonyms / alternative structures but just to KISS will use LD to calculate correctness.

## Setup

### Sync Prisma schema

```
chmod +x prisma-sync.sh
./prisma-sync.sh
```
