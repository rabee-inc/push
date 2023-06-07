package model

import (
	"cloud.google.com/go/firestore"
	"github.com/rabee-inc/go-pkg/stringutil"
	"github.com/rabee-inc/go-pkg/timeutil"
	"github.com/rabee-inc/push/appengine/src/config"
)

func ReserveRef(cFirestore *firestore.Client, appID string) *firestore.CollectionRef {
	return AppRef(cFirestore).Doc(appID).Collection("reserves")
}

type Reserve struct {
	ID         string                 `json:"id"          firestore:"-" cloudfirestore:"id"`
	Ref        *firestore.DocumentRef `json:"-"           firestore:"-" cloudfirestore:"ref"`
	UserIDs    []string               `json:"user_ids"    firestore:"user_ids"`
	Message    *Message               `json:"message"     firestore:"message"`
	ReservedAt int64                  `json:"reserved_at" firestore:"reserved_at"`
	UnManaged  bool                   `json:"unmanaged"   firestore:"unmanaged"`
	Status     config.ReserveStatus   `json:"status"      firestore:"status"`
	CreatedAt  int64                  `json:"created_at"  firestore:"created_at"`
	UpdatedAt  int64                  `json:"updated_at"  firestore:"updated_at"`
}

func NewReserve(
	userIDs []string,
	message *Message,
	reservedAt int64,
	unManaged bool,
) *Reserve {
	now := timeutil.NowUnix()
	return &Reserve{
		ID:         stringutil.UniqueID(),
		UserIDs:    userIDs,
		Message:    message,
		ReservedAt: reservedAt,
		UnManaged:  unManaged,
		Status:     config.ReserveStatusReserved,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
