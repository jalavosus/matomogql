package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/urfave/cli/v3"
	"golang.org/x/sync/errgroup"

	"github.com/jalavosus/matomogql/utils/gql"

	_ "github.com/joho/godotenv/autoload"
)

var portFlag = cli.IntFlag{
	Name:     "port",
	Aliases:  []string{"p"},
	Value:    6700,
	Required: false,
}

func runServer(s *http.Server) func() error {
	return func() error {
		log.Println("connect to http://localhost" + s.Addr + "/ for GraphQL playground")

		if err := s.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}

			return err
		}

		return nil
	}
}

func appMain(ctx context.Context, cmd *cli.Command) error {
	serverPort := cmd.Int(portFlag.Name)
	addr := ":" + strconv.Itoa(serverPort)

	handler := gql.MakeServer(true)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	var g *errgroup.Group
	g, ctx = errgroup.WithContext(ctx)

	s := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	g.Go(runServer(s))

	for {
		select {
		case <-ctx.Done():
			errs := make([]error, 0, 3)

			if err := s.Shutdown(ctx); err != nil {
				errs = append(errs, err)
			}
			if err := ctx.Err(); err != nil {
				if !errors.Is(err, context.Canceled) {
					errs = append(errs, err)
				}
			}
			if err := g.Wait(); err != nil {
				if !errors.Is(err, context.Canceled) {
					errs = append(errs, err)
				}
			}

			return errors.Join(errs...)
		}
	}
}

func main() {
	app := &cli.Command{
		Flags: []cli.Flag{
			&portFlag,
		},
		Action: appMain,
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
