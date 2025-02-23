package main

import (
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Definizione della struttura per EOQInput
type EOQInput struct {
	Demand      float64 `json:"demand"`
	SetupCost   float64 `json:"setup_cost"`
	HoldingCost float64 `json:"holding_cost"`
	Year        int     `json:"year"`
}

// Struttura per EOQResult
type EOQResult struct {
	Year int     `json:"year"`
	EOQ  float64 `json:"eoq"`
}

// Struttura per EOQCostResult
type EOQCostResult struct {
	ID                uint    `json:"id"`
	Year              int     `json:"year"`
	EOQ               float64 `json:"eoq"`
	CostoOrdinazione  float64 `json:"costo_ordinazione"`
	CostoMantenimento float64 `json:"costo_mantenimento"`
	CostoTotale       float64 `json:"costo_totale"`
}

// Modello GORM per il database
type EOQData struct {
	ID          uint    `json:"id"`
	Demand      float64 `json:"demand"`
	SetupCost   float64 `json:"setup_cost"`
	HoldingCost float64 `json:"holding_cost"`
	Year        int     `json:"year"`
}

var db *gorm.DB

func initDB() error {
	var err error
	// Inizializza il database SQLite usando GORM
	db, err = gorm.Open(sqlite.Open("eoq.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	// Puoi fare eventuali operazioni di migrazione qui
	db.AutoMigrate(&EOQData{})
	return nil
}

// Funzione per calcolare l'EOQ e i costi totali associati
func calculateEOQAndCosts() ([]EOQCostResult, error) {
	var data []EOQData
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	var results []EOQCostResult
	for _, record := range data {
		if record.SetupCost > 0 && record.HoldingCost > 0 {
			eoq := math.Sqrt((2.0 * record.Demand * record.SetupCost) / record.HoldingCost)
			costoOrdinazione := (record.Demand / eoq) * record.SetupCost
			costoMantenimento := (eoq / 2) * record.HoldingCost
			costoTotale := costoOrdinazione + costoMantenimento
			results = append(results, EOQCostResult{
				ID:                record.ID,
				Year:              record.Year,
				EOQ:               eoq,
				CostoOrdinazione:  costoOrdinazione,
				CostoMantenimento: costoMantenimento,
				CostoTotale:       costoTotale,
			})
			// Stampa i valori processati nel ciclo
			// log.Printf("ID: %d, Year: %d, EOQ: %f, Costo Ordinazione: %f, Costo Mantenimento: %f, Costo Totale: %f\n",
			// 	record.ID, record.Year, eoq, costoOrdinazione, costoMantenimento, costoTotale)
		} else {
			results = append(results, EOQCostResult{
				ID:                record.ID,
				Year:              record.Year,
				EOQ:               0,
				CostoOrdinazione:  0,
				CostoMantenimento: 0,
				CostoTotale:       0,
			})
			// Stampa i valori processati nel ciclo
			log.Printf("ID: %d, Year: %d, EOQ: %f (SetupCost o HoldingCost sono 0)\n", record.ID, record.Year, 0.0)
		}
	}

	return results, nil
}

// Funzione per calcolare l'EOQ
func calculateEOQ(c *gin.Context) {
	results, err := calculateEOQAndCosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore durante il recupero dei dati"})
		return
	}
	c.JSON(http.StatusOK, results)
}

// Funzione per aggiungere i dati nel database
func addDataHandler(c *gin.Context) {
	// Usa il database qui
	if db == nil {
		c.JSON(500, gin.H{"error": "Database non inizializzato"})
		return
	}
	var input EOQInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Dati non validi"})
		return
	}
	// Imposta l'anno corrente se non specificato
	if input.Year == 0 {
		input.Year = time.Now().Year()
	}
	// Verifica se esistono già dati per l'anno specificato
	var existingData EOQData
	if err := db.Where("year = ?", input.Year).First(&existingData).Error; err == nil {
		c.JSON(400, gin.H{"error": "Dati per l'anno specificato già esistenti"})
		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(500, gin.H{"error": "Errore durante la verifica dei dati esistenti"})
		return
	}
	// Crea un record nel database
	if err := db.Create(&EOQData{
		Demand:      input.Demand,
		SetupCost:   input.SetupCost,
		HoldingCost: input.HoldingCost,
		Year:        input.Year,
	}).Error; err != nil {
		c.JSON(500, gin.H{"error": "Errore durante l'inserimento"})
		return
	}
	c.JSON(200, gin.H{"message": "Dati inseriti con successo"})
}

// Funzione per recuperare tutti i dati dal database
func getDataHandler(c *gin.Context) {
	var data []EOQData
	if err := db.Find(&data).Error; err != nil {
		c.JSON(500, gin.H{"error": "Errore durante il recupero dei dati"})
		return
	}
	c.JSON(200, data)
}

func main() {
	// Creazione del router Gin
	r := gin.Default()

	if err := initDB(); err != nil {
		panic("Database connection failed: " + err.Error())
	}

	// Migrazione del database per creare la tabella (se non esiste)
	db.AutoMigrate(&EOQData{})

	// Definire le route
	r.POST("/api/calculate-eoq", calculateEOQ)
	r.POST("/api/add-data", addDataHandler)
	r.GET("/api/get-data", getDataHandler)

	// Verifica che la directory dei file statici esista
	staticDir := "./frontend"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		panic("Directory dei file statici non trovata: " + staticDir)
	}

	// Servire i file statici (frontend) dalla sottodirectory /static
	r.StaticFile("/static", staticDir+"/home.html")

	// Impostare il database come dipendenza
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Avvio del server sulla porta 8080
	if err := r.Run(":8050"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
