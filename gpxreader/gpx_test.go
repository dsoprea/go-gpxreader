package gpxreader

import (
    "testing"
    "bytes"

    "github.com/dsoprea/go-logging"
)

const (
    testGpxData = `<?xml version="1.0" encoding="UTF-8" ?><gpx version="1.0" creator="GPSLogger 79 - http://gpslogger.mendhak.com/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://www.topografix.com/GPX/1/0" xsi:schemaLocation="http://www.topografix.com/GPX/1/0 http://www.topografix.com/GPX/1/0/gpx.xsd"><time>2016-12-02T08:05:44Z</time><trk><trkseg><trkpt lat="47.61360231405151" lon="-122.33966135543184"><ele>12.178885901563254</ele><time>2016-12-02T08:05:44Z</time><geoidheight>-18.5</geoidheight><src>gps</src><sat>4</sat><hdop>3.8</hdop><vdop>1.0</vdop><pdop>3.9</pdop></trkpt>
<trkpt lat="47.61306246897419" lon="-122.34018970742765"><ele>-36.559077848208034</ele><time>2016-12-02T08:11:35Z</time><speed>0.0</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>6</sat><hdop>2.0</hdop><vdop>6.0</vdop><pdop>6.3</pdop></trkpt>
<trkpt lat="47.611375321719336" lon="-122.34180135324286"><ele>9.95224723728763</ele><time>2016-12-02T08:18:18Z</time><geoidheight>-18.5</geoidheight><src>gps</src><sat>4</sat><hdop>1.5</hdop><vdop>1.0</vdop><pdop>1.8</pdop></trkpt>
<trkpt lat="47.61225667251922" lon="-122.34195911786182"><ele>152.97120701674527</ele><time>2016-12-02T08:24:21Z</time><speed>0.0</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>7</sat><hdop>1.9</hdop><vdop>2.4</vdop><pdop>3.0</pdop></trkpt>
<trkpt lat="47.61029656609932" lon="-122.34047933763608"><ele>134.50190284738613</ele><time>2016-12-02T08:36:31Z</time><course>215.26854</course><speed>0.6004517</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>6</sat><hdop>1.2</hdop><vdop>2.1</vdop><pdop>2.4</pdop></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T08:48:08Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T08:54:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T08:59:22Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T09:04:59Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T09:10:20Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T09:16:20Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T09:22:21Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T09:27:23Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T09:33:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T09:38:03Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T09:44:49Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T09:49:53Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T09:57:25Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T10:02:45Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T10:08:26Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T10:14:27Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T10:20:31Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T10:26:32Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T10:32:33Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T10:38:33Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T10:44:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T10:52:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T10:59:56Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T11:05:19Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T11:14:03Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T11:19:41Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T11:24:49Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T11:30:11Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T11:36:17Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T11:42:18Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T11:47:20Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T11:52:57Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T11:58:08Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T12:03:29Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T12:08:31Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T12:14:04Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T12:20:06Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T12:26:08Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T12:31:29Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T12:36:51Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T12:43:25Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T12:49:26Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T12:54:59Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T13:00:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T13:06:02Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T13:13:05Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T13:19:09Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T13:25:11Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T13:30:27Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T13:36:14Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T13:43:16Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T13:49:17Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T13:55:01Z</time><src>network</src></trkpt>
<trkpt lat="47.61148151435243" lon="-122.34157821647844"><ele>186.34489492647512</ele><time>2016-12-02T14:00:25Z</time><speed>0.0</speed><geoidheight>-18.5</geoidheight><src>gps</src><hdop>0.8</hdop><vdop>1.6</vdop><pdop>1.8</pdop></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T14:05:24Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T14:10:26Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T14:15:27Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T14:22:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T14:27:02Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T14:33:05Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T14:38:07Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T14:43:09Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T14:49:10Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T14:55:23Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T15:01:23Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T15:06:46Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T15:12:20Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T15:18:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T15:23:54Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T15:28:57Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T15:35:04Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T15:40:24Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T15:45:13Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T15:50:14Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T15:55:14Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T16:00:33Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T16:06:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T16:11:21Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T16:16:23Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T16:21:46Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T16:27:08Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T16:33:18Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T16:39:18Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-02T16:44:54Z</time><src>network</src></trkpt>
<trkpt lat="47.61251579786422" lon="-122.34061308434823"><ele>75.33739809965321</ele><time>2016-12-02T16:49:56Z</time><speed>0.0</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>13</sat><hdop>1.9</hdop><vdop>2.5</vdop><pdop>3.2</pdop></trkpt>
<trkpt lat="47.6132103" lon="-122.3401816"><time>2016-12-02T16:54:56Z</time><src>network</src></trkpt>
<trkpt lat="47.613320042004325" lon="-122.34003897612189"><ele>43.76773177464807</ele><time>2016-12-02T17:00:12Z</time><course>147.23445</course><speed>1.5635744</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>9</sat><hdop>1.3</hdop><vdop>1.8</vdop><pdop>2.2</pdop></trkpt>
<trkpt lat="47.6154928" lon="-122.3379666"><time>2016-12-02T17:09:05Z</time><src>network</src></trkpt>
<trkpt lat="47.6154928" lon="-122.3379666"><time>2016-12-02T17:14:08Z</time><src>network</src></trkpt>
<trkpt lat="47.6154928" lon="-122.3379666"><time>2016-12-02T17:20:24Z</time><src>network</src></trkpt>
<trkpt lat="47.6150979" lon="-122.3389477"><time>2016-12-02T17:31:06Z</time><src>network</src></trkpt>
<trkpt lat="47.6150979" lon="-122.3389477"><time>2016-12-02T17:36:08Z</time><src>network</src></trkpt>
<trkpt lat="47.6150979" lon="-122.3389477"><time>2016-12-02T17:45:35Z</time><src>network</src></trkpt>
<trkpt lat="47.6150979" lon="-122.3389477"><time>2016-12-02T17:53:33Z</time><src>network</src></trkpt>
<trkpt lat="47.6153527" lon="-122.3382582"><time>2016-12-02T18:34:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6152872" lon="-122.3382214"><time>2016-12-02T20:04:53Z</time><src>network</src></trkpt>
<trkpt lat="47.6152739" lon="-122.3381891"><time>2016-12-02T20:10:34Z</time><src>network</src></trkpt>
<trkpt lat="47.6152145" lon="-122.3393927"><time>2016-12-02T20:20:24Z</time><src>network</src></trkpt>
<trkpt lat="47.6150922" lon="-122.3392372"><time>2016-12-02T20:29:54Z</time><src>network</src></trkpt>
<trkpt lat="47.6150922" lon="-122.3392372"><time>2016-12-02T20:40:03Z</time><src>network</src></trkpt>
<trkpt lat="47.6153045" lon="-122.3381145"><time>2016-12-02T21:09:54Z</time><src>network</src></trkpt>
<trkpt lat="47.6153362" lon="-122.3382097"><time>2016-12-02T21:20:25Z</time><src>network</src></trkpt>
<trkpt lat="47.6153534" lon="-122.3382632"><time>2016-12-02T21:29:56Z</time><src>network</src></trkpt>
<trkpt lat="47.6153476" lon="-122.3382758"><time>2016-12-02T21:36:10Z</time><src>network</src></trkpt>
<trkpt lat="47.6152931" lon="-122.3381883"><time>2016-12-02T21:44:54Z</time><src>network</src></trkpt>
<trkpt lat="47.6153508" lon="-122.3382583"><time>2016-12-02T21:59:03Z</time><src>network</src></trkpt>
<trkpt lat="47.6152601" lon="-122.3381274"><time>2016-12-02T22:09:27Z</time><src>network</src></trkpt>
<trkpt lat="47.6152899" lon="-122.3381521"><time>2016-12-02T22:14:54Z</time><src>network</src></trkpt>
<trkpt lat="47.615352" lon="-122.3382936"><time>2016-12-02T22:24:55Z</time><src>network</src></trkpt>
<trkpt lat="47.6153489" lon="-122.3382749"><time>2016-12-02T22:30:24Z</time><src>network</src></trkpt>
<trkpt lat="47.6153432" lon="-122.338263"><time>2016-12-02T22:39:53Z</time><src>network</src></trkpt>
<trkpt lat="47.615343" lon="-122.3382612"><time>2016-12-02T22:44:54Z</time><src>network</src></trkpt>
<trkpt lat="47.6153143" lon="-122.3382468"><time>2016-12-02T22:55:30Z</time><src>network</src></trkpt>
<trkpt lat="47.6153068" lon="-122.338241"><time>2016-12-02T23:04:54Z</time><src>network</src></trkpt>
<trkpt lat="47.6153437" lon="-122.3382571"><time>2016-12-02T23:11:27Z</time><src>network</src></trkpt>
<trkpt lat="47.6152428" lon="-122.3382415"><time>2016-12-02T23:17:21Z</time><src>network</src></trkpt>
<trkpt lat="47.6152703" lon="-122.3381705"><time>2016-12-02T23:22:23Z</time><src>network</src></trkpt>
<trkpt lat="47.615236" lon="-122.3381644"><time>2016-12-02T23:28:06Z</time><src>network</src></trkpt>
<trkpt lat="47.6152394" lon="-122.3381819"><time>2016-12-02T23:34:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6152497" lon="-122.3381881"><time>2016-12-02T23:39:22Z</time><src>network</src></trkpt>
<trkpt lat="47.6152638" lon="-122.338192"><time>2016-12-02T23:44:53Z</time><src>network</src></trkpt>
<trkpt lat="47.6152461" lon="-122.3381488"><time>2016-12-02T23:50:26Z</time><src>network</src></trkpt>
<trkpt lat="47.6152537" lon="-122.3381541"><time>2016-12-02T23:59:53Z</time><src>network</src></trkpt>
<trkpt lat="47.6155587" lon="-122.3384755"><time>2016-12-03T00:07:05Z</time><src>network</src></trkpt>
<trkpt lat="47.61298107284424" lon="-122.33871627956734"><ele>142.20351801295345</ele><time>2016-12-03T00:16:08Z</time><course>229.11317</course><speed>1.0412709</speed><geoidheight>-18.5</geoidheight><src>gps</src><hdop>1.6</hdop><vdop>99.0</vdop><pdop>140.0</pdop></trkpt>
<trkpt lat="47.6098334" lon="-122.3407079"><time>2016-12-03T00:22:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6072374" lon="-122.3380962"><time>2016-12-03T00:39:56Z</time><src>network</src></trkpt>
<trkpt lat="47.6078864" lon="-122.3376943"><time>2016-12-03T00:49:29Z</time><src>network</src></trkpt>
<trkpt lat="47.6077661" lon="-122.3391173"><time>2016-12-03T01:04:04Z</time><src>network</src></trkpt>
<trkpt lat="47.6078272" lon="-122.3391951"><time>2016-12-03T01:09:12Z</time><src>network</src></trkpt>
<trkpt lat="47.6099305" lon="-122.3413504"><time>2016-12-03T01:15:55Z</time><src>network</src></trkpt>
<trkpt lat="47.6099305" lon="-122.3413504"><time>2016-12-03T01:25:50Z</time><src>network</src></trkpt>
<trkpt lat="47.6100054" lon="-122.34152"><time>2016-12-03T01:30:52Z</time><src>network</src></trkpt>
<trkpt lat="47.6102222" lon="-122.3416473"><time>2016-12-03T01:36:03Z</time><src>network</src></trkpt>
<trkpt lat="47.6099305" lon="-122.3413504"><time>2016-12-03T01:41:08Z</time><src>network</src></trkpt>
<trkpt lat="47.6099305" lon="-122.3413504"><time>2016-12-03T01:46:24Z</time><src>network</src></trkpt>
<trkpt lat="47.6099305" lon="-122.3413504"><time>2016-12-03T01:54:57Z</time><src>network</src></trkpt>
<trkpt lat="47.6099305" lon="-122.3413504"><time>2016-12-03T02:00:25Z</time><src>network</src></trkpt>
<trkpt lat="47.6116345" lon="-122.3398835"><time>2016-12-03T02:09:57Z</time><src>network</src></trkpt>
<trkpt lat="47.61232014244754" lon="-122.34151466097569"><ele>173.34623878732407</ele><time>2016-12-03T02:15:12Z</time><speed>0.0</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>9</sat><hdop>3.2</hdop><vdop>2.3</vdop><pdop>3.2</pdop></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T02:21:00Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T02:27:16Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T02:32:18Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T02:37:44Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T02:43:43Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T02:49:43Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T02:54:51Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T02:59:57Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T03:05:14Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T03:11:24Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T03:16:34Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T03:22:12Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T03:27:23Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T03:32:30Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T03:37:27Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T03:42:40Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T03:48:41Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T03:54:07Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T03:59:29Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T04:04:59Z</time><src>network</src></trkpt>
<trkpt lat="47.6130798" lon="-122.3396448"><time>2016-12-03T04:10:47Z</time><src>network</src></trkpt>
<trkpt lat="47.6130798" lon="-122.3396448"><time>2016-12-03T04:16:08Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T04:21:10Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T04:26:32Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T04:31:47Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T04:36:48Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T04:42:31Z</time><src>network</src></trkpt>
<trkpt lat="47.6131631" lon="-122.3401957"><time>2016-12-03T04:48:01Z</time><src>network</src></trkpt>
<trkpt lat="47.6132797" lon="-122.3406406"><time>2016-12-03T04:54:57Z</time><src>network</src></trkpt>
<trkpt lat="47.61330576372484" lon="-122.34079956900972"><ele>80.0910692711343</ele><time>2016-12-03T04:59:58Z</time><speed>0.0</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>11</sat><hdop>0.8</hdop><vdop>1.7</vdop><pdop>2.0</pdop></trkpt>
<trkpt lat="47.61978440782248" lon="-122.34897627009876"><ele>21.374163185436377</ele><time>2016-12-03T05:05:14Z</time><course>260.0014</course><speed>3.6852572</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>11</sat><hdop>1.1</hdop><vdop>1.5</vdop><pdop>1.9</pdop></trkpt>
<trkpt lat="47.62043645310058" lon="-122.34757546741446"><ele>21.205325451125926</ele><time>2016-12-03T05:11:08Z</time><geoidheight>-18.5</geoidheight><src>gps</src><hdop>1.6</hdop><vdop>1.7</vdop><pdop>2.4</pdop></trkpt>
<trkpt lat="47.6198321" lon="-122.3467486"><time>2016-12-03T05:19:57Z</time><src>network</src></trkpt>
<trkpt lat="47.62072632728673" lon="-122.3488931669913"><ele>11.863940261453328</ele><time>2016-12-03T05:25:00Z</time><speed>0.0</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>17</sat><hdop>1.2</hdop><vdop>1.4</vdop><pdop>1.7</pdop></trkpt>
<trkpt lat="47.6201983" lon="-122.3488465"><time>2016-12-03T05:30:36Z</time><src>network</src></trkpt>
<trkpt lat="47.621303730503676" lon="-122.3517437806509"><ele>43.59019639053908</ele><time>2016-12-03T05:40:02Z</time><course>7.936195</course><speed>0.5228159</speed><geoidheight>-18.5</geoidheight><src>gps</src><hdop>2.4</hdop><vdop>99.0</vdop><pdop>140.0</pdop></trkpt>
<trkpt lat="47.62159036331387" lon="-122.35166163494243"><ele>39.45286409103766</ele><time>2016-12-03T05:45:15Z</time><course>315.74463</course><speed>0.7217953</speed><geoidheight>-18.5</geoidheight><src>gps</src><hdop>2.4</hdop><vdop>99.0</vdop><pdop>140.0</pdop></trkpt>
<trkpt lat="47.620526" lon="-122.3497085"><time>2016-12-03T05:54:57Z</time><src>network</src></trkpt>
<trkpt lat="47.62117400309288" lon="-122.34963559222493"><ele>16.893907682818714</ele><time>2016-12-03T06:00:02Z</time><course>110.82835</course><speed>1.7929946</speed><geoidheight>-18.5</geoidheight><src>gps</src><hdop>1.6</hdop><vdop>99.0</vdop><pdop>140.0</pdop></trkpt>
<trkpt lat="47.62077514326221" lon="-122.34775468389304"><ele>9.805802846285179</ele><time>2016-12-03T06:05:03Z</time><course>96.2158</course><speed>3.4758363</speed><geoidheight>-18.5</geoidheight><src>gps</src><hdop>1.6</hdop><vdop>99.0</vdop><pdop>140.0</pdop></trkpt>
<trkpt lat="47.62040030200266" lon="-122.3471644040348"><ele>40.27648798481647</ele><time>2016-12-03T06:10:10Z</time><geoidheight>-18.5</geoidheight><src>gps</src><sat>8</sat><hdop>1.2</hdop><vdop>2.1</vdop><pdop>2.4</pdop></trkpt>
<trkpt lat="47.62438526076109" lon="-122.34385781793155"><ele>5.506018253201048</ele><time>2016-12-03T06:19:57Z</time><course>157.47095</course><speed>0.0</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>9</sat><hdop>3.2</hdop><vdop>1.7</vdop><pdop>2.7</pdop></trkpt>
<trkpt lat="47.6250238" lon="-122.340673"><time>2016-12-03T06:25:48Z</time><src>network</src></trkpt>
<trkpt lat="47.63617446260113" lon="-122.3228045129993"><ele>27.439293557701106</ele><time>2016-12-03T06:31:06Z</time><course>6.1948338</course><speed>25.811926</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>9</sat><hdop>0.8</hdop><vdop>1.2</vdop><pdop>1.5</pdop></trkpt>
<trkpt lat="47.720317174661616" lon="-122.3242370460064"><ele>100.4168786473231</ele><time>2016-12-03T06:36:50Z</time><course>0.49949867</course><speed>29.35491</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>17</sat><hdop>1.1</hdop><vdop>1.5</vdop><pdop>1.9</pdop></trkpt>
<trkpt lat="47.802657508360994" lon="-122.31231981802541"><ele>114.78147119005848</ele><time>2016-12-03T06:42:14Z</time><course>35.682568</course><speed>33.44422</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>16</sat><hdop>0.6</hdop><vdop>1.4</vdop><pdop>1.6</pdop></trkpt>
<trkpt lat="47.869487376749596" lon="-122.24456301762936"><ele>148.28819500740258</ele><time>2016-12-03T06:47:32Z</time><course>9.475927</course><speed>31.510231</speed><geoidheight>-18.5</geoidheight><src>gps</src><sat>16</sat><hdop>0.9</hdop><vdop>1.5</vdop><pdop>1.8</pdop></trkpt>
<trkpt lat="47.94483443774716" lon="-122.19732456918234"><ele>50.73247394137797</ele><time>2016-12-03T06:52:52Z</time><course>347.22052</course><speed>29.362377</speed><geoidheight>-18.2</geoidheight><src>gps</src><sat>17</sat><hdop>0.8</hdop><vdop>1.7</vdop><pdop>1.9</pdop></trkpt>
<trkpt lat="48.02697678209969" lon="-122.17633570422109"><ele>-14.269772078513668</ele><time>2016-12-03T06:58:08Z</time><course>352.69046</course><speed>31.765324</speed><geoidheight>-18.2</geoidheight><src>gps</src><sat>13</sat><hdop>0.6</hdop><vdop>1.5</vdop><pdop>1.7</pdop></trkpt>
<trkpt lat="48.11475732919142" lon="-122.18479889947716"><ele>10.834580743006606</ele><time>2016-12-03T07:03:08Z</time><course>358.32947</course><speed>34.87873</speed><geoidheight>-18.2</geoidheight><src>gps</src><sat>17</sat><hdop>0.8</hdop><vdop>1.4</vdop><pdop>1.7</pdop></trkpt>
<trkpt lat="48.20252441851578" lon="-122.21257459392434"><ele>18.256680104632288</ele><time>2016-12-03T07:08:13Z</time><course>342.01358</course><speed>34.46667</speed><geoidheight>-18.0</geoidheight><src>gps</src><sat>16</sat><hdop>1.2</hdop><vdop>2.3</vdop><pdop>2.6</pdop></trkpt>
<trkpt lat="48.28108100565124" lon="-122.28412059084324"><ele>89.98642657773985</ele><time>2016-12-03T07:13:14Z</time><course>321.61</course><speed>34.64136</speed><geoidheight>-18.0</geoidheight><src>gps</src><sat>13</sat><hdop>0.8</hdop><vdop>1.3</vdop><pdop>1.6</pdop></trkpt>
<trkpt lat="48.368254266776354" lon="-122.33560039736219"><ele>-23.237531588452157</ele><time>2016-12-03T07:18:15Z</time><course>1.6491731</course><speed>34.674477</speed><geoidheight>-18.0</geoidheight><src>gps</src><sat>17</sat><hdop>0.6</hdop><vdop>1.8</vdop><pdop>2.0</pdop></trkpt>
<trkpt lat="48.45672936673762" lon="-122.34140644128601"><ele>-3.644368270854601</ele><time>2016-12-03T07:23:50Z</time><course>358.73965</course><speed>30.805782</speed><geoidheight>-17.9</geoidheight><src>gps</src><sat>17</sat><hdop>0.6</hdop><vdop>1.2</vdop><pdop>1.4</pdop></trkpt>
<trkpt lat="48.546254691830875" lon="-122.35022561803628"><ele>46.32385460163984</ele><time>2016-12-03T07:29:20Z</time><course>18.490294</course><speed>28.29401</speed><geoidheight>-17.9</geoidheight><src>gps</src><sat>17</sat><hdop>0.8</hdop><vdop>1.5</vdop><pdop>1.7</pdop></trkpt>
<trkpt lat="48.640913512036875" lon="-122.36255887993427"><ele>71.26325062303715</ele><time>2016-12-03T07:35:09Z</time><course>344.61734</course><speed>28.847834</speed><geoidheight>-17.7</geoidheight><src>gps</src><sat>17</sat><hdop>0.4</hdop><vdop>1.5</vdop><pdop>1.6</pdop></trkpt>
<trkpt lat="48.691469756556415" lon="-122.43486395334077"><ele>127.27657600559903</ele><time>2016-12-03T07:40:31Z</time><course>283.76413</course><speed>28.953133</speed><geoidheight>-17.7</geoidheight><src>gps</src><sat>16</sat><hdop>0.9</hdop><vdop>1.3</vdop><pdop>1.6</pdop></trkpt>
<trkpt lat="48.75857570096825" lon="-122.46162766264509"><ele>23.933414811535805</ele><time>2016-12-03T07:45:53Z</time><course>0.07149661</course><speed>28.433884</speed><geoidheight>-17.7</geoidheight><src>gps</src><sat>12</sat><hdop>0.6</hdop><vdop>1.6</vdop><pdop>1.8</pdop></trkpt>
<trkpt lat="48.81235892086115" lon="-122.54322864921018"><ele>21.725775005140974</ele><time>2016-12-03T07:51:20Z</time><course>327.2857</course><speed>36.942898</speed><geoidheight>-17.7</geoidheight><src>gps</src><sat>17</sat><hdop>0.8</hdop><vdop>1.4</vdop><pdop>1.6</pdop></trkpt>
<trkpt lat="48.89830966221017" lon="-122.60703299088209"><ele>3.395746957184956</ele><time>2016-12-03T07:57:07Z</time><course>312.63947</course><speed>31.326042</speed><geoidheight>-17.6</geoidheight><src>gps</src><sat>16</sat><hdop>0.4</hdop><vdop>1.4</vdop><pdop>1.7</pdop></trkpt>
</trkseg></trk></gpx>
`
)

type GpxPointCollector struct {
    FileVisits int
    FileVisitBalance int
    TrackVisits int
    TrackVisitBalance int
    TrackSegmentVisits int
    TrackSegmentVisitBalance int
    TrackPointVisits int
    TrackPointVisitBalance int
}

func NewGpsPointCollector() *GpxPointCollector {
    return new(GpxPointCollector)
}

func (gpc *GpxPointCollector) GpxOpen(gpx *Gpx) error {
    gpc.FileVisits++
    gpc.FileVisitBalance++

    return nil
}

func (gpc *GpxPointCollector) GpxClose(gpx *Gpx) error {
    gpc.FileVisitBalance--

    return nil
}

func (gpc *GpxPointCollector) TrackOpen(track *Track) error {
    gpc.TrackVisits++
    gpc.TrackVisitBalance++

    return nil
}

func (gpc *GpxPointCollector) TrackClose(track *Track) error {
    gpc.TrackVisitBalance--

    return nil
}

func (gpc *GpxPointCollector) TrackSegmentOpen(trackSegment *TrackSegment) error {
    gpc.TrackSegmentVisits++
    gpc.TrackSegmentVisitBalance++

    return nil
}

func (gpc *GpxPointCollector) TrackSegmentClose(trackSegment *TrackSegment) error {
    gpc.TrackSegmentVisitBalance--

    return nil
}

func (gpc *GpxPointCollector) TrackPointOpen(trackPoint *TrackPoint) error {
    gpc.TrackPointVisits++
    gpc.TrackPointVisitBalance++

    return nil
}

func (gpc *GpxPointCollector) TrackPointClose(trackPoint *TrackPoint) error {
    gpc.TrackPointVisitBalance--

    return nil
}

func TestFullGpxRead(t *testing.T) {
    b := bytes.NewBufferString(testGpxData)
    gpc := NewGpsPointCollector()
    gp := NewGpxParser(b, gpc)

    if err := gp.Parse(); err != nil {
        log.Panic(err)
    }

    if gpc.FileVisits == 0 {
        t.Errorf("No file visits.")
    } else if gpc.FileVisitBalance != 0 {
        t.Errorf("File visits not balanced.")
    }

    if gpc.TrackVisits == 0 {
        t.Errorf("No track visits.")
    } else if gpc.TrackVisitBalance != 0 {
        t.Errorf("Track visits not balanced.")
    }

    if gpc.TrackSegmentVisits == 0 {
        t.Errorf("No track-segment visits.")
    } else if gpc.TrackSegmentVisitBalance != 0 {
        t.Errorf("Track-segment visits not balanced.")
    }

    if gpc.TrackPointVisits == 0 {
        t.Errorf("No track-point visits.")
    } else if gpc.TrackPointVisitBalance != 0 {
        t.Errorf("Track-point visits not balanced.")
    }

    if gpc.TrackPointVisits != 205 {
        t.Errorf("Points not correctly read.")
    }
}
