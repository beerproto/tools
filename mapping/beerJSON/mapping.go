package beerJSON

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/RossMerr/beerjson.go"
	"github.com/beerproto/tools/beerproto"
)

func Mapping(data []byte) (*beerproto.Recipe, error) {

	input := &beerjson.Beerjson{}
	str := &struct {
		Beer *beerjson.Beerjson `json:"beerjson"`
	}{
		Beer: input,
	}

	err := json.Unmarshal(data, &str)
	if err != nil {
		return nil, fmt.Errorf("beerJSON: bad json format %w", err)
	}

	output := &beerproto.Recipe{
		Mashes: []*beerproto.MashProcedureType{},
		Recipes:[]*beerproto.RecipeType{},
		MiscellaneousIngredients:[]*beerproto.MiscellaneousType{},
		Styles: []*beerproto.StyleType{},
	}

	output.Version = float64(input.Version)

	for _, mash := range input.Mashes {
		output.Mashes = append(output.Mashes, ToProtoMashProcedureType(&mash))
	}
	for _, recipe := range input.Recipes {
		output.Recipes = append(output.Recipes, ToProtoRecipeType(recipe))
	}
	for _, ingredients := range input.MiscellaneousIngredients {
		output.MiscellaneousIngredients = append(output.MiscellaneousIngredients, ToProtoMiscellaneousType(ingredients))
	}
	for _, style := range input.Styles {
		output.Styles = append(output.Styles, ToProtoStyleType(style))
	}
	return output, nil
}

func ToProtoStyleType(i beerjson.StyleType) *beerproto.StyleType{
	return &beerproto.StyleType{
		Aroma: *i.Aroma,
		Ingredients: *i.Ingredients,
		CategoryNumber: *i.CategoryNumber,
		Notes: *i.Notes,
		Flavor: *i.Flavor,
		Mouthfeel: *i.Mouthfeel,
		FinalGravity: ToProtoGravityRangeType(i.FinalGravity),
		StyleGuide: *i.StyleGuide,
		Color: ToProtoColorRangeType(i.Color),
		OriginalGravity: ToProtoGravityRangeType(i.OriginalGravity),
		Examples: *i.Examples,
		Name: *i.Name,
		Carbonation: ToProtoCarbonationRangeType(i.Carbonation),
		AlcoholByVolume: ToProtoPercentRangeType(i.AlcoholByVolume),
		InternationalBitternessUnits: ToProtoBitternessRangeType(i.InternationalBitternessUnits),
		Appearance: *i.Appearance,
		Category: *i.Category,
		StyleLetter: *i.StyleLetter,
		Type: ToProtoStyleType_StyleCategories(i.KeyType),
		OverallImpression: *i.OverallImpression,
	}
}

func ToProtoStyleType_StyleCategories(i *beerjson.StyleCategories) beerproto.StyleType_StyleCategories{
	if i == nil {
		return beerproto.StyleType_NULL
	}
	unit := beerproto.StyleType_StyleCategories_value[strings.ToUpper(string(*i))]
	return beerproto.StyleType_StyleCategories(unit)
}

func ToProtoBitternessRangeType(i *beerjson.BitternessRangeType) *beerproto.BitternessRangeType{
	return &beerproto.BitternessRangeType{
		Minimum: ToProtoBitternessType(&i.Minimum),
		Maximum: ToProtoBitternessType(&i.Maximum),
	}
}

func ToProtoBitternessType(i *beerjson.BitternessType) *beerproto.BitternessType{
	return &beerproto.BitternessType{
		Value: i.Value,
		Unit: ToProtoBitternessUnitType(&i.Unit),
	}
}

func ToProtoBitternessUnitType(i *beerjson.BitternessUnitType) beerproto.BitternessType_BitternessUnitType{
	if i == nil {
		return beerproto.BitternessType_NULL
	}
	unit := beerproto.BitternessType_BitternessUnitType_value[strings.ToUpper(string(*i))]
	return beerproto.BitternessType_BitternessUnitType(unit)
}

func ToProtoPercentRangeType(i *beerjson.PercentRangeType) *beerproto.PercentRangeType{
	return &beerproto.PercentRangeType{
		Minimum: ToProtoPercentType(&i.Minimum),
		Maximum: ToProtoPercentType(&i.Maximum),
	}
}

func ToProtoCarbonationRangeType(i *beerjson.CarbonationRangeType) *beerproto.CarbonationRangeType{
	return &beerproto.CarbonationRangeType{
		Minimum: ToProtoCarbonationType(&i.Minimum),
		Maximum: ToProtoCarbonationType(&i.Maximum),
	}
}
func ToProtoCarbonationType(i *beerjson.CarbonationType) *beerproto.CarbonationType{
	return &beerproto.CarbonationType{
		Value: i.Value,
		Unit: ToProtoCarbonationUnitType(&i.Unit),
	}
}

func ToProtoCarbonationUnitType(i *beerjson.CarbonationUnitType) beerproto.CarbonationType_CarbonationUnitType{
	if i == nil {
		return beerproto.CarbonationType_NULL
	}
	unit := beerproto.CarbonationType_CarbonationUnitType_value[strings.ToUpper(string(*i))]
	return beerproto.CarbonationType_CarbonationUnitType(unit)
}

func ToProtoColorRangeType(i *beerjson.ColorRangeType) *beerproto.ColorRangeType{
	return &beerproto.ColorRangeType{
		Minimum: ToProtoColorType(&i.Minimum),
		Maximum: ToProtoColorType(&i.Maximum),
	}
}

func ToProtoGravityRangeType(i *beerjson.GravityRangeType) *beerproto.GravityRangeType{
	if i == nil {
		return nil
	}
	return &beerproto.GravityRangeType{
		Minimum: ToProtoGravityType(&i.Minimum),
		Maximum: ToProtoGravityType(&i.Maximum),
	}
}

func ToProtoMiscellaneousType(i beerjson.MiscellaneousType) *beerproto.MiscellaneousType{
	return &beerproto.MiscellaneousType{
		UseFor: *i.UseFor,
		Notes: *i.Notes,
		Name: *i.Name,
		Producer: *i.Producer,
		ProductId: *i.ProductId,
		Type: ToProtoMiscellaneousBaseType(i.MiscellaneousBaseType),
		Inventory: ToProtoMiscellaneousInventoryType(i.Inventory),
	}
}



func ToProtoMiscellaneousInventoryType(i *beerjson.MiscellaneousInventoryType) *beerproto.MiscellaneousInventoryType{
	if i == nil {
		return nil
	}

	miscellaneousInventoryType := &beerproto.MiscellaneousInventoryType{}

	if mass, ok :=i.Amount.(*beerjson.MassType); ok {
		miscellaneousInventoryType.Amount = &beerproto.MiscellaneousInventoryType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if unit, ok :=i.Amount.(*beerjson.UnitType); ok {
		miscellaneousInventoryType.Amount = &beerproto.MiscellaneousInventoryType_Unit{
			Unit: ToProtoUnitType(unit),
		}
	}
	if volume, ok :=i.Amount.(*beerjson.VolumeType); ok {
		miscellaneousInventoryType.Amount = &beerproto.MiscellaneousInventoryType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return miscellaneousInventoryType
}


func ToProtoRecipeType(i beerjson.RecipeType) *beerproto.RecipeType{
	return &beerproto.RecipeType{
		Efficiency: ToProtoEfficiencyType(i.Efficiency),
		Style: ToProtoRecipeStyleType(i.Style),
		IbuEstimate: ToProtoIBUEstimateType(i.IbuEstimate),
		ColorEstimate: ToProtoColorType(i.ColorEstimate),
		BeerPH: ToProtoAcidityType(i.BeerPH),
		Name: i.Name,
		Type: ToProtoRecipeTypeType(i.RecipeTypeType),
		Coauthor: *i.Coauthor,
		OriginalGravity: ToProtoGravityType(i.OriginalGravity),
		FinalGravity: ToProtoGravityType(i.FinalGravity),
		Carbonation: *i.Carbonation,
		Fermentation: ToProtoFermentationProcedureType(i.Fermentation),
		Author: i.Author,
		Ingredients: ToProtoIngredientsType(i.Ingredients),
		Mash: ToProtoMashProcedureType(i.Mash),
		Packaging: ToProtoPackagingProcedureType(i.Packaging),
		Boil: ToProtoBoilProcedureType(i.Boil),
		Taste: ToProtoTasteType(i.Taste),
		CaloriesPerPint: *i.CaloriesPerPint,
		Created: string(*i.Created),
		BatchSize: ToProtoVolumeType(&i.BatchSize),
		Notes: *i.Notes,
		AlcoholByVolume: ToProtoPercentType(i.AlcoholByVolume),
		ApparentAttenuation: ToProtoPercentType(i.ApparentAttenuation),
	}
}

func ToProtoTasteType(i *beerjson.TasteType) *beerproto.TasteType{
	if i == nil {
		return nil
	}
	return &beerproto.TasteType{
		Notes: i.Notes,
		Rating: i.Rating,
	}
}

func ToProtoBoilProcedureType(i *beerjson.BoilProcedureType) *beerproto.BoilProcedureType{
	if i == nil {
		return nil
	}
	boilSteps := make([]*beerproto.BoilStepType, 0)

	for _, step := range i.BoilSteps {
		boilSteps = append(boilSteps, ToProtoBoilStepType(step))
	}

	return &beerproto.BoilProcedureType{
		PreBoilSize: ToProtoVolumeType(i.PreBoilSize),
		BoilTime: ToProtoTimeType(&i.BoilTime),
		Name: *i.Name,
		Description: *i.Description,
		Notes: *i.Notes,
		BoilSteps: boilSteps,
	}
}

func ToProtoBoilStepType(i beerjson.BoilStepType) *beerproto.BoilStepType{
	return &beerproto.BoilStepType{
		EndGravity: ToProtoGravityType(i.EndGravity),
		ChillingType:ToProtoChillingType(i.BoilStepTypeChillingType),
		Description: *i.Description,
		EndTemperature: ToProtoTemperatureType(i.EndTemperature),
		RampTime: ToProtoTimeType(i.RampTime),
		StepTime: ToProtoTimeType(i.StepTime),
		StartGravity: ToProtoGravityType(i.StartGravity),
		StartPh: ToProtoAcidityType(i.StartPh),
		EndPh: ToProtoAcidityType(i.EndPh),
		Name: i.Name,
		StartTemperature: ToProtoTemperatureType(i.StartTemperature),
	}
}

func ToProtoChillingType(i *beerjson.BoilStepTypeChillingType) beerproto.BoilStepType_BoilStepTypeChillingType{
	if i == nil {
		return beerproto.BoilStepType_NULL
	}
	unit := beerproto.BoilStepType_BoilStepTypeChillingType_value[strings.ToUpper(string(*i))]
	return beerproto.BoilStepType_BoilStepTypeChillingType(unit)
}

func ToProtoPackagingProcedureType(i *beerjson.PackagingProcedureType) *beerproto.PackagingProcedureType{
	if i == nil {
		return nil
	}
	packagingVessels := make([]*beerproto.PackagingVesselType, 0)

	for _, vessels := range i.PackagingVessels {
		packagingVessels = append(packagingVessels, ToProtoPackagingVesselType(vessels))
	}
	return &beerproto.PackagingProcedureType{
		Name: i.Name,
		PackagedVolume: ToProtoVolumeType(i.PackagedVolume),
		Description: *i.Description,
		Notes: *i.Notes,
		PackagingVessels: packagingVessels,
	}
}

func ToProtoPackagingVesselType(i beerjson.PackagingVesselType) *beerproto.PackagingVesselType{
	return &beerproto.PackagingVesselType{
		Type: ToProtoPackagingVesselTypeType(i.PackagingVesselTypeType),
		StartGravity: ToProtoGravityType(i.StartGravity),
		Name: i.Name,
		PackageDate: string(*i.PackageDate),
		StepTime: ToProtoTimeType(i.StepTime),
		EndGravity: ToProtoGravityType(i.EndGravity),
		VesselVolume: ToProtoVolumeType(i.VesselVolume),
		VesselQuantity: *i.VesselQuantity,
		Description: *i.Description,
		StartPh: ToProtoAcidityType(i.StartPh),
		Carbonation: *i.Carbonation,
		StartTemperature: ToProtoTemperatureType(i.StartTemperature),
		EndPh: ToProtoAcidityType(i.EndPh),
		EndTemperature: ToProtoTemperatureType(i.EndTemperature),
	}
}

func ToProtoPackagingVesselTypeType(i *beerjson.PackagingVesselTypeType) beerproto.PackagingVesselType_PackagingVesselTypeType{
	if i == nil {
		return beerproto.PackagingVesselType_NULL
	}
	unit:= beerproto.PackagingVesselType_PackagingVesselTypeType_value[strings.ToUpper(string(*i))]
	return beerproto.PackagingVesselType_PackagingVesselTypeType(unit)
}

func ToProtoIngredientsType(i beerjson.IngredientsType) *beerproto.IngredientsType{
	miscellaneousAdditions := make([]*beerproto.MiscellaneousAdditionType, 0)
	cultureAdditions := make([]*beerproto.CultureAdditionType, 0)
	waterAdditions := make([]*beerproto.WaterAdditionType, 0)
	fermentableAdditions := make([]*beerproto.FermentableAdditionType, 0)
	hopAdditions := make([]*beerproto.HopAdditionType, 0)

	for _, misc := range i.MiscellaneousAdditions {
		miscellaneousAdditions = append(miscellaneousAdditions, ToProtoMiscellaneousAdditionType(misc))
	}
	for _, culture := range i.CultureAdditions {
		cultureAdditions = append(cultureAdditions, ToProtoCultureAdditionType(culture))
	}
	for _, water := range i.WaterAdditions {
		waterAdditions = append(waterAdditions, ToProtoWaterAdditionType(water))
	}
	for _, fermentable := range i.FermentableAdditions {
		fermentableAdditions = append(fermentableAdditions, ToProtoFermentableAdditionType(fermentable))
	}
	for _, hop := range i.HopAdditions {
		hopAdditions = append(hopAdditions, ToProtoHopAdditionType(hop))
	}
	return &beerproto.IngredientsType{
		MiscellaneousAdditions: miscellaneousAdditions,
		CultureAdditions: cultureAdditions,
		WaterAdditions: waterAdditions,
		FermentableAdditions: fermentableAdditions,
		HopAdditions: hopAdditions,
	}
}

func ToProtoHopAdditionType(i beerjson.HopAdditionType) *beerproto.HopAdditionType{
	hopAdditionType := &beerproto.HopAdditionType{
		BetaAcid: ToProtoPercentType(i.BetaAcid),
		Producer: *i.Producer,
		Origin: *i.Origin,
		Year: *i.Year,
		Form: ToProtoHopVarietyBaseForm(i.HopVarietyBaseForm),
		Timing: ToProtoTimingType(&i.Timing),
		Name: *i.Name,
		ProductId: *i.ProductId,
		AlphaAcid: ToProtoPercentType(i.AlphaAcid),
	}

	if mass, ok :=i.Amount.(*beerjson.MassType); ok {
		hopAdditionType.Amount = &beerproto.HopAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if volume, ok :=i.Amount.(*beerjson.VolumeType); ok {
		hopAdditionType.Amount = &beerproto.HopAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return hopAdditionType
}

func ToProtoHopVarietyBaseForm(i *beerjson.HopVarietyBaseForm) beerproto.HopAdditionType_HopVarietyBaseForm{
	if i == nil {
		return beerproto.HopAdditionType_NULL
	}
	unit := beerproto.HopAdditionType_HopVarietyBaseForm_value[strings.ToUpper(string(*i))]
	return beerproto.HopAdditionType_HopVarietyBaseForm(unit)
}

func ToProtoFermentableAdditionType(i beerjson.FermentableAdditionType) *beerproto.FermentableAdditionType {
	fermentableAdditionType := &beerproto.FermentableAdditionType{
		Type: ToProtoFermentableBaseType(i.FermentableBaseType),
		Origin: *i.Origin,
		GrainGroup: ToProtoFermentableBaseGrainGroup(i.FermentableBaseGrainGroup),
		Yield: ToProtoYieldType(i.Yield),
		Color: ToProtoColorType(i.Color),
		Name: *i.Name,
		Producer: *i.Producer,
		ProductId: *i.ProductId,
		Timing: ToProtoTimingType(i.Timing),
	}

	if mass, ok :=i.Amount.(*beerjson.MassType); ok {
		fermentableAdditionType.Amount = &beerproto.FermentableAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if volume, ok :=i.Amount.(*beerjson.VolumeType); ok {
		fermentableAdditionType.Amount = &beerproto.FermentableAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return fermentableAdditionType
}

func ToProtoYieldType(i *beerjson.YieldType) *beerproto.YieldType{
	if i == nil {
		return nil
	}
	return &beerproto.YieldType{
		FineGrind: ToProtoPercentType(i.FineGrind),
		CoarseGrind: ToProtoPercentType(i.CoarseGrind),
		FineCoarseDifference: ToProtoPercentType(i.FineCoarseDifference),
		Potential: ToProtoGravityType(i.Potential),
	}
}

func ToProtoFermentableBaseGrainGroup(i *beerjson.FermentableBaseGrainGroup) beerproto.FermentableAdditionType_FermentableBaseGrainGroup{
	if i == nil {
		return beerproto.FermentableAdditionType_NULL_FermentableBaseGrainGroup
	}
	unit:= beerproto.FermentableAdditionType_FermentableBaseGrainGroup_value[strings.ToUpper(string(*i))]
	return beerproto.FermentableAdditionType_FermentableBaseGrainGroup(unit)
}

func ToProtoFermentableBaseType(i *beerjson.FermentableBaseType) beerproto.FermentableAdditionType_FermentableBaseType{
	if i == nil {
		return beerproto.FermentableAdditionType_NULL_FermentableBaseType
	}
	unit := beerproto.FermentableAdditionType_FermentableBaseType_value[strings.ToUpper(string(*i))]
	return beerproto.FermentableAdditionType_FermentableBaseType(unit)
}

func ToProtoWaterAdditionType(i beerjson.WaterAdditionType) *beerproto.WaterAdditionType{
	return &beerproto.WaterAdditionType{
		Flouride: ToProtoConcentrationType(i.Flouride),
		Sulfate: ToProtoConcentrationType(i.Sulfate),
		Producer: *i.Producer,
		Nitrate: ToProtoConcentrationType(i.Nitrate),
		Nitrite: ToProtoConcentrationType(i.Nitrite),
		Chloride: ToProtoConcentrationType(i.Chloride),
		Amount: ToProtoVolumeType(i.Amount),
		Name: *i.Name,
		Potassium: ToProtoConcentrationType(i.Potassium),
		Magnesium: ToProtoConcentrationType(i.Magnesium),
		Iron: ToProtoConcentrationType(i.Iron),
		Bicarbonate: ToProtoConcentrationType(i.Bicarbonate),
		Calcium: ToProtoConcentrationType(i.Calcium),
		Carbonate: ToProtoConcentrationType(i.Carbonate),
		Sodium: ToProtoConcentrationType(i.Sodium),
	}
}

func ToProtoConcentrationType(i *beerjson.ConcentrationType) *beerproto.ConcentrationType{
	if i == nil {
		return nil
	}
	
	return &beerproto.ConcentrationType{
		Value: i.Value,
		Unit: ToProtoConcentrationUnitType(i.Unit),
	}
}

func ToProtoConcentrationUnitType(i beerjson.ConcentrationUnitType) beerproto.ConcentrationType_ConcentrationUnitType{
	unit := beerproto.ConcentrationType_ConcentrationUnitType_value[strings.ToUpper(string(i))]
	return beerproto.ConcentrationType_ConcentrationUnitType(unit)
}


func ToProtoCultureAdditionType(i beerjson.CultureAdditionType) *beerproto.CultureAdditionType{
	cultureAdditionType := &beerproto.CultureAdditionType{
		Form: ToProtoCultureBaseForm(i.CultureBaseForm),
		ProductId: *i.ProductId,
		Name: *i.Name,
		CellCountBillions: *i.CellCountBillions,
		TimesCultured: *i.TimesCultured,
		Producer: *i.Producer,
		Type: ToProtoCultureBaseType(i.CultureBaseType),
		Attenuation: ToProtoPercentType(i.Attenuation),
		Timing: ToProtoTimingType(i.Timing),
	}

	if mass, ok :=i.Amount.(*beerjson.MassType); ok {
		cultureAdditionType.Amount = &beerproto.CultureAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if unit, ok :=i.Amount.(*beerjson.UnitType); ok {
		cultureAdditionType.Amount = &beerproto.CultureAdditionType_Unit{
			Unit: ToProtoUnitType(unit),
		}
	}
	if volume, ok :=i.Amount.(*beerjson.VolumeType); ok {
		cultureAdditionType.Amount = &beerproto.CultureAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return cultureAdditionType
}

func ToProtoCultureBaseType(i *beerjson.CultureBaseType) beerproto.CultureAdditionType_CultureBaseType{
	if i == nil {
		return beerproto.CultureAdditionType_NULL_CultureBaseType
	}
	unit := beerproto.CultureAdditionType_CultureBaseType_value[strings.ToUpper(string(*i))]
	return beerproto.CultureAdditionType_CultureBaseType(unit)
}

func ToProtoCultureBaseForm(i *beerjson.CultureBaseForm) beerproto.CultureAdditionType_CultureBaseForm{
	if i == nil {
		return beerproto.CultureAdditionType_NULL_CultureBaseForm
	}
	unit := beerproto.CultureAdditionType_CultureBaseForm_value[strings.ToUpper(string(*i))]
	return beerproto.CultureAdditionType_CultureBaseForm(unit)
}

func ToProtoMiscellaneousAdditionType(i beerjson.MiscellaneousAdditionType) *beerproto.MiscellaneousAdditionType{
	miscellaneousAdditionType := &beerproto.MiscellaneousAdditionType{
		Name: *i.Name,
		Producer: *i.Producer,
		Timing: ToProtoTimingType(i.Timing),
		ProductId: *i.ProductId,
		Type: ToProtoMiscellaneousBaseType(i.MiscellaneousBaseType),
	}

	if mass, ok :=i.Amount.(*beerjson.MassType); ok {
		miscellaneousAdditionType.Amount = &beerproto.MiscellaneousAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if unit, ok :=i.Amount.(*beerjson.UnitType); ok {
		miscellaneousAdditionType.Amount = &beerproto.MiscellaneousAdditionType_Unit{
			Unit: ToProtoUnitType(unit),
		}
	}
	if volume, ok :=i.Amount.(*beerjson.VolumeType); ok {
		miscellaneousAdditionType.Amount = &beerproto.MiscellaneousAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return miscellaneousAdditionType
}

func ToProtoUnitType(i *beerjson.UnitType) *beerproto.UnitType{
	if i == nil {
		return nil
	}
	return &beerproto.UnitType{
		Value: i.Value,
		Unit: ToProtoUnitUnitType(i.Unit),
	}
}

func ToProtoUnitUnitType(i beerjson.UnitUnitType) beerproto.UnitType_UnitUnitType{
	unit := beerproto.UnitType_UnitUnitType_value[strings.ToUpper(string(i))]
	return beerproto.UnitType_UnitUnitType(unit)
}


func ToProtoMassType(i *beerjson.MassType) *beerproto.MassType{
	if i == nil {
		return nil
	}
	return &beerproto.MassType{
		Value: i.Value,
		Unit: ToProtoMassUnitType(i.Unit),
	}
}

func ToProtoMassUnitType(i beerjson.MassUnitType) beerproto.MassType_MassUnitType {
	unit := beerproto.MassType_MassUnitType_value[strings.ToUpper(string(i))]
	return beerproto.MassType_MassUnitType(unit)
}


func ToProtoMiscellaneousBaseType(i *beerjson.MiscellaneousBaseType) beerproto.MiscellaneousBaseType{
	unit := beerproto.MiscellaneousBaseType_value[strings.ToUpper(string(*i))]
	return beerproto.MiscellaneousBaseType(unit)
}

func ToProtoTimingType(i *beerjson.TimingType) *beerproto.TimingType{
	if i == nil {
		return nil
	}
	return &beerproto.TimingType{
		Time: ToProtoTimeType(i.Time),
		Duration: ToProtoTimeType(i.Duration),
		Continuous: *i.Continuous,
		SpecificGravity: ToProtoGravityType(i.SpecificGravity),
		Ph: ToProtoAcidityType(i.PH),
		Step: *i.Step,
		Use:ToProtoUseType(i.Use),
	}
}

func ToProtoUseType(i *beerjson.UseType) beerproto.TimingType_UseType{
	if i == nil {
		return beerproto.TimingType_NULL
	}
	unit := beerproto.TimingType_UseType_value[strings.ToUpper(string(*i))]
	return beerproto.TimingType_UseType(unit)
}

func ToProtoFermentationProcedureType(i *beerjson.FermentationProcedureType) *beerproto.FermentationProcedureType{
	if i == nil {
		return nil
	}
	steps := make([]*beerproto.FermentationStepType, 0)
	for _, step := range i.FermentationSteps {
		steps = append(steps, ToProtoFermentationStepType(step))
	}
	return &beerproto.FermentationProcedureType{
		Description: *i.Description,
		Notes: *i.Notes,
		Name: i.Name,
		FermentationSteps: steps,
	}
}

func ToProtoFermentationStepType(i beerjson.FermentationStepType) *beerproto.FermentationStepType{
	return &beerproto.FermentationStepType{
		Name: i.Name,
		EndTemperature: ToProtoTemperatureType(i.EndTemperature),
		StepTime: ToProtoTimeType(i.StepTime),
		FreeRise: *i.FreeRise,
		StartGravity: ToProtoGravityType(i.StartGravity),
		EndGravity: ToProtoGravityType(i.EndGravity),
		StartPh: ToProtoAcidityType(i.StartPh),
		Description: *i.Description,
		StartTemperature: ToProtoTemperatureType(i.StartTemperature),
		EndPh: ToProtoAcidityType(i.EndPh),
		Vessel: *i.Vessel,
	}
}

func ToProtoGravityType(i *beerjson.GravityType) *beerproto.GravityType{
	if i == nil {
		return nil
	}
	return &beerproto.GravityType{
		Value: i.Value,
		Unit: ToProtoGravityUnitType(i.Unit),
	}
}

func ToProtoGravityUnitType(i beerjson.GravityUnitType) beerproto.GravityType_GravityUnitType{
	unit:= beerproto.GravityType_GravityUnitType_value[strings.ToUpper(string(i))]
	return beerproto.GravityType_GravityUnitType(unit)
}

func ToProtoRecipeTypeType(i beerjson.RecipeTypeType) beerproto.RecipeType_RecipeTypeType {
	unit := beerproto.RecipeType_RecipeTypeType_value[strings.ToUpper(string(i))]
	return beerproto.RecipeType_RecipeTypeType(unit)
}

func ToProtoColorType(i *beerjson.ColorType) *beerproto.ColorType{
	if i == nil {
		return nil
	}
	return &beerproto.ColorType{
		Value: i.Value,
		Unit: ToProtoColorUnitType(i.Unit),
	}
}

func ToProtoColorUnitType(i beerjson.ColorUnitType) beerproto.ColorType_ColorUnitType {
	unit := beerproto.ColorType_ColorUnitType_value[strings.ToUpper(string(i))]
	return beerproto.ColorType_ColorUnitType(unit)
}

func ToProtoIBUEstimateType(i *beerjson.IBUEstimateType) *beerproto.IBUEstimateType{
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
	unit := beerproto.IBUEstimateType_IBUMethodType_value[strings.ToUpper(string(*i))]
	return beerproto.IBUEstimateType_IBUMethodType(unit)
}

func ToProtoRecipeStyleType(i *beerjson.RecipeStyleType) *beerproto.RecipeStyleType {
	if i == nil {
		return nil
	}
	return &beerproto.RecipeStyleType{
		Type: ToProtoRecipeStyleType_StyleCategories(i.KeyType),
		Name: *i.Name,
		Category: *i.Category,
		CategoryNumber: *i.CategoryNumber,
		StyleGuide: *i.StyleGuide,
		StyleLetter: *i.StyleLetter,
	}
}

func ToProtoRecipeStyleType_StyleCategories(i *beerjson.StyleCategories) beerproto.RecipeStyleType_StyleCategories {
	if i == nil {
		return beerproto.RecipeStyleType_NULL
	}

	unit := beerproto.RecipeStyleType_StyleCategories_value[strings.ToUpper(string(*i))]
	return beerproto.RecipeStyleType_StyleCategories(unit)
}

func ToProtoEfficiencyType(i beerjson.EfficiencyType) *beerproto.EfficiencyType{
	return &beerproto.EfficiencyType{
		Conversion: ToProtoPercentType(i.Conversion),
		Lauter: ToProtoPercentType(i.Lauter),
		Mash: ToProtoPercentType(i.Mash),
		Brewhouse: ToProtoPercentType(&i.Brewhouse),
	}
}

func ToProtoPercentType(i *beerjson.PercentType) *beerproto.PercentType {
	if i == nil {
		return nil
	}
	return &beerproto.PercentType{
		Value: i.Value,
		Unit: ToProtoPercentUnitType(i.Unit),
	}
}

func ToProtoPercentUnitType(i beerjson.PercentUnitType) beerproto.PercentType_PercentUnitType {
	unit := beerproto.PercentType_PercentUnitType_value[strings.ToUpper(string(i))]
	return beerproto.PercentType_PercentUnitType(unit)
}



func ToProtoMashProcedureType(i *beerjson.MashProcedureType) *beerproto.MashProcedureType {
	if i == nil {
		return nil
	}
	return &beerproto.MashProcedureType{
		Name:             i.Name,
		Notes:            *i.Notes,
		GrainTemperature: ToProtoTemperatureType(&i.GrainTemperature),
	}
}

func ToProtoMashStepType(i beerjson.MashStepType) *beerproto.MashStepType {
	return &beerproto.MashStepType{
		StepTime:          ToProtoTimeType(&i.StepTime),
		RampTime:          ToProtoTimeType(i.RampTime),
		EndTemperature:    ToProtoTemperatureType(i.EndTemperature),
		Description:       *i.Description,
		InfuseTemperature: ToProtoTemperatureType(i.InfuseTemperature),
		StartPH:           ToProtoAcidityType(i.StartPH),
		EndPH:             ToProtoAcidityType(i.EndPH),
		Name:              i.Name,
		Type:              ToProtoMashStepTypeType(i.MashStepTypeType),
		Amount:            ToProtoVolumeType(i.Amount),
		StepTemperature:   ToProtoTemperatureType(&i.StepTemperature),
		WaterGrainRatio:   ToProtoSpecificVolumeType(i.WaterGrainRatio),
	}
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

func ToProtoSpecificVolumeType(i *beerjson.SpecificVolumeType) *beerproto.SpecificVolumeType {
	if i == nil {
		return nil
	}
	return &beerproto.SpecificVolumeType{
		Value: i.Value,
		Unit:  ToProtoSpecificVolumeUnitType(i.Unit),
	}
}

func ToProtoSpecificVolumeUnitType(i beerjson.SpecificVolumeUnitType) beerproto.SpecificVolumeType_SpecificVolumeUnitType {
	unit := beerproto.SpecificVolumeType_SpecificVolumeUnitType_value[strings.ToUpper(string(i))]
	return beerproto.SpecificVolumeType_SpecificVolumeUnitType(unit)
}

func ToProtoMashStepTypeType(i beerjson.MashStepTypeType) beerproto.MashStepType_MashStepTypeType {
	unit := beerproto.MashStepType_MashStepTypeType_value[strings.ToUpper(string(i))]
	return beerproto.MashStepType_MashStepTypeType(unit)
}

func ToProtoAcidityType(i *beerjson.AcidityType) *beerproto.AcidityType {
	if i == nil {
		return nil
	}
	return &beerproto.AcidityType{
		Value: i.Value,
		Unit:  ToProtoAcidityUnitType(i.Unit),
	}
}

func ToProtoAcidityUnitType(i beerjson.AcidityUnitType) beerproto.AcidityType_AcidityUnitType {
	unit := beerproto.AcidityType_AcidityUnitType_value[strings.ToUpper(string(i))]
	return beerproto.AcidityType_AcidityUnitType(unit)
}

func ToProtoTimeType(i *beerjson.TimeType) *beerproto.TimeType {
	if i == nil {
		return nil
	}
	return &beerproto.TimeType{
		Value: i.Value,
		Unit:  ToProtoTimeUnitType(i.Unit),
	}
}

func ToProtoTemperatureType(i *beerjson.TemperatureType) *beerproto.TemperatureType {
	if i == nil {
		return nil
	}
	return &beerproto.TemperatureType{
		Value: i.Value,
		Unit:  ToProtoTemperatureUnitType(i.Unit),
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
