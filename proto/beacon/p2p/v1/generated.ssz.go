// Code generated by fastssz. DO NOT EDIT.
package ethereum_beacon_p2p_v1

import (
	ssz "github.com/ferranbt/fastssz"
	v1alpha1 "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

// MarshalSSZ ssz marshals the BeaconState object
func (b *BeaconState) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconState object to a target array
func (b *BeaconState) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(2687377)

	// Field (0) 'GenesisTime'
	dst = ssz.MarshalUint64(dst, b.GenesisTime)

	// Field (1) 'GenesisValidatorsRoot'
	if len(b.GenesisValidatorsRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, b.GenesisValidatorsRoot...)

	// Field (2) 'Slot'
	dst = ssz.MarshalUint64(dst, b.Slot)

	// Field (3) 'Fork'
	if b.Fork == nil {
		b.Fork = new(Fork)
	}
	if dst, err = b.Fork.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(v1alpha1.BeaconBlockHeader)
	}
	if dst, err = b.LatestBlockHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'BlockRoots'
	if len(b.BlockRoots) != 8192 {
		err = ssz.ErrVectorLength
		return
	}
	for ii := 0; ii < 8192; ii++ {
		if len(b.BlockRoots[ii]) != 32 {
			err = ssz.ErrBytesLength
			return
		}
		dst = append(dst, b.BlockRoots[ii]...)
	}

	// Field (6) 'StateRoots'
	if len(b.StateRoots) != 8192 {
		err = ssz.ErrVectorLength
		return
	}
	for ii := 0; ii < 8192; ii++ {
		if len(b.StateRoots[ii]) != 32 {
			err = ssz.ErrBytesLength
			return
		}
		dst = append(dst, b.StateRoots[ii]...)
	}

	// Offset (7) 'HistoricalRoots'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.HistoricalRoots) * 32

	// Field (8) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(v1alpha1.Eth1Data)
	}
	if dst, err = b.Eth1Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (9) 'Eth1DataVotes'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Eth1DataVotes) * 72

	// Field (10) 'Eth1DepositIndex'
	dst = ssz.MarshalUint64(dst, b.Eth1DepositIndex)

	// Offset (11) 'Validators'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Validators) * 121

	// Offset (12) 'Balances'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Balances) * 8

	// Field (13) 'RandaoMixes'
	if len(b.RandaoMixes) != 65536 {
		err = ssz.ErrVectorLength
		return
	}
	for ii := 0; ii < 65536; ii++ {
		if len(b.RandaoMixes[ii]) != 32 {
			err = ssz.ErrBytesLength
			return
		}
		dst = append(dst, b.RandaoMixes[ii]...)
	}

	// Field (14) 'Slashings'
	if len(b.Slashings) != 8192 {
		err = ssz.ErrVectorLength
		return
	}
	for ii := 0; ii < 8192; ii++ {
		dst = ssz.MarshalUint64(dst, b.Slashings[ii])
	}

	// Offset (15) 'PreviousEpochAttestations'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
		offset += 4
		offset += b.PreviousEpochAttestations[ii].SizeSSZ()
	}

	// Offset (16) 'CurrentEpochAttestations'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.CurrentEpochAttestations); ii++ {
		offset += 4
		offset += b.CurrentEpochAttestations[ii].SizeSSZ()
	}

	// Field (17) 'JustificationBits'
	if len(b.JustificationBits) != 1 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, b.JustificationBits...)

	// Field (18) 'PreviousJustifiedCheckpoint'
	if b.PreviousJustifiedCheckpoint == nil {
		b.PreviousJustifiedCheckpoint = new(v1alpha1.Checkpoint)
	}
	if dst, err = b.PreviousJustifiedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (19) 'CurrentJustifiedCheckpoint'
	if b.CurrentJustifiedCheckpoint == nil {
		b.CurrentJustifiedCheckpoint = new(v1alpha1.Checkpoint)
	}
	if dst, err = b.CurrentJustifiedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (20) 'FinalizedCheckpoint'
	if b.FinalizedCheckpoint == nil {
		b.FinalizedCheckpoint = new(v1alpha1.Checkpoint)
	}
	if dst, err = b.FinalizedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (7) 'HistoricalRoots'
	if len(b.HistoricalRoots) > 16777216 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.HistoricalRoots); ii++ {
		if len(b.HistoricalRoots[ii]) != 32 {
			err = ssz.ErrBytesLength
			return
		}
		dst = append(dst, b.HistoricalRoots[ii]...)
	}

	// Field (9) 'Eth1DataVotes'
	if len(b.Eth1DataVotes) > 2048 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.Eth1DataVotes); ii++ {
		if dst, err = b.Eth1DataVotes[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (11) 'Validators'
	if len(b.Validators) > 1099511627776 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.Validators); ii++ {
		if dst, err = b.Validators[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (12) 'Balances'
	if len(b.Balances) > 1099511627776 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.Balances); ii++ {
		dst = ssz.MarshalUint64(dst, b.Balances[ii])
	}

	// Field (15) 'PreviousEpochAttestations'
	if len(b.PreviousEpochAttestations) > 4096 {
		err = ssz.ErrListTooBig
		return
	}
	{
		offset = 4 * len(b.PreviousEpochAttestations)
		for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.PreviousEpochAttestations[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
		if dst, err = b.PreviousEpochAttestations[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (16) 'CurrentEpochAttestations'
	if len(b.CurrentEpochAttestations) > 4096 {
		err = ssz.ErrListTooBig
		return
	}
	{
		offset = 4 * len(b.CurrentEpochAttestations)
		for ii := 0; ii < len(b.CurrentEpochAttestations); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.CurrentEpochAttestations[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.CurrentEpochAttestations); ii++ {
		if dst, err = b.CurrentEpochAttestations[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconState object
func (b *BeaconState) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 2687377 {
		return ssz.ErrSize
	}

	tail := buf
	var o7, o9, o11, o12, o15, o16 uint64

	// Field (0) 'GenesisTime'
	b.GenesisTime = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'GenesisValidatorsRoot'
	if cap(b.GenesisValidatorsRoot) == 0 {
		b.GenesisValidatorsRoot = make([]byte, 0, len(buf[8:40]))
	}
	b.GenesisValidatorsRoot = append(b.GenesisValidatorsRoot, buf[8:40]...)

	// Field (2) 'Slot'
	b.Slot = ssz.UnmarshallUint64(buf[40:48])

	// Field (3) 'Fork'
	if b.Fork == nil {
		b.Fork = new(Fork)
	}
	if err = b.Fork.UnmarshalSSZ(buf[48:64]); err != nil {
		return err
	}

	// Field (4) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(v1alpha1.BeaconBlockHeader)
	}
	if err = b.LatestBlockHeader.UnmarshalSSZ(buf[64:176]); err != nil {
		return err
	}

	// Field (5) 'BlockRoots'
	b.BlockRoots = make([][]byte, 8192)
	for ii := 0; ii < 8192; ii++ {
		if cap(b.BlockRoots[ii]) == 0 {
			b.BlockRoots[ii] = make([]byte, 0, len(buf[176:262320][ii*32:(ii+1)*32]))
		}
		b.BlockRoots[ii] = append(b.BlockRoots[ii], buf[176:262320][ii*32:(ii+1)*32]...)
	}

	// Field (6) 'StateRoots'
	b.StateRoots = make([][]byte, 8192)
	for ii := 0; ii < 8192; ii++ {
		if cap(b.StateRoots[ii]) == 0 {
			b.StateRoots[ii] = make([]byte, 0, len(buf[262320:524464][ii*32:(ii+1)*32]))
		}
		b.StateRoots[ii] = append(b.StateRoots[ii], buf[262320:524464][ii*32:(ii+1)*32]...)
	}

	// Offset (7) 'HistoricalRoots'
	if o7 = ssz.ReadOffset(buf[524464:524468]); o7 > size {
		return ssz.ErrOffset
	}

	// Field (8) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(v1alpha1.Eth1Data)
	}
	if err = b.Eth1Data.UnmarshalSSZ(buf[524468:524540]); err != nil {
		return err
	}

	// Offset (9) 'Eth1DataVotes'
	if o9 = ssz.ReadOffset(buf[524540:524544]); o9 > size || o7 > o9 {
		return ssz.ErrOffset
	}

	// Field (10) 'Eth1DepositIndex'
	b.Eth1DepositIndex = ssz.UnmarshallUint64(buf[524544:524552])

	// Offset (11) 'Validators'
	if o11 = ssz.ReadOffset(buf[524552:524556]); o11 > size || o9 > o11 {
		return ssz.ErrOffset
	}

	// Offset (12) 'Balances'
	if o12 = ssz.ReadOffset(buf[524556:524560]); o12 > size || o11 > o12 {
		return ssz.ErrOffset
	}

	// Field (13) 'RandaoMixes'
	b.RandaoMixes = make([][]byte, 65536)
	for ii := 0; ii < 65536; ii++ {
		if cap(b.RandaoMixes[ii]) == 0 {
			b.RandaoMixes[ii] = make([]byte, 0, len(buf[524560:2621712][ii*32:(ii+1)*32]))
		}
		b.RandaoMixes[ii] = append(b.RandaoMixes[ii], buf[524560:2621712][ii*32:(ii+1)*32]...)
	}

	// Field (14) 'Slashings'
	b.Slashings = ssz.ExtendUint64(b.Slashings, 8192)
	for ii := 0; ii < 8192; ii++ {
		b.Slashings[ii] = ssz.UnmarshallUint64(buf[2621712:2687248][ii*8 : (ii+1)*8])
	}

	// Offset (15) 'PreviousEpochAttestations'
	if o15 = ssz.ReadOffset(buf[2687248:2687252]); o15 > size || o12 > o15 {
		return ssz.ErrOffset
	}

	// Offset (16) 'CurrentEpochAttestations'
	if o16 = ssz.ReadOffset(buf[2687252:2687256]); o16 > size || o15 > o16 {
		return ssz.ErrOffset
	}

	// Field (17) 'JustificationBits'
	if cap(b.JustificationBits) == 0 {
		b.JustificationBits = make([]byte, 0, len(buf[2687256:2687257]))
	}
	b.JustificationBits = append(b.JustificationBits, buf[2687256:2687257]...)

	// Field (18) 'PreviousJustifiedCheckpoint'
	if b.PreviousJustifiedCheckpoint == nil {
		b.PreviousJustifiedCheckpoint = new(v1alpha1.Checkpoint)
	}
	if err = b.PreviousJustifiedCheckpoint.UnmarshalSSZ(buf[2687257:2687297]); err != nil {
		return err
	}

	// Field (19) 'CurrentJustifiedCheckpoint'
	if b.CurrentJustifiedCheckpoint == nil {
		b.CurrentJustifiedCheckpoint = new(v1alpha1.Checkpoint)
	}
	if err = b.CurrentJustifiedCheckpoint.UnmarshalSSZ(buf[2687297:2687337]); err != nil {
		return err
	}

	// Field (20) 'FinalizedCheckpoint'
	if b.FinalizedCheckpoint == nil {
		b.FinalizedCheckpoint = new(v1alpha1.Checkpoint)
	}
	if err = b.FinalizedCheckpoint.UnmarshalSSZ(buf[2687337:2687377]); err != nil {
		return err
	}

	// Field (7) 'HistoricalRoots'
	{
		buf = tail[o7:o9]
		num, err := ssz.DivideInt2(len(buf), 32, 16777216)
		if err != nil {
			return err
		}
		b.HistoricalRoots = make([][]byte, num)
		for ii := 0; ii < num; ii++ {
			if cap(b.HistoricalRoots[ii]) == 0 {
				b.HistoricalRoots[ii] = make([]byte, 0, len(buf[ii*32:(ii+1)*32]))
			}
			b.HistoricalRoots[ii] = append(b.HistoricalRoots[ii], buf[ii*32:(ii+1)*32]...)
		}
	}

	// Field (9) 'Eth1DataVotes'
	{
		buf = tail[o9:o11]
		num, err := ssz.DivideInt2(len(buf), 72, 2048)
		if err != nil {
			return err
		}
		b.Eth1DataVotes = make([]*v1alpha1.Eth1Data, num)
		for ii := 0; ii < num; ii++ {
			if b.Eth1DataVotes[ii] == nil {
				b.Eth1DataVotes[ii] = new(v1alpha1.Eth1Data)
			}
			if err = b.Eth1DataVotes[ii].UnmarshalSSZ(buf[ii*72 : (ii+1)*72]); err != nil {
				return err
			}
		}
	}

	// Field (11) 'Validators'
	{
		buf = tail[o11:o12]
		num, err := ssz.DivideInt2(len(buf), 121, 1099511627776)
		if err != nil {
			return err
		}
		b.Validators = make([]*v1alpha1.Validator, num)
		for ii := 0; ii < num; ii++ {
			if b.Validators[ii] == nil {
				b.Validators[ii] = new(v1alpha1.Validator)
			}
			if err = b.Validators[ii].UnmarshalSSZ(buf[ii*121 : (ii+1)*121]); err != nil {
				return err
			}
		}
	}

	// Field (12) 'Balances'
	{
		buf = tail[o12:o15]
		num, err := ssz.DivideInt2(len(buf), 8, 1099511627776)
		if err != nil {
			return err
		}
		b.Balances = ssz.ExtendUint64(b.Balances, num)
		for ii := 0; ii < num; ii++ {
			b.Balances[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}

	// Field (15) 'PreviousEpochAttestations'
	{
		buf = tail[o15:o16]
		num, err := ssz.DecodeDynamicLength(buf, 4096)
		if err != nil {
			return err
		}
		b.PreviousEpochAttestations = make([]*PendingAttestation, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.PreviousEpochAttestations[indx] == nil {
				b.PreviousEpochAttestations[indx] = new(PendingAttestation)
			}
			if err = b.PreviousEpochAttestations[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (16) 'CurrentEpochAttestations'
	{
		buf = tail[o16:]
		num, err := ssz.DecodeDynamicLength(buf, 4096)
		if err != nil {
			return err
		}
		b.CurrentEpochAttestations = make([]*PendingAttestation, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.CurrentEpochAttestations[indx] == nil {
				b.CurrentEpochAttestations[indx] = new(PendingAttestation)
			}
			if err = b.CurrentEpochAttestations[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconState object
func (b *BeaconState) SizeSSZ() (size int) {
	size = 2687377

	// Field (7) 'HistoricalRoots'
	size += len(b.HistoricalRoots) * 32

	// Field (9) 'Eth1DataVotes'
	size += len(b.Eth1DataVotes) * 72

	// Field (11) 'Validators'
	size += len(b.Validators) * 121

	// Field (12) 'Balances'
	size += len(b.Balances) * 8

	// Field (15) 'PreviousEpochAttestations'
	for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
		size += 4
		size += b.PreviousEpochAttestations[ii].SizeSSZ()
	}

	// Field (16) 'CurrentEpochAttestations'
	for ii := 0; ii < len(b.CurrentEpochAttestations); ii++ {
		size += 4
		size += b.CurrentEpochAttestations[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the BeaconState object
func (b *BeaconState) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconState object with a hasher
func (b *BeaconState) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'GenesisTime'
	hh.PutUint64(b.GenesisTime)

	// Field (1) 'GenesisValidatorsRoot'
	if len(b.GenesisValidatorsRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(b.GenesisValidatorsRoot)

	// Field (2) 'Slot'
	hh.PutUint64(b.Slot)

	// Field (3) 'Fork'
	if err = b.Fork.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'LatestBlockHeader'
	if err = b.LatestBlockHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'BlockRoots'
	{
		if len(b.BlockRoots) != 8192 {
			err = ssz.ErrVectorLength
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlockRoots {
			if len(i) != 32 {
				err = ssz.ErrBytesLength
				return
			}
			hh.Append(i)
		}
		hh.Merkleize(subIndx)
	}

	// Field (6) 'StateRoots'
	{
		if len(b.StateRoots) != 8192 {
			err = ssz.ErrVectorLength
			return
		}
		subIndx := hh.Index()
		for _, i := range b.StateRoots {
			if len(i) != 32 {
				err = ssz.ErrBytesLength
				return
			}
			hh.Append(i)
		}
		hh.Merkleize(subIndx)
	}

	// Field (7) 'HistoricalRoots'
	{
		if len(b.HistoricalRoots) > 16777216 {
			err = ssz.ErrListTooBig
			return
		}
		subIndx := hh.Index()
		for _, i := range b.HistoricalRoots {
			if len(i) != 32 {
				err = ssz.ErrBytesLength
				return
			}
			hh.Append(i)
		}
		numItems := uint64(len(b.HistoricalRoots))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(16777216, numItems, 32))
	}

	// Field (8) 'Eth1Data'
	if err = b.Eth1Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (9) 'Eth1DataVotes'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Eth1DataVotes))
		if num > 2048 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.Eth1DataVotes[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 2048)
	}

	// Field (10) 'Eth1DepositIndex'
	hh.PutUint64(b.Eth1DepositIndex)

	// Field (11) 'Validators'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Validators))
		if num > 1099511627776 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.Validators[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1099511627776)
	}

	// Field (12) 'Balances'
	{
		if len(b.Balances) > 1099511627776 {
			err = ssz.ErrListTooBig
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Balances {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(b.Balances))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(1099511627776, numItems, 8))
	}

	// Field (13) 'RandaoMixes'
	{
		if len(b.RandaoMixes) != 65536 {
			err = ssz.ErrVectorLength
			return
		}
		subIndx := hh.Index()
		for _, i := range b.RandaoMixes {
			if len(i) != 32 {
				err = ssz.ErrBytesLength
				return
			}
			hh.Append(i)
		}
		hh.Merkleize(subIndx)
	}

	// Field (14) 'Slashings'
	{
		if len(b.Slashings) != 8192 {
			err = ssz.ErrVectorLength
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Slashings {
			hh.AppendUint64(i)
		}
		hh.Merkleize(subIndx)
	}

	// Field (15) 'PreviousEpochAttestations'
	{
		subIndx := hh.Index()
		num := uint64(len(b.PreviousEpochAttestations))
		if num > 4096 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.PreviousEpochAttestations[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 4096)
	}

	// Field (16) 'CurrentEpochAttestations'
	{
		subIndx := hh.Index()
		num := uint64(len(b.CurrentEpochAttestations))
		if num > 4096 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.CurrentEpochAttestations[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 4096)
	}

	// Field (17) 'JustificationBits'
	if len(b.JustificationBits) != 1 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(b.JustificationBits)

	// Field (18) 'PreviousJustifiedCheckpoint'
	if err = b.PreviousJustifiedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (19) 'CurrentJustifiedCheckpoint'
	if err = b.CurrentJustifiedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (20) 'FinalizedCheckpoint'
	if err = b.FinalizedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the Fork object
func (f *Fork) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(f)
}

// MarshalSSZTo ssz marshals the Fork object to a target array
func (f *Fork) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'PreviousVersion'
	if len(f.PreviousVersion) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, f.PreviousVersion...)

	// Field (1) 'CurrentVersion'
	if len(f.CurrentVersion) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, f.CurrentVersion...)

	// Field (2) 'Epoch'
	dst = ssz.MarshalUint64(dst, f.Epoch)

	return
}

// UnmarshalSSZ ssz unmarshals the Fork object
func (f *Fork) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'PreviousVersion'
	if cap(f.PreviousVersion) == 0 {
		f.PreviousVersion = make([]byte, 0, len(buf[0:4]))
	}
	f.PreviousVersion = append(f.PreviousVersion, buf[0:4]...)

	// Field (1) 'CurrentVersion'
	if cap(f.CurrentVersion) == 0 {
		f.CurrentVersion = make([]byte, 0, len(buf[4:8]))
	}
	f.CurrentVersion = append(f.CurrentVersion, buf[4:8]...)

	// Field (2) 'Epoch'
	f.Epoch = ssz.UnmarshallUint64(buf[8:16])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Fork object
func (f *Fork) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the Fork object
func (f *Fork) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(f)
}

// HashTreeRootWith ssz hashes the Fork object with a hasher
func (f *Fork) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'PreviousVersion'
	if len(f.PreviousVersion) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(f.PreviousVersion)

	// Field (1) 'CurrentVersion'
	if len(f.CurrentVersion) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(f.CurrentVersion)

	// Field (2) 'Epoch'
	hh.PutUint64(f.Epoch)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the PendingAttestation object
func (p *PendingAttestation) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the PendingAttestation object to a target array
func (p *PendingAttestation) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(148)

	// Offset (0) 'AggregationBits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(p.AggregationBits)

	// Field (1) 'Data'
	if p.Data == nil {
		p.Data = new(v1alpha1.AttestationData)
	}
	if dst, err = p.Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'InclusionDelay'
	dst = ssz.MarshalUint64(dst, p.InclusionDelay)

	// Field (3) 'ProposerIndex'
	dst = ssz.MarshalUint64(dst, p.ProposerIndex)

	// Field (0) 'AggregationBits'
	if len(p.AggregationBits) > 2048 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, p.AggregationBits...)

	return
}

// UnmarshalSSZ ssz unmarshals the PendingAttestation object
func (p *PendingAttestation) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 148 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'AggregationBits'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Field (1) 'Data'
	if p.Data == nil {
		p.Data = new(v1alpha1.AttestationData)
	}
	if err = p.Data.UnmarshalSSZ(buf[4:132]); err != nil {
		return err
	}

	// Field (2) 'InclusionDelay'
	p.InclusionDelay = ssz.UnmarshallUint64(buf[132:140])

	// Field (3) 'ProposerIndex'
	p.ProposerIndex = ssz.UnmarshallUint64(buf[140:148])

	// Field (0) 'AggregationBits'
	{
		buf = tail[o0:]
		if err = ssz.ValidateBitlist(buf, 2048); err != nil {
			return err
		}
		if cap(p.AggregationBits) == 0 {
			p.AggregationBits = make([]byte, 0, len(buf))
		}
		p.AggregationBits = append(p.AggregationBits, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the PendingAttestation object
func (p *PendingAttestation) SizeSSZ() (size int) {
	size = 148

	// Field (0) 'AggregationBits'
	size += len(p.AggregationBits)

	return
}

// HashTreeRoot ssz hashes the PendingAttestation object
func (p *PendingAttestation) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the PendingAttestation object with a hasher
func (p *PendingAttestation) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'AggregationBits'
	if len(p.AggregationBits) == 0 {
		err = ssz.ErrEmptyBitlist
		return
	}
	hh.PutBitlist(p.AggregationBits, 2048)

	// Field (1) 'Data'
	if err = p.Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'InclusionDelay'
	hh.PutUint64(p.InclusionDelay)

	// Field (3) 'ProposerIndex'
	hh.PutUint64(p.ProposerIndex)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the HistoricalBatch object
func (h *HistoricalBatch) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the HistoricalBatch object to a target array
func (h *HistoricalBatch) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'BlockRoots'
	if len(h.BlockRoots) != 8192 {
		err = ssz.ErrVectorLength
		return
	}
	for ii := 0; ii < 8192; ii++ {
		if len(h.BlockRoots[ii]) != 32 {
			err = ssz.ErrBytesLength
			return
		}
		dst = append(dst, h.BlockRoots[ii]...)
	}

	// Field (1) 'StateRoots'
	if len(h.StateRoots) != 8192 {
		err = ssz.ErrVectorLength
		return
	}
	for ii := 0; ii < 8192; ii++ {
		if len(h.StateRoots[ii]) != 32 {
			err = ssz.ErrBytesLength
			return
		}
		dst = append(dst, h.StateRoots[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the HistoricalBatch object
func (h *HistoricalBatch) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 524288 {
		return ssz.ErrSize
	}

	// Field (0) 'BlockRoots'
	h.BlockRoots = make([][]byte, 8192)
	for ii := 0; ii < 8192; ii++ {
		if cap(h.BlockRoots[ii]) == 0 {
			h.BlockRoots[ii] = make([]byte, 0, len(buf[0:262144][ii*32:(ii+1)*32]))
		}
		h.BlockRoots[ii] = append(h.BlockRoots[ii], buf[0:262144][ii*32:(ii+1)*32]...)
	}

	// Field (1) 'StateRoots'
	h.StateRoots = make([][]byte, 8192)
	for ii := 0; ii < 8192; ii++ {
		if cap(h.StateRoots[ii]) == 0 {
			h.StateRoots[ii] = make([]byte, 0, len(buf[262144:524288][ii*32:(ii+1)*32]))
		}
		h.StateRoots[ii] = append(h.StateRoots[ii], buf[262144:524288][ii*32:(ii+1)*32]...)
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the HistoricalBatch object
func (h *HistoricalBatch) SizeSSZ() (size int) {
	size = 524288
	return
}

// HashTreeRoot ssz hashes the HistoricalBatch object
func (h *HistoricalBatch) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the HistoricalBatch object with a hasher
func (h *HistoricalBatch) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'BlockRoots'
	{
		if len(h.BlockRoots) != 8192 {
			err = ssz.ErrVectorLength
			return
		}
		subIndx := hh.Index()
		for _, i := range h.BlockRoots {
			if len(i) != 32 {
				err = ssz.ErrBytesLength
				return
			}
			hh.Append(i)
		}
		hh.Merkleize(subIndx)
	}

	// Field (1) 'StateRoots'
	{
		if len(h.StateRoots) != 8192 {
			err = ssz.ErrVectorLength
			return
		}
		subIndx := hh.Index()
		for _, i := range h.StateRoots {
			if len(i) != 32 {
				err = ssz.ErrBytesLength
				return
			}
			hh.Append(i)
		}
		hh.Merkleize(subIndx)
	}

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the SigningData object
func (s *SigningData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SigningData object to a target array
func (s *SigningData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ObjectRoot'
	if len(s.ObjectRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, s.ObjectRoot...)

	// Field (1) 'Domain'
	if len(s.Domain) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, s.Domain...)

	return
}

// UnmarshalSSZ ssz unmarshals the SigningData object
func (s *SigningData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'ObjectRoot'
	if cap(s.ObjectRoot) == 0 {
		s.ObjectRoot = make([]byte, 0, len(buf[0:32]))
	}
	s.ObjectRoot = append(s.ObjectRoot, buf[0:32]...)

	// Field (1) 'Domain'
	if cap(s.Domain) == 0 {
		s.Domain = make([]byte, 0, len(buf[32:64]))
	}
	s.Domain = append(s.Domain, buf[32:64]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SigningData object
func (s *SigningData) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the SigningData object
func (s *SigningData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SigningData object with a hasher
func (s *SigningData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ObjectRoot'
	if len(s.ObjectRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(s.ObjectRoot)

	// Field (1) 'Domain'
	if len(s.Domain) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(s.Domain)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the ForkData object
func (f *ForkData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(f)
}

// MarshalSSZTo ssz marshals the ForkData object to a target array
func (f *ForkData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'CurrentVersion'
	if len(f.CurrentVersion) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, f.CurrentVersion...)

	// Field (1) 'GenesisValidatorsRoot'
	if len(f.GenesisValidatorsRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, f.GenesisValidatorsRoot...)

	return
}

// UnmarshalSSZ ssz unmarshals the ForkData object
func (f *ForkData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 36 {
		return ssz.ErrSize
	}

	// Field (0) 'CurrentVersion'
	if cap(f.CurrentVersion) == 0 {
		f.CurrentVersion = make([]byte, 0, len(buf[0:4]))
	}
	f.CurrentVersion = append(f.CurrentVersion, buf[0:4]...)

	// Field (1) 'GenesisValidatorsRoot'
	if cap(f.GenesisValidatorsRoot) == 0 {
		f.GenesisValidatorsRoot = make([]byte, 0, len(buf[4:36]))
	}
	f.GenesisValidatorsRoot = append(f.GenesisValidatorsRoot, buf[4:36]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ForkData object
func (f *ForkData) SizeSSZ() (size int) {
	size = 36
	return
}

// HashTreeRoot ssz hashes the ForkData object
func (f *ForkData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(f)
}

// HashTreeRootWith ssz hashes the ForkData object with a hasher
func (f *ForkData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentVersion'
	if len(f.CurrentVersion) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(f.CurrentVersion)

	// Field (1) 'GenesisValidatorsRoot'
	if len(f.GenesisValidatorsRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(f.GenesisValidatorsRoot)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the Status object
func (s *Status) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the Status object to a target array
func (s *Status) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ForkDigest'
	if len(s.ForkDigest) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, s.ForkDigest...)

	// Field (1) 'FinalizedRoot'
	if len(s.FinalizedRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, s.FinalizedRoot...)

	// Field (2) 'FinalizedEpoch'
	dst = ssz.MarshalUint64(dst, s.FinalizedEpoch)

	// Field (3) 'HeadRoot'
	if len(s.HeadRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, s.HeadRoot...)

	// Field (4) 'HeadSlot'
	dst = ssz.MarshalUint64(dst, s.HeadSlot)

	return
}

// UnmarshalSSZ ssz unmarshals the Status object
func (s *Status) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 84 {
		return ssz.ErrSize
	}

	// Field (0) 'ForkDigest'
	if cap(s.ForkDigest) == 0 {
		s.ForkDigest = make([]byte, 0, len(buf[0:4]))
	}
	s.ForkDigest = append(s.ForkDigest, buf[0:4]...)

	// Field (1) 'FinalizedRoot'
	if cap(s.FinalizedRoot) == 0 {
		s.FinalizedRoot = make([]byte, 0, len(buf[4:36]))
	}
	s.FinalizedRoot = append(s.FinalizedRoot, buf[4:36]...)

	// Field (2) 'FinalizedEpoch'
	s.FinalizedEpoch = ssz.UnmarshallUint64(buf[36:44])

	// Field (3) 'HeadRoot'
	if cap(s.HeadRoot) == 0 {
		s.HeadRoot = make([]byte, 0, len(buf[44:76]))
	}
	s.HeadRoot = append(s.HeadRoot, buf[44:76]...)

	// Field (4) 'HeadSlot'
	s.HeadSlot = ssz.UnmarshallUint64(buf[76:84])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Status object
func (s *Status) SizeSSZ() (size int) {
	size = 84
	return
}

// HashTreeRoot ssz hashes the Status object
func (s *Status) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the Status object with a hasher
func (s *Status) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ForkDigest'
	if len(s.ForkDigest) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(s.ForkDigest)

	// Field (1) 'FinalizedRoot'
	if len(s.FinalizedRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(s.FinalizedRoot)

	// Field (2) 'FinalizedEpoch'
	hh.PutUint64(s.FinalizedEpoch)

	// Field (3) 'HeadRoot'
	if len(s.HeadRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(s.HeadRoot)

	// Field (4) 'HeadSlot'
	hh.PutUint64(s.HeadSlot)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the BeaconBlocksByRangeRequest object
func (b *BeaconBlocksByRangeRequest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlocksByRangeRequest object to a target array
func (b *BeaconBlocksByRangeRequest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'StartSlot'
	dst = ssz.MarshalUint64(dst, b.StartSlot)

	// Field (1) 'Count'
	dst = ssz.MarshalUint64(dst, b.Count)

	// Field (2) 'Step'
	dst = ssz.MarshalUint64(dst, b.Step)

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlocksByRangeRequest object
func (b *BeaconBlocksByRangeRequest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 24 {
		return ssz.ErrSize
	}

	// Field (0) 'StartSlot'
	b.StartSlot = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Count'
	b.Count = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'Step'
	b.Step = ssz.UnmarshallUint64(buf[16:24])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlocksByRangeRequest object
func (b *BeaconBlocksByRangeRequest) SizeSSZ() (size int) {
	size = 24
	return
}

// HashTreeRoot ssz hashes the BeaconBlocksByRangeRequest object
func (b *BeaconBlocksByRangeRequest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlocksByRangeRequest object with a hasher
func (b *BeaconBlocksByRangeRequest) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'StartSlot'
	hh.PutUint64(b.StartSlot)

	// Field (1) 'Count'
	hh.PutUint64(b.Count)

	// Field (2) 'Step'
	hh.PutUint64(b.Step)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the ErrorResponse object
func (e *ErrorResponse) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ErrorResponse object to a target array
func (e *ErrorResponse) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Message'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.Message)

	// Field (0) 'Message'
	if len(e.Message) > 256 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, e.Message...)

	return
}

// UnmarshalSSZ ssz unmarshals the ErrorResponse object
func (e *ErrorResponse) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Message'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Field (0) 'Message'
	{
		buf = tail[o0:]
		if len(buf) > 256 {
			return ssz.ErrBytesLength
		}
		if cap(e.Message) == 0 {
			e.Message = make([]byte, 0, len(buf))
		}
		e.Message = append(e.Message, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ErrorResponse object
func (e *ErrorResponse) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Message'
	size += len(e.Message)

	return
}

// HashTreeRoot ssz hashes the ErrorResponse object
func (e *ErrorResponse) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ErrorResponse object with a hasher
func (e *ErrorResponse) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if len(e.Message) > 256 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(e.Message)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the ENRForkID object
func (e *ENRForkID) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ENRForkID object to a target array
func (e *ENRForkID) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'CurrentForkDigest'
	if len(e.CurrentForkDigest) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, e.CurrentForkDigest...)

	// Field (1) 'NextForkVersion'
	if len(e.NextForkVersion) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, e.NextForkVersion...)

	// Field (2) 'NextForkEpoch'
	dst = ssz.MarshalUint64(dst, e.NextForkEpoch)

	return
}

// UnmarshalSSZ ssz unmarshals the ENRForkID object
func (e *ENRForkID) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'CurrentForkDigest'
	if cap(e.CurrentForkDigest) == 0 {
		e.CurrentForkDigest = make([]byte, 0, len(buf[0:4]))
	}
	e.CurrentForkDigest = append(e.CurrentForkDigest, buf[0:4]...)

	// Field (1) 'NextForkVersion'
	if cap(e.NextForkVersion) == 0 {
		e.NextForkVersion = make([]byte, 0, len(buf[4:8]))
	}
	e.NextForkVersion = append(e.NextForkVersion, buf[4:8]...)

	// Field (2) 'NextForkEpoch'
	e.NextForkEpoch = ssz.UnmarshallUint64(buf[8:16])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ENRForkID object
func (e *ENRForkID) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the ENRForkID object
func (e *ENRForkID) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ENRForkID object with a hasher
func (e *ENRForkID) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentForkDigest'
	if len(e.CurrentForkDigest) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(e.CurrentForkDigest)

	// Field (1) 'NextForkVersion'
	if len(e.NextForkVersion) != 4 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(e.NextForkVersion)

	// Field (2) 'NextForkEpoch'
	hh.PutUint64(e.NextForkEpoch)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the MetaData object
func (m *MetaData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MetaData object to a target array
func (m *MetaData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SeqNumber'
	dst = ssz.MarshalUint64(dst, m.SeqNumber)

	// Field (1) 'Attnets'
	if len(m.Attnets) != 8 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, m.Attnets...)

	return
}

// UnmarshalSSZ ssz unmarshals the MetaData object
func (m *MetaData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'SeqNumber'
	m.SeqNumber = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Attnets'
	if cap(m.Attnets) == 0 {
		m.Attnets = make([]byte, 0, len(buf[8:16]))
	}
	m.Attnets = append(m.Attnets, buf[8:16]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MetaData object
func (m *MetaData) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the MetaData object
func (m *MetaData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MetaData object with a hasher
func (m *MetaData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'SeqNumber'
	hh.PutUint64(m.SeqNumber)

	// Field (1) 'Attnets'
	if len(m.Attnets) != 8 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(m.Attnets)

	hh.Merkleize(indx)
	return
}
