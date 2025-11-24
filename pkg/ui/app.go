package ui

import (
	"context"
	"log"
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/jobs"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

type App struct {
	gui       *gocui.Gui
	ctx       context.Context
	cancel    context.CancelFunc
	State     state.AppState
	intentCh  chan state.Intent
	msgCh     chan state.Message
	disp      *jobs.Dispatcher
}

func Run() error {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	

	ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

	app := &App{
			gui:      g,
			ctx:      ctx,
			cancel:   cancel,
			// State:    state.NewAppState(),
			State:    state.DummyState(),
			intentCh: make(chan state.Intent, 64),
			msgCh:    make(chan state.Message, 64),
		}

	app.disp = jobs.NewDispatcher(app.msgCh)

	g.SetManagerFunc(func(g *gocui.Gui) error {
		return Layout(g, &app.State)
	})

if err := bindKeys(g, app); err != nil {
		return err
	}

	app.emit(state.LoadSession{})

	go app.loop()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Println("main loop:", err)
	}
	return nil
}

func (a *App) loop() {
	for {
		select {
		case intent := <-a.intentCh:
			effects := state.ReduceIntent(&a.State, intent)
			a.runEffects(effects)
			a.requestRender()
		case msg := <-a.msgCh:
			state.ReduceMessage(&a.State, msg)
			a.requestRender()
		case <-a.ctx.Done():
			return
		}
	}
}

func (a *App) emit(i state.Intent) {
	a.intentCh <- i
}

func (a *App) runEffects(effects []state.Effect) error {
	for _, e := range effects {
		switch e.Kind {
		case state.DispatchLoadSession:
			a.disp.DispatchLoadSession(a.ctx)
		case state.FetchProjects:
			a.disp.FetchProjects(a.ctx,)
		case state.CloseProgram:
			// a.disp.CancelAll() // optional: stop workers
			a.cancel() // stop app loop
			a.gui.Update(func(g *gocui.Gui) error { return gocui.ErrQuit })
		}
	}
	return nil
}

func (a *App) requestRender() {
	a.gui.Update(func(*gocui.Gui) error { return nil })
}
