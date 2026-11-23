# VPSAvreme-modul

Modul v jeziku Go za delo z vremenskimi postajami.

- `vreme/` – paket `vreme` (struktura `Postaja`, funkcije `DodajPostajo`, `DodajMeritev`, `IzpisMeritev`, `IzpisPorocila`)
- `main.go` – ukazna vrstica z `github.com/urfave/cli/v3`

## Zagon

```
go run . --minTemp -50 --maxTemp 50 --stMeritev 3
```

## Uvoz

```
go get github.com/Miha-21/VPSAvreme-modul@v1.0.0
```
