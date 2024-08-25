package main

import "regexp"

type ModPVDBScanner struct {
	ModFolder string
}

var (
	pvIDRe    = regexp.MustCompile(`(pv_\d+)`)
	pvExRe    = regexp.MustCompile(`\.extreme\.0\.level=PV_LV_(\d+_\d+)`)
	pvExtraRe = regexp.MustCompile(`\.extreme\.1\.level=PV_LV_(\d+_\d+)`)
)