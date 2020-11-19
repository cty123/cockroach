// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexecagg

import (
	"unsafe"

	"github.com/cockroachdb/apd/v2"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execgen"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
	"github.com/cockroachdb/errors"
)

func newAvgHashAggAlloc(
	allocator *colmem.Allocator, t *types.T, allocSize int64,
) (aggregateFuncAlloc, error) {
	allocBase := aggAllocBase{allocator: allocator, allocSize: allocSize}
	switch t.Family() {
	case types.IntFamily:
		switch t.Width() {
		case 16:
			return &avgInt16HashAggAlloc{aggAllocBase: allocBase}, nil
		case 32:
			return &avgInt32HashAggAlloc{aggAllocBase: allocBase}, nil
		default:
			return &avgInt64HashAggAlloc{aggAllocBase: allocBase}, nil
		}
	case types.DecimalFamily:
		return &avgDecimalHashAggAlloc{aggAllocBase: allocBase}, nil
	case types.FloatFamily:
		return &avgFloat64HashAggAlloc{aggAllocBase: allocBase}, nil
	case types.IntervalFamily:
		return &avgIntervalHashAggAlloc{aggAllocBase: allocBase}, nil
	default:
		return nil, errors.Errorf("unsupported avg agg type %s", t.Name())
	}
}

type avgInt16HashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum apd.Decimal
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper execgen.OverloadHelper
}

var _ AggregateFunc = &avgInt16HashAgg{}

func (a *avgInt16HashAgg) SetOutput(vec coldata.Vec) {
	a.hashAggregateFuncBase.SetOutput(vec)
	a.scratch.vec = vec.Decimal()
}

func (a *avgInt16HashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int16(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.TmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.TmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgInt16HashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {

		a.scratch.vec[outputIdx].SetInt64(a.scratch.curCount)
		if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[outputIdx], &a.scratch.curSum, &a.scratch.vec[outputIdx]); err != nil {
			colexecerror.InternalError(err)
		}
	}
}

type avgInt16HashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgInt16HashAgg
}

var _ aggregateFuncAlloc = &avgInt16HashAggAlloc{}

const sizeOfAvgInt16HashAgg = int64(unsafe.Sizeof(avgInt16HashAgg{}))
const avgInt16HashAggSliceOverhead = int64(unsafe.Sizeof([]avgInt16HashAgg{}))

func (a *avgInt16HashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(avgInt16HashAggSliceOverhead + sizeOfAvgInt16HashAgg*a.allocSize)
		a.aggFuncs = make([]avgInt16HashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type avgInt32HashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum apd.Decimal
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper execgen.OverloadHelper
}

var _ AggregateFunc = &avgInt32HashAgg{}

func (a *avgInt32HashAgg) SetOutput(vec coldata.Vec) {
	a.hashAggregateFuncBase.SetOutput(vec)
	a.scratch.vec = vec.Decimal()
}

func (a *avgInt32HashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int32(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.TmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.TmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgInt32HashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {

		a.scratch.vec[outputIdx].SetInt64(a.scratch.curCount)
		if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[outputIdx], &a.scratch.curSum, &a.scratch.vec[outputIdx]); err != nil {
			colexecerror.InternalError(err)
		}
	}
}

type avgInt32HashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgInt32HashAgg
}

var _ aggregateFuncAlloc = &avgInt32HashAggAlloc{}

const sizeOfAvgInt32HashAgg = int64(unsafe.Sizeof(avgInt32HashAgg{}))
const avgInt32HashAggSliceOverhead = int64(unsafe.Sizeof([]avgInt32HashAgg{}))

func (a *avgInt32HashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(avgInt32HashAggSliceOverhead + sizeOfAvgInt32HashAgg*a.allocSize)
		a.aggFuncs = make([]avgInt32HashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type avgInt64HashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum apd.Decimal
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper execgen.OverloadHelper
}

var _ AggregateFunc = &avgInt64HashAgg{}

func (a *avgInt64HashAgg) SetOutput(vec coldata.Vec) {
	a.hashAggregateFuncBase.SetOutput(vec)
	a.scratch.vec = vec.Decimal()
}

func (a *avgInt64HashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int64(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.TmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.TmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgInt64HashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {

		a.scratch.vec[outputIdx].SetInt64(a.scratch.curCount)
		if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[outputIdx], &a.scratch.curSum, &a.scratch.vec[outputIdx]); err != nil {
			colexecerror.InternalError(err)
		}
	}
}

type avgInt64HashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgInt64HashAgg
}

var _ aggregateFuncAlloc = &avgInt64HashAggAlloc{}

const sizeOfAvgInt64HashAgg = int64(unsafe.Sizeof(avgInt64HashAgg{}))
const avgInt64HashAggSliceOverhead = int64(unsafe.Sizeof([]avgInt64HashAgg{}))

func (a *avgInt64HashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(avgInt64HashAggSliceOverhead + sizeOfAvgInt64HashAgg*a.allocSize)
		a.aggFuncs = make([]avgInt64HashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type avgDecimalHashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum apd.Decimal
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ AggregateFunc = &avgDecimalHashAgg{}

func (a *avgDecimalHashAgg) SetOutput(vec coldata.Vec) {
	a.hashAggregateFuncBase.SetOutput(vec)
	a.scratch.vec = vec.Decimal()
}

func (a *avgDecimalHashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Decimal(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgDecimalHashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {

		a.scratch.vec[outputIdx].SetInt64(a.scratch.curCount)
		if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[outputIdx], &a.scratch.curSum, &a.scratch.vec[outputIdx]); err != nil {
			colexecerror.InternalError(err)
		}
	}
}

type avgDecimalHashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgDecimalHashAgg
}

var _ aggregateFuncAlloc = &avgDecimalHashAggAlloc{}

const sizeOfAvgDecimalHashAgg = int64(unsafe.Sizeof(avgDecimalHashAgg{}))
const avgDecimalHashAggSliceOverhead = int64(unsafe.Sizeof([]avgDecimalHashAgg{}))

func (a *avgDecimalHashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(avgDecimalHashAggSliceOverhead + sizeOfAvgDecimalHashAgg*a.allocSize)
		a.aggFuncs = make([]avgDecimalHashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type avgFloat64HashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum float64
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []float64
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ AggregateFunc = &avgFloat64HashAgg{}

func (a *avgFloat64HashAgg) SetOutput(vec coldata.Vec) {
	a.hashAggregateFuncBase.SetOutput(vec)
	a.scratch.vec = vec.Float64()
}

func (a *avgFloat64HashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Float64(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						a.scratch.curSum = float64(a.scratch.curSum) + float64(col[i])
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{

						a.scratch.curSum = float64(a.scratch.curSum) + float64(col[i])
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgFloat64HashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curSum / float64(a.scratch.curCount)
	}
}

type avgFloat64HashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgFloat64HashAgg
}

var _ aggregateFuncAlloc = &avgFloat64HashAggAlloc{}

const sizeOfAvgFloat64HashAgg = int64(unsafe.Sizeof(avgFloat64HashAgg{}))
const avgFloat64HashAggSliceOverhead = int64(unsafe.Sizeof([]avgFloat64HashAgg{}))

func (a *avgFloat64HashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(avgFloat64HashAggSliceOverhead + sizeOfAvgFloat64HashAgg*a.allocSize)
		a.aggFuncs = make([]avgFloat64HashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type avgIntervalHashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum duration.Duration
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []duration.Duration
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ AggregateFunc = &avgIntervalHashAgg{}

func (a *avgIntervalHashAgg) SetOutput(vec coldata.Vec) {
	a.hashAggregateFuncBase.SetOutput(vec)
	a.scratch.vec = vec.Interval()
}

func (a *avgIntervalHashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Interval(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {
					a.scratch.curSum = a.scratch.curSum.Add(col[i])
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {
					a.scratch.curSum = a.scratch.curSum.Add(col[i])
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgIntervalHashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curSum.Div(int64(a.scratch.curCount))
	}
}

type avgIntervalHashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgIntervalHashAgg
}

var _ aggregateFuncAlloc = &avgIntervalHashAggAlloc{}

const sizeOfAvgIntervalHashAgg = int64(unsafe.Sizeof(avgIntervalHashAgg{}))
const avgIntervalHashAggSliceOverhead = int64(unsafe.Sizeof([]avgIntervalHashAgg{}))

func (a *avgIntervalHashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(avgIntervalHashAggSliceOverhead + sizeOfAvgIntervalHashAgg*a.allocSize)
		a.aggFuncs = make([]avgIntervalHashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}
