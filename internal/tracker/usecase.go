package tracker

import "github.com/google/uuid"

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	err := tracker.AddItem(Item{Name: name, ID: id})
	if err != nil {
		out.Out("error: " + err.Error())
	}
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.Items {
		out.Out(item.toString())
	}
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter item name for deleting:")
	name := in.Get()
	tracker.DeleteItem(name)
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter item id for updating:")
	id := in.Get()
	out.Out("enter a new name for updating:")
	newName := in.Get()
	err := tracker.UpdateItem(Item{Name: newName, ID: id})
	if err != nil {
		out.Out("error: " + err.Error())
	}
}

type FindUsecase struct{}

func (u FindUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter text for finding item:")
	name := in.Get()
	tracker.FindItem(name)
}
