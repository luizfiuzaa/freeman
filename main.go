package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
)

var dirsToRemove = []string{
	"build",
	".dart_tool",
	".gradle",
	".idea",
	".packages",
	"ios/Pods",
	"ios/.symlinks",
	"ios/Flutter/Flutter.framework",
	"ios/Flutter/Flutter.podspec",
	"ios/Flutter/App.framework",
	"android/.gradle",
	"android/.idea",
	"android/.gradle/caches",
	"android/.gradle/daemon",
	"android/.gradle/native",
	"android/.gradle/7.0",
	"android/build",
}

var filesToRemove = []string{
	"pubspec.lock",
}

type Config struct {
	PrioritizeFVM bool `json:"prioritize_fvm"`
}

func configPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".freeman_config.json"
	}
	return filepath.Join(home, ".freeman", "config.json")
}

func loadConfig() Config {
	data, err := os.ReadFile(configPath())
	if err != nil {
		return Config{}
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}
	}
	return cfg
}

func saveConfig(cfg Config) error {
	path := configPath()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func handleConfig(args []string) {
	cfg := loadConfig()

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--prioritize-fvm":
			if i+1 >= len(args) {
				fmt.Fprintln(os.Stderr, "error: --prioritize-fvm requires a value (true/false)")
				os.Exit(1)
			}
			i++
			val, err := strconv.ParseBool(args[i])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: invalid value for --prioritize-fvm: %s\n", args[i])
				os.Exit(1)
			}
			cfg.PrioritizeFVM = val
		default:
			fmt.Fprintf(os.Stderr, "unknown config option: %s\n", args[i])
			os.Exit(1)
		}
	}

	if err := saveConfig(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error saving config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("$ ~ Config saved: prioritize_fvm=%v\n", cfg.PrioritizeFVM)
}

func pubCachePath() string {
	switch runtime.GOOS {
	case "windows":
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData == "" {
			home, _ := os.UserHomeDir()
			localAppData = filepath.Join(home, "AppData", "Local")
		}
		return filepath.Join(localAppData, "Pub", "Cache")
	default:
		home, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		return filepath.Join(home, ".pub-cache")
	}
}

func cleanPubCache() {
	cachePath := pubCachePath()
	if cachePath == "" {
		fmt.Fprintln(os.Stderr, "warning: could not determine pub cache path")
		return
	}
	fmt.Printf("$ ~ CLEANING PUB CACHE: %s\n", cachePath)
	if err := os.RemoveAll(cachePath); err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not remove pub cache: %v\n", err)
		return
	}
	fmt.Println("$ ~ PUB CACHE CLEARED!")
}

func fvmAvailable() bool {
	_, err := exec.LookPath("fvm")
	return err == nil
}

func hasFVMProject() bool {
	_, err := os.Stat(".fvm")
	return err == nil
}

func shouldUseFVM(forceFVM bool) bool {
	if forceFVM || loadConfig().PrioritizeFVM || hasFVMProject() {
		if !fvmAvailable() {
			fmt.Println("$ ~ FVM not found, falling back to global Flutter")
			return false
		}
		return true
	}
	return false
}

func runFlutter(useFVM bool, args ...string) {
	var cmd *exec.Cmd
	if useFVM {
		cmd = exec.Command("fvm", append([]string{"flutter"}, args...)...)
	} else {
		cmd = exec.Command("flutter", args...)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "warning: flutter %v: %v\n", args, err)
	}
}

func logRemove(path string, dryRun bool) {
	if dryRun {
		fmt.Printf("$ ~ [DRY RUN] would remove: %s\n", path)
	} else {
		fmt.Printf("$ ~ removing: %s\n", path)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 && args[0] == "config" {
		handleConfig(args[1:])
		return
	}

	forceFVM := false
	cleanCache := false
	safeMode := false
	noCacheClean := false
	noRepair := false
	keepLockfile := false
	dryRun := false
	verbose := false

	for _, arg := range args {
		switch arg {
		case "--fvm", "--use-fvm":
			forceFVM = true
		case "--clean-cache":
			cleanCache = true
		case "--safe":
			safeMode = true
		case "--no-cache-clean":
			noCacheClean = true
		case "--no-repair":
			noRepair = true
		case "--keep-lockfile":
			keepLockfile = true
		case "--dry-run":
			dryRun = true
		case "--verbose":
			verbose = true
		}
	}

	useFVM := shouldUseFVM(forceFVM)

	// Calculate total steps (flutter clean + optional cache commands)
	totalSteps := 1
	if !safeMode && !noRepair {
		totalSteps++
	}
	if !safeMode && !noCacheClean {
		totalSteps++
	}

	step := 0
	nextStep := func() {
		step++
		fmt.Printf("$ ~ %d/%d\n", step, totalSteps)
		fmt.Println()
	}

	fmt.Println("$ ~ OH! HELLO FREEMAN. LET'S GO...")
	if useFVM {
		fmt.Println("$ ~ USING FVM FOR FLUTTER COMMANDS")
	}
	if safeMode {
		fmt.Println("$ ~ SAFE MODE — skipping cache and directory cleanup")
	}
	if dryRun {
		fmt.Println("$ ~ DRY RUN — no changes will be made")
	}
	fmt.Println()
	fmt.Println("$ ~ FREEMAN DOING WHAT NEEDS TO BE DONE")
	fmt.Println()

	// Optional: remove local pub cache
	if cleanCache && !safeMode {
		if dryRun {
			fmt.Printf("$ ~ [DRY RUN] would remove pub cache: %s\n", pubCachePath())
			fmt.Println()
		} else {
			cleanPubCache()
			fmt.Println()
		}
	}

	// flutter clean
	if dryRun {
		fmt.Println("$ ~ [DRY RUN] flutter clean")
	} else {
		runFlutter(useFVM, "clean")
	}
	nextStep()

	// flutter pub cache repair
	if !safeMode && !noRepair {
		if dryRun {
			fmt.Println("$ ~ [DRY RUN] flutter pub cache repair")
		} else {
			runFlutter(useFVM, "pub", "cache", "repair")
		}
		nextStep()
	}

	// flutter pub cache clean
	if !safeMode && !noCacheClean {
		if dryRun {
			fmt.Println("$ ~ [DRY RUN] flutter pub cache clean")
		} else {
			runFlutter(useFVM, "pub", "cache", "clean")
		}
		nextStep()
	}

	// Directory and file cleanup
	if !safeMode {
		fmt.Println("$ ~ FREEMAN DOING THE SPECIFIC CLEANUP!")

		for _, dir := range dirsToRemove {
			if dryRun || verbose {
				logRemove(dir, dryRun)
			}
			if !dryRun {
				if err := os.RemoveAll(dir); err != nil {
					fmt.Fprintf(os.Stderr, "warning: could not remove %s: %v\n", dir, err)
				}
			}
		}

		if keepLockfile {
			if verbose {
				fmt.Println("$ ~ keeping pubspec.lock (--keep-lockfile)")
			}
		} else {
			for _, file := range filesToRemove {
				if dryRun || verbose {
					logRemove(file, dryRun)
				}
				if !dryRun {
					if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
						fmt.Fprintf(os.Stderr, "warning: could not remove %s: %v\n", file, err)
					}
				}
			}
		}
	}

	fmt.Println()
	fmt.Println("$ ~ FREEMAN RELOADING DEPENDENCIES")
	fmt.Println()

	if dryRun {
		fmt.Println("$ ~ [DRY RUN] flutter pub get")
	} else {
		runFlutter(useFVM, "pub", "get")
	}

	fmt.Println()
	fmt.Println("$ ~ FREEMAN WAS HERE! HAVE A GREAT DAY!")
}
