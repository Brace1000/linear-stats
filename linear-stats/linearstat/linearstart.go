package linearstat

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

// read the data from a file
func ReadFile(fileName string) []float64 {

    file, err := os.Open(fileName)
    if err != nil {
        fmt.Printf("Error opening the file %v\n", err)
        return nil
    }
    defer file.Close()

    var y_values []float64
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            fmt.Printf("the file has been tampered with,the lines should not be empty %v\n", line)
            os.Exit(1)
        }
        value, err := strconv.ParseFloat(line, 64)
        if err != nil {
            fmt.Printf("error while converting to float %v\n", err)
            return nil
        }
        y_values = append(y_values, value)
    }
    return y_values
}

