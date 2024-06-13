// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__html_escape = 192
)

const (
    _stack__html_escape = 72
)

const (
    _size__html_escape = 2064
)

var (
    _pcsp__html_escape = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {2045, 72},
        {2049, 48},
        {2050, 40},
        {2052, 32},
        {2054, 24},
        {2056, 16},
        {2058, 8},
        {2063, 0},
    }
)

var _cfunc_html_escape = []loader.CFunc{
    {"_html_escape_entry", 0,  _entry__html_escape, 0, nil},
    {"_html_escape", _entry__html_escape, _size__html_escape, _stack__html_escape, _pcsp__html_escape},
}