package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type Task struct {
  Title string
  Done  bool
}

var tasks []Task

func showMenu() {
  fmt.Println("~~~ To-Do Лист ~~~")
  fmt.Println("1. Показать все задачи")
  fmt.Println("2. Добавить задачу")
  fmt.Println(`3. Пометить как "Выполнена"`)
  fmt.Println("4. Удалить задачу")
  fmt.Println("5. Выйти")
}

func addTask(reader *bufio.Reader) {
  fmt.Print("Введите название задачи: ")
  title, _ := reader.ReadString('\n')
  title = strings.TrimSpace(title)

  newTask := Task{Title: title, Done: false}
  tasks = append(tasks, newTask)
  fmt.Println("Задача добавлена!")
}

func showTasks() {
    if len(tasks) == 0 {
        fmt.Println("Список задач пуст")
        return
    }
    fmt.Println("~~~ Список Задач ~~~")
    for i, task := range tasks {
        status := "[ ]"
        if task.Done {
            status = "[✓]"
        }
        fmt.Printf("%d. %s %s\n", i+1, status, task.Title)
    }
}

func main() {
  reader := bufio.NewReader(os.Stdin)

  for {
    showMenu()
    fmt.Print("Ваш выбор: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    if input == "5" {
      fmt.Println("До свидания!")
      break
    }
    if input == "2" {
      addTask(reader)
    }
  }

}