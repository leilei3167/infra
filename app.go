package infra

// App is the abstraction of an application
type App struct {
	Name string
}

type Option func(*App)

func NewApp(opts ...Option) *App {
	return &App{}
}

func (a *App) Run() error {
	return nil
}

func WithName(name string) Option {
	return func(app *App) {
		app.Name = name
	}
}
