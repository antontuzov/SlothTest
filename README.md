# SlothTest 🦥

**SlothTest** is the world's most entertaining Go test runner! It adds color, emojis, sound effects, and humor to your test output. Perfect for developers who want to make testing fun again.

---

![Example Output](example/example_output.png)

---

## Features ✨

- **Colorful Output**: Green for passes, red for failures, yellow for skips.
- **Emoji Explosion**: Random emojis for every test result.
- **Sound Effects**: System beep on failures.
- **Snarky Comments**: Humorous messages for failed tests.
- **Victory Dance**: Animated party parrots on success.
- **Watch Mode**: Automatically rerun tests on file changes.

---

## Installation 🛠️

### Option 1: Install Globally
1. Install SlothTest using `go install`:

 ```
   go install github.com/antontuzov/slothtest@latest

 ```

## Install in Your Project

 ```
   go get github.com/antontuzov/slothtest@latest

 ```
## Run Tests with SlothTest
```
# Basic test run
slothtest

# Watch mode (rerun tests on file changes)
slothtest -watch

# Watch mode with emojis and dance party
slothtest -watch -dance

```


## Flags


-watch: Enable file watching mode.

-dance: Enable victory dance party on success.

-emojis: Enable emoji mode (default: true).

-nofun: Disable all fun features (why would you?).



## Example Output 🖥️

```
🦥 SlothTest v1.0 - Go tests but sloooowly awesome!

  ✔ 🍕 mypkg 🎶🌈🪅
  💤😴🌙 🦥 skippedpkg
  ✖ 💩 failingpkg "Maybe it's quantum entanglement? 🌀"

📜 === RUN   TestBroken
📜     broken_test.go:10: Oh nooooo

🎉 Victory Dance Sequence Initiated!
🦜 Dance party complete! You rock! 🤘

```



