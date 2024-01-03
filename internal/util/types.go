package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var ItemTypes = map[string]bool{
	"item":         true,
	"fighting":     true,
	"weapon":       true,
	"meleeWeapon":  true,
	"rangedWeapon": true,
	"armor":        true,
	"shield":       true,
	"food":         true,
	"consumable":   true,
	"money":        true,
	"potion":       true,
	"artifact":     true,
}

type MegaItem struct {
	Id *primitive.ObjectID `json:",omitempty" bson:"_id,omitempty"`

	// util props
	UtilValidTypes []string `json:"utilValidTypes" bson:"utilValidTypes"`
	UtilCreatedAt  string   `json:"utilCreatedAt" bson:"utilCreatedAt"`

	// generic props
	ItemName        *string  `json:"itemName" bson:"itemName"`
	ItemBulkiness   *float32 `json:"itemBulkiness" bson:"itemBulkiness"`
	ItemDescription *string  `json:"itemDescription,omitempty" bson:"itemDescription,omitempty"` // opt
	ItemCost        *float32 `json:"itemCost,omitempty" bson:"itemCost,omitempty"`               // opt
	ItemTypes       []string `json:"itemTypes,omitempty" bson:"itemTypes,omitempty"`             // opt

	// itemType: fighting
	FightingLoad                  *float32 `json:"fightingLoad,omitempty" bson:"fightingLoad,omitempty"`
	FightingProficiencyCategories []string `json:"proficiencyCategories,omitempty" bson:"fightingProficiencyCategories,omitempty"` // opt

	// itemType: weapon (is fighting)
	WeaponStats     []string `json:"weaponStats,omitempty" bson:"weaponStats,omitempty"`
	WeaponHands     *int     `json:"weaponHands,omitempty" bson:"weaponHands,omitempty"`
	WeaponDamage    *string  `json:"weaponDamage,omitempty" bson:"weaponDamage,omitempty"`
	WeaponPotential *int     `json:"weaponPotential,omitempty" bson:"weaponPotential,omitempty"`
	WeaponAttacks   *int     `json:"weaponAttacks,omitempty" bson:"weaponAttacks,omitempty"`
	WeaponAuxiliary *bool    `json:"weaponAuxiliary,omitempty" bson:"weaponAuxiliary,omitempty"`

	// itemType: meleeWeapon (is weapon)
	MeleeAttackTypes   []string `json:"meleeAttackTypes,omitempty" bson:"meleeAttackTypes,omitempty"`
	MeleeRange         *float32 `json:"meleeRange,omitempty" bson:"meleeRange,omitempty"`
	MeleeThrowingRange *float32 `json:"meleeThrowingRange,omitempty" bson:"meleeThrowingRange,omitempty"` // opt

	// itemType: rangedWeapon (is weapon)
	RangedReload *int `json:"rangedReload,omitempty" bson:"rangedReload,omitempty"`

	// itemType: armor (is fighting)
	ArmorBonus *int    `json:"armorBonus,omitempty" bson:"armorBonus,omitempty"`
	ArmorType  *string `json:"armorType,omitempty" bson:"armorType,omitempty"`

	// itemType: shield (is fighting)
	ShieldBonus *int    `json:"shieldBonus,omitempty" bson:"shieldBonus,omitempty"`
	ShieldBody  *string `json:"shieldBody,omitempty" bson:"shieldBody,omitempty"`

	// itemType: food
	FoodTypes []string `json:"foodTypes,omitempty" bson:"foodTypes,omitempty"`

	// itemType: consumable
	ConsumableUnit *string  `json:"consumableUnit,omitempty" bson:"consumableUnit,omitempty"` // opt
	ConsumableTags []string `json:"consumableTags,omitempty" bson:"consumableTags,omitempty"` // opt

	// itemType: money
	MoneyInfluenceZones []string `json:"moneyInfluenceZones,omitempty" bson:"moneyInfluenceZones,omitempty"` // opt
	MoneyValue          *float32 `json:"moneyValue,omitempty" bson:"moneyValue,omitempty"`

	// itemType: potion
	PotionEffect   *string `json:"potionEffect,omitempty" bson:"potionEffect,omitempty"`
	PotionDuration *string `json:"potionDuration,omitempty" bson:"potionDuration,omitempty"`

	// itemType: artifact
	ArtifactEffects []string `json:"artifactEffects,omitempty" bson:"artifactEffects,omitempty"`
}

func (m *MegaItem) checkValidityOnType(itemType string) error {
	var cond bool
	switch itemType {
	case "item":
		cond = m.ItemName != nil &&
			len(*m.ItemName) > 0 &&
			m.ItemBulkiness != nil
	case "fighting":
		cond = m.FightingLoad != nil
	case "weapon":
		cond = m.WeaponStats != nil &&
			len(m.WeaponStats) > 0 &&
			m.WeaponHands != nil &&
			m.WeaponDamage != nil &&
			m.WeaponPotential != nil &&
			m.WeaponAttacks != nil &&
			m.WeaponAuxiliary != nil
	case "meleeWeapon":
		cond = m.MeleeAttackTypes != nil &&
			len(m.MeleeAttackTypes) > 0 &&
			m.MeleeRange != nil
	case "rangedWeapon":
		cond = m.RangedReload != nil
	case "armor":
		cond = m.ArmorBonus != nil &&
			m.ArmorType != nil
	case "shield":
		cond = m.ShieldBonus != nil &&
			m.ShieldBody != nil
	case "food":
		cond = m.FoodTypes != nil &&
			len(m.FoodTypes) > 0
	case "consumable":
		cond = true
	case "money":
		cond = m.MoneyValue != nil
	case "potion":
		cond = m.PotionEffect != nil &&
			m.PotionDuration != nil
	case "artifact":
		cond = m.ArtifactEffects != nil &&
			len(m.ArtifactEffects) > 0
	}
	if !cond {
		return MegaItemIsNotOfSaidTypeError{SaidType: itemType}
	}
	return nil
}

func (m *MegaItem) checkAllTypes() error {
	validTypesMap := make(map[string]bool)
	err := m.checkValidityOnType("item")
	if err != nil {
		return err
	}
	validTypesMap["item"] = true
	if m.ItemTypes == nil {
		m.UtilValidTypes = MapToSlice(validTypesMap)
		return nil
	}
	for _, itemType := range m.ItemTypes {
		if !ItemTypes[itemType] {
			// type does not exist
			return MegaItemTypeDoesNotExistError{InvalidType: itemType}
		}
		err = m.checkValidityOnType(itemType)
		if err != nil {
			return err
		}
		validTypesMap[itemType] = true
	}
	// extra dependency checks
	if !validTypesMap["fighting"] {
		for _, fightPlus := range []string{"weapon", "armor", "shield"} {
			if validTypesMap[fightPlus] {
				return MegaItemIsNotOfSaidTypeError{SaidType: fightPlus}
			}
		}
	}
	if !validTypesMap["weapon"] {
		for _, weaponPlus := range []string{"meleeWeapon", "rangedWeapon"} {
			if validTypesMap[weaponPlus] {
				return MegaItemIsNotOfSaidTypeError{SaidType: weaponPlus}
			}
		}
	}
	m.UtilValidTypes = MapToSlice(validTypesMap)
	return nil
}

func (m *MegaItem) FillUtil() error {
	err := m.checkAllTypes()
	if err != nil {
		return err
	}
	m.UtilCreatedAt = time.Now().Format(time.DateTime)
	return nil
}

type CategoriesToFetch struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

func (c *CategoriesToFetch) MakeValid() error {
	inSet, err := SliceToCatSet(c.Include)
	if err != nil {
		return err
	}
	exSet, err := SliceToCatSet(c.Exclude)
	if err != nil {
		return err
	}
	for k := range inSet {
		if exSet[k] {
			inSet[k] = false
		}
	}
	c.Include = MapToSlice(inSet)
	c.Exclude = MapToSlice(exSet)
	return nil
}
