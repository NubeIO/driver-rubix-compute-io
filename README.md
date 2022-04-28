# nubeio-rubix-lib-pi-gpio-go

## to start
Needs to run as root for the PWMs to work
```
go build app.go && sudo ./app
```

## api

### write one as GET (this should just be used for the devloper)
value from 0-100, for the 2x DOs 0=off 100=on

GET
```
http://0.0.0.0:5001/api/outputs/UO1/1
```

### write all

Write all outputs at once (value from 0-100, for the 2x DOs 0=off 100=on)

```
http://0.0.0.0:5001/api/outputs/all/VALUE
```

### bulk write

POST
```
http://0.0.0.0:5001/api/outputs/bulk
```
BODY
Options for io_num ("UO1", "UO2", "UO3", "UO4", "UO5", "UO6", "DO1", "DO1")
```json
[
  {
    "IONum":"UO1",
    "value":22.2
  },
  {
    "IONum":"UO2",
    "value":22.2
  }
]
```


### write one

Write on at a time (value from 0-100, for the 2x DOs 0=off 100=on)

POST
```
http://0.0.0.0:5001/api/outputs
```
BODY
Options for io_num ("UO1", "UO2", "UO3", "UO4", "UO5", "UO6", "DO1", "DO1")
```json
{
    "io_num": "UO1",
    "value": 100
}
```



Get all input values
```
http://0.0.0.0:5001/api/inputs/all
```

# pigpio

```
sudo apt-get install pigpio
```

to start for dev
```
sudo pigpiod
```

## auto start

`/lib/systemd/system/pigpiod.service`

```
sudo systemctl enable pigpiod
sudo systemctl edit --full pigpiod
sudo systemctl reload-or-restart pigpiod
```

`ExecStart=/usr/bin/pigpiod -l -n 127.0.0.1` to open the soket

```
[Unit]
Description=Daemon required to control GPIO pins via pigpio
[Service]
ExecStart=/usr/bin/pigpiod -l -n 127.0.0.1
ExecStop=/bin/systemctl kill pigpiod
Type=forking
[Install]
WantedBy=multi-user.target
```

## enable hardware PWM (UO3 and UO6 only)

```
sudo nano /boot/config.txt
```

add this in

```
dtoverlay=pwm-2chan
```

On the Raspberry Pi, add dtoverlay=pwm-2chan to /boot/config.txt. This defaults to GPIO_18 as the pin for PWM0 and GPIO_19 as the pin for PWM1.
Alternatively, you can change GPIO_18 to GPIO_12 and GPIO_19 to GPIO_13 using dtoverlay=pwm-2chan,pin=12,func=4,pin2=13,func2=4.
Reboot your Raspberry Pi.
You can check everything is working on running lsmod | grep pwm and looking for pwm_bcm2835


## fix the 485 on the rubix-io

https://www.instructables.com/Raspberry-PI-3-Enable-Serial-Communications-to-Tty/

the bluthooth needs to be disabled on the PI

```
sudo nano /boot/config.txt
```

add this in

```
dtoverlay=disable-bt
```
