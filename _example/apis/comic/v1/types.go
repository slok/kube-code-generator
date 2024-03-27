package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HeroType is the type of a hero.
type HeroType string

const (
	// HeroTypeUnknown is a super hero unknown.
	HeroTypeUnknown = HeroType("unknown")
	// HeroTypeSuperHero is a super hero e.g Batman, Spiderman...
	HeroTypeSuperHero = HeroType("superhero")
	// HeroTypeAntiHero is a anti hero e.g Punisher, Deadpool...
	HeroTypeAntiHero = HeroType("antihero")
	// HeroTypeVillain is a Villain e.g Fisk, Joker...
	HeroTypeVillain = HeroType("villain")
)

// Hero represents a comic hero.
//
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="NAME",type="string",JSONPath=".spec.name"
// +kubebuilder:printcolumn:name="CITY",type="string",JSONPath=".spec.city"
// +kubebuilder:printcolumn:name="KIND",type="string",JSONPath=".spec.kind"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:singular=hero,path=heroes,shortName=he;sh,scope=Namespaced,categories=heroes;superheroes
type Hero struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HeroSpec   `json:"spec,omitempty"`
	Status HeroStatus `json:"status,omitempty"`
}

// HeroSpec is the spec of a Hero.
type HeroSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=128
	Name string `json:"name"`
	// +optional
	City string `json:"city,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=unknown;superhero;antihero;villain
	// +kubebuilder:default=unknown
	Kind HeroType `json:"kind"`
	// +optional
	BirthDate *metav1.Time `json:"birthDate"`
	// +listType=map
	// +optional
	SuperPowers []string `json:"superPowers"`
}

type HeroStatus struct {
	Moving      bool   `json:"moving"`
	CurrentCity string `json:"currentCity"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HeroList is a list of Hero resources.
type HeroList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Hero `json:"items"`
}
