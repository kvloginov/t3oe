package trigger

var triggers = make(map[string]trigger, 0)

type trigger struct {
	owner        TriggerOwner
	hitboxRadius float64
	onHitAction  func(object1 interface{}, object2 interface{})
}

// TODO: m.b. add triggers to gameObjects?
func RegisterTrigger(owner TriggerOwner, hitboxRadius float64, onHitAction func(object1 interface{}, object2 interface{})) {
	tr := trigger{
		owner:        owner,
		hitboxRadius: hitboxRadius,
		onHitAction:  onHitAction,
	}
	triggers[owner.GetName()] = tr
}

func DeleteTriggers(ownerName string) {
	delete(triggers, ownerName)
}

// CheckTriggers TODO: OMG O(n2)
func CheckTriggers() {
	for i := range triggers {
		for j := range triggers {
			if i == j {
				continue
			}
			tr1 := triggers[i]
			tr2 := triggers[j]

			if tr1.owner == nil || tr2.owner == nil {
				continue
			}

			radiuses := tr1.hitboxRadius + tr2.hitboxRadius

			if tr2.owner.GetPosition().Minus(tr1.owner.GetPosition()).Length() < radiuses {
				if tr1.onHitAction != nil {
					tr1.onHitAction(tr1.owner, tr2.owner)
				}
				if tr2.onHitAction != nil {
					tr2.onHitAction(tr2.owner, tr1.owner)
				}
			}

		}
	}
}
