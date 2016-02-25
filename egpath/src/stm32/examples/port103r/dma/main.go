// This example shows how to use DMA for memory to memory transfers.
package main

import (
	"fmt"
	"rtos"
	"unsafe"

	"stm32/hal/dma"
	"stm32/hal/irq"
	"stm32/hal/system"
	"stm32/hal/system/timer/rtc"
)

var (
	ds  *dma.Stream
	tce rtos.EventFlag
)

func init() {
	system.Setup(8, 72/8, false)
	rtc.Setup(32768)

	DMA := dma.DMA1
	DMA.EnableClock(false)
	ds = DMA.Stream(1)
	rtos.IRQ(irq.DMA1_Channel1).Enable()
}

// Try different element types (eg. P [...]uint16, M [...]byte) to see how DMA
// handles such asymetrical cases.
var (
	P = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -2, -3, -4, -5, -6, -7, -8, -9}
	M [len(P)]int
)

func main() {
	ds.Setup(dma.MTM|dma.IncP|dma.IncM, 0)
	ds.SetWordSize(unsafe.Sizeof(P[0]), unsafe.Sizeof(M[0]))
	ds.SetNum(len(P))
	ds.SetAddrP(unsafe.Pointer(&P[0]))
	ds.SetAddrM(unsafe.Pointer(&M[0]))
	ds.EnableInt(dma.TCE) // Simplified (should handle dma.ERR too).
	ds.Enable()
	tce.Wait(0)
	fmt.Println(M[:])
}

func dmaISR() {
	if ds.Events()&dma.TCE != 0 {
		ds.ClearEvents(dma.TCE)
		tce.Set()
	}
}

//emgo:const
//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.RTCAlarm:      rtc.ISR,
	irq.DMA1_Channel1: dmaISR,
}
