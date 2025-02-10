# [`hardkernel` module](https://github.com/viam-modules/hardkernel)

This [hardkernel module](https://app.viam.com/module/viam/hardkernel) implements a hardkernel [Odroid-C4 board](https://www.hardkernel.com/shop/odroid-c4/) using the [`rdk:component:board` API](https://docs.viam.com/appendix/apis/components/board/).

## Setup

Follow the [setup guide](https://docs.viam.com/installation/prepare/odroid-c4-setup/) to prepare your Odroid-C4 for running `viam-server` before configuring an `odroid` board.

> [!NOTE]
> Before configuring your board, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add board / hardkernel:odroid to your machine](https://docs.viam.com/configure/#components).

## Configure your odroid board

### Example configuration
```json
{}
```

### Example configuration with optional analogs and digital interrupts
```json
  {
    "analogs": [
      {
      "name": "current",
      "pin": "1",
      "spi_bus": "1",
      "chip_select": "0"
      }
    ],
      "digital_interrupts": [
      {
        "name": "your-interrupt-1",
        "pin": "15"
      },
      {
        "name": "your-interrupt-2",
        "pin": "16"
      }
    ]
  }
```

### Attributes
The following attributes are available for `viam:hardkernel:odroid` boards:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `digital_interrupts` | object | Optional | Any digital interrupts's pin number and name.|
| `analogs` | object | Optional | Attributes of any pins that can be used as Analog-to-Digital Converter (ADC) inputs.|

### Analogs configuration
An [analog-to-digital converter](https://www.electronics-tutorials.ws/combination/analogue-to-digital-converter.html) (ADC) takes a continuous voltage input (analog signal) and converts it to an discrete integer output (digital signal).

To integrate an ADC into your machine, you must first physically connect the pins on your ADC to your board.

Then, integrate `analogs` into your board by adding the following to your board's `attributes` configuration:

```json {class="line-numbers linkable-line-numbers"}
"analogs": [
  {
    "name": "<your-analog-reader-name>",
    "pin": "<pin-number-on-adc>",
    "spi_bus": "<your-spi-bus-index>",
    "chip_select": "<chip-select-index>",
    "average_over_ms": <int>,
    "samples_per_sec": <int>
  }
]
```

#### Analog Attributes

The following attributes are available for `analogs`:

| Name | Type | Required? | Description |
| ---- | ---- | --------- | ----------- |
|`name` | string | **Required** | Your name for the analog reader. |
|`pin`| string | **Required** | The pin number of the ADC's connection pin, wired to the board. This should be labeled as the physical index of the pin on the ADC.
|`chip_select`| string | **Required** | The chip select index of the board's connection pin, wired to the ADC. |
|`spi_bus` | string | **Required** | The index of the SPI bus connecting the ADC and board. |
| `average_over_ms` | int | Optional | Duration in milliseconds over which the rolling average of the analog input should be taken. |
|`samples_per_sec` | int | Optional | Sampling rate of the analog input in samples per second. |

#### Example configuration

```json {class="line-numbers linkable-line-numbers"}
{
  "components": [
    {
      "name": "<your-hardkernel-odroid-board-name>",
      "model": "viam:hardkernel:odroid",
      "type": "board",
      "namespace": "rdk",
      "attributes": {
        "analogs": [
          {
            "name": "current",
            "pin": "1",
            "spi_bus": "1",
            "chip_select": "0"
          },
          {
            "name": "pressure",
            "pin": "0",
            "spi_bus": "1",
            "chip_select": "0"
          }
        ]
      }
    }
  ]
}
```

### Digital interrupt configuration
[Interrupts](https://en.wikipedia.org/wiki/Interrupt) are a method of signaling precise state changes.
Configuring digital interrupts to monitor GPIO pins on your board is useful when your application needs to know precisely when there is a change in GPIO value between high and low.

- When an interrupt configured on your board processes a change in the state of the GPIO pin it is configured to monitor, it ticks to record the state change.
  You can stream these ticks with the board API's [`StreamTicks()`](https://docs.viam.com/appendix/apis/components/board/#streamticks), or get the current value of the digital interrupt with [`Value()`](https://docs.viam.com/appendix/apis/components/board/#value).
- Calling [`GetGPIO()`](https://docs.viam.com/appendix/apis/components/board/#getgpio) on a GPIO pin, which you can do without configuring interrupts, is useful when you want to know a pin's value at specific points in your program, but is less precise and convenient than using an interrupt.

Integrate `digital_interrupts` into your machine in the `attributes` of your board by adding the following to your board's `attributes` configuration:

```json {class="line-numbers linkable-line-numbers"}
{
  "digital_interrupts": [
    {
      "name": "<your-digital-interrupt-name>",
      "pin": "<your-digital-interrupt-pin-number>"
    }
  ]
}
```

#### Digital Interrupt Attributes

The following attributes are available for `digital_interrupts`:


| Name | Type | Required? | Description |
| ---- | ---- | --------- | ----------- |
|`name` | string | **Required** | Your name for the digital interrupt. |
|`pin`| string | **Required** | The pin number of the board's GPIO pin that you wish to configure the digital interrupt for. |

#### Example configuration

```json {class="line-numbers linkable-line-numbers"}
{
  "components": [
    {
      "name": "<your-hardkernel-odroid-board-name>",
      "model": "viam:hardkernel:odroid",
      "type": "board",
      "namespace": "rdk",
      "attributes": {
        "digital_interrupts": [
          {
            "name": "your-interrupt-1",
            "pin": "15"
          },
          {
            "name": "your-interrupt-2",
            "pin": "16"
          }
        ]
      }
    }
  ]
}
```

## Next Steps
- To test your board, expand the **TEST** section of its configuration pane or go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your board, use one of the [available SDKs](https://docs.viam.com/sdks/).
- To view examples using a board component, explore [these tutorials](https://docs.viam.com/tutorials/).
