package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// InputReader для безопасного чтения ввода
type InputReader struct {
	reader *bufio.Reader
}

func NewInputReader() *InputReader {
	return &InputReader{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (r *InputReader) ReadString(prompt string) string {
	fmt.Print(prompt)
	input, _ := r.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (r *InputReader) ReadInt(prompt string) (int, error) {
	input := r.ReadString(prompt)
	return strconv.Atoi(input)
}

// ShowMainMenu показывает главное меню
func ShowMainMenu() {
	fmt.Println("\n═══════════════════════════════════════════════════════")
	fmt.Println("              ГО ПРАКТИКА - ГЛАВНОЕ МЕНЮ")
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("  1. Сортировка пузырьком")
	fmt.Println("  2. Управление сотрудниками")
	fmt.Println("  3. Выйти из программы")
	fmt.Println("═══════════════════════════════════════════════════════")
}

// RunProgram запускает другую программу Go
func RunProgram(path string) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "go", "run", path)
	} else {
		cmd = exec.Command("go", "run", path)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Ошибка запуска программы: %v\n", err)
	}
}

func main() {
	input := NewInputReader()

	for {
		ShowMainMenu()
		choice, err := input.ReadInt("Выберите программу: ")

		if err != nil {
			fmt.Println("Ошибка: введите число от 1 до 3")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("\nЗапуск сортировки пузырьком...")
			RunProgram("sortirovka/main.go")
		case 2:
			fmt.Println("\nЗапуск управления сотрудниками...")
			RunProgram("sotrudniki/main.go")
		case 3:
			fmt.Println("\nДо свидания!")
			return
		default:
			fmt.Println("Неверный выбор! Пожалуйста, выберите 1-3")
		}
	}
}
