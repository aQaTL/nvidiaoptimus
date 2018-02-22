package nvidiaoptimus

/*
#cgo LDFLAGS: -L${SRCDIR}/flags.def

__declspec(dllexport) unsigned long NvOptimusEnablement = 0x00000001;
__declspec(dllexport) int AmdPowerXpressRequestHighPerformance = 1;
*/
import "C"
