package aoctemplate

import (
	"errors"
	"strconv"
	"strings"

	"github.com/0xleft/aoctemplate/utils"
	"github.com/barweiss/go-tuple"
)

type Solution struct {
	Day      int
	Year     int
	Filename string
}

func (solution *Solution) GetContents() (string, error) {
	return utils.ReadContents(solution.Filename)
}

func (solution *Solution) GetLines(eol EndOfLine) ([]string, error) {
	contents, err := solution.GetContents()
	if err != nil {
		return nil, err
	}

	return strings.Split(contents, string(eol)), nil
}

func (solution *Solution) GetRanges(delimiter string, eol EndOfLine, ignoreErrors bool) ([]tuple.T2[int64, int64], error) {
	lines, err := solution.GetLines(eol)
	if err != nil {
		return nil, err
	}

	ranges := make([]tuple.T2[int64, int64], 0)

	for _, line := range lines {
		split := strings.Split(line, delimiter)
		if len(split) != 2 && !ignoreErrors {
			return nil, errors.New("Invalid length of range line. Expected 2 got: " + strconv.FormatInt(int64(len(split)), 10))
		}

		start, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil && !ignoreErrors {
			return nil, errors.New("Invalid start of range. Expected a number got: " + split[0])
		}
		end, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil && !ignoreErrors {
			return nil, errors.New("Invalid end of range. Expected a number got: " + split[1])
		}

		ranges = append(ranges, tuple.New2(start, end))
	}

	return ranges, nil
}

func (solution *Solution) GetSkippedNumbers(start int, eol EndOfLine, ignoreErrors bool) ([]int64, error) {
	lines, err := solution.GetLines(eol)
	if err != nil {
		return nil, err
	}

	numbers := make([]int64, 0)

	for _, line := range lines {
		numberString := line[start:]
		number, err := strconv.ParseInt(numberString, 10, 64)
		if err != nil && !ignoreErrors {
			return nil, errors.New("Invalid number. Expected a number got: " + numberString)
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}
