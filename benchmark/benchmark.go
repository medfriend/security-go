package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Parámetros configurables
var (
	concurrency int
	requests    int
	url         string
	timeout     int
)

func init() {
	flag.IntVar(&concurrency, "c", 10, "Número de peticiones concurrentes")
	flag.IntVar(&requests, "n", 100, "Número total de peticiones")
	flag.StringVar(&url, "url", "http://localhost:9000/health", "URL a probar")
	flag.IntVar(&timeout, "t", 10, "Timeout en segundos")
	flag.Parse()
}

type RequestResult struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}

func main() {
	fmt.Printf("Iniciando benchmark con %d solicitudes concurrentes para un total de %d peticiones a %s\n",
		concurrency, requests, url)

	// Crear cliente HTTP con timeout
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	// Canal para resultados
	resultChan := make(chan RequestResult, requests)

	// Control de concurrencia
	semaphore := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	startTime := time.Now()

	// Lanzar peticiones
	for i := 0; i < requests; i++ {
		wg.Add(1)
		semaphore <- struct{}{} // Adquirir semáforo
		go func(requestNum int) {
			defer wg.Done()
			defer func() { <-semaphore }() // Liberar semáforo al finalizar

			start := time.Now()
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				resultChan <- RequestResult{StatusCode: 0, Duration: 0, Error: err}
				return
			}

			resp, err := client.Do(req)
			duration := time.Since(start)

			if err != nil {
				resultChan <- RequestResult{StatusCode: 0, Duration: duration, Error: err}
				return
			}
			defer resp.Body.Close()

			resultChan <- RequestResult{
				StatusCode: resp.StatusCode,
				Duration:   duration,
				Error:      nil,
			}

			if requestNum%10 == 0 {
				fmt.Printf("Completadas %d/%d peticiones\n", requestNum, requests)
			}
		}(i + 1)
	}

	// Esperar a que todas las peticiones terminen y cerrar el canal de resultados
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Procesar resultados
	var totalDuration time.Duration
	statusCodes := make(map[int]int)
	errors := 0
	requestCount := 0

	var min, max time.Duration
	min = time.Hour // Valor inicial grande

	for result := range resultChan {
		requestCount++
		if result.Error != nil {
			errors++
			continue
		}

		statusCodes[result.StatusCode]++
		totalDuration += result.Duration

		if result.Duration < min {
			min = result.Duration
		}
		if result.Duration > max {
			max = result.Duration
		}
	}

	// Calcular estadísticas
	totalTime := time.Since(startTime)
	avgDuration := totalDuration / time.Duration(requestCount-errors)
	rps := float64(requestCount) / totalTime.Seconds()

	// Imprimir resultados
	fmt.Println("\n--- Resultados del Benchmark ---")
	fmt.Printf("Tiempo total: %v\n", totalTime)
	fmt.Printf("Peticiones realizadas: %d\n", requestCount)
	fmt.Printf("Peticiones exitosas: %d\n", requestCount-errors)
	fmt.Printf("Peticiones fallidas: %d\n", errors)
	fmt.Printf("Peticiones por segundo: %.2f req/s\n", rps)
	fmt.Printf("Tiempo mínimo de respuesta: %v\n", min)
	fmt.Printf("Tiempo máximo de respuesta: %v\n", max)
	fmt.Printf("Tiempo promedio de respuesta: %v\n", avgDuration)
	fmt.Println("\nDistribución de códigos de estado:")
	for code, count := range statusCodes {
		fmt.Printf("  %d: %d (%.1f%%)\n", code, count, float64(count)/float64(requestCount)*100)
	}
}