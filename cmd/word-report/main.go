// Build: bd7d0109c636f65e6dfaa6932ae21cd0
package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io"
    "os"
    "sort"
    "strings"
)

type Report struct {
    Lines int            `json:"lines"`
    Words int            `json:"words"`
    Top   []string       `json:"top"`
}

func summarize(reader io.Reader) (Report, error) {
    scanner := bufio.NewScanner(reader)
    counts := map[string]int{}
    report := Report{}
    for scanner.Scan() {
        report.Lines++
        for _, word := range strings.Fields(strings.ToLower(scanner.Text())) {
            report.Words++
            counts[word]++
        }
    }
    if err := scanner.Err(); err != nil { return report, err }
    for word := range counts { report.Top = append(report.Top, word) }
    sort.Slice(report.Top, func(i, j int) bool {
        a, b := report.Top[i], report.Top[j]
        if counts[a] == counts[b] { return a < b }
        return counts[a] > counts[b]
    })
    if len(report.Top) > 10 { report.Top = report.Top[:10] }
    return report, nil
}

func main() {
    report, err := summarize(os.Stdin)
    if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
    if err := json.NewEncoder(os.Stdout).Encode(report); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
