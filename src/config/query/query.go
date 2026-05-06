package query

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"belajar-go/src/dto"

	"github.com/VauntDev/tqla"
)

type LoadQuery struct {
	queries  map[string]string
	filePath string
}

func NewLoadQuery(filePath string) (*LoadQuery, error) {
	ql := &LoadQuery{
		queries:  make(map[string]string),
		filePath: filePath,
	}

	if err := ql.load(); err != nil {
		return nil, err
	}

	return ql, nil
}

// Read SQL file
func (ql *LoadQuery) load() error {
	file, err := os.Open(ql.filePath)
	if err != nil {
		return fmt.Errorf("failed to open query file: %w", err)
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}(file)

	var (
		currentName string
		currentSQL  strings.Builder
		scanner     = bufio.NewScanner(file)
	)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "-- ") {
			if currentName != "" {
				ql.queries[currentName] = strings.TrimSpace(currentSQL.String())
				currentSQL.Reset()
			}
			currentName = strings.TrimPrefix(line, "-- ")
			continue
		}

		if currentName != "" {
			currentSQL.WriteString(line + "\n")
		}
	}

	if currentName != "" {
		ql.queries[currentName] = strings.TrimSpace(currentSQL.String())
	}

	return scanner.Err()
}

func (ql *LoadQuery) ExecuteTemplate(name string, filter *dto.UserFilter) (string, []any, error) {
	queryTemplate, ok := ql.Get(name)

	if !ok {
		return "", nil, fmt.Errorf("query %s not found", name)
	}

	orderClause := fmt.Sprintf(
		"ORDER BY %s %s",
		safeSort(filter.SortBy),
		safeDir(filter.SortDir),
	)

	finalQueryTemplate := strings.Replace(
		queryTemplate,
		"__ORDER_BY__",
		orderClause,
		1,
	)

	tqlaCompile, args, err := CompileTqla(finalQueryTemplate, filter)
	if err != nil {
		return "", nil, err
	}

	return tqlaCompile, args, nil
}

func (ql *LoadQuery) Get(name string) (string, bool) {
	query, ok := ql.queries[name]
	return query, ok
}

func CompileTqla(queryTemplate string, data any) (string, []any, error) {
	t, err := tqla.New(tqla.WithPlaceHolder(tqla.Dollar))
	if err != nil {
		return "", nil, fmt.Errorf("failed to create tqla: %w", err)
	}

	query, args, err := t.Compile(queryTemplate, data)
	if err != nil {
		return "", nil, fmt.Errorf("failed to compile query: %w", err)
	}

	return query, args, nil
}

func safeSort(sortBy string) string {
	switch sortBy {
	case "id":
		return "user_id"
	case "name":
		return "name"
	case "email":
		return "email"
	case "username":
		return "username"
	default:
		return "username"
	}
}

func safeDir(dir string) string {
	if strings.ToUpper(dir) == "DESC" {
		return "DESC"
	}
	return "ASC"
}
