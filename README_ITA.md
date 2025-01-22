# Strategy Pattern Application: Demonstrating Task Execution Approaches

Questo applicativo illustra l'uso del **Strategy Pattern** per dimostrare le differenze di performance tra varie modalità di esecuzione dei task. In particolare, confronta:

1. **Esecuzione Sequenziale**
2. **Esecuzione Concorrente con Goroutines e Channels**
3. **Esecuzione Concorrente con Goroutines e Mutex**

---

## Obiettivo del Progetto

L'obiettivo principale è mostrare come l'approccio di esecuzione scelto influenzi:
- **Efficienza**: Quanto velocemente i task vengono completati.
- **Parallelismo**: Quanto bene il sistema gestisce più attività simultaneamente.
- **Overhead**: Il costo di sincronizzazione o gestione di risorse condivise.

---

## Strategie Implementate

### 1. **Esecuzione Sequenziale**
- **Descrizione**: I task vengono eseguiti uno dopo l'altro, senza alcun parallelismo.
- **Caratteristiche**:
  - Ogni task attende il completamento del precedente.
  - Non sfrutta i vantaggi dell'hardware multicore.
- **Uso Reale**: Adatta a flussi di lavoro semplici dove l'ordine è prioritario.

### 2. **Goroutines con Channels**
- **Descrizione**: Ogni task viene eseguito in una goroutine separata. I **channels** coordinano e sincronizzano l'esecuzione.
- **Caratteristiche**:
  - Esecuzione parallela completa.
  - I channels regolano il flusso di comunicazione tra i task.
- **Uso Reale**: Ideale per elaborazione di task indipendenti, come richieste HTTP o job batch.

### 3. **Goroutines con Mutex**
- **Descrizione**: Ogni task è parallelo, ma l'accesso alle risorse condivise è protetto da un **mutex**.
- **Caratteristiche**:
  - Minimizza le condizioni di gara, garantendo sicurezza sui dati condivisi.
  - Introduce blocchi (lock) che possono rallentare l'esecuzione.
- **Uso Reale**: Necessario in applicazioni che manipolano risorse condivise (es. cache).

---

## Analisi delle Performance

Ecco un confronto basato su un carico di lavoro simulato di 10 task:

| Strategia               | Carico Lavoro | Tempo di Esecuzione |
|-------------------------|---------------|----------------------|
| **Sequenziale**         | 10            | 104.7116 ms         |
| **Goroutines + Channel**| 10            | 10.0814 ms          |
| **Goroutines + Mutex**  | 10            | 100.3438 ms         |

---

## Conclusioni

1. **Esecuzione Sequenziale**:
   - Metodo più lento.
   - Evitare per task indipendenti.

2. **Goroutines + Channels**:
   - La scelta migliore per parallelismo puro.
   - Adatta a sistemi ad alta concorrenza senza risorse condivise.

3. **Goroutines + Mutex**:
   - Necessaria per accessi sicuri a dati condivisi.
   - Ottimizzare per ridurre il blocco.

---

## Applicazioni Reali

- **Sequenziale**: Elaborazione di file linea per linea.
- **Goroutines + Channel**: Elaborazione di richieste API o pipeline dati.
- **Goroutines + Mutex**: Sistemi di caching o contatori condivisi.

---

Questo progetto fornisce una base solida per comprendere come scegliere una strategia di esecuzione ottimale in base al contesto e ai requisiti.
