# rtl-sdr-framework
репозиторий с библиотеками и примерами работы с приемником RTL-SDR

## Подготовка к работе

1. Установка библиотеки

    `sudo apt-get install rtl-sdr`
   
2. Возможно, вам потребуется выгрузить модули ядра и внести модули в черный список:

    ```
    #lists modules using the rtl driver
    lsmod |grep rtl
    
    #remove them
    rmmod dvb_usb_rtl28xxu
    rmmod videodev
    rmmod rtl2832_sdr

    #optionally black list for the next reboot
   echo "blacklist dvb_usb_rtl28xxu" >> /etc/modprobe.d/blacklist.conf
   ```

3. 