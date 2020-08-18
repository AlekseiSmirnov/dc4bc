package internal

import (
	"crypto/rsa"
	"time"
)

type ConfirmationParticipantStatus uint8

const (
	SigConfirmationAwaitConfirmation ConfirmationParticipantStatus = iota
	SigConfirmationConfirmed
	SigConfirmationDeclined
	SigConfirmationError
)

func (s ConfirmationParticipantStatus) String() string {
	var str = "undefined"
	switch s {
	case SigConfirmationAwaitConfirmation:
		str = "SigConfirmationAwaitConfirmation"
	case SigConfirmationConfirmed:
		str = "SigConfirmationConfirmed"
	case SigConfirmationDeclined:
		str = "SigConfirmationDeclined"
	case SigConfirmationError:
		str = "SigConfirmationError"
	}
	return str
}

type SignatureConfirmation struct {
	Quorum    SignatureProposalQuorum
	CreatedAt time.Time
	ExpiresAt time.Time
}

type SignatureProposalParticipant struct {
	// Public title for address, such as name, nickname, organization
	ParticipantId int
	Title         string
	PubKey        *rsa.PublicKey
	DkgPubKey     []byte
	// For validation user confirmation: sign(InvitationSecret, PubKey) => user
	InvitationSecret string
	Status           ConfirmationParticipantStatus
	UpdatedAt        time.Time
}

// Unique alias for map iteration - Public Key Fingerprint
// Excludes array merge and rotate operations
type SignatureProposalQuorum map[string]*SignatureProposalParticipant

// DKG proposal

type DKGParticipantStatus uint8

const (
	CommitAwaitConfirmation DKGParticipantStatus = iota
	CommitConfirmed
	CommitConfirmationError
	DealAwaitConfirmation
	DealConfirmed
	DealConfirmationError
	ResponseAwaitConfirmation
	ResponseConfirmed
	ResponseConfirmationError
	MasterKeyAwaitConfirmation
	MasterKeyConfirmed
	MasterKeyConfirmationError
)

type DKGProposalParticipant struct {
	Title     string
	PubKey    []byte
	Commit    []byte
	Deal      []byte
	Response  []byte
	MasterKey []byte
	Status    DKGParticipantStatus
	Error     error
	UpdatedAt time.Time
}

type DKGProposalQuorum map[int]*DKGProposalParticipant

type DKGConfirmation struct {
	Quorum    DKGProposalQuorum
	CreatedAt *time.Time
	ExpiresAt *time.Time
}

type DKGProposalParticipantStatus uint8

func (s DKGParticipantStatus) String() string {
	var str = "undefined"
	switch s {
	case CommitAwaitConfirmation:
		str = "CommitAwaitConfirmation"
	case CommitConfirmed:
		str = "CommitConfirmed"
	case CommitConfirmationError:
		str = "CommitConfirmationError"
	case DealAwaitConfirmation:
		str = "DealAwaitConfirmation"
	case DealConfirmed:
		str = "DealConfirmed"
	case DealConfirmationError:
		str = "DealConfirmationError"
	case ResponseAwaitConfirmation:
		str = "ResponseAwaitConfirmation"
	case ResponseConfirmed:
		str = "ResponseConfirmed"
	case ResponseConfirmationError:
		str = "ResponseConfirmationError"
	case MasterKeyAwaitConfirmation:
		str = "MasterKeyAwaitConfirmation"
	case MasterKeyConfirmed:
		str = "MasterKeyConfirmed"
	case MasterKeyConfirmationError:
		str = "MasterKeyConfirmationError"
	}
	return str
}

// Signing proposal

type SigningConfirmation struct {
	Quorum           SigningProposalQuorum
	RecoveredKey     []byte
	SrcPayload       []byte
	EncryptedPayload []byte
	CreatedAt        time.Time
	ExpiresAt        time.Time
}

type SigningProposalQuorum map[int]*SigningProposalParticipant

type SigningParticipantStatus uint8

const (
	SigningIdle SigningParticipantStatus = iota
	SigningAwaitConfirmation
	SigningConfirmed
	SigningDeclined
	SigningAwaitPartialKeys
	SigningPartialKeysConfirmed
	SigningError
	SigningProcess
)

func (s SigningParticipantStatus) String() string {
	var str = "undefined"
	switch s {
	case SigningIdle:
		str = "SigningIdle"
	case SigningAwaitConfirmation:
		str = "SigningAwaitConfirmation"
	case SigningConfirmed:
		str = "SigningConfirmed"
	case SigningAwaitPartialKeys:
		str = "SigningAwaitPartialKeys"
	case SigningPartialKeysConfirmed:
		str = "SigningPartialKeysConfirmed"
	case SigningError:
		str = "SigningError"
	case SigningProcess:
		str = "SigningProcess"
	}
	return str
}

type SigningProposalParticipant struct {
	Title      string
	Status     SigningParticipantStatus
	PartialKey []byte
	Error      error
	UpdatedAt  time.Time
}
