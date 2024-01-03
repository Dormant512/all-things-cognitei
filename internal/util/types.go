package util

import "time"

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
	// util props
	UtilValidationMap map[string]bool
	UtilCreatedAt     string

	// generic props
	ItemName        *string  `json:"itemName"`
	ItemBulkiness   *float32 `json:"itemBulkiness"`
	ItemDescription *string  `json:"itemDescription"` // optional
	ItemCost        *float32 `json:"itemCost"`        // optional
	ItemTypes       []string `json:"itemTypes"`       // optional

	// itemType: fighting
	FightingLoad                  *float32 `json:"load"`
	FightingProficiencyCategories []string `json:"proficiencyCategories"` // optional

	// itemType: weapon (is fighting)
	WeaponStats     []string `json:"weaponStats"`
	WeaponHands     *int     `json:"weaponHands"`
	WeaponDamage    *string  `json:"weaponDamage"`
	WeaponPotential *int     `json:"weaponPotential"`
	WeaponAttacks   *int     `json:"weaponAttacks"`
	WeaponAuxiliary *bool    `json:"weaponAuxiliary"`

	// itemType: meleeWeapon (is weapon)
	MeleeAttackTypes   []string `json:"meleeAttackTypes"`
	MeleeRange         *float32 `json:"meleeRange"`
	MeleeThrowingRange *float32 `json:"meleeThrowingRange"` // optional

	// itemType: rangedWeapon (is weapon)
	RangedReload *int `json:"rangedReload"`

	// itemType: armor (is fighting)
	ArmorBonus *int    `json:"armorBonus"`
	ArmorType  *string `json:"armorType"`

	// itemType: shield (is fighting)
	ShieldBonus *int    `json:"shieldBonus"`
	ShieldBody  *string `json:"shieldBody"`

	// itemType: food
	FoodTypes []string `json:"foodTypes"`

	// itemType: consumable
	ConsumableUnit *string  `json:"consumableUnit"` // optional
	ConsumableTags []string `json:"consumableTags"` // optional

	// itemType: money
	MoneyInfluenceZones []string `json:"moneyInfluenceZones"` // optional
	MoneyValue          *float32 `json:"moneyValue"`

	// itemType: potion
	PotionEffect   *string `json:"potionEffect"`
	PotionDuration *string `json:"potionDuration"`

	// itemType: artifact
	ArtifactEffects []string `json:"artifactEffects"`
}

// TODO: check if BindJSON binds empty array to nil or []string{} - IMPORTANT!

func (m *MegaItem) isValid(itemType string) bool {
	var cond bool
	switch itemType {
	case "item":
		cond = m.ItemName != nil &&
			m.ItemBulkiness != nil
	case "fighting":
		cond = m.FightingLoad != nil
	case "weapon":
		cond = m.WeaponStats != nil &&
			m.WeaponHands != nil &&
			m.WeaponDamage != nil &&
			m.WeaponPotential != nil &&
			m.WeaponAttacks != nil &&
			m.WeaponAuxiliary != nil
	case "meleeWeapon":
		cond = m.MeleeAttackTypes != nil &&
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
		cond = m.FoodTypes != nil
	case "consumable":
		cond = true
	case "money":
		cond = m.MoneyInfluenceZones != nil &&
			m.MoneyValue != nil
	case "potion":
		cond = m.PotionEffect != nil &&
			m.PotionDuration != nil
	case "artifact":
		cond = m.ArtifactEffects != nil
	}
	return cond
}

func (m *MegaItem) checkAllTypes() error {
	ans := make(map[string]bool)
	itemOk := m.isValid("item")
	if !itemOk {
		return InvalidMegaItemError{}
	}
	if m.ItemTypes == nil {
		ans["item"] = true
		m.UtilValidationMap = ans
		return nil
	}
	for _, itemType := range m.ItemTypes {
		if !ItemTypes[itemType] {
			// type does not exist
			return InvalidMegaItemTypeError{InvalidType: itemType}
		}
		ans[itemType] = itemOk && m.isValid(itemType)
	}
	// extra dependency checks
	if !ans["fighting"] {
		ans["weapon"] = false
		ans["armor"] = false
		ans["shield"] = false
	}
	if !ans["weapon"] {
		ans["meleeWeapon"] = false
		ans["rangedWeapon"] = false
	}
	m.UtilValidationMap = ans
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
