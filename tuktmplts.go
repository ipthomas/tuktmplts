package tuktmplts

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"

	"github.com/ipthomas/tukcnst"
	"github.com/ipthomas/tukdbint"
)

type CGL struct {
	Data struct {
		Client struct {
			BasicDetails struct {
				Address struct {
					AddressLine1 string `json:"addressLine1,omitempty"`
					AddressLine2 string `json:"addressLine2,omitempty"`
					AddressLine3 string `json:"addressLine3,omitempty"`
					AddressLine4 string `json:"addressLine4,omitempty"`
					AddressLine5 string `json:"addressLine5,omitempty"`
					PostCode     string `json:"postCode,omitempty"`
				} `json:"address,omitempty"`
				BirthDate                    string `json:"birthDate,omitempty"`
				Disability                   string `json:"disability,omitempty"`
				LastEngagementByCGLDate      string `json:"lastEngagementByCGLDate,omitempty"`
				LastFaceToFaceEngagementDate string `json:"lastFaceToFaceEngagementDate,omitempty"`
				LocalIdentifier              int    `json:"localIdentifier,omitempty"`
				Name                         struct {
					Family string `json:"family,omitempty"`
					Given  string `json:"given,omitempty"`
				} `json:"name,omitempty"`
				NextCGLAppointmentDate string `json:"nextCGLAppointmentDate,omitempty"`
				NhsNumber              string `json:"nhsNumber,omitempty"`
				SexAtBirth             string `json:"sexAtBirth,omitempty"`
			} `json:"basicDetails,omitempty"`
			BbvInformation struct {
				BbvTested        string `json:"bbvTested,omitempty"`
				HepCLastTestDate string `json:"hepCLastTestDate,omitempty"`
				HepCResult       string `json:"hepCResult,omitempty"`
				HivPositive      string `json:"hivPositive,omitempty"`
			} `json:"bbvInformation,omitempty"`
			DrugTestResults struct {
				DrugTestDate          string `json:"drugTestDate,omitempty"`
				DrugTestSample        string `json:"drugTestSample,omitempty"`
				DrugTestStatus        string `json:"drugTestStatus,omitempty"`
				InstantOrConfirmation string `json:"instantOrConfirmation,omitempty"`
				Results               struct {
					Amphetamine     string `json:"amphetamine,omitempty"`
					Benzodiazepine  string `json:"benzodiazepine,omitempty"`
					Buprenorphine   string `json:"buprenorphine,omitempty"`
					Cannabis        string `json:"cannabis,omitempty"`
					Cocaine         string `json:"cocaine,omitempty"`
					Eddp            string `json:"eddp,omitempty"`
					Fentanyl        string `json:"fentanyl,omitempty"`
					Ketamine        string `json:"ketamine,omitempty"`
					Methadone       string `json:"methadone,omitempty"`
					Methamphetamine string `json:"methamphetamine,omitempty"`
					Morphine        string `json:"morphine,omitempty"`
					Opiates         string `json:"opiates,omitempty"`
					SixMam          string `json:"sixMam,omitempty"`
					Tramadol        string `json:"tramadol,omitempty"`
				} `json:"results,omitempty"`
			} `json:"drugTestResults,omitempty"`
			PrescribingInformation []string `json:"prescribingInformation,omitempty"`
			RiskInformation        struct {
				LastSelfReportedDate string `json:"lastSelfReportedDate,omitempty"`
				MentalHealthDomain   struct {
					AttemptedSuicide                            string `json:"attemptedSuicide,omitempty"`
					CurrentOrPreviousSelfHarm                   string `json:"currentOrPreviousSelfHarm,omitempty"`
					DiagnosedMentalHealthCondition              string `json:"diagnosedMentalHealthCondition,omitempty"`
					FrequentLifeThreateningSelfHarm             string `json:"frequentLifeThreateningSelfHarm,omitempty"`
					Hallucinations                              string `json:"hallucinations,omitempty"`
					HospitalAdmissionsForMentalHealth           string `json:"hospitalAdmissionsForMentalHealth,omitempty"`
					NoIdentifiedRisk                            string `json:"noIdentifiedRisk,omitempty"`
					NotEngagingWithSupport                      string `json:"notEngagingWithSupport,omitempty"`
					NotTakingPrescribedMedicationAsInstructed   string `json:"notTakingPrescribedMedicationAsInstructed,omitempty"`
					PsychiatricOrPreviousCrisisTeamIntervention string `json:"psychiatricOrPreviousCrisisTeamIntervention,omitempty"`
					Psychosis                                   string `json:"psychosis,omitempty"`
					SelfReportedMentalHealthConcerns            string `json:"selfReportedMentalHealthConcerns,omitempty"`
					ThoughtsOfSuicideOrSelfHarm                 string `json:"thoughtsOfSuicideOrSelfHarm,omitempty"`
				} `json:"mentalHealthDomain,omitempty"`
				RiskOfHarmToSelfDomain struct {
					AssessedAsNotHavingMentalCapacity  string `json:"assessedAsNotHavingMentalCapacity,omitempty"`
					BeliefTheyAreWorthless             string `json:"beliefTheyAreWorthless,omitempty"`
					Hoarding                           string `json:"hoarding,omitempty"`
					LearningDisability                 string `json:"learningDisability,omitempty"`
					MeetsSafeguardingAdultsThreshold   string `json:"meetsSafeguardingAdultsThreshold,omitempty"`
					NoIdentifiedRisk                   string `json:"noIdentifiedRisk,omitempty"`
					OngoingConcernsRelatingToOwnSafety string `json:"ongoingConcernsRelatingToOwnSafety,omitempty"`
					ProblemsMaintainingPersonalHygiene string `json:"problemsMaintainingPersonalHygiene,omitempty"`
					ProblemsMeetingNutritionalNeeds    string `json:"problemsMeetingNutritionalNeeds,omitempty"`
					RequiresIndependentAdvocacy        string `json:"requiresIndependentAdvocacy,omitempty"`
					SelfNeglect                        string `json:"selfNeglect,omitempty"`
				} `json:"riskOfHarmToSelfDomain,omitempty"`
				SocialDomain struct {
					FinancialProblems         string `json:"financialProblems,omitempty"`
					HomelessRoughSleepingNFA  string `json:"homelessRoughSleepingNFA,omitempty"`
					HousingAtRisk             string `json:"housingAtRisk,omitempty"`
					NoIdentifiedRisk          string `json:"noIdentifiedRisk,omitempty"`
					SociallyIsolatedNoSupport string `json:"sociallyIsolatedNoSupport,omitempty"`
				} `json:"socialDomain,omitempty"`
				SubstanceMisuseDomain struct {
					ConfusionOrDisorientation string `json:"ConfusionOrDisorientation,omitempty"`
					AdmissionToAandE          string `json:"admissionToAandE,omitempty"`
					BlackoutOrSeizures        string `json:"blackoutOrSeizures,omitempty"`
					ConcurrentUse             string `json:"concurrentUse,omitempty"`
					HigherRiskDrinking        string `json:"higherRiskDrinking,omitempty"`
					InjectedByOthers          string `json:"injectedByOthers,omitempty"`
					Injecting                 string `json:"injecting,omitempty"`
					InjectingInNeckOrGroin    string `json:"injectingInNeckOrGroin,omitempty"`
					NoIdentifiedRisk          string `json:"noIdentifiedRisk,omitempty"`
					PolyDrugUse               string `json:"polyDrugUse,omitempty"`
					PreviousOverDose          string `json:"previousOverDose,omitempty"`
					RecentPrisonRelease       string `json:"recentPrisonRelease,omitempty"`
					ReducedTolerance          string `json:"reducedTolerance,omitempty"`
					SharingWorks              string `json:"sharingWorks,omitempty"`
					Speedballing              string `json:"speedballing,omitempty"`
					UsingOnTop                string `json:"usingOnTop,omitempty"`
				} `json:"substanceMisuseDomain,omitempty"`
			} `json:"riskInformation,omitempty"`
			SafeguardingInformation struct {
				LastReviewDate     string `json:"lastReviewDate,omitempty"`
				RiskHarmFromOthers string `json:"riskHarmFromOthers,omitempty"`
				RiskToAdults       string `json:"riskToAdults,omitempty"`
				RiskToChildrenOrYP string `json:"riskToChildrenOrYP,omitempty"`
				RiskToSelf         string `json:"riskToSelf,omitempty"`
			} `json:"safeguardingInformation,omitempty"`
		} `json:"client,omitempty"`
		KeyWorker struct {
			LocalIdentifier int `json:"localIdentifier,omitempty"`
			Name            struct {
				Family string `json:"family,omitempty"`
				Given  string `json:"given,omitempty"`
			} `json:"name"`
			Telecom string `json:"telecom,omitempty"`
		} `json:"keyWorker,omitempty"`
	} `json:"data,omitempty"`
}
type Delphi struct {
	Data struct {
		LocalIdentifier int    `json:"LocalIdentifier,omitempty"`
		Status          string `json:"Status,omitempty"`
		Title           string `json:"Title,omitempty"`
		Forename        string `json:"Forename,omitempty"`
		Surname         string `json:"Surname,omitempty"`
		GenderAtBirth   string `json:"GenderAtBirth,omitempty"`
		DateOfBirth     string `json:"DateOfBirth,omitempty"`
		Address         struct {
			LocalIdentifier int    `json:"LocalIdentifier,omitempty"`
			AddressLine1    string `json:"AddressLine1,omitempty"`
			AddressLine2    string `json:"AddressLine2,omitempty"`
			AddressLine3    string `json:"AddressLine3,omitempty"`
			AddressLine4    string `json:"AddressLine4,omitempty"`
			PostCode1       string `json:"PostCode1,omitempty"`
			PostCode2       string `json:"PostCode2,omitempty"`
		} `json:"Address,omitempty"`
		Keyworker               string `json:"Keyworker,omitempty"`
		LastAttendedAppointment string `json:"LastAttendedAppointment,omitempty"`
		DrugScreening           []any  `json:"DrugScreening,omitempty"`
		Prescriptions           []any  `json:"Prescriptions,omitempty"`
		Risks                   []any  `json:"Risks,omitempty"`
		Careplans               []struct {
			LocalIdentifier          int    `json:"LocalIdentifier,omitempty"`
			AlcohoUse                bool   `json:"AlcohoUse,omitempty"`
			DrugUse                  bool   `json:"DrugUse,omitempty"`
			EffectsOfAlcoholAndDrugs bool   `json:"EffectsOfAlcoholAndDrugs,omitempty"`
			PreventingRelapse        bool   `json:"PreventingRelapse,omitempty"`
			PreventingOverdose       bool   `json:"PreventingOverdose,omitempty"`
			PersonalCare             bool   `json:"PersonalCare,omitempty"`
			FindingThingsIEnjoy      bool   `json:"FindingThingsIEnjoy,omitempty"`
			ManagingMoney            bool   `json:"ManagingMoney,omitempty"`
			SupportForMyChildren     bool   `json:"SupportForMyChildren,omitempty"`
			EducationOrTraining      bool   `json:"EducationOrTraining,omitempty"`
			Other                    bool   `json:"Other,omitempty"`
			AlcoholDrugUse           bool   `json:"AlcoholDrugUse,omitempty"`
			ManagingCravings         bool   `json:"ManagingCravings,omitempty"`
			MentalEmotionalHealth    bool   `json:"MentalEmotionalHealth,omitempty"`
			AccommodationHousing     bool   `json:"AccommodationHousing,omitempty"`
			LegalProblems            bool   `json:"LegalProblems,omitempty"`
			ParentingHelpSupport     bool   `json:"ParentingHelpSupport,omitempty"`
			PhyscialHealth           bool   `json:"PhyscialHealth,omitempty"`
			ImmediateProblem         string `json:"ImmediateProblem,omitempty"`
			LongTermGoal             string `json:"LongTermGoal,omitempty"`
			StepsToAchievingGoal     string `json:"StepsToAchievingGoal,omitempty"`
			HowDidItGo               string `json:"HowDidItGo,omitempty"`
			NextStepForGoal          string `json:"NextStepForGoal,omitempty"`
			CommunityDetox           bool   `json:"CommunityDetox,omitempty"`
			InpatientDetox           bool   `json:"InpatientDetox,omitempty"`
			OverdoseInformation      bool   `json:"OverdoseInformation,omitempty"`
			NutritionalAdvice        bool   `json:"NutritionalAdvice,omitempty"`
			HepCScreening            bool   `json:"HepCScreening,omitempty"`
			HepAAndBVaccination      bool   `json:"HepAAndBVaccination,omitempty"`
			GroupWork                bool   `json:"GroupWork,omitempty"`
			OneToOneSupport          bool   `json:"OneToOneSupport,omitempty"`
			SupportWorker            bool   `json:"SupportWorker,omitempty"`
			PrescribedMedication     bool   `json:"PrescribedMedication,omitempty"`
			Stabilisation            bool   `json:"Stabilisation,omitempty"`
			MedicalReview            bool   `json:"MedicalReview,omitempty"`
			OtherClinical            bool   `json:"OtherClinical,omitempty"`
			CarePlanGivenToClient    string `json:"CarePlanGivenToClient,omitempty"`
			CareplanStartDate        string `json:"CareplanStartDate,omitempty"`
			CarePlanReviewDate       string `json:"CarePlanReviewDate,omitempty"`
		} `json:"Careplans,omitempty"`
		Discharge any `json:"Discharge,omitempty"`
	} `json:"Data,omitempty"`
}

type PIXm struct {
	ResourceType string `json:"resourceType"`
	ID           string `json:"id"`
	Type         string `json:"type"`
	Total        int    `json:"total"`
	Link         []struct {
		Relation string `json:"relation"`
		URL      string `json:"url"`
	} `json:"link"`
	Entry []struct {
		FullURL  string `json:"fullUrl"`
		Resource struct {
			ResourceType string `json:"resourceType"`
			ID           string `json:"id"`
			Identifier   []struct {
				Use    string `json:"use,omitempty"`
				System string `json:"system"`
				Value  string `json:"value"`
			} `json:"identifier"`
			Active bool `json:"active"`
			Name   []struct {
				Use    string   `json:"use"`
				Family string   `json:"family"`
				Given  []string `json:"given"`
			} `json:"name"`
			Gender    string `json:"gender"`
			BirthDate string `json:"birthDate"`
			Address   []struct {
				Use        string   `json:"use"`
				Line       []string `json:"line"`
				City       string   `json:"city"`
				PostalCode string   `json:"postalCode"`
				Country    string   `json:"country"`
			} `json:"address"`
		} `json:"resource"`
	} `json:"entry"`
}
type Transaction struct {
	Act      string
	Name     string
	Request  string
	Response string
	Template string
}
type Inject struct {
	CGL    CGL
	Delphi Delphi
	PIXm   PIXm
}
type Interface interface {
	execute() error
}

var htmlTemplates *template.Template
var isInit = true

func New_Transaction(i Interface) error {
	return i.execute()
}
func (i *Transaction) execute() error {
	var err error
	if isInit {
		if err = cacheTemplates(); err != nil {
			return err
		}
		isInit = false
	}

	var tplReturn bytes.Buffer
	var inject = Inject{}
	i.Response = i.Request

	switch i.Act {
	case tukcnst.SELECT:
		switch i.Name {
		case tukcnst.PDQ_SERVER_TYPE_CGL:
			err = json.Unmarshal([]byte(i.Request), &inject.CGL)
			if err := htmlTemplates.ExecuteTemplate(&tplReturn, i.Name, inject.CGL); err != nil {
				return err
			}
		case "delphi":
			err = json.Unmarshal([]byte(i.Request), &inject.Delphi)
			if err := htmlTemplates.ExecuteTemplate(&tplReturn, i.Name, inject.Delphi); err != nil {
				return err
			}
		case tukcnst.PDQ_SERVER_TYPE_IHE_PIXM:
			err = json.Unmarshal([]byte(i.Request), &inject.PIXm)
			if err := htmlTemplates.ExecuteTemplate(&tplReturn, i.Name, inject.PIXm); err != nil {
				return err
			}
		}
		i.Response = tplReturn.String()
	case tukcnst.INSERT:
		tmplts := tukdbint.Templates{Action: tukcnst.DELETE}
		tmplt := tukdbint.Template{Name: i.Name}
		tmplts.Templates = append(tmplts.Templates, tmplt)
		tukdbint.NewDBEvent(&tmplts)
		tmplts = tukdbint.Templates{Action: tukcnst.INSERT}
		tmplt = tukdbint.Template{Name: i.Name, IsXML: false, Template: i.Template}
		tmplts.Templates = append(tmplts.Templates, tmplt)
		if err = tukdbint.NewDBEvent(&tmplts); err != nil {
			return err
		}
	}
	return err
}
func cacheTemplates() error {
	var err error
	htmlTemplates = template.New(tukcnst.HTML)
	tmplts := tukdbint.Templates{Action: tukcnst.SELECT}
	tukdbint.NewDBEvent(&tmplts)
	log.Printf("cached %v Templates", tmplts.Count)
	for _, tmplt := range tmplts.Templates {
		if htmlTemplates, err = htmlTemplates.New(tmplt.Name).Parse(tmplt.Template); err != nil {
			return err
		}
	}
	return err
}
