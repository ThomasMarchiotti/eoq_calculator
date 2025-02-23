# EOQ Calculator

## Descrizione del Progetto

Questo progetto implementa un calcolatore del Lotto Economico di Ordinazione (EOQ) utilizzando il framework Gin per il backend e Bootstrap per il frontend. Il calcolatore EOQ aiuta a determinare la quantità ottimale di ordine che minimizza i costi totali di inventario, inclusi i costi di ordinazione e i costi di mantenimento.

## Struttura del Progetto

- **Backend**: Implementato in Go utilizzando il framework Gin e GORM per l'interazione con il database SQLite.
- **Frontend**: Implementato in HTML e CSS con Bootstrap per lo stile e l'allineamento degli elementi.

## Codici Dettagliati

### Backend (`server.go`)

- **Strutture**:
  - `EOQInput`: Struttura per i dati di input dell'EOQ.
  - `EOQResult`: Struttura per i risultati dell'EOQ.
  - `EOQCostResult`: Struttura per i risultati dell'EOQ con i costi totali associati.
  - `EOQData`: Modello GORM per il database.

- **Funzioni**:
  - `initDB()`: Inizializza il database SQLite utilizzando GORM. Questa funzione apre una connessione al database e migra il modello `EOQData` per creare la tabella se non esiste.
  - `calculateEOQAndCosts()`: Calcola l'EOQ e i costi totali associati. Questa funzione recupera i dati dal database, calcola l'EOQ, il costo di ordinazione, il costo di mantenimento e il costo totale per ciascun record, e restituisce i risultati.
  - `calculateEOQ()`: Endpoint per calcolare l'EOQ. Questa funzione chiama `calculateEOQAndCosts()` e restituisce i risultati come JSON.
  - `addDataHandler()`: Endpoint per aggiungere i dati nel database. Questa funzione riceve i dati di input, verifica se esistono già dati per l'anno specificato, e crea un nuovo record nel database.
  - `getDataHandler()`: Endpoint per recuperare tutti i dati dal database. Questa funzione recupera tutti i record dal database e li restituisce come JSON.

### Frontend (`home.html`)

- **Struttura della Pagina**:
  - Navbar: Barra di navigazione in alto.
  - Sezione "Calcolo EOQ": Contiene il pulsante per calcolare l'EOQ e visualizzare i risultati.
  - Sezione "Aggiungi Dati": Contiene il form per aggiungere i dati di input.
  - Modal: Finestra modale per visualizzare i dati recuperati dal database e possibilità di scaricarli in formato CSV.

- **Script JavaScript**:
  - `calculateEOQButton` Event Listener: Gestisce l'evento di click sul pulsante "Calcola". Invia una richiesta POST all'endpoint `/api/calculate-eoq`, riceve i risultati e li visualizza nella sezione "Calcolo EOQ".
  - `addDataButton` Event Listener: Gestisce l'evento di click sul pulsante "Aggiungi". Invia una richiesta POST all'endpoint `/api/add-data` con i dati di input, e visualizza il risultato o l'errore nella sezione "Aggiungi Dati".
  - `getDataButton` Event Listener: Gestisce l'evento di click sul pulsante "Visualizza Dati". Invia una richiesta GET all'endpoint `/api/get-data`, riceve i dati e li visualizza nella finestra modale.
  - `downloadCSVButton` Event Listener: Gestisce l'evento di click sul pulsante "Scarica CSV". Converte i dati della tabella in formato CSV e avvia il download del file.

## Resoconto del Processo di Sviluppo

1. **Inizializzazione del Progetto**:
   - Creazione della struttura del progetto con le directory `backend` e `frontend`.
   - Configurazione del database SQLite utilizzando GORM.

2. **Implementazione del Backend**:
   - Definizione delle strutture per i dati di input e output dell'EOQ.
   - Implementazione delle funzioni per calcolare l'EOQ e i costi totali associati.
   - Creazione degli endpoint per aggiungere e recuperare i dati dal database.

3. **Implementazione del Frontend**:
   - Creazione della pagina HTML con Bootstrap per lo stile e l'allineamento degli elementi.
   - Implementazione delle funzioni JavaScript per gestire gli eventi dei pulsanti e interagire con gli endpoint del backend.

4. **Test e Debug**:
   - Test delle funzionalità del calcolatore EOQ e verifica dei risultati.
   - Debug dei problemi riscontrati durante lo sviluppo e correzione degli errori.

5. **Documentazione**:
   - Compilazione del file `README.md` con i dettagli del progetto e il resoconto del processo di sviluppo.

## Come Eseguire il Progetto

1. **Avviare il Backend**:
   ```sh
   cd backend
   go run src/server.go
   browser at http://localhost:8050/static
   ```

   oppure é possibile utilizzare il comando make.sh per buildare in automatico l''applicativo e farlo partire
   Avverrá in automatico lo spostamento e la creazione delle cartelle di lavoro, verrá creato anche in automatico un database nuovo. A cui bisognerá aggiungere i dati manualmente attraverso l''applicativo.
   (il database con i dati per testare i tre anni si trova all''interno della cartella backend [eoq.db])


   **Github Link**: https://github.com/ThomasMarchiotti/eoq_calculator