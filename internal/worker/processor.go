package worker

import (
	"encoding/csv"
	"excel-processor/internal/entity"
	"excel-processor/internal/repository"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Job struct {
	RowNumber int
	Student   entity.Student
}

type Processor struct {
	repo repository.StudentRepository

	// mutex & counters (for final report)
	mu sync.Mutex
	SuccessCount int
	FailCount int
}

// We need a Repository because the Worker will have to save data to the DB.
func NewProcessor(repo repository.StudentRepository) *Processor  {
	return &Processor{
		repo: repo,
	}
}


func (p *Processor) Start(filename string)  {
	startTime := time.Now()
	fmt.Printf("📂 [Manager] Mulai memproses file: %s\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("❌ Error membuka file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// --> Config Worker Pool <-- //
	const totalWorkers = 5
	jobsChannel := make(chan Job, 100)
	var wg sync.WaitGroup

	// recrut worker (spawn Goroutines)
	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go p.worker(i, jobsChannel, &wg)
	}

	rowNumber := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		rowNumber++

		// mapping csv to entity
		student := entity.Student{
			Name: record[0],
			Email: record[1],
		}

		// send to channel
		jobsChannel <- Job{
			RowNumber: rowNumber,
			Student: student,
		}

	}

	// Close Channel (Task Finished)
	close(jobsChannel)

	// wait all worker go home
	wg.Wait()

	duration := time.Since(startTime)
	fmt.Printf("✅ [Manager] SELESAI! Durasi: %s\n", duration)
	fmt.Printf("📊 Laporan Akhir: Sukses %d | Gagal %d | Total %d\n", p.SuccessCount, p.FailCount, rowNumber)
}

func (p *Processor) worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		err := p.repo.Create(job.Student)

		p.mu.Lock()
		if err != nil {
			p.FailCount++
			fmt.Printf("   ❌ [Worker %d] Gagal Baris %d: %v\n", id, job.RowNumber, err)
		} else {
			p.SuccessCount++
			// Uncomment log ini kalau mau lihat detail per baris (bakal spam kalau 10rb data)
			fmt.Printf("   ✅ [Worker %d] Sukses: %s\n", id, job.Student.Name)
		}
		p.mu.Unlock()
	}
}