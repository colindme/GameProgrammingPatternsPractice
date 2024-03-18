package main

import "fmt"

func main() {
	var testEmitter EventEmitter
	testEmitter.stringEvent.AddListener(Listener{ListenerID: "StringListener1", InvokedFunction: func(params ...any) { fmt.Println("In StringListener1: ", params[0]) }})
	testEmitter.stringEvent.AddListener(Listener{ListenerID: "StringListener2", InvokedFunction: func(params ...any) { fmt.Println("In StringListener2: ", params[0]) }})

	testEmitter.stringEvent.Invoke("Testing string")
	testEmitter.stringEvent.RemoveListener("StringListener1")
	fmt.Println("------------------------------")
	testEmitter.stringEvent.Invoke("Testing string again")
	fmt.Println("------------------------------")

	testEmitter.damageEvent.AddListener(Listener{ListenerID: "PlayerManager",
		InvokedFunction: func(params ...any) {
			if len(params) > 0 {
				damageEvent, ok := params[0].(DamageEvent)
				if ok {
					fmt.Printf("Received damage event for player %s from enemy %s for amount %f\n", damageEvent.Name, damageEvent.InstigatorName, damageEvent.DamageAmount)
				}
			}
		}})
	testEmitter.damageEvent.AddListener(Listener{ListenerID: "PlayerUI",
		InvokedFunction: func(params ...any) {
			if len(params) > 0 {
				damageEvent, ok := params[0].(DamageEvent)
				if ok {
					fmt.Printf("Received damage event for player %s reducing health in UI by %f\n", damageEvent.Name, damageEvent.DamageAmount)
				}
			}
		}})
	testEmitter.damageEvent.Invoke(DamageEvent{Name: "PlayerOne", InstigatorName: "Level 1 goblin", DamageAmount: 5})
}

type EventEmitter struct {
	stringEvent ObservableEvent
	damageEvent ObservableEvent
}

type DamageEvent struct {
	Name           string
	InstigatorName string
	DamageAmount   float32
}

/* Observer / Listener structs and funcs */
type Listener struct {
	ListenerID      string
	InvokedFunction func(params ...any)
}

type ObservableEvent struct {
	invokeList []Listener
}

func (o *ObservableEvent) AddListener(listener Listener) {
	o.invokeList = append(o.invokeList, listener)
}

func (o *ObservableEvent) RemoveListener(listenerID string) {
	for k, v := range o.invokeList {
		if v.ListenerID == listenerID {
			o.invokeList = append(o.invokeList[0:k], o.invokeList[k+1:]...)
			return
		}
	}
}

func (o *ObservableEvent) Invoke(params ...any) {
	for _, v := range o.invokeList {
		v.InvokedFunction(params...)
	}
}
