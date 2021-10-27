# språk

Duolingo-like language learning tool for the CLI.

Go, SQLite with Prisma & [bubbletea](https://github.com/charmbracelet/bubbletea).

- [x] Anki deck sync/import
- [] Modes mirroring Duo:
  - [] Multliple choice (for words, sentences)
  - [] Type sentence to / from target language
  - [] Listen & write from target to native
  - [] Word matching
- [] Google TTS Waveshare voice

Obviously a lot more to be considered like word synonyms / alternative structures but just to KISS will use LD to calculate correctness.

### Cool things it has

Custom Angular-like router, event bus & closure based components as opposed to reciever functions.

## Setup

### Sync Prisma schema

```
chmod +x prisma-sync.sh
./prisma-sync.sh
```

## Debugger

```
dlv debug --headless --listen=:2345 --log --api-version=2 ; clear
```
