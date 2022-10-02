# Contribuire a equation-solver

Se tu vuoi contribuire a questo progetto ci sono dei passaggi che io vorrei che tu seguissi.

## Vuoi aggiungere una nuova feature o fare un aggiustamento?
### Il nome del tuo branch
- Quando crei una nuova branch, il suo nome dovrebbe essere `equation-solver-<author>`.
- Nel caso di una patch, il suo nome dovrebbe essere `equation-solver-patch-<author>`.

### Come dovrebbero essere i commit?
- Preferibilmente frequenti. O almeno evita di fare grandi modifiche senza dei commit nel mezzo.
- Ogni commit in generale dovrebbe riferirsi ad una modifica riguardante un singolo file, a meno che tale modifica sia davvero piccola (cambiare nome di una variabile, aggiungere un commento, ...).

### Dove dovrebbero stare le tue funzioni?
- Usa `equation.go` se le funzioni devono essere accessibili all'utente.
- Usa `internal_functions.go` altrimenti.

### Come dovrebbero essere fatte le funzioni?
- Le funzioni dovrebbero avere una breve descrizione di cosa facciano, usa i commenti sopra i loro nomi:
```go
// Una semplice descrizione di cosa faccia la funzione
//
// Eventualmente puoi scrivere un esempio di input e output
func exampleFunctions() {
    return
} 
```
- Le funzione dovrebbero essere testate.

### How to create a pull request?
- Quando si crea una pull request dai una piccola descrizione delle tue modifiche o aggiunte.

### Extra
- Se parli un'altra lingua, come ad esempio spagnolo, francese, tedesco o altro, valuta magari di scrivere un file `readme.es.md` o `readme.fr.md`.

Grazieeeee!! :)