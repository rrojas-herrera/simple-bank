# 🧪 Test: `TestTransferTx`

## 🧩 Contexto general

Este test forma parte de la capa de base de datos (`package db`).  
Su propósito es **probar una transacción de transferencia (`TransferTx`) entre dos cuentas**, es decir, **mover dinero de `account1` a `account2`**.

---

## 📦 Estructura general

```go
func TestTransferTx(t *testing.T) {
    store := NewStore(testDB)
}
```

`store` es una instancia del **repositorio** (`Store`) conectado a la **base de datos de prueba (`testDB`)**.  
Este `store` tiene un método `TransferTx` que ejecuta una **transferencia como una transacción SQL** (de ahí el “Tx”).

---

## 🧑‍💼 Creación de cuentas de prueba

```go
account1 := createRandomAccount(t)
account2 := createRandomAccount(t)
```

Se crean **dos cuentas aleatorias** para hacer las pruebas:
- `account1`: cuenta emisora  
- `account2`: cuenta receptora

---

## 🔁 Configuración del test concurrente

```go
n := 5
amount := int64(10)

errs := make(chan error)
results := make(chan TransferTxResult)
```

- Se realizarán **5 transferencias concurrentes** (`n = 5`).
- Cada transferencia moverá **10 unidades monetarias** (`amount = 10`).
- Se crean **canales (`chan`)** para recolectar los resultados y errores de cada goroutine (transferencia concurrente).

---

## ⚙️ Ejecución concurrente de transferencias

```go
for i := 0; i < n; i++ {
    go func() {
        result, err := store.TransferTx(context.Background(), TransferTxParams{
            FromAccountID: account1.ID,
            ToAccountID:   account2.ID,
            Amount:        amount,
        })

        errs <- err
        results <- result
    }()
}
```

- Se lanza **una goroutine por cada transferencia**.  
- Cada goroutine ejecuta el método `TransferTx`, que internamente:
  1. **Resta** el saldo de `account1`
  2. **Suma** el saldo a `account2`
  3. **Registra las entradas (entries)** en el libro contable  
  4. **Guarda la transferencia** (`transfer`)

Los errores y resultados se envían por los canales `errs` y `results`.

💡 Esto **prueba la seguridad ante concurrencia**, es decir, verifica que no existan **condiciones de carrera** ni **inconsistencias en los saldos**.

---

## ✅ Verificación de resultados

Después se leen los resultados y se validan uno por uno:

```go
for i := 0; i < n; i++ {
    err := <-errs
    require.NoError(t, err)

    result := <-results
    require.NotEmpty(t, result)
}
```

- `require.NoError` → asegura que **no haya errores**
- `require.NotEmpty` → asegura que el **resultado no esté vacío**

---

## 🔍 Validaciones de la transferencia

Luego, el test revisa que cada elemento del resultado tenga los valores esperados:

```go
transfer := result.Transfer
require.Equal(t, account1.ID, transfer.FromAccountID)
require.Equal(t, account2.ID, transfer.ToAccountID)
require.Equal(t, amount, transfer.Amount)
```

Además, verifica que:
- `transfer.ID` y `transfer.CreatedAt` no sean cero.
- La transferencia exista realmente en la base de datos:

```go
_, err = store.GetTransfer(context.Background(), transfer.ID)
require.NoError(t, err)
```

---

## 🧾 Validaciones de las entradas contables (entries)

Cada transferencia genera **dos movimientos contables:**

1. **Entry negativo** para la cuenta emisora (`fromEntry`)
2. **Entry positivo** para la cuenta receptora (`toEntry`)

```go
fromEntry := result.FromEntry
require.Equal(t, -amount, fromEntry.Amount)

toEntry := result.ToEntry
require.Equal(t, amount, toEntry.Amount)
```

También se verifica que cada entry esté almacenado correctamente en la base de datos:

```go
_, err = store.GetEntry(context.Background(), fromEntry.ID)
require.NoError(t, err)

_, err = store.GetEntry(context.Background(), toEntry.ID)
require.NoError(t, err)
```

---

## 🧮 TODO pendiente

Al final del test hay una nota:

```go
// TODO: check accounts balances
```

Esto indica que aún falta agregar la validación para **verificar los balances finales** de las cuentas después de las transferencias.

---

## 🧠 En resumen

| Etapa | Propósito |
|--------|------------|
| **Crear cuentas** | Tener datos base para la prueba |
| **Lanzar goroutines** | Probar concurrencia en `TransferTx` |
| **Usar canales** | Sincronizar resultados y errores |
| **Validar datos** | Asegurar integridad de transferencias y movimientos |
| **Consultar DB** | Confirmar persistencia de los datos |
| **TODO final** | Verificar balances actualizados |
