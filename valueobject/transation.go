package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transtation is a value obejetc has no identifies and imutable.

type Transtation struct{
	amount int
	from uuid.UUID
	to uuid.UUID
	createdAt time.Time
}