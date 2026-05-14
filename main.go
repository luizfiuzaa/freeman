package main

import (
	"fmt"
	"os"
	"os/exec"
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

func main() {
	fmt.Println("$ ~ OH! HELLO FREEMAN. LET'S GO...")
	fmt.Println()
	fmt.Println("$ ~ FREEMAN DOING WHAT NEEDS TO BE DONE")
	fmt.Println()

	runFlutter("clean")
	fmt.Println("$ ~ 1/3")
	fmt.Println()

	runFlutter("pub", "cache", "repair")
	fmt.Println("$ ~ 2/3")
	fmt.Println()

	runFlutter("pub", "cache", "clean")
	fmt.Println("$ ~ 3/3")
	fmt.Println()

	fmt.Println("$ ~ FREEMAN DOING THE SPECIFIC CLEANUP!")

	for _, dir := range dirsToRemove {
		if err := os.RemoveAll(dir); err != nil {
			fmt.Fprintf(os.Stderr, "warning: could not remove %s: %v\n", dir, err)
		}
	}

	for _, file := range filesToRemove {
		if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "warning: could not remove %s: %v\n", file, err)
		}
	}

	fmt.Println()
	fmt.Println("$ ~ FREEMAN RELOADING DEPENDENCIES")
	fmt.Println()

	runFlutter("pub", "get")

	fmt.Println()
	fmt.Println("$ ~ FREEMAN WAS HERE! HAVE A GREAT DAY!")
}

func runFlutter(args ...string) {
	cmd := exec.Command("flutter", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "warning: flutter %v: %v\n", args, err)
	}
}
