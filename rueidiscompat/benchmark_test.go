package rueidiscompat

import (
"strings"
"testing"
)

// Current implementation using WriteRune
func ReplaceSpacesCurrent(s string) string {
var builder strings.Builder
builder.Grow(len(s))

for _, char := range s {
if char == ' ' {
builder.WriteRune('-')
} else {
builder.WriteRune(char)
}
}

return builder.String()
}

// Optimized implementation using WriteByte
func ReplaceSpacesOptimized(s string) string {
var builder strings.Builder
builder.Grow(len(s))

for i := 0; i < len(s); i++ {
if s[i] == ' ' {
builder.WriteByte('-')
} else {
builder.WriteByte(s[i])
}
}

return builder.String()
}

func BenchmarkReplaceSpacesCurrent(b *testing.B) {
testString := "this is a test string with multiple spaces in it"
b.ResetTimer()
for i := 0; i < b.N; i++ {
_ = ReplaceSpacesCurrent(testString)
}
}

func BenchmarkReplaceSpacesOptimized(b *testing.B) {
testString := "this is a test string with multiple spaces in it"
b.ResetTimer()
for i := 0; i < b.N; i++ {
_ = ReplaceSpacesOptimized(testString)
}
}

func BenchmarkReplaceSpacesCurrentLong(b *testing.B) {
testString := strings.Repeat("this is a test string ", 100)
b.ResetTimer()
for i := 0; i < b.N; i++ {
_ = ReplaceSpacesCurrent(testString)
}
}

func BenchmarkReplaceSpacesOptimizedLong(b *testing.B) {
testString := strings.Repeat("this is a test string ", 100)
b.ResetTimer()
for i := 0; i < b.N; i++ {
_ = ReplaceSpacesOptimized(testString)
}
}
