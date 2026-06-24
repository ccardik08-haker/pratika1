package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// представляет структуру данных о сотруднике
type Employee struct {
	ID       int    //уникальный идентификатор
	Name     string //полное имя
	Age      int    //возраст (должен быть > 0 и < 120)
	Position string //название должности
	Salary   int    //зарплата в рублях
}

// управляет списком сотрудников
type EmployeeManager struct {
	employees []*Employee //список сотрудников
	nextID    int         //следующий доступный ID
	maxSize   int         //максимальное количество сотрудников
}

// создает новый менеджер сотрудников
func NewEmployeeManager(maxSize int) *EmployeeManager {
	return &EmployeeManager{
		employees: make([]*Employee, 0, maxSize),
		nextID:    1,
		maxSize:   maxSize,
	}
}

// добавляет нового сотрудника
func (m *EmployeeManager) Add(name string, age int, position string, salary int) error {
	//проверяем лимит
	if len(m.employees) >= m.maxSize {
		return fmt.Errorf("достигнут лимит сотрудников (%d)", m.maxSize)
	}

	//валидация данных
	if err := ValidateEmployee(name, age, position, salary); err != nil {
		return err
	}

	//создаем нового сотрудника
	emp := &Employee{
		ID:       m.nextID,
		Name:     name,
		Age:      age,
		Position: position,
		Salary:   salary,
	}
	m.nextID++

	//добавляем в список
	m.employees = append(m.employees, emp)
	return nil
}

// удаляет сотрудника по имени
func (m *EmployeeManager) Delete(name string) error {
	//ищем сотрудника
	for i, emp := range m.employees {
		if strings.EqualFold(emp.Name, name) {
			//удаляем элемент из слайса
			m.employees = append(m.employees[:i], m.employees[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("сотрудник '%s' не найден", name)
}

// удаляет сотрудника по ID
func (m *EmployeeManager) DeleteByID(id int) error {
	for i, emp := range m.employees {
		if emp.ID == id {
			m.employees = append(m.employees[:i], m.employees[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("сотрудник с ID %d не найден", id)
}

// возвращает список всех сотрудников
func (m *EmployeeManager) GetList() []*Employee {
	return m.employees
}

// возвращает количество сотрудников
func (m *EmployeeManager) GetCount() int {
	return len(m.employees)
}

// проверяет, заполнен ли список
func (m *EmployeeManager) IsFull() bool {
	return len(m.employees) >= m.maxSize
}

// ищет сотрудника по имени
func (m *EmployeeManager) FindByName(name string) []*Employee {
	var result []*Employee
	for _, emp := range m.employees {
		if strings.Contains(strings.ToLower(emp.Name), strings.ToLower(name)) {
			result = append(result, emp)
		}
	}
	return result
}

// проверяет корректность данных сотрудника
func ValidateEmployee(name string, age int, position string, salary int) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("имя не может быть пустым")
	}
	if age < 0 || age > 120 {
		return fmt.Errorf("возраст должен быть от 0 до 120")
	}
	if strings.TrimSpace(position) == "" {
		return fmt.Errorf("должность не может быть пустой")
	}
	if salary < 0 {
		return fmt.Errorf("зарплата не может быть отрицательной")
	}
	return nil
}

// для безопасного чтения ввода
type InputReader struct {
	reader *bufio.Reader
}

func NewInputReader() *InputReader {
	return &InputReader{
		reader: bufio.NewReader(os.Stdin),
	}
}

// читает строку с консоли
func (r *InputReader) ReadString(prompt string) string {
	fmt.Print(prompt)
	input, _ := r.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// читает целое число с консоли
func (r *InputReader) ReadInt(prompt string) (int, error) {
	input := r.ReadString(prompt)
	return strconv.Atoi(input)
}

// выводит список сотрудников в виде таблицы
func DisplayEmployees(employees []*Employee) {
	if len(employees) == 0 {
		fmt.Println("\nСписок сотрудников пуст")
		return
	}

	fmt.Println("\nСПИСОК СОТРУДНИКОВ")
	fmt.Println("═══════════════════════════════════════════════════════════════════════════")
	fmt.Printf("%-4s │ %-20s │ %-6s │ %-15s │ %-12s\n", "ID", "Имя", "Возраст", "Должность", "Зарплата")
	fmt.Println("─────┼──────────────────────┼────────┼─────────────────┼──────────────")

	for _, emp := range employees {
		fmt.Printf("%-4d │ %-20s │ %-6d │ %-15s │ %-12d\n",
			emp.ID, emp.Name, emp.Age, emp.Position, emp.Salary)
	}
	fmt.Println("═══════════════════════════════════════════════════════════════════════════")
	fmt.Printf("Всего сотрудников: %d\n", len(employees))
}

// отображает главное меню
func ShowMenu() {
	fmt.Println("\n═══════════════════════════════════════════════════════")
	fmt.Println("           УПРАВЛЕНИЕ СОТРУДНИКАМИ v1.0")
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("  1. Добавить нового сотрудника")
	fmt.Println("  2. Удалить сотрудника по имени")
	fmt.Println("  3. Удалить сотрудника по ID")
	fmt.Println("  4. Показать всех сотрудников")
	fmt.Println("  5. Найти сотрудника по имени")
	fmt.Println("  6. Показать статистику")
	fmt.Println("  7. Выйти из программы")
	fmt.Println("═══════════════════════════════════════════════════════")
}

// показывает статистику по сотрудникам
func ShowStatistics(manager *EmployeeManager) {
	employees := manager.GetList()
	if len(employees) == 0 {
		fmt.Println("\nНет данных для статистики")
		return
	}

	var totalSalary int
	var totalAge int
	var maxSalary int
	var minSalary int = employees[0].Salary

	for _, emp := range employees {
		totalSalary += emp.Salary
		totalAge += emp.Age
		if emp.Salary > maxSalary {
			maxSalary = emp.Salary
		}
		if emp.Salary < minSalary {
			minSalary = emp.Salary
		}
	}

	avgSalary := totalSalary / len(employees)
	avgAge := totalAge / len(employees)

	fmt.Println("\nСТАТИСТИКА")
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Printf("Всего сотрудников: %d\n", len(employees))
	fmt.Printf("Средняя зарплата: %d руб.\n", avgSalary)
	fmt.Printf("Максимальная зарплата: %d руб.\n", maxSalary)
	fmt.Printf("Минимальная зарплата: %d руб.\n", minSalary)
	fmt.Printf("Средний возраст: %d лет\n", avgAge)
	fmt.Printf("Занятость: %d/%d (%.1f%%)\n",
		len(employees), manager.maxSize,
		float64(len(employees))/float64(manager.maxSize)*100)
	fmt.Println("═══════════════════════════════════════════════════════")
}

func main() {
	//создаем менеджер сотрудников с лимитом 512
	manager := NewEmployeeManager(512)
	input := NewInputReader()

	fmt.Println("Добро пожаловать в систему управления сотрудниками!")

	for {
		ShowMenu()
		choice, err := input.ReadInt("Выберите действие: ")

		if err != nil {
			fmt.Println("Ошибка: введите число от 1 до 7")
			continue
		}

		switch choice {
		case 1:
			//добавление сотрудника
			fmt.Println("\nВведите данные сотрудника:")

			name := input.ReadString("Имя: ")
			age, _ := input.ReadInt("Возраст: ")
			position := input.ReadString("Должность: ")
			salary, _ := input.ReadInt("Зарплата: ")

			if err := manager.Add(name, age, position, salary); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("Сотрудник '%s' успешно добавлен! (ID: %d)\n", name, manager.nextID-1)
			}

		case 2:
			//удаление по имени
			if manager.GetCount() == 0 {
				fmt.Println("Нет сотрудников для удаления")
				continue
			}
			name := input.ReadString("Введите имя сотрудника для удаления: ")
			if err := manager.Delete(name); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("Сотрудник '%s' успешно удален!\n", name)
			}

		case 3:
			//удаление по ID
			if manager.GetCount() == 0 {
				fmt.Println("Нет сотрудников для удаления")
				continue
			}
			id, err := input.ReadInt("Введите ID сотрудника для удаления: ")
			if err != nil {
				fmt.Println("Ошибка: введите корректный ID")
				continue
			}
			if err := manager.DeleteByID(id); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("Сотрудник с ID %d успешно удален!\n", id)
			}

		case 4:
			//вывод списка
			DisplayEmployees(manager.GetList())

		case 5:
			//поиск по имени
			if manager.GetCount() == 0 {
				fmt.Println("Нет сотрудников для поиска")
				continue
			}
			name := input.ReadString("Введите имя для поиска: ")
			results := manager.FindByName(name)
			if len(results) == 0 {
				fmt.Printf("Сотрудников с именем '%s' не найдено\n", name)
			} else {
				fmt.Printf("Найдено сотрудников: %d\n", len(results))
				DisplayEmployees(results)
			}

		case 6:
			//статистика
			ShowStatistics(manager)

		case 7:
			//выход
			fmt.Println("\nДо свидания! Спасибо за использование программы!")
			return

		default:
			fmt.Println("Неверный выбор! Пожалуйста, выберите 1-7")
		}
	}
}
