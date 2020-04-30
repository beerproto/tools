package beerXML

import (
	"strconv"
	"strings"

	"github.com/RossMerr/beerjson.go"
	beerXML "github.com/RossMerr/beerxml.go"
	"github.com/beerproto/tools/beerproto"
)

func MapToProto(i *beerXML.Recipe) *beerproto.Recipe {
	output := &beerproto.Recipe{
		Mashes:                   []*beerproto.MashProcedureType{},
		Recipes:                  []*beerproto.RecipeType{},
		MiscellaneousIngredients: []*beerproto.MiscellaneousType{},
		Styles:                   []*beerproto.StyleType{},
		Fermentations:            []*beerproto.FermentationProcedureType{},
		Boil:                     []*beerproto.BoilProcedureType{},
		Version:                  float64(i.Version),
		Fermentables:             []*beerproto.FermentableType{},
		// TimingObject:             ToProtoTimingType(i.TimingObject),
		Cultures:     []*beerproto.CultureInformation{},
		Equipments:   []*beerproto.EquipmentType{},
		Packaging:    []*beerproto.PackagingProcedureType{},
		HopVarieties: []*beerproto.VarietyInformation{},
		Profiles:     []*beerproto.WaterBase{},
	}

	if i.Mash != nil {
		output.Mashes = append(output.Mashes, ToProtoMashProcedureType(i.Mash))
	}

	if i.Style != nil {
		output.Recipes = append(output.Recipes, ToProtoRecipeType(i, i.Style))
	}
	if i.Style != nil {
		output.Styles = append(output.Styles, ToProtoStyleType(i.Style))
	}

	if i.Yeasts != nil {
		for _, culture := range i.Yeasts.Yeast {
			output.Cultures = append(output.Cultures, ToProtoCultureInformation(culture))
		}
	}

	if i.Equipment != nil {
		output.Equipments = append(output.Equipments, ToProtoEquipmentType(i.Equipment))
	}
	//
	// for _, ingredients := range i.MiscellaneousIngredients {
	// 	output.MiscellaneousIngredients = append(output.MiscellaneousIngredients, ToProtoMiscellaneousType(&ingredients))
	// }
	//

	//
	// for _, fermentation := range i.Fermentations {
	// 	output.Fermentations = append(output.Fermentations, ToProtoFermentationProcedureType(&fermentation))
	// }
	//
	// for _, boil := range i.Boil {
	// 	output.Boil = append(output.Boil, ToProtoBoilProcedureType(&boil))
	// }
	//
	// for _, fermentable := range i.Fermentables {
	// 	output.Fermentables = append(output.Fermentables, ToProtoFermentableType(&fermentable))
	// }
	//

	//

	//
	// for _, packaging := range i.Packaging {
	// 	output.Packaging = append(output.Packaging, ToProtoPackagingProcedureType(&packaging))
	// }
	//
	// for _, hopVariety := range i.HopVarieties {
	// 	output.HopVarieties = append(output.HopVarieties, ToProtoVarietyInformation(&hopVariety))
	// }
	//
	// for _, profile := range i.Profiles {
	// 	output.Profiles = append(output.Profiles, ToProtoWaterBase(&profile))
	// }

	return output
}

func ToProtoRecipeType(i *beerXML.Recipe, s *beerXML.Style) *beerproto.RecipeType {
	return &beerproto.RecipeType{
		// Efficiency:          ToProtoEfficiencyType(&i.Efficiency),
		Style: ToProtoRecipeStyleType(i.Style),
		//IbuEstimate:         ToProtoIBUEstimateType(s.IBUMAX),
		ColorEstimate:       ToProtoColorType(s.Colormax),
		// BeerPH:              ToProtoAcidityType(i.BeerPH),
		Name:                i.Name,
		Type:                ToProtoRecipeTypeType(i.Type),
		Coauthor:            i.AsstBrewer,
		OriginalGravity:     ToProtoGravityType(s.OGMAX),
		FinalGravity:        ToProtoGravityType(s.FGMAX),
		Carbonation:         float64(s.Carbmax),
		Fermentation:        ToProtoFermentationProcedureType(i),
		Author:              i.Brewer,
		//Ingredients:         ToProtoIngredientsType(&i.Ingredients),
		Mash:                ToProtoMashProcedureType(i.Mash),
		//Packaging:           ToProtoPackagingProcedureType(i.Packaging),
		//Boil:                ToProtoBoilProcedureType(i.Boil),
		Taste:               ToProtoTasteType(s.Profile),
		//CaloriesPerPint:     UseFloat(i.CaloriesPerPint),
		//Created:             created,
		//BatchSize:           ToProtoVolumeType(&i.BatchSize),
		Notes:               i.Notes,
		AlcoholByVolume:     ToProtoPercentType(s.ABVMAX),
		//ApparentAttenuation: ToProtoPercentType(i.ApparentAttenuation),
	}
}

func ToProtoTasteType(i string) *beerproto.TasteType {
	return &beerproto.TasteType{
		Notes:  i,
	}
}
func ToProtoEquipmentItemType(i *beerjson.EquipmentItemType) *beerproto.EquipmentItemType {
	if i == nil {
		return nil
	}

	return &beerproto.EquipmentItemType{
		BoilRatePerHour:     ToProtoVolumeType(i.BoilRatePerHour),
		Type:                UseString(i.KeyType),
		Form:                ToProtoEquipmentBaseForm(i.EquipmentBaseForm),
		DrainRatePerMinute:  ToProtoVolumeType(i.DrainRatePerMinute),
		SpecificHeat:        ToProtoSpecificHeatType(i.SpecificHeat),
		GrainAbsorptionRate: ToProtoSpecificVolumeType(i.GrainAbsorptionRate),
		Name:                UseString(i.Name),
		MaximumVolume:       ToProtoVolumeType(i.MaximumVolume),
		Weight:              ToProtoMassType(i.Weight),
		Loss:                ToProtoVolumeType(&i.Loss),
	}
}

func ToProtoWaterBase(i *beerjson.WaterBase) *beerproto.WaterBase {
	if i == nil {
		return nil
	}

	return &beerproto.WaterBase{
		Calcium:     ToProtoConcentrationType(&i.Calcium),
		Nitrite:     ToProtoConcentrationType(i.Nitrite),
		Chloride:    ToProtoConcentrationType(&i.Chloride),
		Name:        i.Name,
		Potassium:   ToProtoConcentrationType(i.Potassium),
		Carbonate:   ToProtoConcentrationType(i.Carbonate),
		Iron:        ToProtoConcentrationType(i.Iron),
		Flouride:    ToProtoConcentrationType(i.Flouride),
		Sulfate:     ToProtoConcentrationType(&i.Sulfate),
		Magnesium:   ToProtoConcentrationType(&i.Magnesium),
		Producer:    UseString(i.Producer),
		Bicarbonate: ToProtoConcentrationType(&i.Bicarbonate),
		Nitrate:     ToProtoConcentrationType(i.Nitrate),
		Sodium:      ToProtoConcentrationType(&i.Sodium),
	}
}

func ToProtoVarietyInformation(i *beerjson.VarietyInformation) *beerproto.VarietyInformation {
	return &beerproto.VarietyInformation{
		Inventory:   ToProtoHopInventoryType(i.Inventory),
		Type:        ToProtoVarietyInformationType(i.VarietyInformationType),
		OilContent:  ToProtoOilContentType(i.OilContent),
		PercentLost: ToProtoPercentType(i.PercentLost),
		ProductId:   UseString(i.ProductId),
		AlphaAcid:   ToProtoPercentType(i.AlphaAcid),
		BetaAcid:    ToProtoPercentType(i.BetaAcid),
		Name:        UseString(i.Name),
		Origin:      UseString(i.Origin),
		Substitutes: UseString(i.Substitutes),
		Year:        UseString(i.Year),
		Form:        ToProtoHopVarietyBaseForm(i.HopVarietyBaseForm),
		Producer:    UseString(i.Producer),
		Notes:       UseString(i.Notes),
	}
}

func ToProtoOilContentType(i *beerjson.OilContentType) *beerproto.OilContentType {
	if i == nil {
		return nil
	}

	return &beerproto.OilContentType{
		Polyphenols:        ToProtoPercentType(i.Polyphenols),
		TotalOilMlPer_100G: UseFloat(i.TotalOilMlPer100g),
		Farnesene:          ToProtoPercentType(i.Farnesene),
		Limonene:           ToProtoPercentType(i.Limonene),
		Nerol:              ToProtoPercentType(i.Nerol),
		Geraniol:           ToProtoPercentType(i.Geraniol),
		BPinene:            ToProtoPercentType(i.BPinene),
		Linalool:           ToProtoPercentType(i.Linalool),
		Caryophyllene:      ToProtoPercentType(i.Caryophyllene),
		Cohumulone:         ToProtoPercentType(i.Cohumulone),
		Xanthohumol:        ToProtoPercentType(i.Xanthohumol),
		Humulene:           ToProtoPercentType(i.Humulene),
		Myrcene:            ToProtoPercentType(i.Myrcene),
		Pinene:             ToProtoPercentType(i.Pinene),
	}
}

func ToProtoVarietyInformationType(i *beerjson.VarietyInformationType) beerproto.VarietyInformation_VarietyInformationType {
	if i == nil {
		return beerproto.VarietyInformation_NULL_VarietyInformationType
	}

	unit := beerproto.VarietyInformation_VarietyInformationType_value[strings.ToUpper(string(*i))]
	return beerproto.VarietyInformation_VarietyInformationType(unit)
}

func ToProtoHopInventoryType(i *beerjson.HopInventoryType) *beerproto.HopInventoryType {
	if i == nil {
		return nil
	}

	hopInventoryType := &beerproto.HopInventoryType{}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		hopInventoryType.Amount = &beerproto.HopInventoryType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		hopInventoryType.Amount = &beerproto.HopInventoryType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return hopInventoryType
}

func ToProtoEquipmentType(i *beerXML.Equipment) *beerproto.EquipmentType {
	if i == nil {
		return nil
	}

	equipmentItemType := []*beerproto.EquipmentItemType{}


	equipmentItemType = append(equipmentItemType, &beerproto.EquipmentItemType{
		SpecificHeat: 	 ToProtoSpecificHeatType(i.Tunspecificheat),

	})

	// for _, item := range i.EquipmentItems {
	// 	equipmentItemType = append(equipmentItemType, ToProtoEquipmentItemType(&item))
	// }

	return &beerproto.EquipmentType{
		Name:           i.Name,
		EquipmentItems: equipmentItemType,
	}
}

func ToProtoSpecificHeatType(i float32) *beerproto.SpecificHeatType {
	if i == 0 {
		return nil
	}

	return &beerproto.SpecificHeatType{
		Value: float64(i),
		Unit:  beerproto.SpecificHeatType_CALGC,
	}
}

func ToProtoSpecificHeatUnitType(i *beerjson.SpecificHeatUnitType) beerproto.SpecificHeatType_SpecificHeatUnitType {
	if i == nil {
		return beerproto.SpecificHeatType_NULL
	}

	unit := beerproto.SpecificHeatType_SpecificHeatUnitType_value[strings.ToUpper(string(*i))]
	return beerproto.SpecificHeatType_SpecificHeatUnitType(unit)
}

func ToProtoEquipmentBaseForm(i *beerjson.EquipmentBaseForm) beerproto.EquipmentItemType_EquipmentBaseForm {
	if i == nil {
		return beerproto.EquipmentItemType_NULL
	}

	switch *i {
	case beerjson.EquipmentBaseForm_HLT:
		return beerproto.EquipmentItemType_HLT
	case beerjson.EquipmentBaseForm_MashTun:
		return beerproto.EquipmentItemType_MASH_TUN
	case beerjson.EquipmentBaseForm_LauterTun:
		return beerproto.EquipmentItemType_LAUTER_TUN
	case beerjson.EquipmentBaseForm_BrewKettle:
		return beerproto.EquipmentItemType_BREW_KETTLE
	case beerjson.EquipmentBaseForm_Fermenter:
		return beerproto.EquipmentItemType_FERMENTER
	case beerjson.EquipmentBaseForm_AgingVessel:
		return beerproto.EquipmentItemType_AGING_VESSEL
	case beerjson.EquipmentBaseForm_PackagingVessel:
		return beerproto.EquipmentItemType_PACKAGING_VESSEL
	}

	return beerproto.EquipmentItemType_NULL
}

func ToProtoCultureInformation(i beerXML.Yeast) *beerproto.CultureInformation {
	return &beerproto.CultureInformation{
		Form:             ToProtoCultureBaseForm(i.Form),
		Producer:         i.Laboratory,
		TemperatureRange: ToProtoTemperatureRangeType(i.Maxtemperature, i.Mintemperature),
		Notes:            i.Notes,
		BestFor:          i.Bestfor,
		//Inventory:        ToProtoCultureInventoryType(i.Inventory),
		ProductId:        i.Productid,
		Name:             i.Name,
		//AlcoholTolerance: ToProtoPercentType(i.AlcoholTolerance),
		//Glucoamylase:     UseBool(i.Glucoamylase),
		Type:             ToProtoCultureBaseType(i.Type),
		Flocculation:     ToProtoQualitativeRangeType(i.Flocculation),
		AttenuationRange: ToProtoPercentRangeType(i.Attenuation),
		MaxReuse:         int32(i.Maxreuse),
	}
}

func ToProtoZymocide(i *beerjson.Zymocide) *beerproto.Zymocide {
	if i == nil {
		return nil
	}
	return &beerproto.Zymocide{
		No1:     UseBool(i.No1),
		No2:     UseBool(i.No2),
		No28:    UseBool(i.No28),
		Klus:    UseBool(i.Klus),
		Neutral: UseBool(i.Neutral),
	}
}
func ToProtoQualitativeRangeType(i string) beerproto.CultureInformation_QualitativeRangeType {
	if i == "" {
		return beerproto.CultureInformation_NULL_QualitativeRangeType
	}
	if unit, ok := beerproto.CultureInformation_QualitativeRangeType_value[strings.ReplaceAll(strings.ToUpper(i), " ", "_")]; ok {
		return beerproto.CultureInformation_QualitativeRangeType(unit)
	}
	return beerproto.CultureInformation_NULL_QualitativeRangeType
}

func ToProtoCultureBaseType(i string) beerproto.CultureBaseType {
	if i == "" {
		return beerproto.CultureBaseType_NULL_CultureBaseType
	}

	if unit, ok := beerproto.CultureBaseType_value[strings.ToUpper(i)]; ok {
		return beerproto.CultureBaseType(unit)
	}

	return beerproto.CultureBaseType_OTHER_CultureBaseType
}

func ToProtoCultureInventoryType(i *beerjson.CultureInventoryType) *beerproto.CultureInventoryType {
	if i == nil {
		return nil
	}
	return &beerproto.CultureInventoryType{
		Liquid:  ToProtoVolumeType(i.Liquid),
		Dry:     ToProtoMassType(i.Dry),
		Slant:   ToProtoVolumeType(i.Slant),
		Culture: ToProtoVolumeType(i.Culture),
	}
}

func ToProtoTemperatureRangeType(max, min float64) *beerproto.TemperatureRangeType {
	if max == 0 && min == 0 {
		return nil
	}

	return &beerproto.TemperatureRangeType{
		Minimum: ToProtoTemperatureType(min),
		Maximum: ToProtoTemperatureType(max),
	}
}

func ToProtoFermentableType(i *beerjson.FermentableType) *beerproto.FermentableType {
	if i == nil {
		return nil
	}
	return &beerproto.FermentableType{
		MaxInBatch:     ToProtoPercentType(i.MaxInBatch),
		RecommendMash:  UseBool(i.RecommendMash),
		Protein:        ToProtoPercentType(i.Protein),
		ProductId:      UseString(i.ProductId),
		GrainGroup:     ToProtoGrainGroup(i.FermentableBaseGrainGroup),
		Yield:          ToProtoYieldType(i.Yield),
		Type:           ToProtoFermentableBaseType(i.FermentableBaseType),
		Producer:       UseString(i.Producer),
		AlphaAmylase:   UseFloat(i.AlphaAmylase),
		Color:          ToProtoColorType(i.Color),
		Name:           UseString(i.Name),
		DiastaticPower: ToProtoDiastaticPowerType(i.DiastaticPower),
		Moisture:       ToProtoPercentType(i.Moisture),
		Origin:         UseString(i.Origin),
		Inventory:      ToProtoFermentableInventoryType(i.Inventory),
		KolbachIndex:   UseFloat(i.KolbachIndex),
		Notes:          UseString(i.Notes),
	}
}

func ToProtoFermentableInventoryType(i *beerjson.FermentableInventoryType) *beerproto.FermentableInventoryType {
	if i == nil {
		return nil
	}

	fermentableInventoryType := &beerproto.FermentableInventoryType{}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		fermentableInventoryType.Amount = &beerproto.FermentableInventoryType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		fermentableInventoryType.Amount = &beerproto.FermentableInventoryType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return fermentableInventoryType
}

func ToProtoDiastaticPowerType(i *beerjson.DiastaticPowerType) *beerproto.DiastaticPowerType {
	if i == nil {
		return nil
	}
	return &beerproto.DiastaticPowerType{
		Value: i.Value,
		Unit:  ToProtoDiastaticPowerUnitType(&i.Unit),
	}
}

func ToProtoDiastaticPowerUnitType(i *beerjson.DiastaticPowerUnitType) beerproto.DiastaticPowerType_DiastaticPowerUnitType {
	if i == nil {
		return beerproto.DiastaticPowerType_NULL
	}

	unit := beerproto.DiastaticPowerType_DiastaticPowerUnitType_value[strings.ToUpper(string(*i))]
	return beerproto.DiastaticPowerType_DiastaticPowerUnitType(unit)
}

func ToProtoStyleType(i *beerXML.Style) *beerproto.StyleType {
	if i == nil {
		return nil
	}

	categoryNumber := int32(0)
	if no, err := strconv.ParseInt(i.Categorynumber, 10, 32); err == nil {
		categoryNumber = int32(no)
	}

	return &beerproto.StyleType{
		Aroma:                        i.Profile,
		Ingredients:                  i.Ingredients,
		CategoryNumber:               categoryNumber,
		Notes:                        i.Notes,
		Flavor:                       i.Profile,
		Mouthfeel:                    i.Profile,
		FinalGravity:                 ToProtoGravityRangeType(i.FGMAX, i.FGMIN),
		StyleGuide:                   i.Styleguide,
		Color:                        ToProtoColorRangeType(i.Colormax, i.Colormin),
		OriginalGravity:              ToProtoGravityRangeType(i.OGMAX, i.OGMIN),
		Examples:                     i.Examples,
		Name:                         i.Name,
		Carbonation:                  ToProtoCarbonationRangeType(i.Carbmax, i.Carbmin),
		AlcoholByVolume:              ToProtoPercentRangeType(i.ABVMAX, i.ABVMIN),
		InternationalBitternessUnits: ToProtoBitternessRangeType(i.IBUMAX, i.IBUMIN),
		//Appearance:                   UseString(i.Appearance),
		Category:                     i.Category,
		StyleLetter:                  i.Styleletter,
		Type:                         ToProtoStyleType_StyleCategories(i.Type),
		//OverallImpression:            UseString(i.OverallImpression),
	}
}


func ToProtoColorRangeType(max, min float32) *beerproto.ColorRangeType {
	if max == 0 && min == 0 {
		return nil
	}

	return &beerproto.ColorRangeType{
		Minimum: ToProtoColorType(min),
		Maximum: ToProtoColorType(max),
	}
}

func ToProtoGravityRangeType(max, min float32) *beerproto.GravityRangeType {
	if max == 0 && min == 0 {
		return nil
	}

	return &beerproto.GravityRangeType{
		Minimum: ToProtoGravityType(min),
		Maximum: ToProtoGravityType(max),
	}
}


func ToProtoStyleType_StyleCategories(i string) beerproto.StyleType_StyleCategories {
	if i == "" {
		return beerproto.StyleType_NULL
	}

	switch strings.ToLower(i) {
	case "lager", "ale", "wheat":
		return beerproto.StyleType_BEER
	}

	if unit, ok := beerproto.StyleType_StyleCategories_value[strings.ToUpper(i)]; ok {
		return beerproto.StyleType_StyleCategories(unit)
	}
	return beerproto.StyleType_OTHER
}

func ToProtoBitternessRangeType(max, min float32) *beerproto.BitternessRangeType {
	if max == 0 && min == 0 {
		return nil
	}

	return &beerproto.BitternessRangeType{
		Minimum: ToProtoBitternessType(min),
		Maximum: ToProtoBitternessType(max),
	}
}

func ToProtoBitternessType(i float32) *beerproto.BitternessType {
	if i == 0 {
		return nil
	}
	return &beerproto.BitternessType{
		Value: float64(i),
		Unit:  beerproto.BitternessType_IBUs,
	}
}

func ToProtoBitternessUnitType(i *beerjson.BitternessUnitType) beerproto.BitternessType_BitternessUnitType {
	if i == nil {
		return beerproto.BitternessType_NULL
	}
	unit := beerproto.BitternessType_BitternessUnitType_value[strings.ToUpper(string(*i))]
	return beerproto.BitternessType_BitternessUnitType(unit)
}

func ToProtoPercentRangeType(max, min float32) *beerproto.PercentRangeType {
	if max == 0 && min == 0 {
		return nil
	}

	return &beerproto.PercentRangeType{
		Minimum: ToProtoPercentType(min),
		Maximum: ToProtoPercentType(max),
	}
}

func ToProtoCarbonationRangeType(max, min float32) *beerproto.CarbonationRangeType {
	if max == 0 && min == 0 {
		return nil
	}

	return &beerproto.CarbonationRangeType{
		Minimum: ToProtoCarbonationType(min),
		Maximum: ToProtoCarbonationType(max),
	}
}
func ToProtoCarbonationType(i float32) *beerproto.CarbonationType {
	if i == 0 {
		return nil
	}
	return &beerproto.CarbonationType{
		Value: float64(i),
		Unit:  beerproto.CarbonationType_VOLS,
	}
}

func ToProtoIngredientsType(i *beerjson.IngredientsType) *beerproto.IngredientsType {
	if i == nil {
		return nil
	}

	miscellaneousAdditions := make([]*beerproto.MiscellaneousAdditionType, 0)
	cultureAdditions := make([]*beerproto.CultureAdditionType, 0)
	waterAdditions := make([]*beerproto.WaterAdditionType, 0)
	fermentableAdditions := make([]*beerproto.FermentableAdditionType, 0)
	hopAdditions := make([]*beerproto.HopAdditionType, 0)

	for _, misc := range i.MiscellaneousAdditions {
		miscellaneousAdditions = append(miscellaneousAdditions, ToProtoMiscellaneousAdditionType(&misc))
	}
	for _, culture := range i.CultureAdditions {
		cultureAdditions = append(cultureAdditions, ToProtoCultureAdditionType(&culture))
	}
	for _, water := range i.WaterAdditions {
		waterAdditions = append(waterAdditions, ToProtoWaterAdditionType(&water))
	}
	for _, fermentable := range i.FermentableAdditions {
		fermentableAdditions = append(fermentableAdditions, ToProtoFermentableAdditionType(&fermentable))
	}
	for _, hop := range i.HopAdditions {
		hopAdditions = append(hopAdditions, ToProtoHopAdditionType(&hop))
	}
	return &beerproto.IngredientsType{
		MiscellaneousAdditions: miscellaneousAdditions,
		CultureAdditions:       cultureAdditions,
		WaterAdditions:         waterAdditions,
		FermentableAdditions:   fermentableAdditions,
		HopAdditions:           hopAdditions,
	}
}

func ToProtoHopAdditionType(i *beerjson.HopAdditionType) *beerproto.HopAdditionType {
	if i == nil {
		return nil
	}

	hopAdditionType := &beerproto.HopAdditionType{
		BetaAcid:  ToProtoPercentType(i.BetaAcid),
		Producer:  UseString(i.Producer),
		Origin:    UseString(i.Origin),
		Year:      UseString(i.Year),
		Form:      ToProtoHopVarietyBaseForm(i.HopVarietyBaseForm),
		Timing:    ToProtoTimingType(&i.Timing),
		Name:      UseString(i.Name),
		ProductId: UseString(i.ProductId),
		AlphaAcid: ToProtoPercentType(i.AlphaAcid),
	}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		hopAdditionType.Amount = &beerproto.HopAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		hopAdditionType.Amount = &beerproto.HopAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return hopAdditionType
}

func ToProtoHopVarietyBaseForm(i *beerjson.HopVarietyBaseForm) beerproto.HopVarietyBaseForm {
	if i == nil {
		return beerproto.HopVarietyBaseForm_NULL_HopVarietyBaseForm
	}

	switch *i {
	case beerjson.HopVarietyBaseForm_Extract:
		return beerproto.HopVarietyBaseForm_EXTRACT_HopVarietyBaseForm
	case beerjson.HopVarietyBaseForm_Leaf:
		return beerproto.HopVarietyBaseForm_LEAF
	case beerjson.HopVarietyBaseForm_LeafWet:
		return beerproto.HopVarietyBaseForm_LEAFWET
	case beerjson.HopVarietyBaseForm_Pellet:
		return beerproto.HopVarietyBaseForm_PELLET
	case beerjson.HopVarietyBaseForm_Powder:
		return beerproto.HopVarietyBaseForm_POWDER
	case beerjson.HopVarietyBaseForm_Plug:
		return beerproto.HopVarietyBaseForm_PLUG
	}

	unit := beerproto.HopVarietyBaseForm_value[strings.ToUpper(string(*i))]
	return beerproto.HopVarietyBaseForm(unit)
}

func ToProtoFermentableAdditionType(i *beerjson.FermentableAdditionType) *beerproto.FermentableAdditionType {
	if i == nil {
		return nil
	}

	fermentableAdditionType := &beerproto.FermentableAdditionType{
		Type:       ToProtoFermentableBaseType(i.FermentableBaseType),
		Origin:     UseString(i.Origin),
		GrainGroup: ToProtoGrainGroup(i.FermentableBaseGrainGroup),
		Yield:      ToProtoYieldType(i.Yield),
		Color:      ToProtoColorType(i.Color),
		Name:       UseString(i.Name),
		Producer:   UseString(i.Producer),
		ProductId:  UseString(i.ProductId),
		Timing:     ToProtoTimingType(i.Timing),
	}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		fermentableAdditionType.Amount = &beerproto.FermentableAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		fermentableAdditionType.Amount = &beerproto.FermentableAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return fermentableAdditionType
}

func ToProtoYieldType(i *beerjson.YieldType) *beerproto.YieldType {
	if i == nil {
		return nil
	}

	return &beerproto.YieldType{
		FineGrind:            ToProtoPercentType(i.FineGrind),
		CoarseGrind:          ToProtoPercentType(i.CoarseGrind),
		FineCoarseDifference: ToProtoPercentType(i.FineCoarseDifference),
		Potential:            ToProtoGravityType(i.Potential),
	}
}

func ToProtoGrainGroup(i *beerjson.FermentableBaseGrainGroup) beerproto.GrainGroup {
	if i == nil {
		return beerproto.GrainGroup_NULL_GrainGroup
	}

	unit := beerproto.GrainGroup_value[strings.ToUpper(string(*i))]
	return beerproto.GrainGroup(unit)
}

func ToProtoFermentableBaseType(i *beerjson.FermentableBaseType) beerproto.FermentableBaseType {
	if i == nil {
		return beerproto.FermentableBaseType_NULL_FermentableBaseType
	}

	unit := beerproto.FermentableBaseType_value[strings.ToUpper(string(*i))]
	return beerproto.FermentableBaseType(unit)
}

func ToProtoWaterAdditionType(i *beerjson.WaterAdditionType) *beerproto.WaterAdditionType {
	if i == nil {
		return nil
	}

	return &beerproto.WaterAdditionType{
		Flouride:    ToProtoConcentrationType(i.Flouride),
		Sulfate:     ToProtoConcentrationType(i.Sulfate),
		Producer:    UseString(i.Producer),
		Nitrate:     ToProtoConcentrationType(i.Nitrate),
		Nitrite:     ToProtoConcentrationType(i.Nitrite),
		Chloride:    ToProtoConcentrationType(i.Chloride),
		Amount:      ToProtoVolumeType(i.Amount),
		Name:        UseString(i.Name),
		Potassium:   ToProtoConcentrationType(i.Potassium),
		Magnesium:   ToProtoConcentrationType(i.Magnesium),
		Iron:        ToProtoConcentrationType(i.Iron),
		Bicarbonate: ToProtoConcentrationType(i.Bicarbonate),
		Calcium:     ToProtoConcentrationType(i.Calcium),
		Carbonate:   ToProtoConcentrationType(i.Carbonate),
		Sodium:      ToProtoConcentrationType(i.Sodium),
	}
}

func ToProtoConcentrationType(i *beerjson.ConcentrationType) *beerproto.ConcentrationType {
	if i == nil {
		return nil
	}

	return &beerproto.ConcentrationType{
		Value: i.Value,
		Unit:  ToProtoConcentrationUnitType(i.Unit),
	}
}

func ToProtoConcentrationUnitType(i beerjson.ConcentrationUnitType) beerproto.ConcentrationType_ConcentrationUnitType {
	unit := beerproto.ConcentrationType_ConcentrationUnitType_value[strings.ToUpper(string(i))]
	return beerproto.ConcentrationType_ConcentrationUnitType(unit)
}

func ToProtoCultureAdditionType(i *beerjson.CultureAdditionType) *beerproto.CultureAdditionType {
	if i == nil {
		return nil
	}

	cultureAdditionType := &beerproto.CultureAdditionType{
		Form:              ToProtoCultureBaseForm(i.CultureBaseForm),
		ProductId:         UseString(i.ProductId),
		Name:              UseString(i.Name),
		CellCountBillions: UseInt(i.CellCountBillions),
		TimesCultured:     UseInt(i.TimesCultured),
		Producer:          UseString(i.Producer),
		Type:              ToProtoCultureBaseType(i.CultureBaseType),
		Attenuation:       ToProtoPercentType(i.Attenuation),
		Timing:            ToProtoTimingType(i.Timing),
	}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		cultureAdditionType.Amount = &beerproto.CultureAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if unit, ok := i.Amount.(*beerjson.UnitType); ok {
		cultureAdditionType.Amount = &beerproto.CultureAdditionType_Unit{
			Unit: ToProtoUnitType(unit),
		}
	}
	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		cultureAdditionType.Amount = &beerproto.CultureAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return cultureAdditionType
}

func ToProtoCultureBaseForm(i string) beerproto.CultureBaseForm {
	if i == "" {
		return beerproto.CultureBaseForm_NULL_CultureBaseForm
	}

	unit := beerproto.CultureBaseForm_value[strings.ToUpper(i)]
	return beerproto.CultureBaseForm(unit)
}

func ToProtoMiscellaneousAdditionType(i *beerjson.MiscellaneousAdditionType) *beerproto.MiscellaneousAdditionType {
	if i == nil {
		return nil
	}

	miscellaneousAdditionType := &beerproto.MiscellaneousAdditionType{
		Name:      UseString(i.Name),
		Producer:  UseString(i.Producer),
		Timing:    ToProtoTimingType(i.Timing),
		ProductId: UseString(i.ProductId),
		Type:      ToProtoMiscellaneousBaseType(i.MiscellaneousBaseType),
	}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		miscellaneousAdditionType.Amount = &beerproto.MiscellaneousAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if unit, ok := i.Amount.(*beerjson.UnitType); ok {
		miscellaneousAdditionType.Amount = &beerproto.MiscellaneousAdditionType_Unit{
			Unit: ToProtoUnitType(unit),
		}
	}
	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		miscellaneousAdditionType.Amount = &beerproto.MiscellaneousAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return miscellaneousAdditionType
}

func ToProtoUnitType(i *beerjson.UnitType) *beerproto.UnitType {
	if i == nil {
		return nil
	}
	return &beerproto.UnitType{
		Value: i.Value,
		Unit:  ToProtoUnitUnitType(i.Unit),
	}
}

func ToProtoUnitUnitType(i beerjson.UnitUnitType) beerproto.UnitType_UnitUnitType {
	unit := beerproto.UnitType_UnitUnitType_value[strings.ToUpper(string(i))]
	return beerproto.UnitType_UnitUnitType(unit)
}

func ToProtoMassType(i *beerjson.MassType) *beerproto.MassType {
	if i == nil {
		return nil
	}
	return &beerproto.MassType{
		Value: i.Value,
		Unit:  ToProtoMassUnitType(i.Unit),
	}
}

func ToProtoMassUnitType(i beerjson.MassUnitType) beerproto.MassType_MassUnitType {
	unit := beerproto.MassType_MassUnitType_value[strings.ToUpper(string(i))]
	return beerproto.MassType_MassUnitType(unit)
}

func ToProtoMiscellaneousBaseType(i *beerjson.MiscellaneousBaseType) beerproto.MiscellaneousBaseType {
	unit := beerproto.MiscellaneousBaseType_value[strings.ToUpper(string(*i))]
	return beerproto.MiscellaneousBaseType(unit)
}

func ToProtoTimingType(i *beerjson.TimingType) *beerproto.TimingType {
	if i == nil {
		return nil
	}
	return &beerproto.TimingType{
		Time:            ToProtoTimeType(i.Time),
		Duration:        ToProtoTimeType(i.Duration),
		Continuous:      UseBool(i.Continuous),
		SpecificGravity: ToProtoGravityType(i.SpecificGravity),
		Ph:              ToProtoAcidityType(i.PH),
		Step:            UseInt(i.Step),
		Use:             ToProtoUseType(i.Use),
	}
}

func ToProtoUseType(i *beerjson.UseType) beerproto.TimingType_UseType {
	if i == nil {
		return beerproto.TimingType_NULL
	}
	unit := beerproto.TimingType_UseType_value[strings.ToUpper(string(*i))]
	return beerproto.TimingType_UseType(unit)
}

func ToProtoFermentationProcedureType(i *beerXML.Recipe) *beerproto.FermentationProcedureType {
	steps := make([]*beerproto.FermentationStepType, 0)

	steps = append(steps, &beerproto.FermentationStepType{
		Name:           "Primary",
		EndTemperature: ToProtoTemperatureType(i.Primarytemp),
		StepTime: &beerproto.TimeType{
			Value: float64(i.Primaryage),
			Unit:  beerproto.TimeType_DAY,
		},
		StartTemperature: ToProtoTemperatureType(i.Primarytemp),
	})

	steps = append(steps, &beerproto.FermentationStepType{
		Name:           "Secondary",
		EndTemperature: ToProtoTemperatureType(i.Secondarytemp),
		StepTime: &beerproto.TimeType{
			Value: float64(i.Secondaryage),
			Unit:  beerproto.TimeType_DAY,
		},
		StartTemperature: ToProtoTemperatureType(i.Secondarytemp),
	})


	steps = append(steps, &beerproto.FermentationStepType{
		Name:           "Tertiary",
		EndTemperature: ToProtoTemperatureType(i.Tertiarytemp),
		StepTime: &beerproto.TimeType{
			Value: float64(i.Tertiaryage),
			Unit:  beerproto.TimeType_DAY,
		},
		StartTemperature: ToProtoTemperatureType(i.Tertiarytemp),
	})

	return &beerproto.FermentationProcedureType{
		Notes:             i.Notes,
		Name:              i.Name,
		FermentationSteps: steps,
	}
}

func ToProtoGravityType(i float32) *beerproto.GravityType {
	return &beerproto.GravityType{
		Value: float64(i),
		Unit:  beerproto.GravityType_SG,
	}
}

func ToProtoGravityUnitType(i beerjson.GravityUnitType) beerproto.GravityType_GravityUnitType {
	unit := beerproto.GravityType_GravityUnitType_value[strings.ToUpper(string(i))]
	return beerproto.GravityType_GravityUnitType(unit)
}

func ToProtoRecipeTypeType(i string) beerproto.RecipeType_RecipeTypeType {
	switch i {
	case "Extract":
		return beerproto.RecipeType_EXTRACT
	case "Partial Mash":
		return beerproto.RecipeType_PARTIAL_MASH
	case "All Grain":
		return beerproto.RecipeType_ALL_GRAIN
	}
	return beerproto.RecipeType_NULL
}

func ToProtoColorType(i float32) *beerproto.ColorType {
	return &beerproto.ColorType{
		Value: float64(i),
		Unit: beerproto.ColorType_SRM,
	}
}

func ToProtoColorUnitType(i beerjson.ColorUnitType) beerproto.ColorType_ColorUnitType {
	unit := beerproto.ColorType_ColorUnitType_value[strings.ToUpper(string(i))]

	return beerproto.ColorType_ColorUnitType(unit)
}

func ToProtoIBUEstimateType(i *beerjson.IBUEstimateType) *beerproto.IBUEstimateType {
	if i == nil {
		return nil
	}

	return &beerproto.IBUEstimateType{
		Method: ToProtoIBUMethodType(i.Method),
	}
}

func ToProtoIBUMethodType(i *beerjson.IBUMethodType) beerproto.IBUEstimateType_IBUMethodType {
	if i == nil {
		return beerproto.IBUEstimateType_NULL
	}
	unit := beerproto.IBUEstimateType_IBUMethodType_value[string(*i)]
	return beerproto.IBUEstimateType_IBUMethodType(unit)
}

func ToProtoRecipeStyleType(i *beerXML.Style) *beerproto.RecipeStyleType {
	if i == nil {
		return nil
	}

	categoryNumber := int32(0)
	if no, err := strconv.ParseInt(i.Categorynumber, 10, 32); err == nil {
		categoryNumber = int32(no)
	}

	return &beerproto.RecipeStyleType{
		Type:           ToProtoRecipeStyleType_StyleCategories(i.Type),
		Name:           i.Name,
		Category:       i.Category,
		CategoryNumber: categoryNumber,
		StyleGuide:     i.Styleguide,
		StyleLetter:    i.Styleletter,
	}
}

func ToProtoRecipeStyleType_StyleCategories(i string) beerproto.RecipeStyleType_StyleCategories {
	if i == "" {
		return beerproto.RecipeStyleType_NULL
	}

	switch strings.ToLower(i) {
	case "lager", "ale", "wheat":
		return beerproto.RecipeStyleType_BEER
	case "mead":
		return beerproto.RecipeStyleType_MEAD
	case "mixed":
		return beerproto.RecipeStyleType_OTHER
	case "cider":
		return beerproto.RecipeStyleType_cider
	}

	unit := beerproto.RecipeStyleType_StyleCategories_value[strings.ToUpper(i)]
	return beerproto.RecipeStyleType_StyleCategories(unit)
}

func ToProtoEfficiencyType(i *beerjson.EfficiencyType) *beerproto.EfficiencyType {
	if i == nil {
		return nil
	}

	return &beerproto.EfficiencyType{
		Conversion: ToProtoPercentType(i.Conversion),
		Lauter:     ToProtoPercentType(i.Lauter),
		Mash:       ToProtoPercentType(i.Mash),
		Brewhouse:  ToProtoPercentType(&i.Brewhouse),
	}
}

func ToProtoPercentType(i float32) *beerproto.PercentType {
	return &beerproto.PercentType{
		Value: float64(i),
		Unit:  beerproto.PercentType_PERCENT_SIGN,
	}
}

func ToProtoPercentUnitType(i beerjson.PercentUnitType) beerproto.PercentType_PercentUnitType {
	unit := beerproto.PercentType_PercentUnitType_value[strings.ToUpper(string(i))]
	return beerproto.PercentType_PercentUnitType(unit)
}

func ToProtoMashProcedureType(i *beerXML.Mash) *beerproto.MashProcedureType {

	mashSteps := []*beerproto.MashStepType{}

	for _, step := range i.Mashsteps.Mashstep {
		mashSteps = append(mashSteps, ToProtoMashStepType(step))
	}

	if i.Spargetemp > 0 {
		sparge := ToProtoMashStepType(beerXML.MashStep{
			Type:     "Sparge",
			Steptemp: i.Spargetemp,
			Endtemp:  i.Spargetemp,
		})
		sparge.StartPH = ToProtoAcidityType(i.PH)
		mashSteps = append(mashSteps, sparge)
	}

	if i.Tuntemp > 0 {
		mashSteps = append(mashSteps, ToProtoMashStepType(beerXML.MashStep{
			Type:     "Tun",
			Steptemp: i.Tuntemp,
			Endtemp:  i.Tuntemp,
		}))
	}

	return &beerproto.MashProcedureType{
		Name:             i.Name,
		Notes:            i.Notes,
		GrainTemperature: ToProtoTemperatureType(i.Graintemp),
		MashSteps:        mashSteps,
	}
}

func ToProtoMashStepType(i beerXML.MashStep) *beerproto.MashStepType {
	t := ToProtoMashStepTypeType(i.Type)
	mashStep := &beerproto.MashStepType{
		StepTime:        ToProtoTimeType(i.Steptime),
		RampTime:        ToProtoTimeType(i.Ramptime),
		EndTemperature:  ToProtoTemperatureType(i.Endtemp),
		Description:     i.Description,
		Name:            i.Name,
		Type:            t,
		StepTemperature: ToProtoTemperatureType(i.Steptemp),
		WaterGrainRatio: ToProtoSpecificVolumeType(i.Infuseamount),
	}

	if t == beerproto.MashStepType_INFUSION{
		mashStep.Amount = &beerproto.VolumeType{
			Value: float64(i.Infuseamount),
			Unit: beerproto.VolumeType_L,
		}
	}

	return mashStep
}

func ToProtoVolumeType(i *beerjson.VolumeType) *beerproto.VolumeType {
	if i == nil {
		return nil
	}
	return &beerproto.VolumeType{
		Value: i.Value,
		Unit:  ToProtoVolumeUnitType(i.Unit),
	}
}

func ToProtoVolumeUnitType(i beerjson.VolumeUnitType) beerproto.VolumeType_VolumeUnitType {
	unit := beerproto.VolumeType_VolumeUnitType_value[strings.ToUpper(string(i))]
	return beerproto.VolumeType_VolumeUnitType(unit)
}

func ToProtoSpecificVolumeType(i float32) *beerproto.SpecificVolumeType {
	return &beerproto.SpecificVolumeType{
		Value: float64(i),
		Unit:  beerproto.SpecificVolumeType_LG,
	}
}

func ToProtoSpecificVolumeUnitType(i beerjson.SpecificVolumeUnitType) beerproto.SpecificVolumeType_SpecificVolumeUnitType {
	unit := beerproto.SpecificVolumeType_SpecificVolumeUnitType_value[strings.ToUpper(string(i))]

	switch i {
	case beerjson.SpecificVolumeUnitType_QtLb:
		return beerproto.SpecificVolumeType_QTLB
	case beerjson.SpecificVolumeUnitType_GalLb:
		return beerproto.SpecificVolumeType_GALLB
	case beerjson.SpecificVolumeUnitType_GalOz:
		return beerproto.SpecificVolumeType_GALOZ
	case beerjson.SpecificVolumeUnitType_LG:
		return beerproto.SpecificVolumeType_LG
	case beerjson.SpecificVolumeUnitType_LKg:
		return beerproto.SpecificVolumeType_LKG
	case beerjson.SpecificVolumeUnitType_FlozOz:
		return beerproto.SpecificVolumeType_FLOZOZ
	case beerjson.SpecificVolumeUnitType_M3Kg:
		return beerproto.SpecificVolumeType_M3KG
	case beerjson.SpecificVolumeUnitType_Ft3Lb:
		return beerproto.SpecificVolumeType_FT3LB
	}

	return beerproto.SpecificVolumeType_SpecificVolumeUnitType(unit)
}

func ToProtoMashStepTypeType(i string) beerproto.MashStepType_MashStepTypeType {
	unit := beerproto.MashStepType_MashStepTypeType_value[strings.ToUpper(i)]

	switch i {
	case "Infusion":
		return beerproto.MashStepType_INFUSION
	case "Temperature":
		return beerproto.MashStepType_TEMPERATURE
	case "Decoction":
		return beerproto.MashStepType_DECOCTION
	case "Sparge":
		return beerproto.MashStepType_SPARGE
	case "Tun":
		return beerproto.MashStepType_DRAIN_MASH_TUN
	}

	return beerproto.MashStepType_MashStepTypeType(unit)
}

func ToProtoAcidityType(i float32) *beerproto.AcidityType {
	return &beerproto.AcidityType{
		Value: float64(i),
		Unit:  beerproto.AcidityType_PH,
	}
}

func ToProtoAcidityUnitType(i beerjson.AcidityUnitType) beerproto.AcidityType_AcidityUnitType {
	unit := beerproto.AcidityType_AcidityUnitType_value[strings.ToUpper(string(i))]
	return beerproto.AcidityType_AcidityUnitType(unit)
}

func ToProtoTimeType(i int64) *beerproto.TimeType {
	return &beerproto.TimeType{
		Value: float64(i),
		Unit:  beerproto.TimeType_MIN,
	}
}

func ToProtoTemperatureType(i float64) *beerproto.TemperatureType {
	return &beerproto.TemperatureType{
		Value: i,
		Unit:  beerproto.TemperatureType_C,
	}
}

func ToProtoTimeUnitType(i beerjson.TimeUnitType) beerproto.TimeType_TimeUnitType {
	unit := beerproto.TimeType_TimeUnitType_value[strings.ToUpper(string(i))]
	return beerproto.TimeType_TimeUnitType(unit)
}

func ToProtoTemperatureUnitType(i beerjson.TemperatureUnitType) beerproto.TemperatureType_TemperatureUnitType {
	unit := beerproto.TemperatureType_TemperatureUnitType_value[strings.ToUpper(string(i))]
	return beerproto.TemperatureType_TemperatureUnitType(unit)
}

func UseString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func UseFloat(s *float64) float64 {
	if s == nil {
		return 0
	}
	return *s
}

func UseBool(s *bool) bool {
	if s == nil {
		return false
	}
	return *s
}

func UseInt(s *int32) int32 {
	if s == nil {
		return 0
	}
	return *s
}
