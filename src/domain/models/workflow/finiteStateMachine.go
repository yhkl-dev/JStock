package workflow

import (
	"JStock/src/domain/models/repos"
	"fmt"
)

type State struct {
	StateName string `json:"state_name" gorm:"state_name"`
	StateCode string `json:"state_code" gorm:"state_code"`
}

func NewState(name, code string) *State {
	return &State{StateName: name, StateCode: code}
}

type character string

type FARule struct {
	State     *State             `json:"current_state" gorm:"embedded"`
	Character character          `json:"character" gorm:"character"`
	NextState *State             `json:"next_state" gorm:"embedded"`
	Repo      repos.IFARuleModel `gorm:"-"`
}

func NewFARule(s *State, c character, ns *State) *FARule {
	return &FARule{State: s, Character: c, NextState: ns}
}

func (s *FARule) Follow() *State {
	return s.NextState
}

func (s *FARule) AppliesTo(state *State, char character) bool {
	return s.State.StateCode == state.StateCode && s.Character == char
}

func (s *FARule) String() string {
	return fmt.Sprintf("#FARule %s -- %s --> %s", s.State.StateName, s.Character, s.NextState.StateName)
}

type DFARuleBook struct {
	Rules []*FARule `json:"rules"`
}

func NewDFARuleBook(rules []*FARule) *DFARuleBook {
	return &DFARuleBook{Rules: rules}
}

func (s *DFARuleBook) RuleFor(state *State, char character) *FARule {
	r := &FARule{}
	for _, rule := range s.Rules {
		if rule.AppliesTo(state, char) {
			r = rule
			break
		}
	}
	return r
}

func (s *DFARuleBook) NextState(state *State, char character) *FARule {
	return s.RuleFor(state, char)
}
