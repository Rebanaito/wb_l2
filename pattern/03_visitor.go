package main

type visitor interface {
	visitWindows(windows)
	visitLinux(linux)
	visitMac(mac)
}

type monitor struct {
	data string
}

func (mon *monitor) visitWindows(w windows) {
	mon.data = w.data
}

func (mon *monitor) visitLinux(l linux) {
	mon.data = l.data
}

func (mon *monitor) visitMac(m mac) {
	mon.data = m.data
}

type system interface {
	getOSInfo() string
	accept(visitor)
}

type windows struct {
	data string
}

func (w windows) getOSInfo() string {
	return "Windows" + w.data
}

func (w windows) accept(v visitor) {
	v.visitWindows(w)
}

type linux struct {
	data string
}

func (l linux) getOSInfo() string {
	return "Windows" + l.data
}

func (l linux) accept(v visitor) {
	v.visitLinux(l)
}

type mac struct {
	data string
}

func (m mac) getOSInfo() string {
	return "Windows" + m.data
}

func (m mac) accept(v visitor) {
	v.visitMac(m)
}
