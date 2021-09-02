package ui

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type food struct {
	id    int
	label string
}

type foodList struct {
	app.Compo
	foods []food
}

func NewFoodList() *foodList {
	return &foodList{}
}

func (l *foodList) OnMount(ctx app.Context) {
	l.initFood(ctx)
}

func (l *foodList) initFood(ctx app.Context) {
	l.foods = []food{
		{
			id:    1,
			label: "French fries",
		},
		{
			id:    2,
			label: "Pasta",
		},
		{
			id:    3,
			label: "Rice",
		},
		{
			id:    4,
			label: "Steak",
		},
		{
			id:    5,
			label: "Fish",
		},
	}
	l.Update()
}

func (l *foodList) Render() app.UI {
	return app.P().Body(
		app.H3().Text("Dealing with List"),
		app.Range(l.foods).Slice(func(i int) app.UI {
			f := l.foods[i]

			return app.Button().
				Text("Remove "+f.label).
				OnClick(l.eat(f), f.id)
		}),
		app.Button().
			Text("Reset").
			OnClick(l.reset),
	)
}

func (l *foodList) eat(f food) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		for i, fi := range l.foods {
			if fi.id == f.id {
				// Removing food item:
				copy(l.foods[i:], l.foods[i+1:])
				l.foods = l.foods[:len(l.foods)-1]
				l.Update()
				return
			}
		}
	}
}

func (l *foodList) reset(ctx app.Context, e app.Event) {
	l.initFood(ctx)
}
