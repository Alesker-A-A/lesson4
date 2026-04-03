package main

import "fmt"

type Task struct{
    Title string
    Done bool
}
var tasks []Task

func showMenu() {
    fmt.Printf("~~~ To-Do Лист ~~~")
    fmt.Printf("1. Показать все задачи")
    fmt.Printf("2. Добавить задачу")
    fmt.Printf("3. Пометить как 'Выполнена'")
    fmt.Printf("4. Удалить задачу")
    fmt.Printf("5. Выйти")
}

func main() {
    showMenu

}
