// Code generated by "stringer -type=APIErrorCode -output=types_err.go"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrInvalidRequest-4400]
	_ = x[ErrUnauthorizedAccess-4401]
	_ = x[ErrSystemTemporarilyDisabled-4001]
	_ = x[ErrTenantsOutOfDatacap-4009]
	_ = x[ErrStorageProviderSuspended-4010]
	_ = x[ErrStorageProviderIneligibleToMine-4011]
	_ = x[ErrStorageProviderAboveMaxPending-4012]
	_ = x[ErrUnclaimedPieceCID-4020]
	_ = x[ErrOversizedPiece-4021]
	_ = x[ErrExternalReservationRefused-4029]
	_ = x[ErrReplicaAlreadyPending-4032]
	_ = x[ErrReplicaAlreadyActive-4031]
	_ = x[ErrTooManyReplicas-4033]
}

const (
	_APIErrorCode_name_0 = "ErrSystemTemporarilyDisabled"
	_APIErrorCode_name_1 = "ErrTenantsOutOfDatacapErrStorageProviderSuspendedErrStorageProviderIneligibleToMineErrStorageProviderAboveMaxPending"
	_APIErrorCode_name_2 = "ErrUnclaimedPieceCIDErrOversizedPiece"
	_APIErrorCode_name_3 = "ErrExternalReservationRefused"
	_APIErrorCode_name_4 = "ErrReplicaAlreadyActiveErrReplicaAlreadyPendingErrTooManyReplicas"
	_APIErrorCode_name_5 = "ErrInvalidRequestErrUnauthorizedAccess"
)

var (
	_APIErrorCode_index_1 = [...]uint8{0, 22, 49, 83, 116}
	_APIErrorCode_index_2 = [...]uint8{0, 20, 37}
	_APIErrorCode_index_4 = [...]uint8{0, 23, 47, 65}
	_APIErrorCode_index_5 = [...]uint8{0, 17, 38}
)

func (i APIErrorCode) String() string {
	switch {
	case i == 4001:
		return _APIErrorCode_name_0
	case 4009 <= i && i <= 4012:
		i -= 4009
		return _APIErrorCode_name_1[_APIErrorCode_index_1[i]:_APIErrorCode_index_1[i+1]]
	case 4020 <= i && i <= 4021:
		i -= 4020
		return _APIErrorCode_name_2[_APIErrorCode_index_2[i]:_APIErrorCode_index_2[i+1]]
	case i == 4029:
		return _APIErrorCode_name_3
	case 4031 <= i && i <= 4033:
		i -= 4031
		return _APIErrorCode_name_4[_APIErrorCode_index_4[i]:_APIErrorCode_index_4[i+1]]
	case 4400 <= i && i <= 4401:
		i -= 4400
		return _APIErrorCode_name_5[_APIErrorCode_index_5[i]:_APIErrorCode_index_5[i+1]]
	default:
		return "APIErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
