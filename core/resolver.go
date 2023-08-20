package core

import (
	"bufio"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/imroc/req/v3"
	log "github.com/toilet-tools/toilet-resolver/logger"
	"github.com/toilet-tools/toilet-resolver/utils"
)

var (
	UserAgent      string
	agentsFromFile []string

	logger = log.New("<TIME>")
)

func HandleAgents() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		utils.Error("Failed getting path: "+err.Error(), 1)
	}
	file, err := os.Open(dir + "/agents.txt")

	if err != nil {
		utils.Error("Failed opening file: "+err.Error(), 1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if !strings.Contains(scanner.Text(), "//") {
			agentsFromFile = append(agentsFromFile, scanner.Text())
		}
	}

	file.Close()
}

func ProcessUserAgent(userAgent string) {
	if len(agentsFromFile) == 0 {
		HandleAgents()
	}

	if userAgent == "random" {
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := len(agentsFromFile)
		UserAgent = agentsFromFile[rand.Intn(max-min+1)+min]
	} else {
		UserAgent = userAgent
	}
}

func Resolve(domain string, verbose bool) {
	client := req.C().SetUserAgent(UserAgent)
	for i := 1; i < 15; i++ {
		resp, err := client.R().Get("https://google.com")
		if err != nil {
			utils.Error("error: "+err.Error(), 1)
			return
		}

		if resp.IsSuccessState() { // Status code is between 200 and 299.
			logger.Success("successfully sent GET req to google.com")
		}
	}
}
