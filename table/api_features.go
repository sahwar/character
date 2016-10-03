// Copyright © 2015,2016 Phil Pennock.
// All rights reserved, except as granted under license.
// Licensed per file LICENSE.txt

package table

// Alignment indicates our table column alignments.  We only use per-column
// alignment and do not pass through any per-cell alignments.
type Alignment int

// These constants define how a given column of the table should have each
// cell aligned.
const (
	LEFT Alignment = iota
	CENTER
	RIGHT
)
