package mollie

import "time"

// Capabilities describes what capabilities the organization can perform.
type Capabilities struct {
	Resource     string                    `json:"resource,omitempty"`
	Name         string                    `json:"name,omitempty"`
	Status       CapabilitiesStatus        `json:"status,omitempty"`
	StatusReason *CapabilitiesStatusReason `json:"statusReason,omitempty"`
	Requirements []CapabilityRequirement   `json:"requirements,omitempty"`
}

// CapabilitiesStatus describes status of capability.
type CapabilitiesStatus string

// Valid capability status.
const (
	CapabilityUnrequested CapabilitiesStatus = "unrequested"
	CapabilityEnabled     CapabilitiesStatus = "enabled"
	CapabilityDisabled    CapabilitiesStatus = "disabled"
	CapabilityPending     CapabilitiesStatus = "pending"
)

// CapabilitiesStatusReason describes reason for status of capability.
type CapabilitiesStatusReason string

// Valid capability status reasons.
const (
	CapabilitiesStatusReasonRequirementPastDue        CapabilitiesStatusReason = "requirement-past-due"
	CapabilityStatusReasonOnboardingInformationNeeded CapabilitiesStatusReason = "onboarding-information-needed"
)

// CapabilityRequirement referring to the task to be fulfilled by the organization to enable or re-enable the capability
type CapabilityRequirement struct {
	Id      string                      `json:"id"`
	Status  CapabilityRequirementStatus `json:"status"`
	DueDate *time.Time                  `json:"dueDate,omitempty"`
	Links   CapabilityRequirementLinks  `json:"_links,omitempty"`
}

// CapabilityRequirementStatus describes status of capability requirement.
type CapabilityRequirementStatus string

// Valid capability requirement status.
const (
	CapabilityRequirementCurrentlyDue CapabilityRequirementStatus = "currently-due"
	CapabilityRequirementPastDue      CapabilityRequirementStatus = "past-due"
	CapabilityRequirementRequested    CapabilityRequirementStatus = "requested"
)

// CapabilityRequirementLinks contains URL objects relevant to the requirements.
type CapabilityRequirementLinks struct {
	Dashboard *URL `json:"dashboard,omitempty"`
}
