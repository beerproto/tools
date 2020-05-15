package beerJSON

import (
	"strings"

	"github.com/beerproto/beerjson.go"
	"github.com/beerproto/beerproto.go"
)

func MapToJSON(i *beerproto.Recipe) *beerjson.Beerjson {
	output := &beerjson.Beerjson{
		Mashes:                   []beerjson.MashProcedureType{},
		Recipes:                  []beerjson.RecipeType{},
		MiscellaneousIngredients: []beerjson.MiscellaneousType{},
		Styles:                   []beerjson.StyleType{},
		Fermentations:            []beerjson.FermentationProcedureType{},
		Boil:                     []beerjson.BoilProcedureType{},
		Version:                  beerjson.VersionType(i.Version),
		Fermentables:             []beerjson.FermentableType{},
		TimingObject:             ToJSONTimingType(i.TimingObject),
		Cultures:                 []beerjson.CultureInformation{},
		Equipments:               []beerjson.EquipmentType{},
		Packaging:                []beerjson.PackagingProcedureType{},
		HopVarieties:             []beerjson.VarietyInformation{},
		Profiles:                 []beerjson.WaterBase{},
	}

	for _, mash := range i.Mashes {
		output.Mashes = append(output.Mashes, *ToJSONMashProcedureType(mash))
	}

	for _, recipe := range i.Recipes {
		output.Recipes = append(output.Recipes, *ToJSONRecipeType(recipe))
	}

	for _, ingredients := range i.MiscellaneousIngredients {
		output.MiscellaneousIngredients = append(output.MiscellaneousIngredients, *ToJSONMiscellaneousType(ingredients))
	}

	for _, style := range i.Styles {
		output.Styles = append(output.Styles, *ToJSONStyleType(style))
	}

	for _, fermentation := range i.Fermentations {
		output.Fermentations = append(output.Fermentations, *ToJSONFermentationProcedureType(fermentation))
	}

	for _, boil := range i.Boil {
		output.Boil = append(output.Boil, *ToJSONBoilProcedureType(boil))
	}

	for _, fermentable := range i.Fermentables {
		output.Fermentables = append(output.Fermentables, *ToJSONFermentableType(fermentable))
	}

	for _, culture := range i.Cultures {
		output.Cultures = append(output.Cultures, *ToJSONCultureInformation(culture))
	}

	for _, equipment := range i.Equipments {
		output.Equipments = append(output.Equipments, *ToJSONEquipmentType(equipment))
	}

	for _, packaging := range i.Packaging {
		output.Packaging = append(output.Packaging, *ToJSONPackagingProcedureType(packaging))
	}

	for _, hopVariety := range i.HopVarieties {
		output.HopVarieties = append(output.HopVarieties, *ToJSONVarietyInformation(hopVariety))
	}

	for _, profile := range i.Profiles {
		output.Profiles = append(output.Profiles, *ToJSONWaterBase(profile))
	}

	return output
}

func ToJSONWaterBase(i *beerproto.WaterBase) *beerjson.WaterBase {
	if i == nil {
		return nil
	}

	return &beerjson.WaterBase{
		Calcium:     *ToJSONConcentrationType(i.Calcium),
		Nitrite:     ToJSONConcentrationType(i.Nitrite),
		Chloride:    *ToJSONConcentrationType(i.Chloride),
		Name:        i.Name,
		Potassium:   ToJSONConcentrationType(i.Potassium),
		Carbonate:   ToJSONConcentrationType(i.Carbonate),
		Iron:        ToJSONConcentrationType(i.Iron),
		Flouride:    ToJSONConcentrationType(i.Flouride),
		Sulfate:     *ToJSONConcentrationType(i.Sulfate),
		Magnesium:   *ToJSONConcentrationType(i.Magnesium),
		Producer:    &i.Producer,
		Bicarbonate: *ToJSONConcentrationType(i.Bicarbonate),
		Nitrate:     ToJSONConcentrationType(i.Nitrate),
		Sodium:      *ToJSONConcentrationType(i.Sodium),
	}
}

func ToJSONVarietyInformation(i *beerproto.VarietyInformation) *beerjson.VarietyInformation {
	return &beerjson.VarietyInformation{
		Inventory:              ToJSONHopInventoryType(i.Inventory),
		VarietyInformationType: ToJSONVarietyInformationType(i.Type),
		OilContent:             ToJSONOilContentType(i.OilContent),
		PercentLost:            ToJSONPercentType(i.PercentLost),
		ProductId:              &i.ProductId,
		AlphaAcid:              ToJSONPercentType(i.AlphaAcid),
		BetaAcid:               ToJSONPercentType(i.BetaAcid),
		Name:                   &i.Name,
		Origin:                 &i.Origin,
		Substitutes:            &i.Substitutes,
		Year:                   &i.Year,
		HopVarietyBaseForm:     ToJSONHopVarietyBaseForm(i.Form),
		Producer:               &i.Producer,
		Notes:                  &i.Notes,
	}
}

func ToJSONOilContentType(i *beerproto.OilContentType) *beerjson.OilContentType {
	if i == nil {
		return nil
	}

	return &beerjson.OilContentType{
		Polyphenols:       ToJSONPercentType(i.Polyphenols),
		TotalOilMlPer100g: &i.TotalOilMlPer_100G,
		Farnesene:         ToJSONPercentType(i.Farnesene),
		Limonene:          ToJSONPercentType(i.Limonene),
		Nerol:             ToJSONPercentType(i.Nerol),
		Geraniol:          ToJSONPercentType(i.Geraniol),
		BPinene:           ToJSONPercentType(i.BPinene),
		Linalool:          ToJSONPercentType(i.Linalool),
		Caryophyllene:     ToJSONPercentType(i.Caryophyllene),
		Cohumulone:        ToJSONPercentType(i.Cohumulone),
		Xanthohumol:       ToJSONPercentType(i.Xanthohumol),
		Humulene:          ToJSONPercentType(i.Humulene),
		Myrcene:           ToJSONPercentType(i.Myrcene),
		Pinene:            ToJSONPercentType(i.Pinene),
	}
}

func ToJSONVarietyInformationType(i beerproto.VarietyInformation_VarietyInformationType) *beerjson.VarietyInformationType {
	if i == beerproto.VarietyInformation_NULL_VarietyInformationType {
		return nil
	}

	unit := beerproto.VarietyInformation_VarietyInformationType_name[int32(i)]
	t := beerjson.VarietyInformationType(strings.ToLower(unit))
	return &t
}

func ToJSONHopInventoryType(i *beerproto.HopInventoryType) *beerjson.HopInventoryType {
	if i == nil {
		return nil
	}

	hopInventoryType := &beerjson.HopInventoryType{}

	if mass, ok := i.Amount.(*beerproto.HopInventoryType_Mass); ok {
		hopInventoryType.Amount = ToJSONMassType(mass.Mass)
	}

	if volume, ok := i.Amount.(*beerproto.HopInventoryType_Volume); ok {
		hopInventoryType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return hopInventoryType
}

func ToJSONEquipmentType(i *beerproto.EquipmentType) *beerjson.EquipmentType {
	if i == nil {
		return nil
	}

	equipmentItemType := []beerjson.EquipmentItemType{}
	for _, item := range i.EquipmentItems {
		equipmentItemType = append(equipmentItemType, ToJSONEquipmentItemType(item))
	}

	return &beerjson.EquipmentType{
		Name:           i.Name,
		EquipmentItems: equipmentItemType,
	}
}

func ToJSONEquipmentItemType(i *beerproto.EquipmentItemType) beerjson.EquipmentItemType {
	return beerjson.EquipmentItemType{
		BoilRatePerHour:     ToJSONVolumeType(i.BoilRatePerHour),
		KeyType:             &i.Type,
		EquipmentBaseForm:   ToJSONEquipmentBaseForm(i.Form),
		DrainRatePerMinute:  ToJSONVolumeType(i.DrainRatePerMinute),
		SpecificHeat:        ToJSONSpecificHeatType(i.SpecificHeat),
		GrainAbsorptionRate: ToJSONSpecificVolumeType(i.GrainAbsorptionRate),
		Name:                &i.Name,
		MaximumVolume:       ToJSONVolumeType(i.MaximumVolume),
		Weight:              ToJSONMassType(i.Weight),
		Loss:                *ToJSONVolumeType(i.Loss),
	}
}

func ToJSONSpecificHeatType(i *beerproto.SpecificHeatType) *beerjson.SpecificHeatType {
	if i == nil {
		return nil
	}

	return &beerjson.SpecificHeatType{
		Value: i.Value,
		Unit:  *ToJSONSpecificHeatUnitType(i.Unit),
	}
}

func ToJSONSpecificHeatUnitType(i beerproto.SpecificHeatType_SpecificHeatUnitType) *beerjson.SpecificHeatUnitType {
	if i == beerproto.SpecificHeatType_NULL {
		return nil
	}

	unit := beerproto.SpecificHeatType_SpecificHeatUnitType_name[int32(i)]
	t := beerjson.SpecificHeatUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONEquipmentBaseForm(i beerproto.EquipmentItemType_EquipmentBaseForm) *beerjson.EquipmentBaseForm {
	if i == beerproto.EquipmentItemType_NULL {
		return nil
	}

	var t beerjson.EquipmentBaseForm
	switch i {
	case beerproto.EquipmentItemType_HLT:
		t = beerjson.EquipmentBaseForm_HLT
	case beerproto.EquipmentItemType_MASH_TUN:
		t = beerjson.EquipmentBaseForm_MashTun
	case beerproto.EquipmentItemType_LAUTER_TUN:
		t = beerjson.EquipmentBaseForm_LauterTun
	case beerproto.EquipmentItemType_BREW_KETTLE:
		t = beerjson.EquipmentBaseForm_BrewKettle
	case beerproto.EquipmentItemType_FERMENTER:
		t = beerjson.EquipmentBaseForm_Fermenter
	case beerproto.EquipmentItemType_AGING_VESSEL:
		t = beerjson.EquipmentBaseForm_AgingVessel
	case beerproto.EquipmentItemType_PACKAGING_VESSEL:
		t = beerjson.EquipmentBaseForm_PackagingVessel
	}

	return &t
}

func ToJSONCultureInformation(i *beerproto.CultureInformation) *beerjson.CultureInformation {
	if i == nil {
		return nil
	}

	return &beerjson.CultureInformation{
		CultureBaseForm:  ToJSONCultureBaseForm(i.Form),
		Producer:         &i.Producer,
		TemperatureRange: ToJSONTemperatureRangeType(i.TemperatureRange),
		Notes:            &i.Notes,
		BestFor:          &i.BestFor,
		Inventory:        ToJSONCultureInventoryType(i.Inventory),
		ProductId:        &i.ProductId,
		Name:             &i.Name,
		AlcoholTolerance: ToJSONPercentType(i.AlcoholTolerance),
		Glucoamylase:     &i.Glucoamylase,
		CultureBaseType:  ToJSONCultureBaseType(i.Type),
		Flocculation:     ToJSONQualitativeRangeType(i.Flocculation),
		AttenuationRange: ToJSONPercentRangeType(i.AttenuationRange),
		MaxReuse:         &i.MaxReuse,
		Pof:              &i.Pof,
		Zymocide:         ToJSONZymocide(i.Zymocide),
	}
}

func ToJSONZymocide(i *beerproto.Zymocide) *beerjson.Zymocide {
	if i == nil {
		return nil
	}
	return &beerjson.Zymocide{
		No1:     &i.No1,
		No2:     &i.No2,
		No28:    &i.No28,
		Klus:    &i.Klus,
		Neutral: &i.Neutral,
	}
}
func ToJSONQualitativeRangeType(i beerproto.CultureInformation_QualitativeRangeType) *beerjson.QualitativeRangeType {
	if i == beerproto.CultureInformation_NULL_QualitativeRangeType {
		return nil
	}

	unit := beerproto.CultureInformation_QualitativeRangeType_name[int32(i)]
	t := beerjson.QualitativeRangeType(strings.ToLower(unit))
	return &t
}

func ToJSONCultureBaseType(i beerproto.CultureBaseType) *beerjson.CultureBaseType {
	if i == beerproto.CultureBaseType_NULL_CultureBaseType {
		return nil
	}

	unit := beerproto.CultureBaseType_name[int32(i)]
	t := beerjson.CultureBaseType(strings.ToLower(unit))
	return &t
}

func ToJSONCultureInventoryType(i *beerproto.CultureInventoryType) *beerjson.CultureInventoryType {
	if i == nil {
		return nil
	}
	return &beerjson.CultureInventoryType{
		Liquid:  ToJSONVolumeType(i.Liquid),
		Dry:     ToJSONMassType(i.Dry),
		Slant:   ToJSONVolumeType(i.Slant),
		Culture: ToJSONVolumeType(i.Culture),
	}
}

func ToJSONTemperatureRangeType(i *beerproto.TemperatureRangeType) *beerjson.TemperatureRangeType {
	if i == nil {
		return nil
	}
	return &beerjson.TemperatureRangeType{
		Minimum: *ToJSONTemperatureType(i.Minimum),
		Maximum: *ToJSONTemperatureType(i.Maximum),
	}
}

func ToJSONFermentableType(i *beerproto.FermentableType) *beerjson.FermentableType {
	if i == nil {
		return nil
	}
	return &beerjson.FermentableType{
		MaxInBatch:                ToJSONPercentType(i.MaxInBatch),
		RecommendMash:             &i.RecommendMash,
		Protein:                   ToJSONPercentType(i.Protein),
		ProductId:                 &i.ProductId,
		FermentableBaseGrainGroup: ToJSONGrainGroup(i.GrainGroup),
		Yield:                     ToJSONYieldType(i.Yield),
		FermentableBaseType:       ToJSONFermentableBaseType(i.Type),
		Producer:                  &i.Producer,
		AlphaAmylase:              &i.AlphaAmylase,
		Color:                     ToJSONColorType(i.Color),
		Name:                      &i.Name,
		DiastaticPower:            ToJSONDiastaticPowerType(i.DiastaticPower),
		Moisture:                  ToJSONPercentType(i.Moisture),
		Origin:                    &i.Origin,
		Inventory:                 ToJSONFermentableInventoryType(i.Inventory),
		KolbachIndex:              &i.KolbachIndex,
		Notes:                     &i.Notes,
	}
}

func ToJSONFermentableInventoryType(i *beerproto.FermentableInventoryType) *beerjson.FermentableInventoryType {
	if i == nil {
		return nil
	}

	fermentableInventoryType := &beerjson.FermentableInventoryType{}

	if mass, ok := i.Amount.(*beerproto.FermentableInventoryType_Mass); ok {
		fermentableInventoryType.Amount = ToJSONMassType(mass.Mass)
	}

	if volume, ok := i.Amount.(*beerproto.FermentableInventoryType_Volume); ok {
		fermentableInventoryType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return fermentableInventoryType
}

func ToJSONDiastaticPowerType(i *beerproto.DiastaticPowerType) *beerjson.DiastaticPowerType {
	if i == nil {
		return nil
	}
	return &beerjson.DiastaticPowerType{
		Value: i.Value,
		Unit:  ToJSONDiastaticPowerUnitType(i.Unit),
	}
}

func ToJSONDiastaticPowerUnitType(i beerproto.DiastaticPowerType_DiastaticPowerUnitType) beerjson.DiastaticPowerUnitType {
	if i == beerproto.DiastaticPowerType_NULL {
		return beerjson.DiastaticPowerUnitType_Lintner
	}

	unit := beerproto.DiastaticPowerType_DiastaticPowerUnitType_name[int32(i)]
	return beerjson.DiastaticPowerUnitType(strings.ToLower(unit))
}

func ToJSONStyleType(i *beerproto.StyleType) *beerjson.StyleType {
	if i == nil {
		return nil
	}

	return &beerjson.StyleType{
		Aroma:                        &i.Aroma,
		Ingredients:                  &i.Ingredients,
		CategoryNumber:               &i.CategoryNumber,
		Notes:                        &i.Notes,
		Flavor:                       &i.Flavor,
		Mouthfeel:                    &i.Mouthfeel,
		FinalGravity:                 ToJSONGravityRangeType(i.FinalGravity),
		StyleGuide:                   &i.StyleGuide,
		Color:                        ToJSONColorRangeType(i.Color),
		OriginalGravity:              ToJSONGravityRangeType(i.OriginalGravity),
		Examples:                     &i.Examples,
		Name:                         &i.Name,
		Carbonation:                  ToJSONCarbonationRangeType(i.Carbonation),
		AlcoholByVolume:              ToJSONPercentRangeType(i.AlcoholByVolume),
		InternationalBitternessUnits: ToJSONBitternessRangeType(i.InternationalBitternessUnits),
		Appearance:                   &i.Appearance,
		Category:                     &i.Category,
		StyleLetter:                  &i.StyleLetter,
		KeyType:                      ToJSONStyleType_StyleCategories(i.Type),
		OverallImpression:            &i.OverallImpression,
	}
}

func ToJSONStyleType_StyleCategories(i beerproto.StyleType_StyleCategories) *beerjson.StyleCategories {
	if i == beerproto.StyleType_NULL {
		return nil
	}

	unit := beerproto.StyleType_StyleCategories_name[int32(i)]
	t := beerjson.StyleCategories(strings.ToLower(unit))
	return &t
}

func ToJSONBitternessRangeType(i *beerproto.BitternessRangeType) *beerjson.BitternessRangeType {
	if i == nil {
		return nil
	}
	return &beerjson.BitternessRangeType{
		Minimum: *ToJSONBitternessType(i.Minimum),
		Maximum: *ToJSONBitternessType(i.Maximum),
	}
}

func ToJSONBitternessType(i *beerproto.BitternessType) *beerjson.BitternessType {
	if i == nil {
		return nil
	}
	return &beerjson.BitternessType{
		Value: i.Value,
		Unit:  ToJSONBitternessUnitType(i.Unit),
	}
}

func ToJSONBitternessUnitType(i beerproto.BitternessType_BitternessUnitType) beerjson.BitternessUnitType {
	if i == beerproto.BitternessType_NULL {
		return beerjson.BitternessUnitType_IBUs
	}

	switch i {
	case beerproto.BitternessType_IBUs:
		return beerjson.BitternessUnitType_IBUs
	}

	return beerjson.BitternessUnitType_IBUs
}

func ToJSONPercentRangeType(i *beerproto.PercentRangeType) *beerjson.PercentRangeType {
	if i == nil {
		return nil
	}
	return &beerjson.PercentRangeType{
		Minimum: *ToJSONPercentType(i.Minimum),
		Maximum: *ToJSONPercentType(i.Maximum),
	}
}

func ToJSONCarbonationRangeType(i *beerproto.CarbonationRangeType) *beerjson.CarbonationRangeType {
	if i == nil {
		return nil
	}
	return &beerjson.CarbonationRangeType{
		Minimum: *ToJSONCarbonationType(i.Minimum),
		Maximum: *ToJSONCarbonationType(i.Maximum),
	}
}
func ToJSONCarbonationType(i *beerproto.CarbonationType) *beerjson.CarbonationType {
	if i == nil {
		return nil
	}
	return &beerjson.CarbonationType{
		Value: i.Value,
		Unit:  *ToJSONCarbonationUnitType(i.Unit),
	}
}

func ToJSONCarbonationUnitType(i beerproto.CarbonationType_CarbonationUnitType) *beerjson.CarbonationUnitType {
	if i == beerproto.CarbonationType_NULL {
		return nil
	}

	unit := beerproto.CarbonationType_CarbonationUnitType_name[int32(i)]
	t := beerjson.CarbonationUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONColorRangeType(i *beerproto.ColorRangeType) *beerjson.ColorRangeType {
	if i == nil {
		return nil
	}
	return &beerjson.ColorRangeType{
		Minimum: *ToJSONColorType(i.Minimum),
		Maximum: *ToJSONColorType(i.Maximum),
	}
}

func ToJSONGravityRangeType(i *beerproto.GravityRangeType) *beerjson.GravityRangeType {
	if i == nil {
		return nil
	}

	return &beerjson.GravityRangeType{
		Minimum: *ToJSONGravityType(i.Minimum),
		Maximum: *ToJSONGravityType(i.Maximum),
	}
}

func ToJSONMiscellaneousType(i *beerproto.MiscellaneousType) *beerjson.MiscellaneousType {
	if i == nil {
		return nil
	}

	return &beerjson.MiscellaneousType{
		UseFor:                &i.UseFor,
		Notes:                 &i.Notes,
		Name:                  &i.Name,
		Producer:              &i.Producer,
		ProductId:             &i.ProductId,
		MiscellaneousBaseType: ToJSONMiscellaneousBaseType(i.Type),
		Inventory:             ToJSONMiscellaneousInventoryType(i.Inventory),
	}
}

func ToJSONMiscellaneousInventoryType(i *beerproto.MiscellaneousInventoryType) *beerjson.MiscellaneousInventoryType {
	if i == nil {
		return nil
	}

	miscellaneousInventoryType := &beerjson.MiscellaneousInventoryType{}

	if mass, ok := i.Amount.(*beerproto.MiscellaneousInventoryType_Mass); ok {
		miscellaneousInventoryType.Amount = ToJSONMassType(mass.Mass)
	}

	if unit, ok := i.Amount.(*beerproto.MiscellaneousInventoryType_Unit); ok {
		miscellaneousInventoryType.Amount = ToJSONUnitType(unit.Unit)
	}
	if volume, ok := i.Amount.(*beerproto.MiscellaneousInventoryType_Volume); ok {
		miscellaneousInventoryType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return miscellaneousInventoryType
}

func ToJSONRecipeType(i *beerproto.RecipeType) *beerjson.RecipeType {
	if i == nil {
		return nil
	}

	var created beerjson.DateType
	if i.Created != "" {
		created = beerjson.DateType(i.Created)
	}
	return &beerjson.RecipeType{
		Efficiency:          *ToJSONEfficiencyType(i.Efficiency),
		Style:               ToJSONRecipeStyleType(i.Style),
		IbuEstimate:         ToJSONIBUEstimateType(i.IbuEstimate),
		ColorEstimate:       ToJSONColorType(i.ColorEstimate),
		BeerPH:              ToJSONAcidityType(i.BeerPH),
		Name:                i.Name,
		RecipeTypeType:      ToJSONRecipeTypeType(i.Type),
		Coauthor:            &i.Coauthor,
		OriginalGravity:     ToJSONGravityType(i.OriginalGravity),
		FinalGravity:        ToJSONGravityType(i.FinalGravity),
		Carbonation:         &i.Carbonation,
		Fermentation:        ToJSONFermentationProcedureType(i.Fermentation),
		Author:              i.Author,
		Ingredients:         *ToJSONIngredientsType(i.Ingredients),
		Mash:                ToJSONMashProcedureType(i.Mash),
		Packaging:           ToJSONPackagingProcedureType(i.Packaging),
		Boil:                ToJSONBoilProcedureType(i.Boil),
		Taste:               ToJSONTasteType(i.Taste),
		CaloriesPerPint:     &i.CaloriesPerPint,
		Created:             &created,
		BatchSize:           *ToJSONVolumeType(i.BatchSize),
		Notes:               &i.Notes,
		AlcoholByVolume:     ToJSONPercentType(i.AlcoholByVolume),
		ApparentAttenuation: ToJSONPercentType(i.ApparentAttenuation),
	}
}

func ToJSONTasteType(i *beerproto.TasteType) *beerjson.TasteType {
	if i == nil {
		return nil
	}
	return &beerjson.TasteType{
		Notes:  i.Notes,
		Rating: i.Rating,
	}
}

func ToJSONBoilProcedureType(i *beerproto.BoilProcedureType) *beerjson.BoilProcedureType {
	if i == nil {
		return nil
	}
	boilSteps := make([]beerjson.BoilStepType, 0)

	for _, step := range i.BoilSteps {
		boilSteps = append(boilSteps, *ToJSONBoilStepType(step))
	}

	return &beerjson.BoilProcedureType{
		PreBoilSize: ToJSONVolumeType(i.PreBoilSize),
		BoilTime:    *ToJSONTimeType(i.BoilTime),
		Name:        &i.Name,
		Description: &i.Description,
		Notes:       &i.Notes,
		BoilSteps:   boilSteps,
	}
}

func ToJSONBoilStepType(i *beerproto.BoilStepType) *beerjson.BoilStepType {
	if i == nil {
		return nil
	}

	return &beerjson.BoilStepType{
		EndGravity:               ToJSONGravityType(i.EndGravity),
		BoilStepTypeChillingType: ToJSONChillingType(i.ChillingType),
		Description:              &i.Description,
		EndTemperature:           ToJSONTemperatureType(i.EndTemperature),
		RampTime:                 ToJSONTimeType(i.RampTime),
		StepTime:                 ToJSONTimeType(i.StepTime),
		StartGravity:             ToJSONGravityType(i.StartGravity),
		StartPh:                  ToJSONAcidityType(i.StartPh),
		EndPh:                    ToJSONAcidityType(i.EndPh),
		Name:                     i.Name,
		StartTemperature:         ToJSONTemperatureType(i.StartTemperature),
	}
}

func ToJSONChillingType(i beerproto.BoilStepType_BoilStepTypeChillingType) *beerjson.BoilStepTypeChillingType {
	if i == beerproto.BoilStepType_NULL {
		return nil
	}

	unit := beerproto.BoilStepType_BoilStepTypeChillingType_name[int32(i)]
	t := beerjson.BoilStepTypeChillingType(strings.ToLower(unit))
	return &t
}

func ToJSONPackagingProcedureType(i *beerproto.PackagingProcedureType) *beerjson.PackagingProcedureType {
	if i == nil {
		return nil
	}
	packagingVessels := make([]beerjson.PackagingVesselType, 0)

	for _, vessels := range i.PackagingVessels {
		packagingVessels = append(packagingVessels, *ToJSONPackagingVesselType(vessels))
	}
	return &beerjson.PackagingProcedureType{
		Name:             i.Name,
		PackagedVolume:   ToJSONVolumeType(i.PackagedVolume),
		Description:      &i.Description,
		Notes:            &i.Notes,
		PackagingVessels: packagingVessels,
	}
}

func ToJSONPackagingVesselType(i *beerproto.PackagingVesselType) *beerjson.PackagingVesselType {
	if i == nil {
		return nil
	}

	var packageDate beerjson.DateType
	if i.PackageDate != "" {
		packageDate = beerjson.DateType(i.PackageDate)
	}
	return &beerjson.PackagingVesselType{
		PackagingVesselTypeType: ToJSONPackagingVesselTypeType(i.Type),
		StartGravity:            ToJSONGravityType(i.StartGravity),
		Name:                    i.Name,
		PackageDate:             &packageDate,
		StepTime:                ToJSONTimeType(i.StepTime),
		EndGravity:              ToJSONGravityType(i.EndGravity),
		VesselVolume:            ToJSONVolumeType(i.VesselVolume),
		VesselQuantity:          &i.VesselQuantity,
		Description:             &i.Description,
		StartPh:                 ToJSONAcidityType(i.StartPh),
		Carbonation:             &i.Carbonation,
		StartTemperature:        ToJSONTemperatureType(i.StartTemperature),
		EndPh:                   ToJSONAcidityType(i.EndPh),
		EndTemperature:          ToJSONTemperatureType(i.EndTemperature),
	}
}

func ToJSONPackagingVesselTypeType(i beerproto.PackagingVesselType_PackagingVesselTypeType) *beerjson.PackagingVesselTypeType {
	if i == beerproto.PackagingVesselType_NULL {
		return nil
	}

	unit := beerproto.PackagingVesselType_PackagingVesselTypeType_name[int32(i)]
	t := beerjson.PackagingVesselTypeType(strings.ToLower(unit))
	return &t
}

func ToJSONIngredientsType(i *beerproto.IngredientsType) *beerjson.IngredientsType {
	if i == nil {
		return nil
	}

	miscellaneousAdditions := make([]beerjson.MiscellaneousAdditionType, 0)
	cultureAdditions := make([]beerjson.CultureAdditionType, 0)
	waterAdditions := make([]beerjson.WaterAdditionType, 0)
	fermentableAdditions := make([]beerjson.FermentableAdditionType, 0)
	hopAdditions := make([]beerjson.HopAdditionType, 0)

	for _, misc := range i.MiscellaneousAdditions {
		miscellaneousAdditions = append(miscellaneousAdditions, *ToJSONMiscellaneousAdditionType(misc))
	}
	for _, culture := range i.CultureAdditions {
		cultureAdditions = append(cultureAdditions, *ToJSONCultureAdditionType(culture))
	}
	for _, water := range i.WaterAdditions {
		waterAdditions = append(waterAdditions, *ToJSONWaterAdditionType(water))
	}
	for _, fermentable := range i.FermentableAdditions {
		fermentableAdditions = append(fermentableAdditions, *ToJSONFermentableAdditionType(fermentable))
	}
	for _, hop := range i.HopAdditions {
		hopAdditions = append(hopAdditions, *ToJSONHopAdditionType(hop))
	}
	return &beerjson.IngredientsType{
		MiscellaneousAdditions: miscellaneousAdditions,
		CultureAdditions:       cultureAdditions,
		WaterAdditions:         waterAdditions,
		FermentableAdditions:   fermentableAdditions,
		HopAdditions:           hopAdditions,
	}
}

func ToJSONHopAdditionType(i *beerproto.HopAdditionType) *beerjson.HopAdditionType {
	if i == nil {
		return nil
	}

	hopAdditionType := &beerjson.HopAdditionType{
		BetaAcid:           ToJSONPercentType(i.BetaAcid),
		Producer:           &i.Producer,
		Origin:             &i.Origin,
		Year:               &i.Year,
		HopVarietyBaseForm: ToJSONHopVarietyBaseForm(i.Form),
		Timing:             *ToJSONTimingType(i.Timing),
		Name:               &i.Name,
		ProductId:          &i.ProductId,
		AlphaAcid:          ToJSONPercentType(i.AlphaAcid),
	}

	if mass, ok := i.Amount.(*beerproto.HopAdditionType_Mass); ok {
		hopAdditionType.Amount = ToJSONMassType(mass.Mass)
	}

	if volume, ok := i.Amount.(*beerproto.HopAdditionType_Volume); ok {
		hopAdditionType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return hopAdditionType
}

func ToJSONHopVarietyBaseForm(i beerproto.HopVarietyBaseForm) *beerjson.HopVarietyBaseForm {
	if i == beerproto.HopVarietyBaseForm_NULL_HopVarietyBaseForm {
		return nil
	}

	var t beerjson.HopVarietyBaseForm

	switch i {
	case beerproto.HopVarietyBaseForm_EXTRACT_HopVarietyBaseForm:
		t = beerjson.HopVarietyBaseForm_Extract
	case beerproto.HopVarietyBaseForm_LEAF:
		t = beerjson.HopVarietyBaseForm_Leaf
	case beerproto.HopVarietyBaseForm_LEAFWET:
		t = beerjson.HopVarietyBaseForm_LeafWet
	case beerproto.HopVarietyBaseForm_PELLET:
		t = beerjson.HopVarietyBaseForm_Pellet
	case beerproto.HopVarietyBaseForm_POWDER:
		t = beerjson.HopVarietyBaseForm_Powder
	case beerproto.HopVarietyBaseForm_PLUG:
		t = beerjson.HopVarietyBaseForm_Plug
	}

	return &t
}

func ToJSONFermentableAdditionType(i *beerproto.FermentableAdditionType) *beerjson.FermentableAdditionType {
	if i == nil {
		return nil
	}

	fermentableAdditionType := &beerjson.FermentableAdditionType{
		FermentableBaseType:       ToJSONFermentableBaseType(i.Type),
		Origin:                    &i.Origin,
		FermentableBaseGrainGroup: ToJSONGrainGroup(i.GrainGroup),
		Yield:                     ToJSONYieldType(i.Yield),
		Color:                     ToJSONColorType(i.Color),
		Name:                      &i.Name,
		Producer:                  &i.Producer,
		ProductId:                 &i.ProductId,
		Timing:                    ToJSONTimingType(i.Timing),
	}

	if mass, ok := i.Amount.(*beerproto.FermentableAdditionType_Mass); ok {
		fermentableAdditionType.Amount = ToJSONMassType(mass.Mass)
	}

	if volume, ok := i.Amount.(*beerproto.FermentableAdditionType_Volume); ok {
		fermentableAdditionType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return fermentableAdditionType
}

func ToJSONYieldType(i *beerproto.YieldType) *beerjson.YieldType {
	if i == nil {
		return nil
	}

	return &beerjson.YieldType{
		FineGrind:            ToJSONPercentType(i.FineGrind),
		CoarseGrind:          ToJSONPercentType(i.CoarseGrind),
		FineCoarseDifference: ToJSONPercentType(i.FineCoarseDifference),
		Potential:            ToJSONGravityType(i.Potential),
	}
}

func ToJSONGrainGroup(i beerproto.GrainGroup) *beerjson.FermentableBaseGrainGroup {
	if i == beerproto.GrainGroup_NULL_GrainGroup {
		return nil
	}

	unit := beerproto.GrainGroup_name[int32(i)]
	t := beerjson.FermentableBaseGrainGroup(strings.ToLower(unit))
	return &t
}

func ToJSONFermentableBaseType(i beerproto.FermentableBaseType) *beerjson.FermentableBaseType {
	if i == beerproto.FermentableBaseType_NULL_FermentableBaseType {
		return nil
	}

	unit := beerproto.FermentableBaseType_name[int32(i)]
	t := beerjson.FermentableBaseType(strings.ToLower(unit))
	return &t
}

func ToJSONWaterAdditionType(i *beerproto.WaterAdditionType) *beerjson.WaterAdditionType {
	if i == nil {
		return nil
	}

	return &beerjson.WaterAdditionType{
		Flouride:    ToJSONConcentrationType(i.Flouride),
		Sulfate:     ToJSONConcentrationType(i.Sulfate),
		Producer:    &i.Producer,
		Nitrate:     ToJSONConcentrationType(i.Nitrate),
		Nitrite:     ToJSONConcentrationType(i.Nitrite),
		Chloride:    ToJSONConcentrationType(i.Chloride),
		Amount:      ToJSONVolumeType(i.Amount),
		Name:        &i.Name,
		Potassium:   ToJSONConcentrationType(i.Potassium),
		Magnesium:   ToJSONConcentrationType(i.Magnesium),
		Iron:        ToJSONConcentrationType(i.Iron),
		Bicarbonate: ToJSONConcentrationType(i.Bicarbonate),
		Calcium:     ToJSONConcentrationType(i.Calcium),
		Carbonate:   ToJSONConcentrationType(i.Carbonate),
		Sodium:      ToJSONConcentrationType(i.Sodium),
	}
}

func ToJSONConcentrationType(i *beerproto.ConcentrationType) *beerjson.ConcentrationType {
	if i == nil {
		return nil
	}

	return &beerjson.ConcentrationType{
		Value: i.Value,
		Unit:  *ToJSONConcentrationUnitType(i.Unit),
	}
}

func ToJSONConcentrationUnitType(i beerproto.ConcentrationType_ConcentrationUnitType) *beerjson.ConcentrationUnitType {
	if i == beerproto.ConcentrationType_NULL {
		return nil
	}

	unit := beerproto.ConcentrationType_ConcentrationUnitType_name[int32(i)]
	t := beerjson.ConcentrationUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONCultureAdditionType(i *beerproto.CultureAdditionType) *beerjson.CultureAdditionType {
	if i == nil {
		return nil
	}

	cultureAdditionType := &beerjson.CultureAdditionType{
		CultureBaseForm:   ToJSONCultureBaseForm(i.Form),
		ProductId:         &i.ProductId,
		Name:              &i.Name,
		CellCountBillions: &i.CellCountBillions,
		TimesCultured:     &i.TimesCultured,
		Producer:          &i.Producer,
		CultureBaseType:   ToJSONCultureBaseType(i.Type),
		Attenuation:       ToJSONPercentType(i.Attenuation),
		Timing:            ToJSONTimingType(i.Timing),
	}

	if mass, ok := i.Amount.(*beerproto.CultureAdditionType_Mass); ok {
		cultureAdditionType.Amount = ToJSONMassType(mass.Mass)
	}

	if unit, ok := i.Amount.(*beerproto.CultureAdditionType_Unit); ok {
		cultureAdditionType.Amount = ToJSONUnitType(unit.Unit)
	}
	if volume, ok := i.Amount.(*beerproto.CultureAdditionType_Volume); ok {
		cultureAdditionType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return cultureAdditionType
}

func ToJSONCultureBaseForm(i beerproto.CultureBaseForm) *beerjson.CultureBaseForm {
	if i == beerproto.CultureBaseForm_NULL_CultureBaseForm {
		return nil
	}

	unit := beerproto.CultureBaseForm_name[int32(i)]
	t := beerjson.CultureBaseForm(strings.ToLower(unit))
	return &t
}

func ToJSONMiscellaneousAdditionType(i *beerproto.MiscellaneousAdditionType) *beerjson.MiscellaneousAdditionType {
	if i == nil {
		return nil
	}

	miscellaneousAdditionType := &beerjson.MiscellaneousAdditionType{
		Name:                  &i.Name,
		Producer:              &i.Producer,
		Timing:                ToJSONTimingType(i.Timing),
		ProductId:             &i.ProductId,
		MiscellaneousBaseType: ToJSONMiscellaneousBaseType(i.Type),
	}

	if mass, ok := i.Amount.(*beerproto.MiscellaneousAdditionType_Mass); ok {
		miscellaneousAdditionType.Amount = ToJSONMassType(mass.Mass)
	}

	if unit, ok := i.Amount.(*beerproto.MiscellaneousAdditionType_Unit); ok {
		miscellaneousAdditionType.Amount = ToJSONUnitType(unit.Unit)
	}
	if volume, ok := i.Amount.(*beerproto.MiscellaneousAdditionType_Volume); ok {
		miscellaneousAdditionType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return miscellaneousAdditionType
}

func ToJSONUnitType(i *beerproto.UnitType) *beerjson.UnitType {
	if i == nil {
		return nil
	}
	return &beerjson.UnitType{
		Value: i.Value,
		Unit:  *ToJSONUnitUnitType(i.Unit),
	}
}

func ToJSONUnitUnitType(i beerproto.UnitType_UnitUnitType) *beerjson.UnitUnitType {
	if i == beerproto.UnitType_NULL {
		return nil
	}

	unit := beerproto.UnitType_UnitUnitType_name[int32(i)]
	t := beerjson.UnitUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONMassType(i *beerproto.MassType) *beerjson.MassType {
	if i == nil {
		return nil
	}
	return &beerjson.MassType{
		Value: i.Value,
		Unit:  *ToJSONMassUnitType(i.Unit),
	}
}

func ToJSONMassUnitType(i beerproto.MassType_MassUnitType) *beerjson.MassUnitType {
	if i == beerproto.MassType_NULL {
		return nil
	}

	unit := beerproto.MassType_MassUnitType_name[int32(i)]
	t := beerjson.MassUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONMiscellaneousBaseType(i beerproto.MiscellaneousBaseType) *beerjson.MiscellaneousBaseType {
	if i == beerproto.MiscellaneousBaseType_NULL {
		return nil
	}

	unit := beerproto.MiscellaneousBaseType_name[int32(i)]
	t := beerjson.MiscellaneousBaseType(strings.ToLower(unit))
	return &t
}

func ToJSONTimingType(i *beerproto.TimingType) *beerjson.TimingType {
	if i == nil {
		return nil
	}
	return &beerjson.TimingType{
		Time:            ToJSONTimeType(i.Time),
		Duration:        ToJSONTimeType(i.Duration),
		Continuous:      &i.Continuous,
		SpecificGravity: ToJSONGravityType(i.SpecificGravity),
		PH:              ToJSONAcidityType(i.Ph),
		Step:            &i.Step,
		Use:             ToJSONUseType(i.Use),
	}
}

func ToJSONUseType(i beerproto.TimingType_UseType) *beerjson.UseType {
	if i == beerproto.TimingType_NULL {
		return nil
	}

	unit := beerproto.TimingType_UseType_name[int32(i)]
	t := beerjson.UseType(strings.ToLower(unit))
	return &t
}

func ToJSONFermentationProcedureType(i *beerproto.FermentationProcedureType) *beerjson.FermentationProcedureType {
	if i == nil {
		return nil
	}
	steps := make([]beerjson.FermentationStepType, 0)
	for _, step := range i.FermentationSteps {
		steps = append(steps, *ToJSONFermentationStepType(step))
	}
	return &beerjson.FermentationProcedureType{
		Description:       &i.Description,
		Notes:             &i.Notes,
		Name:              i.Name,
		FermentationSteps: steps,
	}
}

func ToJSONFermentationStepType(i *beerproto.FermentationStepType) *beerjson.FermentationStepType {
	if i == nil {
		return nil
	}

	return &beerjson.FermentationStepType{
		Name:             i.Name,
		EndTemperature:   ToJSONTemperatureType(i.EndTemperature),
		StepTime:         ToJSONTimeType(i.StepTime),
		FreeRise:         &i.FreeRise,
		StartGravity:     ToJSONGravityType(i.StartGravity),
		EndGravity:       ToJSONGravityType(i.EndGravity),
		StartPh:          ToJSONAcidityType(i.StartPh),
		Description:      &i.Description,
		StartTemperature: ToJSONTemperatureType(i.StartTemperature),
		EndPh:            ToJSONAcidityType(i.EndPh),
		Vessel:           &i.Vessel,
	}
}

func ToJSONGravityType(i *beerproto.GravityType) *beerjson.GravityType {
	if i == nil {
		return nil
	}
	return &beerjson.GravityType{
		Value: i.Value,
		Unit:  *ToJSONGravityUnitType(i.Unit),
	}
}

func ToJSONGravityUnitType(i beerproto.GravityType_GravityUnitType) *beerjson.GravityUnitType {
	if i == beerproto.GravityType_NULL {
		return nil
	}

	unit := beerproto.GravityType_GravityUnitType_name[int32(i)]
	t := beerjson.GravityUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONRecipeTypeType(i beerproto.RecipeType_RecipeTypeType) beerjson.RecipeTypeType {
	if i == beerproto.RecipeType_NULL {
		return beerjson.RecipeTypeType_AllGrain
	}

	return beerjson.RecipeTypeType(strings.ToLower(i.String()))
}

func ToJSONColorType(i *beerproto.ColorType) *beerjson.ColorType {
	if i == nil {
		return nil
	}
	return &beerjson.ColorType{
		Value: i.Value,
		Unit:  *ToJSONColorUnitType(i.Unit),
	}
}

func ToJSONColorUnitType(i beerproto.ColorType_ColorUnitType) *beerjson.ColorUnitType {
	if i == beerproto.ColorType_NULL {
		return nil
	}

	unit := beerproto.ColorType_ColorUnitType_name[int32(i)]

	if i == beerproto.ColorType_LOVI {
		unit = strings.Title(strings.ToLower(unit))
	}
	t := beerjson.ColorUnitType(unit)
	return &t
}

func ToJSONIBUEstimateType(i *beerproto.IBUEstimateType) *beerjson.IBUEstimateType {
	if i == nil {
		return nil
	}

	return &beerjson.IBUEstimateType{
		Method: ToJSONIBUMethodType(i.Method),
	}
}

func ToJSONIBUMethodType(i beerproto.IBUEstimateType_IBUMethodType) *beerjson.IBUMethodType {
	if i == beerproto.IBUEstimateType_NULL {
		return nil
	}

	unit := beerproto.IBUEstimateType_IBUMethodType_name[int32(i)]
	t := beerjson.IBUMethodType(unit)
	return &t
}

func ToJSONRecipeStyleType(i *beerproto.RecipeStyleType) *beerjson.RecipeStyleType {
	if i == nil {
		return nil
	}
	return &beerjson.RecipeStyleType{
		KeyType:        ToJSONRecipeStyleType_StyleCategories(i.Type),
		Name:           &i.Name,
		Category:       &i.Category,
		CategoryNumber: &i.CategoryNumber,
		StyleGuide:     &i.StyleGuide,
		StyleLetter:    &i.StyleLetter,
	}
}

func ToJSONRecipeStyleType_StyleCategories(i beerproto.RecipeStyleType_StyleCategories) *beerjson.StyleCategories {
	if i == beerproto.RecipeStyleType_NULL {
		return nil
	}

	unit := beerproto.RecipeStyleType_StyleCategories_name[int32(i)]
	t := beerjson.StyleCategories(strings.ToLower(unit))
	return &t
}

func ToJSONEfficiencyType(i *beerproto.EfficiencyType) *beerjson.EfficiencyType {
	if i == nil {
		return nil
	}

	return &beerjson.EfficiencyType{
		Conversion: ToJSONPercentType(i.Conversion),
		Lauter:     ToJSONPercentType(i.Lauter),
		Mash:       ToJSONPercentType(i.Mash),
		Brewhouse:  *ToJSONPercentType(i.Brewhouse),
	}
}

func ToJSONPercentType(i *beerproto.PercentType) *beerjson.PercentType {
	if i == nil {
		return nil
	}

	return &beerjson.PercentType{
		Value: i.Value,
		Unit:  ToJSONPercentUnitType(i.Unit),
	}
}

func ToJSONPercentUnitType(i beerproto.PercentType_PercentUnitType) beerjson.PercentUnitType {
	if i == beerproto.PercentType_NULL {
		return beerjson.PercentUnitType_No
	}

	return beerjson.PercentUnitType(strings.ToLower(i.String()))
}

func ToJSONMashProcedureType(i *beerproto.MashProcedureType) *beerjson.MashProcedureType {
	if i == nil {
		return nil
	}

	mashSteps := []beerjson.MashStepType{}
	for _, step := range i.MashSteps {
		mashSteps = append(mashSteps, *ToJSONMashStepType(step))
	}

	return &beerjson.MashProcedureType{
		Name:             i.Name,
		Notes:            &i.Notes,
		GrainTemperature: *ToJSONTemperatureType(i.GrainTemperature),
		MashSteps:        mashSteps,
	}
}

func ToJSONMashStepType(i *beerproto.MashStepType) *beerjson.MashStepType {
	if i == nil {
		return nil
	}

	return &beerjson.MashStepType{
		StepTime:          *ToJSONTimeType(i.StepTime),
		RampTime:          ToJSONTimeType(i.RampTime),
		EndTemperature:    ToJSONTemperatureType(i.EndTemperature),
		Description:       &i.Description,
		InfuseTemperature: ToJSONTemperatureType(i.InfuseTemperature),
		StartPH:           ToJSONAcidityType(i.StartPH),
		EndPH:             ToJSONAcidityType(i.EndPH),
		Name:              i.Name,
		MashStepTypeType:  *ToJSONMashStepTypeType(i.Type),
		Amount:            ToJSONVolumeType(i.Amount),
		StepTemperature:   *ToJSONTemperatureType(i.StepTemperature),
		WaterGrainRatio:   ToJSONSpecificVolumeType(i.WaterGrainRatio),
	}
}

func ToJSONVolumeType(i *beerproto.VolumeType) *beerjson.VolumeType {
	if i == nil {
		return nil
	}
	return &beerjson.VolumeType{
		Value: i.Value,
		Unit:  ToJSONVolumeUnitType(i.Unit),
	}
}

func ToJSONVolumeUnitType(i beerproto.VolumeType_VolumeUnitType) beerjson.VolumeUnitType {
	if i == beerproto.VolumeType_NULL {
		return beerjson.VolumeUnitType_Ml
	}

	return beerjson.VolumeUnitType(strings.ToLower(i.String()))
}

func ToJSONSpecificVolumeType(i *beerproto.SpecificVolumeType) *beerjson.SpecificVolumeType {
	if i == nil {
		return nil
	}
	return &beerjson.SpecificVolumeType{
		Value: i.Value,
		Unit:  ToJSONSpecificVolumeUnitType(i.Unit),
	}
}

func ToJSONSpecificVolumeUnitType(i beerproto.SpecificVolumeType_SpecificVolumeUnitType) beerjson.SpecificVolumeUnitType {
	if i == beerproto.SpecificVolumeType_NULL {
		return beerjson.SpecificVolumeUnitType_LG
	}

	switch i {
	case beerproto.SpecificVolumeType_QTLB:
		return beerjson.SpecificVolumeUnitType_QtLb
	case beerproto.SpecificVolumeType_GALLB:
		return beerjson.SpecificVolumeUnitType_GalLb
	case beerproto.SpecificVolumeType_GALOZ:
		return beerjson.SpecificVolumeUnitType_GalOz
	case beerproto.SpecificVolumeType_LG:
		return beerjson.SpecificVolumeUnitType_LG
	case beerproto.SpecificVolumeType_LKG:
		return beerjson.SpecificVolumeUnitType_LKg
	case beerproto.SpecificVolumeType_FLOZOZ:
		return beerjson.SpecificVolumeUnitType_FlozOz
	case beerproto.SpecificVolumeType_M3KG:
		return beerjson.SpecificVolumeUnitType_M3Kg
	case beerproto.SpecificVolumeType_FT3LB:
		return beerjson.SpecificVolumeUnitType_Ft3Lb
	}

	return beerjson.SpecificVolumeUnitType(i.String())
}

func ToJSONMashStepTypeType(i beerproto.MashStepType_MashStepTypeType) *beerjson.MashStepTypeType {
	if i == beerproto.MashStepType_NULL {
		return nil
	}

	unit := beerproto.MashStepType_MashStepTypeType_name[int32(i)]
	t := beerjson.MashStepTypeType(strings.ToLower(unit))
	return &t
}

func ToJSONAcidityType(i *beerproto.AcidityType) *beerjson.AcidityType {
	if i == nil {
		return nil
	}
	return &beerjson.AcidityType{
		Value: i.Value,
		Unit:  *ToJSONAcidityUnitType(i.Unit),
	}
}

func ToJSONAcidityUnitType(i beerproto.AcidityType_AcidityUnitType) *beerjson.AcidityUnitType {
	if i == beerproto.AcidityType_NULL {
		return nil
	}

	unit := beerproto.AcidityType_AcidityUnitType_name[int32(i)]
	t := beerjson.AcidityUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONTimeType(i *beerproto.TimeType) *beerjson.TimeType {
	if i == nil {
		return nil
	}
	return &beerjson.TimeType{
		Value: i.Value,
		Unit:  *ToJSONTimeUnitType(i.Unit),
	}
}

func ToJSONTemperatureType(i *beerproto.TemperatureType) *beerjson.TemperatureType {
	if i == nil {
		return nil
	}
	return &beerjson.TemperatureType{
		Value: i.Value,
		Unit:  *ToJSONTemperatureUnitType(i.Unit),
	}
}

func ToJSONTimeUnitType(i beerproto.TimeType_TimeUnitType) *beerjson.TimeUnitType {
	if i == beerproto.TimeType_NULL {
		return nil
	}

	unit := beerproto.TimeType_TimeUnitType_name[int32(i)]
	t := beerjson.TimeUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONTemperatureUnitType(i beerproto.TemperatureType_TemperatureUnitType) *beerjson.TemperatureUnitType {
	if i == beerproto.TemperatureType_NULL {
		return nil
	}

	unit := beerproto.TemperatureType_TemperatureUnitType_name[int32(i)]
	t := beerjson.TemperatureUnitType(unit)
	return &t
}
