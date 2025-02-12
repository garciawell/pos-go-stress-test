package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/garciawell/pos-go-stress-test/internal/web"
	"github.com/pkg/browser"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type VegetaType struct {
	Url         string
	Requests    float64
	Concurrency float64
}

type TestResults struct {
	Requests    uint64  `json:"requests"`
	MeanLatency float64 `json:"mean_latency_ms"`
	SuccessRate float64 `json:"success_rate"`
	ErrorRate   float64 `json:"error_rate"`
}

func VegetaRun(params VegetaType) {
	startTime := time.Now()
	rate := vegeta.Rate{Freq: int(params.Concurrency), Per: time.Second}
	duration := time.Duration(params.Requests/params.Concurrency) * time.Second
	target := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    params.Url,
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(target, rate, duration, "Load Test") {
		metrics.Add(res)
	}
	metrics.Close()
	disableAutoOpen := os.Getenv("DISABLE_AUTO_OPEN")

	totalDuration := time.Since(startTime).String()
	requestResult := metrics.Requests
	latencyResult := metrics.Latencies.Mean
	successResult := (float64(metrics.StatusCodes["200"]) / float64(metrics.Requests)) * 100
	errorResult := ((float64(metrics.StatusCodes["404"]) + float64(metrics.StatusCodes["500"])) / float64(metrics.Requests)) * 100
	fmt.Printf("\n🔹 URL Testada: %s\n", params.Url)
	fmt.Printf("📌 Total de Requisições: %d\n", requestResult)
	fmt.Printf("⚡ Tempo Médio de Resposta: %s\n", latencyResult)
	fmt.Printf("✅ Status Code 200: %.2f%%\n", successResult)
	fmt.Printf("❌ Status error: %.2f%%\n", errorResult)
	fmt.Printf("🕙 Tempo total:  %s\n", totalDuration)

	var filePath string

	if disableAutoOpen == "true" {
		filePath = "/app/reports/result.html"
	} else {
		filePath = "result.html"
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	result := TestResults{
		Requests:    requestResult,
		MeanLatency: float64(latencyResult.Milliseconds()),
		SuccessRate: successResult,
		ErrorRate:   errorResult,
	}

	tmpl := template.Must(template.New("dashboard").Parse(web.HtmlTemplate))
	jsonData, _ := json.Marshal(result)
	tmpl.Execute(file, string(jsonData))

	if disableAutoOpen == "true" {
		fmt.Println("🔹 Resultados disponíveis em: result.html")
		return
	}

	err = browser.OpenFile("result.html")
	if err != nil {
		fmt.Printf("Erro ao abrir o navegador: %v\n", err)
	}
}
