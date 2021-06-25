package lti

type definition string

type Claims struct {
	Context                  definition
	ResourceLink             definition
	ToolPlatform             definition
	Ags                      definition
	Mentor                   definition
	Roles                    definition
	Custom                   definition
	Extension                definition
	Lis                      definition
	TargetLinkUri            definition
	Lti11LegacyUserId        definition
	DeepLinking              definition
	DeepLinkingData          definition
	DeepLinkingToolMsg       definition
	DeepLinkingToolLog       definition
	ContentItem              definition
	NamesAndRoles            definition
	ToolLaunchCaliperContext definition
	ToolUseCaliperContext    definition
	Caliper                  definition
}

type Scopes struct {
	AgsLineItem   definition
	AgsResult     definition
	AgsScore      definition
	NamesAndRoles definition
	Caliper       definition
	Student       definition
	Instructor    definition
	Learner       definition
	Mentor        definition
	MentorRole    definition
}

type Contexts struct {
	Course  definition
	Account definition
}

type Roles struct {
	CanvasPublicLtiKeysUrl definition
	CanvasOidcUrl          definition
	CanvasAuthTokenUrl     definition

	CanvasBetaPublicLtiKeysUrl definition
	CanvasBetaAuthTokenUrl     definition
	CanvasBetaOidcUrl          definition

	CanvasSubmissionType definition

	AdministratorSystemRole           definition
	NoneSystemRole                    definition
	AccountAdminSystemRole            definition
	CreatorSystemRole                 definition
	SysAdminSystemRole                definition
	SysSupportSystemRole              definition
	UserSystemRole                    definition
	AdministratorInstitutionRole      definition
	FacultyInstitutionRole            definition
	GuestInstitutionRole              definition
	NoneInstitutionRole               definition
	OtherInstitutionRole              definition
	StaffInstitutionRole              definition
	StudentInstitutionRole            definition
	AlumniInstitutionRole             definition
	InstructorInstitutionRole         definition
	LearnerInstitutionRole            definition
	MemberInstitutionRole             definition
	MentorInstitutionRole             definition
	ObserverInstitutionRole           definition
	ProspectiveStudentInstitutionRole definition
	AdministratorContextRole          definition
	ContentDeveloperContextRole       definition
	InstructorContextRole             definition
	LearnerContextRole                definition
	MentorContextRole                 definition
	ManagerContextRole                definition
	MemberContextRole                 definition
	OfficerContextRole                definition
}

type LtiDefinitions struct {
	LtiVersion         definition
	LaunchPresentation definition
	DeploymentId       definition
	MessageType        definition

	Claims   Claims
	Scopes   Scopes
	Contexts Contexts
	Roles    Roles
}

var Definitions LtiDefinitions = LtiDefinitions{
	LtiVersion:         "https://purl.imsglobal.org/spec/lti/claim/version",
	LaunchPresentation: "https://purl.imsglobal.org/spec/lti/claim/launch_presentation",
	DeploymentId:       "https://purl.imsglobal.org/spec/lti/claim/deployment_id",
	MessageType:        "https://purl.imsglobal.org/spec/lti/claim/message_type",

	Claims: Claims{
		Context:                  "https://purl.imsglobal.org/spec/lti/claim/context",
		ResourceLink:             "https://purl.imsglobal.org/spec/lti/claim/resource_link",
		ToolPlatform:             "https://purl.imsglobal.org/spec/lti/claim/tool_platform",
		Ags:                      "https://purl.imsglobal.org/spec/lti-ags/claim/endpoint",
		Mentor:                   "https://purl.imsglobal.org/spec/lti/claim/role_scope_mentor",
		Roles:                    "https://purl.imsglobal.org/spec/lti/claim/roles",
		Custom:                   "https://purl.imsglobal.org/spec/lti/claim/custom",
		Extension:                "http://www.ExamplePlatformVendor.com/session",
		Lis:                      "https://purl.imsglobal.org/spec/lti/claim/lis",
		TargetLinkUri:            "https://purl.imsglobal.org/spec/lti/claim/target_link_uri",
		Lti11LegacyUserId:        "https://purl.imsglobal.org/spec/lti/claim/lti11_legacy_user_id",
		DeepLinking:              "https://purl.imsglobal.org/spec/lti-dl/claim/deep_linking_settings",
		DeepLinkingData:          "https://purl.imsglobal.org/spec/lti-dl/claim/data",
		DeepLinkingToolMsg:       "https://purl.imsglobal.org/spec/lti-dl/claim/msg",
		DeepLinkingToolLog:       "https://purl.imsglobal.org/spec/lti-dl/claim/log",
		ContentItem:              "https://purl.imsglobal.org/spec/lti-dl/claim/content_items",
		NamesAndRoles:            "https://purl.imsglobal.org/spec/lti-nrps/claim/namesroleservice",
		Caliper:                  "https://purl.imsglobal.org/spec/lti-ces/claim/caliper-endpoint-service",
		ToolLaunchCaliperContext: "http://purl.imsglobal.org/ctx/caliper/v1p1/ToolLaunchProfile-extension",
		ToolUseCaliperContext:    "http://purl.imsglobal.org/ctx/caliper/v1p1",
	},

	Scopes: Scopes{
		AgsLineItem:   "https://purl.imsglobal.org/spec/lti-ags/scope/lineitem",
		AgsResult:     "https://purl.imsglobal.org/spec/lti-ags/scope/result.readonly",
		AgsScore:      "https://purl.imsglobal.org/spec/lti-ags/scope/score",
		NamesAndRoles: "https://purl.imsglobal.org/spec/lti-nrps/scope/contextmembership.readonly",
		Caliper:       "https://purl.imsglobal.org/spec/lti-ces/v1p0/scope/send",

		Student:    "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Student",
		Instructor: "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Instructor",
		Learner:    "http://purl.imsglobal.org/vocab/lis/v2/membership#Learner",
		Mentor:     "http://purl.imsglobal.org/vocab/lis/v2/membership#Mentor",
		MentorRole: "a62c52c02ba262003f5e",
	},

	Contexts: Contexts{
		Course:  "http://purl.imsglobal.org/vocab/lis/v2/course#courseoffering",
		Account: "Account",
	},

	Roles: Roles{
		AdministratorSystemRole:           "http://purl.imsglobal.org/vocab/lis/v2/system/person#Administrator",
		NoneSystemRole:                    "http://purl.imsglobal.org/vocab/lis/v2/system/person#None",
		AccountAdminSystemRole:            "http://purl.imsglobal.org/vocab/lis/v2/system/person#AccountAdmin",
		CreatorSystemRole:                 "http://purl.imsglobal.org/vocab/lis/v2/system/person#Creator",
		SysAdminSystemRole:                "http://purl.imsglobal.org/vocab/lis/v2/system/person#SysAdmin",
		SysSupportSystemRole:              "http://purl.imsglobal.org/vocab/lis/v2/system/person#SysSupport",
		UserSystemRole:                    "http://purl.imsglobal.org/vocab/lis/v2/system/person#User",
		AdministratorInstitutionRole:      "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Administrator",
		FacultyInstitutionRole:            "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Faculty",
		GuestInstitutionRole:              "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Guest",
		NoneInstitutionRole:               "http://purl.imsglobal.org/vocab/lis/v2/institution/person#None",
		OtherInstitutionRole:              "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Other",
		StaffInstitutionRole:              "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Staff",
		StudentInstitutionRole:            "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Student",
		AlumniInstitutionRole:             "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Alumni",
		InstructorInstitutionRole:         "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Instructor",
		LearnerInstitutionRole:            "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Learner",
		MemberInstitutionRole:             "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Member",
		MentorInstitutionRole:             "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Mentor",
		ObserverInstitutionRole:           "http://purl.imsglobal.org/vocab/lis/v2/institution/person#Observer",
		ProspectiveStudentInstitutionRole: "http://purl.imsglobal.org/vocab/lis/v2/institution/person#ProspectiveStudent",
		AdministratorContextRole:          "http://purl.imsglobal.org/vocab/lis/v2/membership#Administrator",
		ContentDeveloperContextRole:       "http://purl.imsglobal.org/vocab/lis/v2/membership#ContentDeveloper",
		InstructorContextRole:             "http://purl.imsglobal.org/vocab/lis/v2/membership#Instructor",
		LearnerContextRole:                "http://purl.imsglobal.org/vocab/lis/v2/membership#Learner",
		MentorContextRole:                 "http://purl.imsglobal.org/vocab/lis/v2/membership#Mentor",
		ManagerContextRole:                "http://purl.imsglobal.org/vocab/lis/v2/membership#Manager",
		MemberContextRole:                 "http://purl.imsglobal.org/vocab/lis/v2/membership#Member",
		OfficerContextRole:                "http://purl.imsglobal.org/vocab/lis/v2/membership#Officer",
	},
}
