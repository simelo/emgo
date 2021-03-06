// Peripheral: GPIO_Periph  General Purpose I/O.
// Instances:
//  GPIOA  mmap.GPIOA_BASE
//  GPIOB  mmap.GPIOB_BASE
//  GPIOC  mmap.GPIOC_BASE
//  GPIOD  mmap.GPIOD_BASE
//  GPIOE  mmap.GPIOE_BASE
//  GPIOH  mmap.GPIOH_BASE
// Registers:
//  0x00 32  MODER   Port mode register.
//  0x04 32  OTYPER  Port output type register.
//  0x08 32  OSPEEDR Port output speed register.
//  0x0C 32  PUPDR   Port pull-up/pull-down register.
//  0x10 32  IDR     Port input data register.
//  0x14 32  ODR     Port output data register.
//  0x18 32  BSRR    Port bit set/reset register.
//  0x1C 32  LCKR    Port configuration lock register.
//  0x20 32  AFR[2]  Alternate function registers.
// Import:
//  stm32/o/f411xe/mmap
package gpio

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MODE0   MODER = 0x03 << 0  //+
	MODE1   MODER = 0x03 << 2  //+
	MODE2   MODER = 0x03 << 4  //+
	MODE3   MODER = 0x03 << 6  //+
	MODE4   MODER = 0x03 << 8  //+
	MODE5   MODER = 0x03 << 10 //+
	MODE6   MODER = 0x03 << 12 //+
	MODE7   MODER = 0x03 << 14 //+
	MODE8   MODER = 0x03 << 16 //+
	MODE9   MODER = 0x03 << 18 //+
	MODE10  MODER = 0x03 << 20 //+
	MODE11  MODER = 0x03 << 22 //+
	MODE12  MODER = 0x03 << 24 //+
	MODE13  MODER = 0x03 << 26 //+
	MODE14  MODER = 0x03 << 28 //+
	MODE15  MODER = 0x03 << 30 //+
	MODER0  MODER = 0x03 << 0
	MODER1  MODER = 0x03 << 2
	MODER2  MODER = 0x03 << 4
	MODER3  MODER = 0x03 << 6
	MODER4  MODER = 0x03 << 8
	MODER5  MODER = 0x03 << 10
	MODER6  MODER = 0x03 << 12
	MODER7  MODER = 0x03 << 14
	MODER8  MODER = 0x03 << 16
	MODER9  MODER = 0x03 << 18
	MODER10 MODER = 0x03 << 20
	MODER11 MODER = 0x03 << 22
	MODER12 MODER = 0x03 << 24
	MODER13 MODER = 0x03 << 26
	MODER14 MODER = 0x03 << 28
	MODER15 MODER = 0x03 << 30
)

const (
	MODE0n  = 0
	MODE1n  = 2
	MODE2n  = 4
	MODE3n  = 6
	MODE4n  = 8
	MODE5n  = 10
	MODE6n  = 12
	MODE7n  = 14
	MODE8n  = 16
	MODE9n  = 18
	MODE10n = 20
	MODE11n = 22
	MODE12n = 24
	MODE13n = 26
	MODE14n = 28
	MODE15n = 30
)

const (
	OT0  OTYPER = 0x01 << 0  //+
	OT1  OTYPER = 0x01 << 1  //+
	OT2  OTYPER = 0x01 << 2  //+
	OT3  OTYPER = 0x01 << 3  //+
	OT4  OTYPER = 0x01 << 4  //+
	OT5  OTYPER = 0x01 << 5  //+
	OT6  OTYPER = 0x01 << 6  //+
	OT7  OTYPER = 0x01 << 7  //+
	OT8  OTYPER = 0x01 << 8  //+
	OT9  OTYPER = 0x01 << 9  //+
	OT10 OTYPER = 0x01 << 10 //+
	OT11 OTYPER = 0x01 << 11 //+
	OT12 OTYPER = 0x01 << 12 //+
	OT13 OTYPER = 0x01 << 13 //+
	OT14 OTYPER = 0x01 << 14 //+
	OT15 OTYPER = 0x01 << 15 //+
)

const (
	OT0n  = 0
	OT1n  = 1
	OT2n  = 2
	OT3n  = 3
	OT4n  = 4
	OT5n  = 5
	OT6n  = 6
	OT7n  = 7
	OT8n  = 8
	OT9n  = 9
	OT10n = 10
	OT11n = 11
	OT12n = 12
	OT13n = 13
	OT14n = 14
	OT15n = 15
)

const (
	OSPEED0  OSPEEDR = 0x03 << 0  //+
	OSPEED1  OSPEEDR = 0x03 << 2  //+
	OSPEED2  OSPEEDR = 0x03 << 4  //+
	OSPEED3  OSPEEDR = 0x03 << 6  //+
	OSPEED4  OSPEEDR = 0x03 << 8  //+
	OSPEED5  OSPEEDR = 0x03 << 10 //+
	OSPEED6  OSPEEDR = 0x03 << 12 //+
	OSPEED7  OSPEEDR = 0x03 << 14 //+
	OSPEED8  OSPEEDR = 0x03 << 16 //+
	OSPEED9  OSPEEDR = 0x03 << 18 //+
	OSPEED10 OSPEEDR = 0x03 << 20 //+
	OSPEED11 OSPEEDR = 0x03 << 22 //+
	OSPEED12 OSPEEDR = 0x03 << 24 //+
	OSPEED13 OSPEEDR = 0x03 << 26 //+
	OSPEED14 OSPEEDR = 0x03 << 28 //+
	OSPEED15 OSPEEDR = 0x03 << 30 //+
)

const (
	OSPEED0n  = 0
	OSPEED1n  = 2
	OSPEED2n  = 4
	OSPEED3n  = 6
	OSPEED4n  = 8
	OSPEED5n  = 10
	OSPEED6n  = 12
	OSPEED7n  = 14
	OSPEED8n  = 16
	OSPEED9n  = 18
	OSPEED10n = 20
	OSPEED11n = 22
	OSPEED12n = 24
	OSPEED13n = 26
	OSPEED14n = 28
	OSPEED15n = 30
)

const (
	PUPD0  PUPDR = 0x03 << 0  //+
	PUPD1  PUPDR = 0x03 << 2  //+
	PUPD2  PUPDR = 0x03 << 4  //+
	PUPD3  PUPDR = 0x03 << 6  //+
	PUPD4  PUPDR = 0x03 << 8  //+
	PUPD5  PUPDR = 0x03 << 10 //+
	PUPD6  PUPDR = 0x03 << 12 //+
	PUPD7  PUPDR = 0x03 << 14 //+
	PUPD8  PUPDR = 0x03 << 16 //+
	PUPD9  PUPDR = 0x03 << 18 //+
	PUPD10 PUPDR = 0x03 << 20 //+
	PUPD11 PUPDR = 0x03 << 22 //+
	PUPD12 PUPDR = 0x03 << 24 //+
	PUPD13 PUPDR = 0x03 << 26 //+
	PUPD14 PUPDR = 0x03 << 28 //+
	PUPD15 PUPDR = 0x03 << 30 //+
)

const (
	PUPD0n  = 0
	PUPD1n  = 2
	PUPD2n  = 4
	PUPD3n  = 6
	PUPD4n  = 8
	PUPD5n  = 10
	PUPD6n  = 12
	PUPD7n  = 14
	PUPD8n  = 16
	PUPD9n  = 18
	PUPD10n = 20
	PUPD11n = 22
	PUPD12n = 24
	PUPD13n = 26
	PUPD14n = 28
	PUPD15n = 30
)

const (
	ID0  IDR = 0x01 << 0  //+
	ID1  IDR = 0x01 << 1  //+
	ID2  IDR = 0x01 << 2  //+
	ID3  IDR = 0x01 << 3  //+
	ID4  IDR = 0x01 << 4  //+
	ID5  IDR = 0x01 << 5  //+
	ID6  IDR = 0x01 << 6  //+
	ID7  IDR = 0x01 << 7  //+
	ID8  IDR = 0x01 << 8  //+
	ID9  IDR = 0x01 << 9  //+
	ID10 IDR = 0x01 << 10 //+
	ID11 IDR = 0x01 << 11 //+
	ID12 IDR = 0x01 << 12 //+
	ID13 IDR = 0x01 << 13 //+
	ID14 IDR = 0x01 << 14 //+
	ID15 IDR = 0x01 << 15 //+
)

const (
	ID0n  = 0
	ID1n  = 1
	ID2n  = 2
	ID3n  = 3
	ID4n  = 4
	ID5n  = 5
	ID6n  = 6
	ID7n  = 7
	ID8n  = 8
	ID9n  = 9
	ID10n = 10
	ID11n = 11
	ID12n = 12
	ID13n = 13
	ID14n = 14
	ID15n = 15
)

const (
	OD0  ODR = 0x01 << 0  //+
	OD1  ODR = 0x01 << 1  //+
	OD2  ODR = 0x01 << 2  //+
	OD3  ODR = 0x01 << 3  //+
	OD4  ODR = 0x01 << 4  //+
	OD5  ODR = 0x01 << 5  //+
	OD6  ODR = 0x01 << 6  //+
	OD7  ODR = 0x01 << 7  //+
	OD8  ODR = 0x01 << 8  //+
	OD9  ODR = 0x01 << 9  //+
	OD10 ODR = 0x01 << 10 //+
	OD11 ODR = 0x01 << 11 //+
	OD12 ODR = 0x01 << 12 //+
	OD13 ODR = 0x01 << 13 //+
	OD14 ODR = 0x01 << 14 //+
	OD15 ODR = 0x01 << 15 //+
)

const (
	OD0n  = 0
	OD1n  = 1
	OD2n  = 2
	OD3n  = 3
	OD4n  = 4
	OD5n  = 5
	OD6n  = 6
	OD7n  = 7
	OD8n  = 8
	OD9n  = 9
	OD10n = 10
	OD11n = 11
	OD12n = 12
	OD13n = 13
	OD14n = 14
	OD15n = 15
)

const (
	BS0  BSRR = 0x01 << 0  //+
	BS1  BSRR = 0x01 << 1  //+
	BS2  BSRR = 0x01 << 2  //+
	BS3  BSRR = 0x01 << 3  //+
	BS4  BSRR = 0x01 << 4  //+
	BS5  BSRR = 0x01 << 5  //+
	BS6  BSRR = 0x01 << 6  //+
	BS7  BSRR = 0x01 << 7  //+
	BS8  BSRR = 0x01 << 8  //+
	BS9  BSRR = 0x01 << 9  //+
	BS10 BSRR = 0x01 << 10 //+
	BS11 BSRR = 0x01 << 11 //+
	BS12 BSRR = 0x01 << 12 //+
	BS13 BSRR = 0x01 << 13 //+
	BS14 BSRR = 0x01 << 14 //+
	BS15 BSRR = 0x01 << 15 //+
	BR0  BSRR = 0x01 << 16 //+
	BR1  BSRR = 0x01 << 17 //+
	BR2  BSRR = 0x01 << 18 //+
	BR3  BSRR = 0x01 << 19 //+
	BR4  BSRR = 0x01 << 20 //+
	BR5  BSRR = 0x01 << 21 //+
	BR6  BSRR = 0x01 << 22 //+
	BR7  BSRR = 0x01 << 23 //+
	BR8  BSRR = 0x01 << 24 //+
	BR9  BSRR = 0x01 << 25 //+
	BR10 BSRR = 0x01 << 26 //+
	BR11 BSRR = 0x01 << 27 //+
	BR12 BSRR = 0x01 << 28 //+
	BR13 BSRR = 0x01 << 29 //+
	BR14 BSRR = 0x01 << 30 //+
	BR15 BSRR = 0x01 << 31 //+
)

const (
	BS0n  = 0
	BS1n  = 1
	BS2n  = 2
	BS3n  = 3
	BS4n  = 4
	BS5n  = 5
	BS6n  = 6
	BS7n  = 7
	BS8n  = 8
	BS9n  = 9
	BS10n = 10
	BS11n = 11
	BS12n = 12
	BS13n = 13
	BS14n = 14
	BS15n = 15
	BR0n  = 16
	BR1n  = 17
	BR2n  = 18
	BR3n  = 19
	BR4n  = 20
	BR5n  = 21
	BR6n  = 22
	BR7n  = 23
	BR8n  = 24
	BR9n  = 25
	BR10n = 26
	BR11n = 27
	BR12n = 28
	BR13n = 29
	BR14n = 30
	BR15n = 31
)

const (
	LCK0  LCKR = 0x01 << 0  //+
	LCK1  LCKR = 0x01 << 1  //+
	LCK2  LCKR = 0x01 << 2  //+
	LCK3  LCKR = 0x01 << 3  //+
	LCK4  LCKR = 0x01 << 4  //+
	LCK5  LCKR = 0x01 << 5  //+
	LCK6  LCKR = 0x01 << 6  //+
	LCK7  LCKR = 0x01 << 7  //+
	LCK8  LCKR = 0x01 << 8  //+
	LCK9  LCKR = 0x01 << 9  //+
	LCK10 LCKR = 0x01 << 10 //+
	LCK11 LCKR = 0x01 << 11 //+
	LCK12 LCKR = 0x01 << 12 //+
	LCK13 LCKR = 0x01 << 13 //+
	LCK14 LCKR = 0x01 << 14 //+
	LCK15 LCKR = 0x01 << 15 //+
	LCKK  LCKR = 0x01 << 16 //+
)

const (
	LCK0n  = 0
	LCK1n  = 1
	LCK2n  = 2
	LCK3n  = 3
	LCK4n  = 4
	LCK5n  = 5
	LCK6n  = 6
	LCK7n  = 7
	LCK8n  = 8
	LCK9n  = 9
	LCK10n = 10
	LCK11n = 11
	LCK12n = 12
	LCK13n = 13
	LCK14n = 14
	LCK15n = 15
	LCKKn  = 16
)
