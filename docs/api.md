# API specification
## handles

- GET `items/id?id={}`:
    - response:
        - 200: valid json for item with id `id`
        - 400: invalid `id`
        - 404: items with id `id` not found
- GET `items/category?type={}`:
    - response:
        - 200: array of jsons for items with given `type`
        - 400: invalid `type`
        - 404: items with type `type` not found
- POST `items/new`:
    - takes json body with correct params
    - validates json
    - adds item to the storage if it does not already exist
    - response:
        - 201: item added, outputs created `id`
        - 400: invalid input json
        - 500: internal server error
- POST `items/fix?id={}`:
    - takes json body with correct params
    - validates json
    - updates item with id `id` if it exists
    - response:
        - 200: item updated, old and new json output
        - 400: invalid input json
        - 404: item with id `id` not found
- DELETE `items/delete?id={}`:
    - response:
        - 200: item deleted
        - 400: invalid `id`
        - 404: items with id `id` not found

## schemas

### generic props
- ItemName (string, required)
- ItemBulkiness (float, required)
- ItemDescription (string, optional)
- ItemCost (float, optional)
- ItemTypes ([]string, optional)

### itemType: fighting
- FightingLoad (float, required)
- FightingProficiencyCategories ([]string, optional)

### itemType: weapon (is fighting)
- WeaponStats ([]string, required)
- WeaponHands (int, required)
- WeaponDamage (string, required)
- WeaponPotential (int, required)
- WeaponAttacks (int, required)
- WeaponAuxiliary (bool, required)

### itemType: meleeWeapon (is weapon)
- MeleeAttackTypes ([]string, required)
- MeleeRange (float, required)
- MeleeThrowingRange (float, optional)

### itemType: rangedWeapon (is weapon)
- RangedReload (int, required)

### itemType: armor (is fighting)
- ArmorBonus (int, required)
- ArmorType (string, required)

### itemType: shield (is fighting)
- ShieldBonus (int, required)
- ShieldBody (string, required)

### itemType: food
- FoodTypes ([]string, required)

### itemType: consumable
- ConsumableUnit (string, optional)
- ConsumableTags ([]string, optional)

### itemType: money
- MoneyInfluenceZones ([]string, optional)
- MoneyValue (float, required)

### itemType: potion
- PotionEffect (string, required)
- PotionDuration (string, required)

### itemType: artifact
- ArtifactEffects ([]string, required)
