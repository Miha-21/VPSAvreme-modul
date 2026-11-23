package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Miha-21/VPSAvreme-modul/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "VPSAvreme-modul",
		Usage: "obdelava meritev vremenskih postaj",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Value: -50,
				Usage: "najnižja dovoljena temperatura meritve (°C)",
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "najvišja dovoljena temperatura meritve (°C)",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "najmanjše število meritev za izračun statistike",
			},
		},
		Action: zazeni,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func zazeni(ctx context.Context, cmd *cli.Command) error {
	minTemp := cmd.Float("minTemp")
	maxTemp := cmd.Float("maxTemp")
	stMeritev := int(cmd.Int("stMeritev"))

	if minTemp > maxTemp {
		return fmt.Errorf("minTemp (%.1f) je večji od maxTemp (%.1f)", minTemp, maxTemp)
	}

	postaje := make(map[string]vreme.Postaja)
	vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
	vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")
	vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
	vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")

	dodaj := func(oznaka string, temperatura float64) {
		vreme.DodajMeritev(postaje, oznaka, temperatura, minTemp, maxTemp)
	}

	dodaj("KP01", 24.5)
	dodaj("KP01", 27)
	dodaj("KP01", 26.1)
	dodaj("KP01", 28.4)

	dodaj("KR01", 2)
	dodaj("KR01", -1.5)
	dodaj("KR01", 3.5)

	dodaj("LJ01", 18.2)
	dodaj("LJ01", 21.5)
	dodaj("LJ01", 19.8)
	dodaj("LJ01", 16.5)

	dodaj("LJ01", 55.3)
	dodaj("lj01", 43.3)

	dodaj("MB01", 12.4)

	vreme.IzpisMeritev(postaje)
	vreme.IzpisPorocila(postaje, stMeritev)
	return nil
}
