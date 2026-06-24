package main

import (
<<<<<<< HEAD
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
=======
	"fmt"
	"time"
)

// выполняет сортировку пузырьком с оптимизацией
// принимает срез целых чисел и возвращает отсортированный срез
func BubbleSort(arr []int) []int {
	//создаем копию, чтобы не изменять исходный массив
	result := make([]int, len(arr))
	copy(result, arr)

	n := len(result)
	if n <= 1 {
		return result
	}

	//внешний цикл - количество проходов
	for i := 0; i < n-1; i++ {
		swapped := false

		//внутренний цикл - сравниваем соседние элементы
		//с каждым проходом уменьшаем количество проверяемых элементов
		for j := 0; j < n-i-1; j++ {
			//если текущий элемент больше следующего - меняем местами
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}

		//если за проход не было обменов - массив отсортирован
		if !swapped {
			break
		}
	}

	return result
}

// выводит массив в консоль
func PrintArray(arr []int, title string) {
	fmt.Printf("\n%s:\n", title)
	fmt.Println("─────────────────────────────────────────────")

	//выводим по 10 элементов в строке для удобства чтения
	for i, val := range arr {
		fmt.Printf("%5d ", val)
		if (i+1)%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("─────────────────────────────────────────────")
	fmt.Printf("Всего элементов: %d\n", len(arr))
}

// возвращает массив из 100 чисел для примера
func GetExampleArray() []int {
	return []int{
		542, -565, 531, -294, -56, 14, 270, -51, -914, 605,
		-117, -768, 331, 708, -603, 84, -548, 579, 434, 751,
		592, -349, 408, -602, 721, 909, 170, -432, -970, -171,
		-972, 316, 405, -676, -929, -795, -682, -646, 46, -609,
		-84, 180, -158, -662, -384, 854, -721, 39, 180, -197,
		-818, -946, -529, -555, -36, -853, -322, 540, -936, -919,
		473, 978, 782, 586, 869, 333, -977, -548, -789, 988,
		-393, 807, -609, 997, 824, -480, -205, -576, 856, 494,
		131, 40, -601, 467, 221, -640, 34, -220, 482, 948,
		523, -27, -771, -914, 438, 957, 205, -411, -749, -723,
	}
}

// измеряет время выполнения функции
func MeasureExecutionTime(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

// проверяет, отсортирован ли массив
func CheckSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func main() {
	//получаем пример массива
	original := GetExampleArray()

	//выводим исходный массив
	PrintArray(original, "Исходный массив")

	//измеряем время сортировки
	duration := MeasureExecutionTime(func() {
		//сортируем и сохраняем результат
		sorted := BubbleSort(original)

		//выводим отсортированный массив
		PrintArray(sorted, "Отсортированный массив")

		//проверяем, что массив отсортирован
		if CheckSorted(sorted) {
			fmt.Println("Массив успешно отсортирован!")
		} else {
			fmt.Println("Ошибка: массив не отсортирован!")
		}
	})

	//выводим время выполнения
	fmt.Printf("\nВремя сортировки: %v\n", duration)
>>>>>>> f0f463ef9e897691dd71fb50b52d693aaeec033b
}
