package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const NumericChars = "0123456789"

type CalibrationPart struct {
	char  string
	index int
}

type CalibrationLine struct {
	parts []CalibrationPart
}

func (cl *CalibrationLine) Add(part CalibrationPart) {
	cl.parts = append(cl.parts, part)
}

func (cl *CalibrationLine) Sum() int64 {
	var b bytes.Buffer
	switch len(cl.parts) {
	case 1:
		b.WriteString(cl.parts[0].char)
		b.WriteString(cl.parts[0].char)
	case 0:
		break
	default:
		for _, p := range cl.parts {
			b.WriteString(p.char)
		}
	}
	val, _ := strconv.Atoi(b.String())
	return int64(val)
}

func NewCalibrationLine() *CalibrationLine {
	return &CalibrationLine{parts: []CalibrationPart{}}
}

type Calibration struct {
	lines []*CalibrationLine
}

func NewCalibration() *Calibration {
	return &Calibration{lines: []*CalibrationLine{}}
}

func (c *Calibration) AddLine(line *CalibrationLine) {
	c.lines = append(c.lines, line)
}

func (c *Calibration) Sums() []int64 {
	var sums []int64
	for _, line := range c.lines {
		sums = append(sums, line.Sum())
	}
	return sums
}

func (c *Calibration) Sum() int64 {
	var sum int64
	for _, s := range c.Sums() {
		sum = sum + s
	}
	return sum
}

func main() {
	var filepath string
	flag.StringVar(&filepath, "file", "input", "filepath to input file")
	flag.Parse()

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("failed to open file %s : %v", filepath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int64
	for scanner.Scan() {
		line := scanner.Bytes()
		fmt.Println(string(line))
		cLine := FirstLastParts(line)

		s := cLine.Sum()
		total += s
	}

	log.Default().Printf("Sum: %v", total)
	os.Exit(0)
}

// FirstLastParts create a CalibrationLine from byte slice
func FirstLastParts(line []byte) *CalibrationLine {
	cLine := NewCalibrationLine()
	fIdx := bytes.IndexAny(line, NumericChars)
	cLine.Add(CalibrationPart{char: string(line[fIdx]), index: fIdx})
	lIdx := bytes.LastIndexAny(line, NumericChars)
	if fIdx != lIdx {
		cLine.Add(CalibrationPart{char: string(line[lIdx]), index: lIdx})
	}
	return cLine
}
