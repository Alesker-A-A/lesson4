package main

import "fmt"

type Task struct{
    Title string
    Done bool
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

func main() {
    showMenu()

}
