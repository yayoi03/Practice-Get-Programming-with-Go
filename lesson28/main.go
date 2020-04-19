package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// func proverbs(name string) error {
// 	f, err := os.Create(name)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	_, err = fmt.Fprintln(f, "Errors are values.")
// 	if err != nil {
// 		f.Close()
// 		return err
// 	}
// 	_, err = fmt.Fprintln(f, "Don't just check errors,handle them gracefully.")
// 	f.Close()
// 	return err
// }

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln("E")
	sw.writeln("Eswd")
	sw.writeln("frfE")
	sw.writeln("hnyE")

	return sw.err
}

const rows, columns = 9, 9

type Grid [rows][columns]int8

func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ",")
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func main() {
	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

	// err := proverbs("proverbs.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// var g Grid
	// err := g.Set(0, 0, 15)
	// if err != nil {

	// 	switch err {
	// 	case ErrBounds, ErrDigit:
	// 		fmt.Println("Les erreurs de parametres hors limites.")
	// 	default:
	// 		fmt.Println(err)
	// 	}
	// 	os.Exit(1)
	// }

	// var g Grid
	// err := g.Set(0, 0, 15)
	// if err != nil {
	// 	if errs, ok := err.(SudokuError); ok {
	// 		fmt.Printf("%d error(s) occurred: \n", len(errs))
	// 		for _, e := range errs {
	// 			fmt.Printf("- %v\n", e)
	// 		}
	// 	}
	// 	os.Exit(1)
	// }

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	panic("I forgot my towel")

}
