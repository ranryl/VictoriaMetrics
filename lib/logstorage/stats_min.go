package logstorage

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"unsafe"
)

type statsMin struct {
	fields       []string
	containsStar bool
}

func (sm *statsMin) String() string {
	return "min(" + fieldNamesString(sm.fields) + ")"
}

func (sm *statsMin) neededFields() []string {
	return sm.fields
}

func (sm *statsMin) newStatsProcessor() (statsProcessor, int) {
	smp := &statsMinProcessor{
		sm:  sm,
		min: nan,
	}
	return smp, int(unsafe.Sizeof(*smp))
}

type statsMinProcessor struct {
	sm *statsMin

	min float64
}

func (smp *statsMinProcessor) updateStatsForAllRows(br *blockResult) int {
	if smp.sm.containsStar {
		// Find the minimum value across all the columns
		for _, c := range br.getColumns() {
			f := c.getMinValue(br)
			if f < smp.min || math.IsNaN(smp.min) {
				smp.min = f
			}
		}
		return 0
	}

	// Find the minimum value across the requested columns
	for _, field := range smp.sm.fields {
		c := br.getColumnByName(field)
		f := c.getMinValue(br)
		if f < smp.min || math.IsNaN(smp.min) {
			smp.min = f
		}
	}
	return 0
}

func (smp *statsMinProcessor) updateStatsForRow(br *blockResult, rowIdx int) int {
	if smp.sm.containsStar {
		// Find the minimum value across all the fields for the given row
		for _, c := range br.getColumns() {
			f := c.getFloatValueAtRow(rowIdx)
			if f < smp.min || math.IsNaN(smp.min) {
				smp.min = f
			}
		}
		return 0
	}

	// Find the minimum value across the requested fields for the given row
	for _, field := range smp.sm.fields {
		c := br.getColumnByName(field)
		f := c.getFloatValueAtRow(rowIdx)
		if f < smp.min || math.IsNaN(smp.min) {
			smp.min = f
		}
	}
	return 0
}

func (smp *statsMinProcessor) mergeState(sfp statsProcessor) {
	src := sfp.(*statsMinProcessor)
	if src.min < smp.min {
		smp.min = src.min
	}
}

func (smp *statsMinProcessor) finalizeStats() string {
	return strconv.FormatFloat(smp.min, 'g', -1, 64)
}

func parseStatsMin(lex *lexer) (*statsMin, error) {
	lex.nextToken()
	fields, err := parseFieldNamesInParens(lex)
	if err != nil {
		return nil, fmt.Errorf("cannot parse 'min' args: %w", err)
	}
	if len(fields) == 0 {
		return nil, fmt.Errorf("'min' must contain at least one arg")
	}
	sm := &statsMin{
		fields:       fields,
		containsStar: slices.Contains(fields, "*"),
	}
	return sm, nil
}
