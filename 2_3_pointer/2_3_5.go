//Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль. Ввод данных уже написан за вас.

// считайте что fmt уже импортирован и main объявлен
func test(x1 *int, x2 *int) {
	// здесь пишите ваш код
    fmt.Print(*x1 * *x2)
}
