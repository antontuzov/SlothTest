package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

// ======================
//
//	SlothTest Configuration
//
// ======================
var (
	green  = color.New(color.FgGreen).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()

	emojis = []rune("🦄🐧🚀🎸🍕🦥👑🎩🐒🦆🥑⚡🌈🍔🎮📌🔑🎲")

	soundEffect = map[string]string{
		"fail": "💥🔥📢 BZZT!",
		"pass": "🎶🌈🪅",
		"skip": "💤😴🌙",
	}

	snarkyComments = []string{
		"Was that test written by a cat? 🐾",
		"Have you tried sacrificing a goat? 🐐",
		"Maybe it's quantum entanglement? 🌀",
		"Congratulations! New error discovered! 🏆",
	}

	partyParrot = []string{"🎉", "🦜", "💃", "🕺", "👯", "✨"}
)

// =================
//
//	Command Line Flags
//
// =================
var (
	watchMode  = flag.Bool("watch", false, "Enable sloth-style file watching")
	emojiMode  = flag.Bool("emojis", true, "Enable hyperactive emoji mode")
	danceParty = flag.Bool("dance", false, "Enable victory dance party")
	noFun      = flag.Bool("nofun", false, "Disable all fun (why would you?)")
)

// ================
//
//	Test Event Struct
//
// ================
type TestEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Elapsed float64   `json:"Elapsed"`
	Output  string    `json:"Output"`
}

// ==============
//
//	Initialization
//
// ==============
func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(cyan("\n🦥 SlothTest v1.0 - Go tests but sloooowly awesome!"))
}

// ===================
//
//	Watch Mode Function
//
// ===================
func runWatchMode() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !isHiddenDirectory(path) {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	var debounceTimer *time.Timer
	debounceDuration := 500 * time.Millisecond

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	runTests()

	for {
		select {
		case <-sigChan:
			fmt.Println(cyan("\n🦥 Sloth says goodbye! Catch you on the flip side!"))
			return
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if shouldHandleEvent(event) {
				if debounceTimer != nil {
					debounceTimer.Stop()
				}
				debounceTimer = time.AfterFunc(debounceDuration, func() {
					fmt.Print("\033[H\033[2J") // Clear screen
					fmt.Println(blue("🦥 Sloth detected changes! Re-running tests..."))
					runTests()
				})
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("😱 Watcher error:", err)
		}
	}
}

// =====================
//
//	Test Runner Function
//
// =====================
func runTests() {
	var passCount, failCount, skipCount int

	cmd := exec.Command("go", "test", "./...", "-json")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("😡 Error creating stdout pipe:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("😱 Error starting command:", err)
		return
	}

	// Progress indicator
	go func() {
		for i := 0; ; i++ {
			if *emojiMode {
				fmt.Printf("\r%s Running tests... %s",
					[]string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"}[i%8],
					[]string{"|", "/", "-", "\\"}[i%4],
				)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	decoder := json.NewDecoder(stdout)
	for decoder.More() {
		var event TestEvent
		if err := decoder.Decode(&event); err != nil {
			fmt.Println("🤯 Error decoding JSON:", err)
			break
		}

		switch event.Action {
		case "pass":
			passCount++
		case "fail":
			failCount++
		case "skip":
			skipCount++
		}

		processEvent(event)
	}

	// Clear progress indicator
	fmt.Print("\r")

	// Print summary
	printSummary(passCount, failCount, skipCount)
	printBingoMessage(failCount)

	if failCount == 0 && *danceParty {
		printVictoryDance()
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(red("😭 Tests failed. Sloth is disappointed."))
	}
}

// ======================
//
//	Event Processing Logic
//
// ======================
func processEvent(event TestEvent) {
	if !*emojiMode {
		// Fallback to simple output if emoji mode is disabled
		fmt.Println(event.Output)
		return
	}

	switch event.Action {
	case "pass":
		if event.Test == "" {
			fmt.Printf("%s %s %s\n",
				green("✔"),
				randomEmoji(),
				event.Package,
			)
		} else {
			fmt.Printf("    %s %s %s\n",
				green("✔"),
				randomEmoji(),
				event.Test,
			)
		}

	case "fail":
		fmt.Printf("\a") // System alert sound
		if event.Test == "" {
			fmt.Printf("%s %s %s\n%s\n",
				red("✖"),
				"💩",
				event.Package,
				snarkyComments[rand.Intn(len(snarkyComments))],
			)
		} else {
			fmt.Printf("    %s %s %s %s\n",
				red("✖"),
				"💩",
				event.Test,
				soundEffect["fail"],
			)
		}

	case "skip":
		fmt.Printf("%s %s %s %s\n",
			yellow("⚠"),
			"🦥",
			event.Package,
			soundEffect["skip"],
		)

	case "output":
		if event.Output != "" {
			fmt.Printf("%s %s", "📜", event.Output)
		}
	}
}

// ===============
//
//	Helper Functions
//
// ===============
func randomEmoji() string {
	return string(emojis[rand.Intn(len(emojis))])
}

func isHiddenDirectory(path string) bool {
	return len(path) > 1 && path[0] == '.' && path != "." && path != ".."
}

func shouldHandleEvent(event fsnotify.Event) bool {
	return (event.Op.Has(fsnotify.Write) || event.Op.Has(fsnotify.Create)) &&
		filepath.Ext(event.Name) == ".go"
}

func printSummary(pass, fail, skip int) {
	fmt.Printf("\n%s\n", cyan("📊 Test Summary:"))
	fmt.Printf("  %s Passed: %d\n", green("✔"), pass)
	fmt.Printf("  %s Failed: %d\n", red("✖"), fail)
	fmt.Printf("  %s Skipped: %d\n", yellow("⚠"), skip)
}

func printBingoMessage(fails int) {
	if fails > 3 {
		fmt.Printf("\n  %s BINGO! %d failures! Reward: %s\n",
			"🎰",
			fails,
			[]string{
				"Take a coffee break! ☕",
				"Do 10 pushups! 💪",
				"Pet your nearest animal! 🐶",
			}[rand.Intn(3)],
		)
	}
}

func printVictoryDance() {
	fmt.Println("\n🎉 Victory Dance Sequence Initiated!")
	for i := 0; i < 3; i++ {
		fmt.Printf("\r%s ", partyParrot[rand.Intn(len(partyParrot))])
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("\r%s ", partyParrot[rand.Intn(len(partyParrot))])
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("\n🦜 Dance party complete! You rock! 🤘")
}

// ============
//
//	Entry Point
//
// ============
func main() {
	flag.Parse()

	if *noFun {
		*emojiMode = false
		*danceParty = false
		fmt.Println("😢 Fun disabled. You monster.")
	}

	if *watchMode {
		runWatchMode()
	} else {
		runTests()
	}
}
